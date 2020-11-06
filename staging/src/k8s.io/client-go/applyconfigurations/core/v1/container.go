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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// ContainerApplyConfiguration represents an declarative configuration of the Container type for use
// with apply.
type ContainerApplyConfiguration struct {
	Name                     *string                                 `json:"name,omitempty"`
	Image                    *string                                 `json:"image,omitempty"`
	Command                  *[]string                               `json:"command,omitempty"`
	Args                     *[]string                               `json:"args,omitempty"`
	WorkingDir               *string                                 `json:"workingDir,omitempty"`
	Ports                    *ContainerPortList                      `json:"ports,omitempty"`
	EnvFrom                  *EnvFromSourceList                      `json:"envFrom,omitempty"`
	Env                      *EnvVarList                             `json:"env,omitempty"`
	Resources                *ResourceRequirementsApplyConfiguration `json:"resources,omitempty"`
	VolumeMounts             *VolumeMountList                        `json:"volumeMounts,omitempty"`
	VolumeDevices            *VolumeDeviceList                       `json:"volumeDevices,omitempty"`
	LivenessProbe            *ProbeApplyConfiguration                `json:"livenessProbe,omitempty"`
	ReadinessProbe           *ProbeApplyConfiguration                `json:"readinessProbe,omitempty"`
	StartupProbe             *ProbeApplyConfiguration                `json:"startupProbe,omitempty"`
	Lifecycle                *LifecycleApplyConfiguration            `json:"lifecycle,omitempty"`
	TerminationMessagePath   *string                                 `json:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy *corev1.TerminationMessagePolicy        `json:"terminationMessagePolicy,omitempty"`
	ImagePullPolicy          *corev1.PullPolicy                      `json:"imagePullPolicy,omitempty"`
	SecurityContext          *SecurityContextApplyConfiguration      `json:"securityContext,omitempty"`
	Stdin                    *bool                                   `json:"stdin,omitempty"`
	StdinOnce                *bool                                   `json:"stdinOnce,omitempty"`
	TTY                      *bool                                   `json:"tty,omitempty"`
}

// ContainerApplyConfiguration constructs an declarative configuration of the Container type for use with
// apply.
func Container() *ContainerApplyConfiguration {
	return &ContainerApplyConfiguration{}
}

// SetName sets the Name field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetName(value string) *ContainerApplyConfiguration {
	b.Name = &value
	return b
}

// RemoveName removes the Name field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveName() *ContainerApplyConfiguration {
	b.Name = nil
	return b
}

// GetName gets the Name field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetName() (value string, ok bool) {
	if v := b.Name; v != nil {
		return *v, true
	}
	return value, false
}

// SetImage sets the Image field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetImage(value string) *ContainerApplyConfiguration {
	b.Image = &value
	return b
}

// RemoveImage removes the Image field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveImage() *ContainerApplyConfiguration {
	b.Image = nil
	return b
}

// GetImage gets the Image field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetImage() (value string, ok bool) {
	if v := b.Image; v != nil {
		return *v, true
	}
	return value, false
}

// SetCommand sets the Command field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetCommand(value []string) *ContainerApplyConfiguration {
	b.Command = &value
	return b
}

// RemoveCommand removes the Command field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveCommand() *ContainerApplyConfiguration {
	b.Command = nil
	return b
}

// GetCommand gets the Command field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetCommand() (value []string, ok bool) {
	if v := b.Command; v != nil {
		return *v, true
	}
	return value, false
}

// SetArgs sets the Args field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetArgs(value []string) *ContainerApplyConfiguration {
	b.Args = &value
	return b
}

// RemoveArgs removes the Args field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveArgs() *ContainerApplyConfiguration {
	b.Args = nil
	return b
}

// GetArgs gets the Args field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetArgs() (value []string, ok bool) {
	if v := b.Args; v != nil {
		return *v, true
	}
	return value, false
}

// SetWorkingDir sets the WorkingDir field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetWorkingDir(value string) *ContainerApplyConfiguration {
	b.WorkingDir = &value
	return b
}

// RemoveWorkingDir removes the WorkingDir field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveWorkingDir() *ContainerApplyConfiguration {
	b.WorkingDir = nil
	return b
}

// GetWorkingDir gets the WorkingDir field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetWorkingDir() (value string, ok bool) {
	if v := b.WorkingDir; v != nil {
		return *v, true
	}
	return value, false
}

// SetPorts sets the Ports field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetPorts(value ContainerPortList) *ContainerApplyConfiguration {
	b.Ports = &value
	return b
}

// RemovePorts removes the Ports field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemovePorts() *ContainerApplyConfiguration {
	b.Ports = nil
	return b
}

// GetPorts gets the Ports field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetPorts() (value ContainerPortList, ok bool) {
	if v := b.Ports; v != nil {
		return *v, true
	}
	return value, false
}

// SetEnvFrom sets the EnvFrom field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetEnvFrom(value EnvFromSourceList) *ContainerApplyConfiguration {
	b.EnvFrom = &value
	return b
}

// RemoveEnvFrom removes the EnvFrom field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveEnvFrom() *ContainerApplyConfiguration {
	b.EnvFrom = nil
	return b
}

// GetEnvFrom gets the EnvFrom field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetEnvFrom() (value EnvFromSourceList, ok bool) {
	if v := b.EnvFrom; v != nil {
		return *v, true
	}
	return value, false
}

// SetEnv sets the Env field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetEnv(value EnvVarList) *ContainerApplyConfiguration {
	b.Env = &value
	return b
}

// RemoveEnv removes the Env field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveEnv() *ContainerApplyConfiguration {
	b.Env = nil
	return b
}

// GetEnv gets the Env field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetEnv() (value EnvVarList, ok bool) {
	if v := b.Env; v != nil {
		return *v, true
	}
	return value, false
}

// SetResources sets the Resources field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetResources(value *ResourceRequirementsApplyConfiguration) *ContainerApplyConfiguration {
	b.Resources = value
	return b
}

// RemoveResources removes the Resources field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveResources() *ContainerApplyConfiguration {
	b.Resources = nil
	return b
}

// GetResources gets the Resources field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetResources() (value *ResourceRequirementsApplyConfiguration, ok bool) {
	return b.Resources, b.Resources != nil
}

// SetVolumeMounts sets the VolumeMounts field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetVolumeMounts(value VolumeMountList) *ContainerApplyConfiguration {
	b.VolumeMounts = &value
	return b
}

// RemoveVolumeMounts removes the VolumeMounts field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveVolumeMounts() *ContainerApplyConfiguration {
	b.VolumeMounts = nil
	return b
}

// GetVolumeMounts gets the VolumeMounts field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetVolumeMounts() (value VolumeMountList, ok bool) {
	if v := b.VolumeMounts; v != nil {
		return *v, true
	}
	return value, false
}

// SetVolumeDevices sets the VolumeDevices field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetVolumeDevices(value VolumeDeviceList) *ContainerApplyConfiguration {
	b.VolumeDevices = &value
	return b
}

// RemoveVolumeDevices removes the VolumeDevices field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveVolumeDevices() *ContainerApplyConfiguration {
	b.VolumeDevices = nil
	return b
}

// GetVolumeDevices gets the VolumeDevices field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetVolumeDevices() (value VolumeDeviceList, ok bool) {
	if v := b.VolumeDevices; v != nil {
		return *v, true
	}
	return value, false
}

// SetLivenessProbe sets the LivenessProbe field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetLivenessProbe(value *ProbeApplyConfiguration) *ContainerApplyConfiguration {
	b.LivenessProbe = value
	return b
}

// RemoveLivenessProbe removes the LivenessProbe field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveLivenessProbe() *ContainerApplyConfiguration {
	b.LivenessProbe = nil
	return b
}

// GetLivenessProbe gets the LivenessProbe field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetLivenessProbe() (value *ProbeApplyConfiguration, ok bool) {
	return b.LivenessProbe, b.LivenessProbe != nil
}

// SetReadinessProbe sets the ReadinessProbe field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetReadinessProbe(value *ProbeApplyConfiguration) *ContainerApplyConfiguration {
	b.ReadinessProbe = value
	return b
}

// RemoveReadinessProbe removes the ReadinessProbe field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveReadinessProbe() *ContainerApplyConfiguration {
	b.ReadinessProbe = nil
	return b
}

// GetReadinessProbe gets the ReadinessProbe field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetReadinessProbe() (value *ProbeApplyConfiguration, ok bool) {
	return b.ReadinessProbe, b.ReadinessProbe != nil
}

// SetStartupProbe sets the StartupProbe field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetStartupProbe(value *ProbeApplyConfiguration) *ContainerApplyConfiguration {
	b.StartupProbe = value
	return b
}

// RemoveStartupProbe removes the StartupProbe field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveStartupProbe() *ContainerApplyConfiguration {
	b.StartupProbe = nil
	return b
}

// GetStartupProbe gets the StartupProbe field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetStartupProbe() (value *ProbeApplyConfiguration, ok bool) {
	return b.StartupProbe, b.StartupProbe != nil
}

// SetLifecycle sets the Lifecycle field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetLifecycle(value *LifecycleApplyConfiguration) *ContainerApplyConfiguration {
	b.Lifecycle = value
	return b
}

// RemoveLifecycle removes the Lifecycle field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveLifecycle() *ContainerApplyConfiguration {
	b.Lifecycle = nil
	return b
}

// GetLifecycle gets the Lifecycle field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetLifecycle() (value *LifecycleApplyConfiguration, ok bool) {
	return b.Lifecycle, b.Lifecycle != nil
}

// SetTerminationMessagePath sets the TerminationMessagePath field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetTerminationMessagePath(value string) *ContainerApplyConfiguration {
	b.TerminationMessagePath = &value
	return b
}

// RemoveTerminationMessagePath removes the TerminationMessagePath field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveTerminationMessagePath() *ContainerApplyConfiguration {
	b.TerminationMessagePath = nil
	return b
}

// GetTerminationMessagePath gets the TerminationMessagePath field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetTerminationMessagePath() (value string, ok bool) {
	if v := b.TerminationMessagePath; v != nil {
		return *v, true
	}
	return value, false
}

// SetTerminationMessagePolicy sets the TerminationMessagePolicy field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetTerminationMessagePolicy(value corev1.TerminationMessagePolicy) *ContainerApplyConfiguration {
	b.TerminationMessagePolicy = &value
	return b
}

// RemoveTerminationMessagePolicy removes the TerminationMessagePolicy field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveTerminationMessagePolicy() *ContainerApplyConfiguration {
	b.TerminationMessagePolicy = nil
	return b
}

// GetTerminationMessagePolicy gets the TerminationMessagePolicy field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetTerminationMessagePolicy() (value corev1.TerminationMessagePolicy, ok bool) {
	if v := b.TerminationMessagePolicy; v != nil {
		return *v, true
	}
	return value, false
}

// SetImagePullPolicy sets the ImagePullPolicy field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetImagePullPolicy(value corev1.PullPolicy) *ContainerApplyConfiguration {
	b.ImagePullPolicy = &value
	return b
}

// RemoveImagePullPolicy removes the ImagePullPolicy field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveImagePullPolicy() *ContainerApplyConfiguration {
	b.ImagePullPolicy = nil
	return b
}

// GetImagePullPolicy gets the ImagePullPolicy field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetImagePullPolicy() (value corev1.PullPolicy, ok bool) {
	if v := b.ImagePullPolicy; v != nil {
		return *v, true
	}
	return value, false
}

// SetSecurityContext sets the SecurityContext field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetSecurityContext(value *SecurityContextApplyConfiguration) *ContainerApplyConfiguration {
	b.SecurityContext = value
	return b
}

// RemoveSecurityContext removes the SecurityContext field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveSecurityContext() *ContainerApplyConfiguration {
	b.SecurityContext = nil
	return b
}

// GetSecurityContext gets the SecurityContext field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetSecurityContext() (value *SecurityContextApplyConfiguration, ok bool) {
	return b.SecurityContext, b.SecurityContext != nil
}

// SetStdin sets the Stdin field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetStdin(value bool) *ContainerApplyConfiguration {
	b.Stdin = &value
	return b
}

// RemoveStdin removes the Stdin field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveStdin() *ContainerApplyConfiguration {
	b.Stdin = nil
	return b
}

// GetStdin gets the Stdin field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetStdin() (value bool, ok bool) {
	if v := b.Stdin; v != nil {
		return *v, true
	}
	return value, false
}

// SetStdinOnce sets the StdinOnce field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetStdinOnce(value bool) *ContainerApplyConfiguration {
	b.StdinOnce = &value
	return b
}

// RemoveStdinOnce removes the StdinOnce field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveStdinOnce() *ContainerApplyConfiguration {
	b.StdinOnce = nil
	return b
}

// GetStdinOnce gets the StdinOnce field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetStdinOnce() (value bool, ok bool) {
	if v := b.StdinOnce; v != nil {
		return *v, true
	}
	return value, false
}

// SetTTY sets the TTY field in the declarative configuration to the given value.
func (b *ContainerApplyConfiguration) SetTTY(value bool) *ContainerApplyConfiguration {
	b.TTY = &value
	return b
}

// RemoveTTY removes the TTY field from the declarative configuration.
func (b *ContainerApplyConfiguration) RemoveTTY() *ContainerApplyConfiguration {
	b.TTY = nil
	return b
}

// GetTTY gets the TTY field from the declarative configuration.
func (b *ContainerApplyConfiguration) GetTTY() (value bool, ok bool) {
	if v := b.TTY; v != nil {
		return *v, true
	}
	return value, false
}

// ContainerList represents a listAlias of ContainerApplyConfiguration.
type ContainerList []*ContainerApplyConfiguration

// ContainerList represents a map of ContainerApplyConfiguration.
type ContainerMap map[string]ContainerApplyConfiguration
