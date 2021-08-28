/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strings"

	"k8s.io/klog/v2"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/features"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/apiserver/pkg/storage/value"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
)

// Backend describes the storage servers, the information here should be enough
// for health validations.
type Backend struct {
	// the url of storage backend like: https://etcd.domain:2379
	Server string
	// the required tls config
	TLSConfig *tls.Config
}

// StorageFactory is the interface to locate the storage for a given GroupResource
type StorageFactory interface {
	// New finds the storage destination for the given group and resource. It will
	// return an error if the group has no storage destination configured.
	NewConfig(groupResource schema.GroupResource) (*storagebackend.Config, error)

	// ResourcePrefix returns the overridden resource prefix for the GroupResource
	// This allows for cohabitation of resources with different native types and provides
	// centralized control over the shape of etcd directories
	ResourcePrefix(groupResource schema.GroupResource) string

	// Backends gets all backends for all registered storage destinations.
	// Used for getting all instances for health validations.
	Backends() []Backend
}

// DefaultStorageFactory takes a GroupResource and returns back its storage interface.  This result includes:
// 1. Merged etcd config, including: auth, server locations, prefixes
// 2. Resource encodings for storage: group,version,kind to store as
// 3. Cohabitating default: some resources like hpa are exposed through multiple APIs.  They must agree on 1 and 2
type DefaultStorageFactory struct {
	// StorageConfig describes how to create a storage backend in general.
	// Its authentication information will be used for every storage.Interface returned.
	StorageConfig storagebackend.FactoryConfig

	Overrides map[schema.GroupResource]groupResourceOverrides

	DefaultResourcePrefixes map[schema.GroupResource]string

	// DefaultMediaType is the media type used to store resources. If it is not set, "application/json" is used.
	DefaultMediaType string

	// DefaultSerializer is used to create encoders and decoders for the storage.Interface.
	DefaultSerializer runtime.StorageSerializer

	// ResourceEncodingConfig describes how to encode a particular GroupVersionResource
	ResourceEncodingConfig ResourceEncodingConfig

	// APIResourceConfigSource indicates whether the *storage* is enabled, NOT the API
	// This is discrete from resource enablement because those are separate concerns.  How this source is configured
	// is left to the caller.
	APIResourceConfigSource APIResourceConfigSource

	// newStorageCodecFn exists to be overwritten for unit testing.
	newStorageCodecFn func(opts StorageCodecConfig) (codec runtime.Codec, encodeVersioner runtime.GroupVersioner, err error)
}

type groupResourceOverrides struct {
	// etcdLocation contains the list of "special" locations that are used for particular GroupResources
	// These are merged on top of the StorageConfig when requesting the storage.Interface for a given GroupResource
	etcdLocation []string
	// etcdPrefix is the base location for a GroupResource.
	etcdPrefix string
	// etcdResourcePrefix is the location to use to store a particular type under the `etcdPrefix` location
	// If empty, the default mapping is used.  If the default mapping doesn't contain an entry, it will use
	// the ToLowered name of the resource, not including the group.
	etcdResourcePrefix string
	// mediaType is the desired serializer to choose. If empty, the default is chosen.
	mediaType string
	// serializer contains the list of "special" serializers for a GroupResource.  Resource=* means for the entire group
	serializer runtime.StorageSerializer
	// cohabitatingResources keeps track of which resources must be stored together.  This happens when we have multiple ways
	// of exposing one set of concepts.  autoscaling.HPA and extensions.HPA as a for instance
	// The order of the slice matters!  It is the priority order of lookup for finding a storage location
	cohabitatingResources []schema.GroupResource
	// encoderDecoratorFn is optional and may wrap the provided encoder prior to being serialized.
	encoderDecoratorFn func(runtime.Encoder) runtime.Encoder
	// decoderDecoratorFn is optional and may wrap the provided decoders (can add new decoders). The order of
	// returned decoders will be priority for attempt to decode.
	decoderDecoratorFn func([]runtime.Decoder) []runtime.Decoder
	// transformer is optional and shall encrypt that resource at rest.
	transformer value.Transformer
	// disablePaging will prevent paging on the provided resource.
	disablePaging bool
}

