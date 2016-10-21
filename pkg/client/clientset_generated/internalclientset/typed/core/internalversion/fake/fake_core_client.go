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

package fake

import (
	internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/core/internalversion"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeCoreInternalversion struct {
	*core.Fake
}

func (c *FakeCoreInternalversion) ComponentStatuses() internalversion.ComponentStatusInterface {
	return &FakeComponentStatuses{c}
}

func (c *FakeCoreInternalversion) ConfigMaps(namespace string) internalversion.ConfigMapInterface {
	return &FakeConfigMaps{c, namespace}
}

func (c *FakeCoreInternalversion) Endpoints(namespace string) internalversion.EndpointsInterface {
	return &FakeEndpoints{c, namespace}
}

func (c *FakeCoreInternalversion) Events(namespace string) internalversion.EventInterface {
	return &FakeEvents{c, namespace}
}

func (c *FakeCoreInternalversion) LimitRanges(namespace string) internalversion.LimitRangeInterface {
	return &FakeLimitRanges{c, namespace}
}

func (c *FakeCoreInternalversion) Namespaces() internalversion.NamespaceInterface {
	return &FakeNamespaces{c}
}

func (c *FakeCoreInternalversion) Nodes() internalversion.NodeInterface {
	return &FakeNodes{c}
}

func (c *FakeCoreInternalversion) PersistentVolumes() internalversion.PersistentVolumeInterface {
	return &FakePersistentVolumes{c}
}

func (c *FakeCoreInternalversion) PersistentVolumeClaims(namespace string) internalversion.PersistentVolumeClaimInterface {
	return &FakePersistentVolumeClaims{c, namespace}
}

func (c *FakeCoreInternalversion) Pods(namespace string) internalversion.PodInterface {
	return &FakePods{c, namespace}
}

func (c *FakeCoreInternalversion) PodTemplates(namespace string) internalversion.PodTemplateInterface {
	return &FakePodTemplates{c, namespace}
}

func (c *FakeCoreInternalversion) ReplicationControllers(namespace string) internalversion.ReplicationControllerInterface {
	return &FakeReplicationControllers{c, namespace}
}

func (c *FakeCoreInternalversion) ResourceQuotas(namespace string) internalversion.ResourceQuotaInterface {
	return &FakeResourceQuotas{c, namespace}
}

func (c *FakeCoreInternalversion) Secrets(namespace string) internalversion.SecretInterface {
	return &FakeSecrets{c, namespace}
}

func (c *FakeCoreInternalversion) Services(namespace string) internalversion.ServiceInterface {
	return &FakeServices{c, namespace}
}

func (c *FakeCoreInternalversion) ServiceAccounts(namespace string) internalversion.ServiceAccountInterface {
	return &FakeServiceAccounts{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreInternalversion) RESTClient() restclient.Interface {
	var ret *restclient.RESTClient
	return ret
}
