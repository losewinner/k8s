/*
Copyright 2024 The Kubernetes Authors.

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

package generic

import (
	"context"
	"errors"
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/apiserver/pkg/admission/plugin/policy/matching"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

// H is the Hook type generated by the source and consumed by the dispatcher.
type sourceFactory[H any] func(informers.SharedInformerFactory, kubernetes.Interface, dynamic.Interface, meta.RESTMapper) Source[H]
type dispatcherFactory[H any] func(authorizer.Authorizer, *matching.Matcher) Dispatcher[H]

// AdmissionPolicyManager is an abstract admission plugin with all the
// infrastructure to define Admit or Validate on-top.
type Plugin[H any] struct {
	*admission.Handler

	sourceFactory     sourceFactory[H]
	dispatcherFactory dispatcherFactory[H]

	source     Source[H]
	dispatcher Dispatcher[H]
	matcher    *matching.Matcher

	informerFactory informers.SharedInformerFactory
	client          kubernetes.Interface
	restMapper      meta.RESTMapper
	dynamicClient   dynamic.Interface
	stopCh          <-chan struct{}
	authorizer      authorizer.Authorizer
	enabled         bool
}

var (
	_ initializer.WantsExternalKubeInformerFactory = &Plugin[any]{}
	_ initializer.WantsExternalKubeClientSet       = &Plugin[any]{}
	_ initializer.WantsRESTMapper                  = &Plugin[any]{}
	_ initializer.WantsDynamicClient               = &Plugin[any]{}
	_ initializer.WantsDrainedNotification         = &Plugin[any]{}
	_ initializer.WantsAuthorizer                  = &Plugin[any]{}
	_ admission.InitializationValidator            = &Plugin[any]{}
)

func NewPlugin[H any](
	handler *admission.Handler,
	sourceFactory sourceFactory[H],
	dispatcherFactory dispatcherFactory[H],
) *Plugin[H] {
	return &Plugin[H]{
		Handler:           handler,
		sourceFactory:     sourceFactory,
		dispatcherFactory: dispatcherFactory,
	}
}

func (c *Plugin[H]) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	c.informerFactory = f
}

func (c *Plugin[H]) SetExternalKubeClientSet(client kubernetes.Interface) {
	c.client = client
}

func (c *Plugin[H]) SetRESTMapper(mapper meta.RESTMapper) {
	c.restMapper = mapper
}

func (c *Plugin[H]) SetDynamicClient(client dynamic.Interface) {
	c.dynamicClient = client
}

func (c *Plugin[H]) SetDrainedNotification(stopCh <-chan struct{}) {
	c.stopCh = stopCh
}

func (c *Plugin[H]) SetAuthorizer(authorizer authorizer.Authorizer) {
	c.authorizer = authorizer
}

func (c *Plugin[H]) SetMatcher(matcher *matching.Matcher) {
	c.matcher = matcher
}

func (c *Plugin[H]) SetEnabled(enabled bool) {
	c.enabled = enabled
}

// ValidateInitialization - once clientset and informer factory are provided, creates and starts the admission controller
func (c *Plugin[H]) ValidateInitialization() error {
	// By default enabled is set to false. It is up to types which embed this
	// struct to set it to true (if feature gate is enabled, or other conditions)
	if !c.enabled {
		return nil
	}
	if c.Handler == nil {
		return errors.New("missing handler")
	}
	if c.informerFactory == nil {
		return errors.New("missing informer factory")
	}
	if c.client == nil {
		return errors.New("missing kubernetes client")
	}
	if c.restMapper == nil {
		return errors.New("missing rest mapper")
	}
	if c.dynamicClient == nil {
		return errors.New("missing dynamic client")
	}
	if c.stopCh == nil {
		return errors.New("missing stop channel")
	}
	if c.authorizer == nil {
		return errors.New("missing authorizer")
	}

	// Use default matcher
	namespaceInformer := c.informerFactory.Core().V1().Namespaces()
	c.matcher = matching.NewMatcher(namespaceInformer.Lister(), c.client)

	if err := c.matcher.ValidateInitialization(); err != nil {
		return err
	}

	c.source = c.sourceFactory(c.informerFactory, c.client, c.dynamicClient, c.restMapper)
	c.dispatcher = c.dispatcherFactory(c.authorizer, c.matcher)

	pluginContext, pluginContextCancel := context.WithCancel(context.Background())
	go func() {
		defer pluginContextCancel()
		<-c.stopCh
	}()

	go func() {
		err := c.source.Run(pluginContext)
		if err != nil && !errors.Is(err, context.Canceled) {
			utilruntime.HandleError(fmt.Errorf("policy source context unexpectedly closed: %v", err))
		}
	}()

	c.SetReadyFunc(func() bool {
		return namespaceInformer.Informer().HasSynced() && c.source.HasSynced()
	})
	return nil
}

func (c *Plugin[H]) Dispatch(
	ctx context.Context,
	a admission.Attributes,
	o admission.ObjectInterfaces,
) (err error) {
	if !c.enabled {
		return nil
	} else if isPolicyResource(a) {
		return nil
	} else if !c.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}

	return c.dispatcher.Dispatch(ctx, a, o, c.source.Hooks())
}

func isPolicyResource(attr admission.Attributes) bool {
	gvk := attr.GetResource()
	if gvk.Group == "admissionregistration.k8s.io" {
		if gvk.Resource == "validatingadmissionpolicies" || gvk.Resource == "validatingadmissionpolicybindings" {
			return true
		}
	}
	return false
}
