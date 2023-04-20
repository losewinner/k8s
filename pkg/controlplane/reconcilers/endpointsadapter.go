/*
Copyright 2019 The Kubernetes Authors.

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

package reconcilers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	discoveryclient "k8s.io/client-go/kubernetes/typed/discovery/v1"
	"k8s.io/klog/v2"
	utilnet "k8s.io/utils/net"
)

// EndpointsAdapter provides a simple interface for reading and writing both
// Endpoints and EndpointSlices for an EndpointReconciler.
type EndpointsAdapter struct {
	endpointClient      corev1client.EndpointsGetter
	endpointSliceClient discoveryclient.EndpointSlicesGetter

	serviceNamespace string
	serviceName      string
}

// NewEndpointsAdapter returns a new EndpointsAdapter
func NewEndpointsAdapter(endpointClient corev1client.EndpointsGetter, endpointSliceClient discoveryclient.EndpointSlicesGetter, serviceNamespace, serviceName string) *EndpointsAdapter {
	return &EndpointsAdapter{
		endpointClient:      endpointClient,
		endpointSliceClient: endpointSliceClient,

		serviceNamespace: serviceNamespace,
		serviceName:      serviceName,
	}
}

// Get returns the IPs from the existing apiserver Endpoints/EndpointSlice objects. If an
// error (beside "not found") occurs fetching the data, that error will be returned.
func (adapter *EndpointsAdapter) Get() (sets.Set[string], error) {
	ips := sets.New[string]()

	endpoints, err := adapter.endpointClient.Endpoints(adapter.serviceNamespace).Get(context.TODO(), adapter.serviceName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return ips, nil
		}
		return nil, err
	}

	if len(endpoints.Subsets) == 1 {
		for _, addr := range endpoints.Subsets[0].Addresses {
			ips.Insert(addr.IP)
		}
	}
	return ips, nil
}

// Sync updates the apiserver Endpoints/EndpointSlice objects with the new set of IPs. If
// reconcilePorts is true it will also ensure that the objects have the correct ports. If
// an error occurs while updating, that error will be returned.
func (adapter *EndpointsAdapter) Sync(ips sets.Set[string], endpointPorts []corev1.EndpointPort, reconcilePorts bool) error {
	// Sync Endpoints
	currentEndpoints, err := adapter.endpointClient.Endpoints(adapter.serviceNamespace).Get(context.TODO(), adapter.serviceName, metav1.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		currentEndpoints = nil
	}
	endpoints := adapter.updateEndpoints(currentEndpoints, ips, endpointPorts, reconcilePorts)

	if currentEndpoints == nil {
		_, err = adapter.endpointClient.Endpoints(endpoints.Namespace).Create(context.TODO(), endpoints, metav1.CreateOptions{})
	} else if !apiequality.Semantic.DeepEqual(currentEndpoints.Subsets, endpoints.Subsets) ||
		!apiequality.Semantic.DeepEqual(currentEndpoints.Labels, endpoints.Labels) {
		klog.Warningf("Resetting endpoints for master service %q to %v", endpoints.Name, endpoints)
		_, err = adapter.endpointClient.Endpoints(endpoints.Namespace).Update(context.TODO(), endpoints, metav1.UpdateOptions{})
	}
	if err != nil {
		return err
	}

	// Sync EndpointSlice
	currentEndpointSlice, err := adapter.endpointSliceClient.EndpointSlices(adapter.serviceNamespace).Get(context.TODO(), adapter.serviceName, metav1.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		currentEndpointSlice = nil
	}

	endpointSlice := endpointSliceFromEndpoints(endpoints)

	// required for transition from IP to IPv4 address type.
	if currentEndpointSlice != nil && currentEndpointSlice.AddressType != endpointSlice.AddressType {
		err = adapter.endpointSliceClient.EndpointSlices(endpointSlice.Namespace).Delete(context.TODO(), endpointSlice.Name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
		currentEndpointSlice = nil
	}

	if currentEndpointSlice == nil {
		_, err = adapter.endpointSliceClient.EndpointSlices(endpointSlice.Namespace).Create(context.TODO(), endpointSlice, metav1.CreateOptions{})
	} else if !apiequality.Semantic.DeepEqual(currentEndpointSlice.Endpoints, endpointSlice.Endpoints) ||
		!apiequality.Semantic.DeepEqual(currentEndpointSlice.Ports, endpointSlice.Ports) ||
		!apiequality.Semantic.DeepEqual(currentEndpointSlice.Labels, endpointSlice.Labels) {
		_, err = adapter.endpointSliceClient.EndpointSlices(endpointSlice.Namespace).Update(context.TODO(), endpointSlice, metav1.UpdateOptions{})
	}
	return err
}

// updateEndpoints updates endpoints to reflect ips and (optionally) endpointPorts
func (adapter *EndpointsAdapter) updateEndpoints(currentEndpoints *corev1.Endpoints, ips sets.Set[string], endpointPorts []corev1.EndpointPort, reconcilePorts bool) *corev1.Endpoints {
	var endpoints *corev1.Endpoints

	if currentEndpoints != nil {
		endpoints = currentEndpoints.DeepCopy()
	} else {
		endpoints = &corev1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:      adapter.serviceName,
				Namespace: adapter.serviceNamespace,
			},
		}
	}

	// Ensure correct labels
	endpoints.Labels = map[string]string{
		discovery.LabelSkipMirror: "true",
	}
	// Ensure correct format
	if len(endpoints.Subsets) != 1 {
		endpoints.Subsets = make([]corev1.EndpointSubset, 1)
	}

	// Set addresses
	sortedIPs := sets.List(ips)
	endpoints.Subsets[0].Addresses = make([]corev1.EndpointAddress, len(sortedIPs))
	for i := range sortedIPs {
		endpoints.Subsets[0].Addresses[i].IP = sortedIPs[i]
	}

	// Set ports
	if len(endpoints.Subsets[0].Ports) == 0 || (reconcilePorts && !apiequality.Semantic.DeepEqual(endpoints.Subsets[0].Ports, endpointPorts)) {
		endpoints.Subsets[0].Ports = endpointPorts
	}

	return endpoints
}

// endpointSliceFromEndpoints generates an EndpointSlice from an Endpoints
// resource.
func endpointSliceFromEndpoints(endpoints *corev1.Endpoints) *discovery.EndpointSlice {
	endpointSlice := &discovery.EndpointSlice{}
	endpointSlice.Name = endpoints.Name
	endpointSlice.Namespace = endpoints.Namespace
	endpointSlice.Labels = map[string]string{discovery.LabelServiceName: endpoints.Name}

	// TODO: Add support for dual stack here (and in the rest of
	// EndpointsAdapter).
	endpointSlice.AddressType = discovery.AddressTypeIPv4

	if len(endpoints.Subsets) > 0 {
		subset := endpoints.Subsets[0]
		for i := range subset.Ports {
			endpointSlice.Ports = append(endpointSlice.Ports, discovery.EndpointPort{
				Port:     &subset.Ports[i].Port,
				Name:     &subset.Ports[i].Name,
				Protocol: &subset.Ports[i].Protocol,
			})
		}

		if allAddressesIPv6(append(subset.Addresses, subset.NotReadyAddresses...)) {
			endpointSlice.AddressType = discovery.AddressTypeIPv6
		}

		endpointSlice.Endpoints = append(endpointSlice.Endpoints, getEndpointsFromAddresses(subset.Addresses, endpointSlice.AddressType, true)...)
		endpointSlice.Endpoints = append(endpointSlice.Endpoints, getEndpointsFromAddresses(subset.NotReadyAddresses, endpointSlice.AddressType, false)...)
	}

	return endpointSlice
}

// getEndpointsFromAddresses returns a list of Endpoints from addresses that
// match the provided address type.
func getEndpointsFromAddresses(addresses []corev1.EndpointAddress, addressType discovery.AddressType, ready bool) []discovery.Endpoint {
	endpoints := []discovery.Endpoint{}
	isIPv6AddressType := addressType == discovery.AddressTypeIPv6

	for _, address := range addresses {
		if utilnet.IsIPv6String(address.IP) == isIPv6AddressType {
			endpoints = append(endpoints, endpointFromAddress(address, ready))
		}
	}

	return endpoints
}

// endpointFromAddress generates an Endpoint from an EndpointAddress resource.
func endpointFromAddress(address corev1.EndpointAddress, ready bool) discovery.Endpoint {
	ep := discovery.Endpoint{
		Addresses:  []string{address.IP},
		Conditions: discovery.EndpointConditions{Ready: &ready},
	}

	if address.NodeName != nil {
		ep.NodeName = address.NodeName
	}

	return ep
}

// allAddressesIPv6 returns true if all provided addresses are IPv6.
func allAddressesIPv6(addresses []corev1.EndpointAddress) bool {
	if len(addresses) == 0 {
		return false
	}

	for _, address := range addresses {
		if !utilnet.IsIPv6String(address.IP) {
			return false
		}
	}

	return true
}
