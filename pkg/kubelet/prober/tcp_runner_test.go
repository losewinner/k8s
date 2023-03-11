/*
Copyright 2023 The Kubernetes Authors.

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

package prober

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"testing"
)

func TestNewTCPProbeRunner(t *testing.T) {
	runner := newTCPProbeRunner()

	pod := newPod(v1.ProbeHandler{TCPSocket: &v1.TCPSocketAction{
		Port: intstr.FromInt(5000),
	}})

	_ = runner.sync(pod.Spec.Containers[0], pod.Status, liveness)
	assert.Equal(t, runner.host, pod.Status.PodIP)

	// change ip address
	pod.Status.PodIP = "192.168.1.11"
	_ = runner.sync(pod.Spec.Containers[0], pod.Status, liveness)
	assert.Equal(t, runner.host, pod.Status.PodIP)
}
