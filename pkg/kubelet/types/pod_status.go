/*
Copyright 2018 The Kubernetes Authors.

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

package types

import (
	v1 "k8s.io/api/core/v1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/features"
)

// PodConditionsByKubelet is the list of pod conditions owned by kubelet
var PodConditionsByKubelet = []v1.PodConditionType{
	v1.PodScheduled,
	v1.PodReady,
	v1.PodInitialized,
	v1.ContainersReady,
}

// PodConditionByKubelet returns if the pod condition type is owned by kubelet
func PodConditionByKubelet(conditionType v1.PodConditionType) bool {
	for _, c := range PodConditionsByKubelet {
		if c == conditionType {
			return true
		}
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.PodHasNetworkCondition) {
		if conditionType == PodHasNetwork {
			return true
		}
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.PodDisruptionConditions) {
		if PodFailureCondition(conditionType) {
			return true
		}
	}
	return false
}

// PodFailureCondition returns the if the pod condition type is a pod failure condition
func PodFailureCondition(conditionType v1.PodConditionType) bool {
	for _, podFailureType := range PodFailureConditions() {
		if podFailureType == conditionType {
			return true
		}
	}
	return false
}

// PodFailureConditions returns the list of pod failure conditions
func PodFailureConditions() []v1.PodConditionType {
	if utilfeature.DefaultFeatureGate.Enabled(features.PodDisruptionConditions) {
		return []v1.PodConditionType{
			v1.AlphaNoCompatGuaranteeDisruptionTarget,
			ResourceExhausted,
		}
	}
	return []v1.PodConditionType{}
}
