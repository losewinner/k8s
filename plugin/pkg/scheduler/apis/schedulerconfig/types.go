/*
Copyright 2015 The Kubernetes Authors.

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

package schedulerconfig

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KubeSchedulerConfiguration struct {
	metav1.TypeMeta

	// port is the port that the scheduler's http service runs on.
	Port int32
	// address is the IP address to serve on.
	Address string
	// algorithmProvider is the scheduling algorithm provider to use.
	AlgorithmProvider string
	// policyConfigFile is the filepath to the scheduler policy configuration.
	PolicyConfigFile string
	// enableProfiling enables profiling via web interface.
	EnableProfiling bool
	// enableContentionProfiling enables lock contention profiling, if enableProfiling is true.
	EnableContentionProfiling bool
	// contentType is contentType of requests sent to apiserver.
	ContentType string
	// kubeAPIQPS is the QPS to use while talking with kubernetes apiserver.
	KubeAPIQPS float32
	// kubeAPIBurst is the QPS burst to use while talking with kubernetes apiserver.
	KubeAPIBurst int32
	// schedulerName is name of the scheduler, used to select which pods
	// will be processed by this scheduler, based on pod's "spec.SchedulerName".
	SchedulerName string
	// RequiredDuringScheduling affinity is not symmetric, but there is an implicit PreferredDuringScheduling affinity rule
	// corresponding to every RequiredDuringScheduling affinity rule.
	// HardPodAffinitySymmetricWeight represents the weight of implicit PreferredDuringScheduling affinity rule, in the range 0-100.
	HardPodAffinitySymmetricWeight int
	// Indicate the "all topologies" set for empty topologyKey when it's used for PreferredDuringScheduling pod anti-affinity.
	// DEPRECATED: This is no longer used.
	FailureDomains string
	// leaderElection defines the configuration of leader election client.
	LeaderElection componentconfig.LeaderElectionConfiguration
	// LockObjectNamespace defines the namespace of the lock object
	LockObjectNamespace string
	// LockObjectName defines the lock object name
	LockObjectName string
	// PolicyConfigMapName is the name of the ConfigMap object that specifies
	// the scheduler's policy config. If UseLegacyPolicyConfig is true, scheduler
	// uses PolicyConfigFile. If UseLegacyPolicyConfig is false and
	// PolicyConfigMapName is not empty, the ConfigMap object with this name must
	// exist in PolicyConfigMapNamespace before scheduler initialization.
	PolicyConfigMapName string
	// PolicyConfigMapNamespace is the namespace where the above policy config map
	// is located. If none is provided default system namespace ("kube-system")
	// will be used.
	PolicyConfigMapNamespace string
	// UseLegacyPolicyConfig tells the scheduler to ignore Policy ConfigMap and
	// to use PolicyConfigFile if available.
	UseLegacyPolicyConfig bool
}

const (
	// "kube-system" is the default scheduler lock object namespace
	SchedulerDefaultLockObjectNamespace string = "kube-system"

	// "kube-scheduler" is the default scheduler lock object name
	SchedulerDefaultLockObjectName = "kube-scheduler"
)