// Apply overrides the provided config and options if the override has a value in that position
func (o groupResourceOverrides) Apply(config *storagebackend.Config, options *StorageCodecConfig) {
	if len(o.etcdLocation) > 0 {
		config.Transport.ServerList = o.etcdLocation
	}
	if len(o.etcdPrefix) > 0 {
		config.Prefix = o.etcdPrefix
	}

	if len(o.mediaType) > 0 {
		options.StorageMediaType = o.mediaType
	}
	if o.serializer != nil {
		options.StorageSerializer = o.serializer
	}
	if o.encoderDecoratorFn != nil {
		options.EncoderDecoratorFn = o.encoderDecoratorFn
	}
	if o.decoderDecoratorFn != nil {
		options.DecoderDecoratorFn = o.decoderDecoratorFn
	}
	if o.transformer != nil {
		config.Transformer = o.transformer
	}
	if o.disablePaging {
		config.Paging = false
	}
}

var _ StorageFactory = &DefaultStorageFactory{}

const AllResources = "*"

func NewDefaultStorageFactory(
	config storagebackend.FactoryConfig,
	defaultMediaType string,
	defaultSerializer runtime.StorageSerializer,
	resourceEncodingConfig ResourceEncodingConfig,
	resourceConfig APIResourceConfigSource,
	specialDefaultResourcePrefixes map[schema.GroupResource]string,
) *DefaultStorageFactory {
	config.Paging = utilfeature.DefaultFeatureGate.Enabled(features.APIListChunking)
	if len(defaultMediaType) == 0 {
		defaultMediaType = runtime.ContentTypeJSON
	}
	return &DefaultStorageFactory{
		StorageConfig:           config,
		Overrides:               map[schema.GroupResource]groupResourceOverrides{},
		DefaultMediaType:        defaultMediaType,
		DefaultSerializer:       defaultSerializer,
		ResourceEncodingConfig:  resourceEncodingConfig,
		APIResourceConfigSource: resourceConfig,
		DefaultResourcePrefixes: specialDefaultResourcePrefixes,

		newStorageCodecFn: NewStorageCodec,
	}
}

func (s *DefaultStorageFactory) SetEtcdLocation(groupResource schema.GroupResource, location []string) {
	overrides := s.Overrides[groupResource]
	overrides.etcdLocation = location
	s.Overrides[groupResource] = overrides
}

func (s *DefaultStorageFactory) SetEtcdPrefix(groupResource schema.GroupResource, prefix string) {
	overrides := s.Overrides[groupResource]
	overrides.etcdPrefix = prefix
	s.Overrides[groupResource] = overrides
}

// SetDisableAPIListChunking allows a specific resource to disable paging at the storage layer, to prevent
// exposure of key names in continuations. This may be overridden by feature gates.
func (s *DefaultStorageFactory) SetDisableAPIListChunking(groupResource schema.GroupResource) {
	overrides := s.Overrides[groupResource]
	overrides.disablePaging = true
	s.Overrides[groupResource] = overrides
}

// SetResourceEtcdPrefix sets the prefix for a resource, but not the base-dir.  You'll end up in `etcdPrefix/resourceEtcdPrefix`.
func (s *DefaultStorageFactory) SetResourceEtcdPrefix(groupResource schema.GroupResource, prefix string) {
	overrides := s.Overrides[groupResource]
	overrides.etcdResourcePrefix = prefix
	s.Overrides[groupResource] = overrides
}

func (s *DefaultStorageFactory) SetSerializer(groupResource schema.GroupResource, mediaType string, serializer runtime.StorageSerializer) {
	overrides := s.Overrides[groupResource]
	overrides.mediaType = mediaType
	overrides.serializer = serializer
	s.Overrides[groupResource] = overrides
}

func (s *DefaultStorageFactory) SetTransformer(groupResource schema.GroupResource, transformer value.Transformer) {
	overrides := s.Overrides[groupResource]
	overrides.transformer = transformer
	s.Overrides[groupResource] = overrides
}

// AddCohabitatingResources links resources together the order of the slice matters!  its the priority order of lookup for finding a storage location
func (s *DefaultStorageFactory) AddCohabitatingResources(groupResources ...schema.GroupResource) {
	for _, groupResource := range groupResources {
		overrides := s.Overrides[groupResource]
		overrides.cohabitatingResources = groupResources
		s.Overrides[groupResource] = overrides
	}
}

func (s *DefaultStorageFactory) AddSerializationChains(encoderDecoratorFn func(runtime.Encoder) runtime.Encoder, decoderDecoratorFn func([]runtime.Decoder) []runtime.Decoder, groupResources ...schema.GroupResource) {
	for _, groupResource := range groupResources {
		overrides := s.Overrides[groupResource]
		overrides.encoderDecoratorFn = encoderDecoratorFn
		overrides.decoderDecoratorFn = decoderDecoratorFn
		s.Overrides[groupResource] = overrides
	}
}

