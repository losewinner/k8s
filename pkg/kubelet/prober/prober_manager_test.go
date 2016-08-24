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

package prober

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api"
	kubecontainer "k8s.io/kubernetes/pkg/kubelet/container"
	"k8s.io/kubernetes/pkg/kubelet/prober/results"
	"k8s.io/kubernetes/pkg/probe"
	"k8s.io/kubernetes/pkg/types"
	"k8s.io/kubernetes/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/util/wait"
)

func init() {
	runtime.ReallyCrash = true
}

var defaultProbe *api.Probe = &api.Probe{
	Handler: api.Handler{
		Exec: &api.ExecAction{},
	},
	TimeoutSeconds:   1,
	PeriodSeconds:    1,
	SuccessThreshold: 1,
	FailureThreshold: 3,
}

func TestAddRemovePods(t *testing.T) {
	noProbePod := api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID: "no_probe_pod",
		},
		Spec: api.PodSpec{
			Containers: []api.Container{{
				Name: "no_probe1",
			}, {
				Name: "no_probe2",
			}},
		},
	}

	probePod := api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID: "probe_pod",
		},
		Spec: api.PodSpec{
			Containers: []api.Container{{
				Name: "no_probe1",
			}, {
				Name:           "readiness",
				ReadinessProbe: defaultProbe,
			}, {
				Name: "no_probe2",
			}, {
				Name:          "liveness",
				LivenessProbe: defaultProbe,
			}},
		},
	}

	m := newTestManager()
	defer cleanup(t, m)
	if err := expectProbes(m, nil); err != nil {
		t.Error(err)
	}

	// Adding a pod with no probes should be a no-op.
	m.AddPod(&noProbePod)
	if err := expectProbes(m, nil); err != nil {
		t.Error(err)
	}

	// Adding a pod with probes.
	m.AddPod(&probePod)
	probePaths := []probeKey{
		{"probe_pod", "readiness", readiness},
		{"probe_pod", "liveness", liveness},
	}
	if err := expectProbes(m, probePaths); err != nil {
		t.Error(err)
	}

	// Removing un-probed pod.
	m.RemovePod(&noProbePod)
	if err := expectProbes(m, probePaths); err != nil {
		t.Error(err)
	}

	// Removing probed pod.
	m.RemovePod(&probePod)
	if err := waitForWorkerExit(m, probePaths); err != nil {
		t.Fatal(err)
	}
	if err := expectProbes(m, nil); err != nil {
		t.Error(err)
	}

	// Removing already removed pods should be a no-op.
	m.RemovePod(&probePod)
	if err := expectProbes(m, nil); err != nil {
		t.Error(err)
	}
}

func TestCleanupPods(t *testing.T) {
	m := newTestManager()
	defer cleanup(t, m)
	podToCleanup := api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID: "pod_cleanup",
		},
		Spec: api.PodSpec{
			Containers: []api.Container{{
				Name:           "prober1",
				ReadinessProbe: defaultProbe,
			}, {
				Name:          "prober2",
				LivenessProbe: defaultProbe,
			}},
		},
	}
	podToKeep := api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID: "pod_keep",
		},
		Spec: api.PodSpec{
			Containers: []api.Container{{
				Name:           "prober1",
				ReadinessProbe: defaultProbe,
			}, {
				Name:          "prober2",
				LivenessProbe: defaultProbe,
			}},
		},
	}
	m.AddPod(&podToCleanup)
	m.AddPod(&podToKeep)

	m.CleanupPods([]*api.Pod{&podToKeep})

	removedProbes := []probeKey{
		{"pod_cleanup", "prober1", readiness},
		{"pod_cleanup", "prober2", liveness},
	}
	expectedProbes := []probeKey{
		{"pod_keep", "prober1", readiness},
		{"pod_keep", "prober2", liveness},
	}
	if err := waitForWorkerExit(m, removedProbes); err != nil {
		t.Fatal(err)
	}
	if err := expectProbes(m, expectedProbes); err != nil {
		t.Error(err)
	}
}

