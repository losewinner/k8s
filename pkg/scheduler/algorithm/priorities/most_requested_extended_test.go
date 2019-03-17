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

package priorities

import (
	"reflect"
	"testing"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

const (
	NvidiaGpu              = "nvidia.com/gpu"
	AnyUserDefinedResource = "foo.com/bar"
)

func TestMostRequestedExtended(t *testing.T) {
	labels1 := map[string]string{
		"foo": "bar",
		"baz": "blah",
	}
	labels2 := map[string]string{
		"bar": "foo",
		"baz": "blah",
	}
	noResources := v1.PodSpec{
		Containers: []v1.Container{},
	}
	cpuOnly := v1.PodSpec{
		NodeName: "machine1",
		Containers: []v1.Container{
			{
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("1000m"),
						v1.ResourceMemory: resource.MustParse("0"),
					},
				},
			},
			{
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("2000m"),
						v1.ResourceMemory: resource.MustParse("0"),
					},
				},
			},
		},
	}
	cpuOnly2 := cpuOnly
	cpuOnly2.NodeName = "machine2"
	cpuAndMemoryAndGpu := v1.PodSpec{
		NodeName: "machine2",
		Containers: []v1.Container{
			{
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("1000m"),
						v1.ResourceMemory: resource.MustParse("2000"),
					},
				},
			},
			{
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("2000m"),
						v1.ResourceMemory: resource.MustParse("3000"),
						NvidiaGpu:         resource.MustParse("1"),
					},
				},
			},
		},
	}
	bigCPUAndMemory := v1.PodSpec{
		NodeName: "machine1",
		Containers: []v1.Container{
			{
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("2000m"),
						v1.ResourceMemory: resource.MustParse("4000"),
					},
				},
			},
			{
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("3000m"),
						v1.ResourceMemory: resource.MustParse("5000"),
						NvidiaGpu:         resource.MustParse("4"),
					},
				},
			},
		},
	}
	tests := []struct {
		pod          *v1.Pod
		pods         []*v1.Pod
		nodes        []*v1.Node
		expectedList schedulerapi.HostPriorityList
		name         string
	}{
		{
			/*
			   Node1 scores (used resources) on 0-10 scale
			   CPU Score: (0 * 10) / 4000 = 0
			   Memory Score: (0 * 10) / 10000 = 0
			   GPU Score: (0 * 10) / 4 = 0
			   Node1 Score: (0 + 0 + 0) / 3 = 0

			   Node2 scores (used resources) on 0-10 scale
			   CPU Score: (0 * 10) / 4000 = 0
			   Memory Score: (0 * 10) / 10000 = 0
			   GPU Score: (0 * 10) / 4 = 0
			   Node2 Score: (0 + 0 + 0) / 3 = 0
			*/
			pod: &v1.Pod{Spec: noResources},
			nodes: []*v1.Node{
				makeNodeWithAllResources("machine1", 4000, 10000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
				makeNodeWithAllResources("machine2", 4000, 10000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
			},
			expectedList: []schedulerapi.HostPriority{{Host: "machine1", Score: 0}, {Host: "machine2", Score: 0}},
			name:         "nothing scheduled, nothing requested",
		},
		{
			/*
			   Node1 scores on 0-10 scale
			   CPU Score: (3000 * 10) / 4000 = 7.5
			   Memory Score: (5000 * 10) / 10000 = 5
			   GPU Score: (1 * 10) / 4 = 2.5
			   Node1 Score: (7.5 + 5 + 2.5) / 3 = 5

			   Node2 scores on 0-10 scale
			   CPU Score: (3000 * 10) / 6000 = 5
			   Memory Score: (5000 * 10) / 10000 = 5
			   GPU Score: (1 * 10) / 4 = 2.5
			   Node2 Score: (5 + 5 + 2.5) / 3 = 4.5 -> 4
			*/
			pod: &v1.Pod{Spec: cpuAndMemoryAndGpu},
			nodes: []*v1.Node{
				makeNodeWithAllResources("machine1", 4000, 10000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
				makeNodeWithAllResources("machine2", 6000, 10000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
			},
			expectedList: []schedulerapi.HostPriority{{Host: "machine1", Score: 5}, {Host: "machine2", Score: 4}},
			name:         "nothing scheduled, resources requested, differently sized machines",
		},
		{
			/*
			   Node1 scores on 0-10 scale
			   CPU Score: (6000 * 10) / 10000 = 6
			   Memory Score: (0 * 10) / 20000 = 0
			   GPU Score: (0 * 10) / 4 = 0
			   Node1 Score: (6 + 0 + 0) / 3 = 2

			   Node2 scores on 0-10 scale
			   CPU Score: (6000 * 10) / 10000 = 6
			   Memory Score: (5000 * 10) / 20000 = 2.5
			   GPU Score: (1 * 10) / 4 = 2.5
			   Node2 Score: (6 + 2.5 + 2.5) / 3 = 3.6667 -> 3
			*/
			pod: &v1.Pod{Spec: noResources},
			nodes: []*v1.Node{
				makeNodeWithAllResources("machine1", 10000, 20000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
				makeNodeWithAllResources("machine2", 10000, 20000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
			},
			expectedList: []schedulerapi.HostPriority{{Host: "machine1", Score: 2}, {Host: "machine2", Score: 3}},
			name:         "no resources requested, pods scheduled with resources",
			pods: []*v1.Pod{
				{Spec: cpuOnly, ObjectMeta: metav1.ObjectMeta{Labels: labels2}},
				{Spec: cpuOnly, ObjectMeta: metav1.ObjectMeta{Labels: labels1}},
				{Spec: cpuOnly2, ObjectMeta: metav1.ObjectMeta{Labels: labels1}},
				{Spec: cpuAndMemoryAndGpu, ObjectMeta: metav1.ObjectMeta{Labels: labels1}},
			},
		},
		{
			/*
			   Node1 scores on 0-10 scale
			   CPU Score: (6000 * 10) / 10000 = 6
			   Memory Score: (5000 * 10) / 20000 = 2.5
			   GPU Score: (1 * 10) / 4 = 2.5
			   Node1 Score: (6 + 2.5 + 2.5) / 3 = 3.6667 -> 3

			   Node2 scores on 0-10 scale
			   CPU Score: (6000 * 10) / 10000 = 6
			   Memory Score: (10000 * 10) / 20000 = 5
			   GPU Score: (2 * 10) / 4 = 5
			   Node2 Score: (6 + 5 + 5) / 3 = 5.33333 -> 5
			*/
			pod: &v1.Pod{Spec: cpuAndMemoryAndGpu},
			nodes: []*v1.Node{
				makeNodeWithAllResources("machine1", 10000, 20000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
				makeNodeWithAllResources("machine2", 10000, 20000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
			},
			expectedList: []schedulerapi.HostPriority{{Host: "machine1", Score: 3}, {Host: "machine2", Score: 5}},
			name:         "resources requested, pods scheduled with resources",
			pods: []*v1.Pod{
				{Spec: cpuOnly},
				{Spec: cpuAndMemoryAndGpu},
			},
		},
		{
			/*
			   Node1 scores on 0-10 scale
			   CPU Score: 5000 > 4000 return 0
			   Memory Score: (9000 * 10) / 10000 = 9
			   GPU Score: (4 * 10) / 4 = 10
			   Node1 Score: (0 + 9 + 10) / 3 = 6.33333 -> 6

			   Node2 scores on 0-10 scale
			   CPU Score: (5000 * 10) / 10000 = 5
			   Memory Score: 9000 > 8000 return 0
			   GPU Score: (4 * 10) / 4 = 10
			   Node2 Score: (5 + 0 + 10) / 3 = 5
			*/
			pod: &v1.Pod{Spec: bigCPUAndMemory},
			nodes: []*v1.Node{
				makeNodeWithAllResources("machine1", 4000, 10000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
				makeNodeWithAllResources("machine2", 10000, 8000, 0, map[v1.ResourceName]int64{NvidiaGpu: 4}),
			},
			expectedList: []schedulerapi.HostPriority{{Host: "machine1", Score: 6}, {Host: "machine2", Score: 5}},
			name:         "resources requested, nothing scheduled",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nodeNameToInfo := schedulernodeinfo.CreateNodeNameToInfoMap(test.pods, test.nodes)
			list, err := priorityFunction(MostExtendedRequestedPriorityMap, nil, nil)(test.pod, nodeNameToInfo, test.nodes)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.expectedList, list) {
				t.Errorf("expected %#v, got %#v", test.expectedList, list)
			}
		})
	}
}
