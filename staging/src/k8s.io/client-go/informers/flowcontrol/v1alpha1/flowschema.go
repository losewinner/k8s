/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	flowcontrolv1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	wait "k8s.io/apimachinery/pkg/util/wait"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/flowcontrol/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
)

// FlowSchemaInformer provides access to a shared informer and lister for
// FlowSchemas.
type FlowSchemaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FlowSchemaLister
}

type flowSchemaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	backoff          wait.BackoffManager
}

// NewFlowSchemaInformer constructs a new informer for FlowSchema type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFlowSchemaInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, backoff wait.BackoffManager) cache.SharedIndexInformer {
	return NewFilteredFlowSchemaInformer(client, resyncPeriod, indexers, nil, backoff)
}

// NewFilteredFlowSchemaInformer constructs a new informer for FlowSchema type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFlowSchemaInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc, backoff wait.BackoffManager) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1alpha1().FlowSchemas().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1alpha1().FlowSchemas().Watch(context.TODO(), options)
			},
		},
		&flowcontrolv1alpha1.FlowSchema{},
		resyncPeriod,
		indexers,
		backoff,
	)
}

func (f *flowSchemaInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration, backoff wait.BackoffManager) cache.SharedIndexInformer {
	return NewFilteredFlowSchemaInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions, f.backoff)
}

func (f *flowSchemaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flowcontrolv1alpha1.FlowSchema{}, f.defaultInformer)
}

func (f *flowSchemaInformer) Lister() v1alpha1.FlowSchemaLister {
	return v1alpha1.NewFlowSchemaLister(f.Informer().GetIndexer())
}