func TestCleanupRepeated(t *testing.T) {
	m := newTestManager()
	defer cleanup(t, m)
	podTemplate := api.Pod{
		Spec: api.PodSpec{
			Containers: []api.Container{{
				Name:           "prober1",
				ReadinessProbe: defaultProbe,
				LivenessProbe:  defaultProbe,
			}},
		},
	}

	const numTestPods = 100
	for i := 0; i < numTestPods; i++ {
		pod := podTemplate
		pod.UID = types.UID(strconv.Itoa(i))
		m.AddPod(&pod)
	}

	for i := 0; i < 10; i++ {
		m.CleanupPods([]*api.Pod{})
	}
}

func TestUpdatePodStatus(t *testing.T) {
	unprobed := api.ContainerStatus{
		Name:        "unprobed_container",
		ContainerID: "test://unprobed_container_id",
		State: api.ContainerState{
			Running: &api.ContainerStateRunning{},
		},
	}
	probedReady := api.ContainerStatus{
		Name:        "probed_container_ready",
		ContainerID: "test://probed_container_ready_id",
		State: api.ContainerState{
			Running: &api.ContainerStateRunning{},
		},
	}
	probedPending := api.ContainerStatus{
		Name:        "probed_container_pending",
		ContainerID: "test://probed_container_pending_id",
		State: api.ContainerState{
			Running: &api.ContainerStateRunning{},
		},
	}
	probedUnready := api.ContainerStatus{
		Name:        "probed_container_unready",
		ContainerID: "test://probed_container_unready_id",
		State: api.ContainerState{
			Running: &api.ContainerStateRunning{},
		},
	}
	terminated := api.ContainerStatus{
		Name:        "terminated_container",
		ContainerID: "test://terminated_container_id",
		State: api.ContainerState{
			Terminated: &api.ContainerStateTerminated{},
		},
	}
	podStatus := api.PodStatus{
		Phase: api.PodRunning,
		ContainerStatuses: []api.ContainerStatus{
			unprobed, probedReady, probedPending, probedUnready, terminated,
		},
	}

	m := newTestManager()
	// no cleanup: using fake workers.

	// Setup probe "workers" and cached results.
	m.workers = map[probeKey]*worker{
		probeKey{testPodUID, unprobed.Name, liveness}:       {},
		probeKey{testPodUID, probedReady.Name, readiness}:   {},
		probeKey{testPodUID, probedPending.Name, readiness}: {},
		probeKey{testPodUID, probedUnready.Name, readiness}: {},
		probeKey{testPodUID, terminated.Name, readiness}:    {},
	}
	m.readinessManager.Set(kubecontainer.ParseContainerID(probedReady.ContainerID), results.Success, &api.Pod{})
	m.readinessManager.Set(kubecontainer.ParseContainerID(probedUnready.ContainerID), results.Failure, &api.Pod{})
	m.readinessManager.Set(kubecontainer.ParseContainerID(terminated.ContainerID), results.Success, &api.Pod{})

	m.UpdatePodStatus(testPodUID, &podStatus)

	expectedReadiness := map[probeKey]bool{
		probeKey{testPodUID, unprobed.Name, readiness}:      true,
		probeKey{testPodUID, probedReady.Name, readiness}:   true,
		probeKey{testPodUID, probedPending.Name, readiness}: false,
		probeKey{testPodUID, probedUnready.Name, readiness}: false,
		probeKey{testPodUID, terminated.Name, readiness}:    false,
	}
	for _, c := range podStatus.ContainerStatuses {
		expected, ok := expectedReadiness[probeKey{testPodUID, c.Name, readiness}]
		if !ok {
			t.Fatalf("Missing expectation for test case: %v", c.Name)
		}
		if expected != c.Ready {
			t.Errorf("Unexpected readiness for container %v: Expected %v but got %v",
				c.Name, expected, c.Ready)
		}
	}
}