func getAllResourcesAlias(resource schema.GroupResource) schema.GroupResource {
	return schema.GroupResource{Group: resource.Group, Resource: AllResources}
}

func (s *DefaultStorageFactory) getStorageGroupResource(groupResource schema.GroupResource) schema.GroupResource {
	for _, potentialStorageResource := range s.Overrides[groupResource].cohabitatingResources {
		if s.APIResourceConfigSource.AnyVersionForGroupEnabled(potentialStorageResource.Group) {
			return potentialStorageResource
		}
	}

	return groupResource
}

// New finds the storage destination for the given group and resource. It will
// return an error if the group has no storage destination configured.
func (s *DefaultStorageFactory) NewConfig(groupResource schema.GroupResource) (*storagebackend.Config, error) {
	chosenStorageResource := s.getStorageGroupResource(groupResource)

	// operate on copy
	storageConfig := storagebackend.Config{FactoryConfig: s.StorageConfig, GroupResourceString: groupResource.String()}
	codecConfig := StorageCodecConfig{
		StorageMediaType:  s.DefaultMediaType,
		StorageSerializer: s.DefaultSerializer,
	}

	if override, ok := s.Overrides[getAllResourcesAlias(chosenStorageResource)]; ok {
		override.Apply(&storageConfig, &codecConfig)
	}
	if override, ok := s.Overrides[chosenStorageResource]; ok {
		override.Apply(&storageConfig, &codecConfig)
	}

	var err error
	codecConfig.StorageVersion, err = s.ResourceEncodingConfig.StorageEncodingFor(chosenStorageResource)
	if err != nil {
		return nil, err
	}
	codecConfig.MemoryVersion, err = s.ResourceEncodingConfig.InMemoryEncodingFor(groupResource)
	if err != nil {
		return nil, err
	}
	codecConfig.Config = storageConfig

	storageConfig.Codec, storageConfig.EncodeVersioner, err = s.newStorageCodecFn(codecConfig)
	if err != nil {
		return nil, err
	}
	klog.V(3).Infof("storing %v in %v, reading as %v from %#v", groupResource, codecConfig.StorageVersion, codecConfig.MemoryVersion, codecConfig.Config)

	return &storageConfig, nil
}

// Backends returns all backends for all registered storage destinations.
// Used for getting all instances for health validations.
func (s *DefaultStorageFactory) Backends() []Backend {
	servers := sets.NewString(s.StorageConfig.Transport.ServerList...)

	for _, overrides := range s.Overrides {
		servers.Insert(overrides.etcdLocation...)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	if len(s.StorageConfig.Transport.CertFile) > 0 && len(s.StorageConfig.Transport.KeyFile) > 0 {
		cert, err := tls.LoadX509KeyPair(s.StorageConfig.Transport.CertFile, s.StorageConfig.Transport.KeyFile)
		if err != nil {
			klog.Errorf("failed to load key pair while getting backends: %s", err)
		} else {
			tlsConfig.Certificates = []tls.Certificate{cert}
		}
	}
	if len(s.StorageConfig.Transport.TrustedCAFile) > 0 {
		if caCert, err := ioutil.ReadFile(s.StorageConfig.Transport.TrustedCAFile); err != nil {
			klog.Errorf("failed to read ca file while getting backends: %s", err)
		} else {
			caPool := x509.NewCertPool()
			caPool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = caPool
			tlsConfig.InsecureSkipVerify = false
		}
	}

	backends := []Backend{}
	for server := range servers {
		backends = append(backends, Backend{
			Server: server,
			// We can't share TLSConfig across different backends to avoid races.
			// For more details see: http://pr.k8s.io/59338
			TLSConfig: tlsConfig.Clone(),
		})
	}
	return backends
}

func (s *DefaultStorageFactory) ResourcePrefix(groupResource schema.GroupResource) string {
	chosenStorageResource := s.getStorageGroupResource(groupResource)
	groupOverride := s.Overrides[getAllResourcesAlias(chosenStorageResource)]
	exactResourceOverride := s.Overrides[chosenStorageResource]

	etcdResourcePrefix := s.DefaultResourcePrefixes[chosenStorageResource]
	if len(groupOverride.etcdResourcePrefix) > 0 {
		etcdResourcePrefix = groupOverride.etcdResourcePrefix
	}
	if len(exactResourceOverride.etcdResourcePrefix) > 0 {
		etcdResourcePrefix = exactResourceOverride.etcdResourcePrefix
	}
	if len(etcdResourcePrefix) == 0 {
		etcdResourcePrefix = strings.ToLower(chosenStorageResource.Resource)
	}

	return etcdResourcePrefix
}
