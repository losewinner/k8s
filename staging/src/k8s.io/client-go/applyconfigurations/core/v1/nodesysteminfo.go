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

// NodeSystemInfoApplyConfiguration represents an declarative configuration of the NodeSystemInfo type for use
// with apply.
type NodeSystemInfoApplyConfiguration struct {
	MachineID               *string `json:"machineID,omitempty"`
	SystemUUID              *string `json:"systemUUID,omitempty"`
	BootID                  *string `json:"bootID,omitempty"`
	KernelVersion           *string `json:"kernelVersion,omitempty"`
	OSImage                 *string `json:"osImage,omitempty"`
	ContainerRuntimeVersion *string `json:"containerRuntimeVersion,omitempty"`
	KubeletVersion          *string `json:"kubeletVersion,omitempty"`
	KubeProxyVersion        *string `json:"kubeProxyVersion,omitempty"`
	OperatingSystem         *string `json:"operatingSystem,omitempty"`
	Architecture            *string `json:"architecture,omitempty"`
}

// NodeSystemInfoApplyConfiguration constructs an declarative configuration of the NodeSystemInfo type for use with
// apply.
func NodeSystemInfo() *NodeSystemInfoApplyConfiguration {
	return &NodeSystemInfoApplyConfiguration{}
}

// SetMachineID sets the MachineID field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetMachineID(value string) *NodeSystemInfoApplyConfiguration {
	b.MachineID = &value
	return b
}

// RemoveMachineID removes the MachineID field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveMachineID() *NodeSystemInfoApplyConfiguration {
	b.MachineID = nil
	return b
}

// GetMachineID gets the MachineID field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetMachineID() (value string, ok bool) {
	if v := b.MachineID; v != nil {
		return *v, true
	}
	return value, false
}

// SetSystemUUID sets the SystemUUID field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetSystemUUID(value string) *NodeSystemInfoApplyConfiguration {
	b.SystemUUID = &value
	return b
}

// RemoveSystemUUID removes the SystemUUID field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveSystemUUID() *NodeSystemInfoApplyConfiguration {
	b.SystemUUID = nil
	return b
}

// GetSystemUUID gets the SystemUUID field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetSystemUUID() (value string, ok bool) {
	if v := b.SystemUUID; v != nil {
		return *v, true
	}
	return value, false
}

// SetBootID sets the BootID field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetBootID(value string) *NodeSystemInfoApplyConfiguration {
	b.BootID = &value
	return b
}

// RemoveBootID removes the BootID field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveBootID() *NodeSystemInfoApplyConfiguration {
	b.BootID = nil
	return b
}

// GetBootID gets the BootID field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetBootID() (value string, ok bool) {
	if v := b.BootID; v != nil {
		return *v, true
	}
	return value, false
}

// SetKernelVersion sets the KernelVersion field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetKernelVersion(value string) *NodeSystemInfoApplyConfiguration {
	b.KernelVersion = &value
	return b
}

// RemoveKernelVersion removes the KernelVersion field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveKernelVersion() *NodeSystemInfoApplyConfiguration {
	b.KernelVersion = nil
	return b
}

// GetKernelVersion gets the KernelVersion field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetKernelVersion() (value string, ok bool) {
	if v := b.KernelVersion; v != nil {
		return *v, true
	}
	return value, false
}

// SetOSImage sets the OSImage field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetOSImage(value string) *NodeSystemInfoApplyConfiguration {
	b.OSImage = &value
	return b
}

// RemoveOSImage removes the OSImage field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveOSImage() *NodeSystemInfoApplyConfiguration {
	b.OSImage = nil
	return b
}

// GetOSImage gets the OSImage field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetOSImage() (value string, ok bool) {
	if v := b.OSImage; v != nil {
		return *v, true
	}
	return value, false
}

// SetContainerRuntimeVersion sets the ContainerRuntimeVersion field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetContainerRuntimeVersion(value string) *NodeSystemInfoApplyConfiguration {
	b.ContainerRuntimeVersion = &value
	return b
}

// RemoveContainerRuntimeVersion removes the ContainerRuntimeVersion field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveContainerRuntimeVersion() *NodeSystemInfoApplyConfiguration {
	b.ContainerRuntimeVersion = nil
	return b
}

// GetContainerRuntimeVersion gets the ContainerRuntimeVersion field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetContainerRuntimeVersion() (value string, ok bool) {
	if v := b.ContainerRuntimeVersion; v != nil {
		return *v, true
	}
	return value, false
}

// SetKubeletVersion sets the KubeletVersion field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetKubeletVersion(value string) *NodeSystemInfoApplyConfiguration {
	b.KubeletVersion = &value
	return b
}

// RemoveKubeletVersion removes the KubeletVersion field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveKubeletVersion() *NodeSystemInfoApplyConfiguration {
	b.KubeletVersion = nil
	return b
}

// GetKubeletVersion gets the KubeletVersion field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetKubeletVersion() (value string, ok bool) {
	if v := b.KubeletVersion; v != nil {
		return *v, true
	}
	return value, false
}

// SetKubeProxyVersion sets the KubeProxyVersion field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetKubeProxyVersion(value string) *NodeSystemInfoApplyConfiguration {
	b.KubeProxyVersion = &value
	return b
}

// RemoveKubeProxyVersion removes the KubeProxyVersion field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveKubeProxyVersion() *NodeSystemInfoApplyConfiguration {
	b.KubeProxyVersion = nil
	return b
}

// GetKubeProxyVersion gets the KubeProxyVersion field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetKubeProxyVersion() (value string, ok bool) {
	if v := b.KubeProxyVersion; v != nil {
		return *v, true
	}
	return value, false
}

// SetOperatingSystem sets the OperatingSystem field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetOperatingSystem(value string) *NodeSystemInfoApplyConfiguration {
	b.OperatingSystem = &value
	return b
}

// RemoveOperatingSystem removes the OperatingSystem field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveOperatingSystem() *NodeSystemInfoApplyConfiguration {
	b.OperatingSystem = nil
	return b
}

// GetOperatingSystem gets the OperatingSystem field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetOperatingSystem() (value string, ok bool) {
	if v := b.OperatingSystem; v != nil {
		return *v, true
	}
	return value, false
}

// SetArchitecture sets the Architecture field in the declarative configuration to the given value.
func (b *NodeSystemInfoApplyConfiguration) SetArchitecture(value string) *NodeSystemInfoApplyConfiguration {
	b.Architecture = &value
	return b
}

// RemoveArchitecture removes the Architecture field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) RemoveArchitecture() *NodeSystemInfoApplyConfiguration {
	b.Architecture = nil
	return b
}

// GetArchitecture gets the Architecture field from the declarative configuration.
func (b *NodeSystemInfoApplyConfiguration) GetArchitecture() (value string, ok bool) {
	if v := b.Architecture; v != nil {
		return *v, true
	}
	return value, false
}

// NodeSystemInfoList represents a listAlias of NodeSystemInfoApplyConfiguration.
type NodeSystemInfoList []*NodeSystemInfoApplyConfiguration

// NodeSystemInfoList represents a map of NodeSystemInfoApplyConfiguration.
type NodeSystemInfoMap map[string]NodeSystemInfoApplyConfiguration