func TestUpdateReadiness(t *testing.T) {
	testPod := getTestPod()
	setTestProbe(testPod, readiness, api.Probe{})
	m := newTestManager()
	defer cleanup(t, m)

	// Start syncing readiness without leaking goroutine.
	stopCh := make(chan struct{})
	go wait.Until(m.updateReadiness, 0, stopCh)
	defer func() {
		close(stopCh)
		// Send an update to exit updateReadiness()
		m.readinessManager.Set(kubecontainer.ContainerID{}, results.Success, &api.Pod{})
	}()

	exec := syncExecProber{}
	exec.set(probe.Success, nil)
	m.prober.exec = &exec

	m.statusManager.SetPodStatus(testPod, getTestRunningStatus(), nil)

	m.AddPod(testPod)
	probePaths := []probeKey{{testPodUID, testContainerName, readiness}}
	if err := expectProbes(m, probePaths); err != nil {
		t.Error(err)
	}

	// Wait for ready status.
	if err := waitForReadyStatus(m, true); err != nil {
		t.Error(err)
	}

	// Prober fails.
	exec.set(probe.Failure, nil)

	// Wait for failed status.
	if err := waitForReadyStatus(m, false); err != nil {
		t.Error(err)
	}
}

func expectProbes(m *manager, expectedProbes []probeKey) error {
	m.workerLock.RLock()
	defer m.workerLock.RUnlock()

	var unexpected []probeKey
	missing := make([]probeKey, len(expectedProbes))
	copy(missing, expectedProbes)

outer:
	for probePath := range m.workers {
		for i, expectedPath := range missing {
			if probePath == expectedPath {
				missing = append(missing[:i], missing[i+1:]...)
				continue outer
			}
		}
		unexpected = append(unexpected, probePath)
	}

	if len(missing) == 0 && len(unexpected) == 0 {
		return nil // Yay!
	}

	return fmt.Errorf("Unexpected probes: %v; Missing probes: %v;", unexpected, missing)
}

const interval = 1 * time.Second

// Wait for the given workers to exit & clean up.
func waitForWorkerExit(m *manager, workerPaths []probeKey) error {
	for _, w := range workerPaths {
		condition := func() (bool, error) {
			_, exists := m.getWorker(w.podUID, w.containerName, w.probeType)
			return !exists, nil
		}
		if exited, _ := condition(); exited {
			continue // Already exited, no need to poll.
		}
		glog.Infof("Polling %v", w)
		if err := wait.Poll(interval, wait.ForeverTestTimeout, condition); err != nil {
			return err
		}
	}

	return nil
}

// Wait for the given workers to exit & clean up.
func waitForReadyStatus(m *manager, ready bool) error {
	condition := func() (bool, error) {
		status, ok := m.statusManager.GetPodStatus(testPodUID)
		if !ok {
			return false, fmt.Errorf("status not found: %q", testPodUID)
		}
		if len(status.ContainerStatuses) != 1 {
			return false, fmt.Errorf("expected single container, found %d", len(status.ContainerStatuses))
		}
		if status.ContainerStatuses[0].ContainerID != testContainerID.String() {
			return false, fmt.Errorf("expected container %q, found %q",
				testContainerID, status.ContainerStatuses[0].ContainerID)
		}
		return status.ContainerStatuses[0].Ready == ready, nil
	}
	glog.Infof("Polling for ready state %v", ready)
	if err := wait.Poll(interval, wait.ForeverTestTimeout, condition); err != nil {
		return err
	}

	return nil
}

// cleanup running probes to avoid leaking goroutines.
func cleanup(t *testing.T, m *manager) {
	m.CleanupPods(nil)

	condition := func() (bool, error) {
		workerCount := m.workerCount()
		if workerCount > 0 {
			glog.Infof("Waiting for %d workers to exit...", workerCount)
		}
		return workerCount == 0, nil
	}
	if exited, _ := condition(); exited {
		return // Already exited, no need to poll.
	}
	if err := wait.Poll(interval, wait.ForeverTestTimeout, condition); err != nil {
		t.Fatalf("Error during cleanup: %v", err)
	}
}
