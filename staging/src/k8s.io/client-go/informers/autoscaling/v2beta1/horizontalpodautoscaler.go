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

package v2beta1

import (
	"context"
	time "time"

	autoscalingv2beta1 "k8s.io/api/autoscaling/v2beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	wait "k8s.io/apimachinery/pkg/util/wait"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v2beta1 "k8s.io/client-go/listers/autoscaling/v2beta1"
	cache "k8s.io/client-go/tools/cache"
)

// HorizontalPodAutoscalerInformer provides access to a shared informer and lister for
// HorizontalPodAutoscalers.
type HorizontalPodAutoscalerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2beta1.HorizontalPodAutoscalerLister
}

type horizontalPodAutoscalerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	backoff          wait.BackoffManager
	namespace        string
}

// NewHorizontalPodAutoscalerInformer constructs a new informer for HorizontalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHorizontalPodAutoscalerInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, backoff wait.BackoffManager) cache.SharedIndexInformer {
	return NewFilteredHorizontalPodAutoscalerInformer(client, namespace, resyncPeriod, indexers, nil, backoff)
}

// NewFilteredHorizontalPodAutoscalerInformer constructs a new informer for HorizontalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHorizontalPodAutoscalerInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc, backoff wait.BackoffManager) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).Watch(context.TODO(), options)
			},
		},
		&autoscalingv2beta1.HorizontalPodAutoscaler{},
		resyncPeriod,
		indexers,
		backoff,
	)
}

func (f *horizontalPodAutoscalerInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration, backoff wait.BackoffManager) cache.SharedIndexInformer {
	return NewFilteredHorizontalPodAutoscalerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions, f.backoff)
}

func (f *horizontalPodAutoscalerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&autoscalingv2beta1.HorizontalPodAutoscaler{}, f.defaultInformer)
}

func (f *horizontalPodAutoscalerInformer) Lister() v2beta1.HorizontalPodAutoscalerLister {
	return v2beta1.NewHorizontalPodAutoscalerLister(f.Informer().GetIndexer())
}
