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

// VolumeSourceApplyConfiguration represents an declarative configuration of the VolumeSource type for use
// with apply.
type VolumeSourceApplyConfiguration struct {
	HostPath              *HostPathVolumeSourceApplyConfiguration              `json:"hostPath,omitempty"`
	EmptyDir              *EmptyDirVolumeSourceApplyConfiguration              `json:"emptyDir,omitempty"`
	GCEPersistentDisk     *GCEPersistentDiskVolumeSourceApplyConfiguration     `json:"gcePersistentDisk,omitempty"`
	AWSElasticBlockStore  *AWSElasticBlockStoreVolumeSourceApplyConfiguration  `json:"awsElasticBlockStore,omitempty"`
	GitRepo               *GitRepoVolumeSourceApplyConfiguration               `json:"gitRepo,omitempty"`
	Secret                *SecretVolumeSourceApplyConfiguration                `json:"secret,omitempty"`
	NFS                   *NFSVolumeSourceApplyConfiguration                   `json:"nfs,omitempty"`
	ISCSI                 *ISCSIVolumeSourceApplyConfiguration                 `json:"iscsi,omitempty"`
	Glusterfs             *GlusterfsVolumeSourceApplyConfiguration             `json:"glusterfs,omitempty"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSourceApplyConfiguration `json:"persistentVolumeClaim,omitempty"`
	RBD                   *RBDVolumeSourceApplyConfiguration                   `json:"rbd,omitempty"`
	FlexVolume            *FlexVolumeSourceApplyConfiguration                  `json:"flexVolume,omitempty"`
	Cinder                *CinderVolumeSourceApplyConfiguration                `json:"cinder,omitempty"`
	CephFS                *CephFSVolumeSourceApplyConfiguration                `json:"cephfs,omitempty"`
	Flocker               *FlockerVolumeSourceApplyConfiguration               `json:"flocker,omitempty"`
	DownwardAPI           *DownwardAPIVolumeSourceApplyConfiguration           `json:"downwardAPI,omitempty"`
	FC                    *FCVolumeSourceApplyConfiguration                    `json:"fc,omitempty"`
	AzureFile             *AzureFileVolumeSourceApplyConfiguration             `json:"azureFile,omitempty"`
	ConfigMap             *ConfigMapVolumeSourceApplyConfiguration             `json:"configMap,omitempty"`
	VsphereVolume         *VsphereVirtualDiskVolumeSourceApplyConfiguration    `json:"vsphereVolume,omitempty"`
	Quobyte               *QuobyteVolumeSourceApplyConfiguration               `json:"quobyte,omitempty"`
	AzureDisk             *AzureDiskVolumeSourceApplyConfiguration             `json:"azureDisk,omitempty"`
	PhotonPersistentDisk  *PhotonPersistentDiskVolumeSourceApplyConfiguration  `json:"photonPersistentDisk,omitempty"`
	Projected             *ProjectedVolumeSourceApplyConfiguration             `json:"projected,omitempty"`
	PortworxVolume        *PortworxVolumeSourceApplyConfiguration              `json:"portworxVolume,omitempty"`
	ScaleIO               *ScaleIOVolumeSourceApplyConfiguration               `json:"scaleIO,omitempty"`
	StorageOS             *StorageOSVolumeSourceApplyConfiguration             `json:"storageos,omitempty"`
	CSI                   *CSIVolumeSourceApplyConfiguration                   `json:"csi,omitempty"`
	Ephemeral             *EphemeralVolumeSourceApplyConfiguration             `json:"ephemeral,omitempty"`
}

// VolumeSourceApplyConfiguration constructs an declarative configuration of the VolumeSource type for use with
// apply.
func VolumeSource() *VolumeSourceApplyConfiguration {
	return &VolumeSourceApplyConfiguration{}
}

// SetHostPath sets the HostPath field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetHostPath(value *HostPathVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.HostPath = value
	return b
}

// RemoveHostPath removes the HostPath field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveHostPath() *VolumeSourceApplyConfiguration {
	b.HostPath = nil
	return b
}

// GetHostPath gets the HostPath field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetHostPath() (value *HostPathVolumeSourceApplyConfiguration, ok bool) {
	return b.HostPath, b.HostPath != nil
}

// SetEmptyDir sets the EmptyDir field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetEmptyDir(value *EmptyDirVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.EmptyDir = value
	return b
}

// RemoveEmptyDir removes the EmptyDir field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveEmptyDir() *VolumeSourceApplyConfiguration {
	b.EmptyDir = nil
	return b
}

// GetEmptyDir gets the EmptyDir field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetEmptyDir() (value *EmptyDirVolumeSourceApplyConfiguration, ok bool) {
	return b.EmptyDir, b.EmptyDir != nil
}

// SetGCEPersistentDisk sets the GCEPersistentDisk field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetGCEPersistentDisk(value *GCEPersistentDiskVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.GCEPersistentDisk = value
	return b
}

// RemoveGCEPersistentDisk removes the GCEPersistentDisk field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveGCEPersistentDisk() *VolumeSourceApplyConfiguration {
	b.GCEPersistentDisk = nil
	return b
}

// GetGCEPersistentDisk gets the GCEPersistentDisk field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetGCEPersistentDisk() (value *GCEPersistentDiskVolumeSourceApplyConfiguration, ok bool) {
	return b.GCEPersistentDisk, b.GCEPersistentDisk != nil
}

// SetAWSElasticBlockStore sets the AWSElasticBlockStore field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetAWSElasticBlockStore(value *AWSElasticBlockStoreVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.AWSElasticBlockStore = value
	return b
}

// RemoveAWSElasticBlockStore removes the AWSElasticBlockStore field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveAWSElasticBlockStore() *VolumeSourceApplyConfiguration {
	b.AWSElasticBlockStore = nil
	return b
}

// GetAWSElasticBlockStore gets the AWSElasticBlockStore field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetAWSElasticBlockStore() (value *AWSElasticBlockStoreVolumeSourceApplyConfiguration, ok bool) {
	return b.AWSElasticBlockStore, b.AWSElasticBlockStore != nil
}

// SetGitRepo sets the GitRepo field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetGitRepo(value *GitRepoVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.GitRepo = value
	return b
}

// RemoveGitRepo removes the GitRepo field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveGitRepo() *VolumeSourceApplyConfiguration {
	b.GitRepo = nil
	return b
}

// GetGitRepo gets the GitRepo field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetGitRepo() (value *GitRepoVolumeSourceApplyConfiguration, ok bool) {
	return b.GitRepo, b.GitRepo != nil
}

// SetSecret sets the Secret field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetSecret(value *SecretVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Secret = value
	return b
}

// RemoveSecret removes the Secret field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveSecret() *VolumeSourceApplyConfiguration {
	b.Secret = nil
	return b
}

// GetSecret gets the Secret field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetSecret() (value *SecretVolumeSourceApplyConfiguration, ok bool) {
	return b.Secret, b.Secret != nil
}

// SetNFS sets the NFS field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetNFS(value *NFSVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.NFS = value
	return b
}

// RemoveNFS removes the NFS field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveNFS() *VolumeSourceApplyConfiguration {
	b.NFS = nil
	return b
}

// GetNFS gets the NFS field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetNFS() (value *NFSVolumeSourceApplyConfiguration, ok bool) {
	return b.NFS, b.NFS != nil
}

// SetISCSI sets the ISCSI field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetISCSI(value *ISCSIVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.ISCSI = value
	return b
}

// RemoveISCSI removes the ISCSI field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveISCSI() *VolumeSourceApplyConfiguration {
	b.ISCSI = nil
	return b
}

// GetISCSI gets the ISCSI field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetISCSI() (value *ISCSIVolumeSourceApplyConfiguration, ok bool) {
	return b.ISCSI, b.ISCSI != nil
}

// SetGlusterfs sets the Glusterfs field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetGlusterfs(value *GlusterfsVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Glusterfs = value
	return b
}

// RemoveGlusterfs removes the Glusterfs field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveGlusterfs() *VolumeSourceApplyConfiguration {
	b.Glusterfs = nil
	return b
}

// GetGlusterfs gets the Glusterfs field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetGlusterfs() (value *GlusterfsVolumeSourceApplyConfiguration, ok bool) {
	return b.Glusterfs, b.Glusterfs != nil
}

// SetPersistentVolumeClaim sets the PersistentVolumeClaim field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetPersistentVolumeClaim(value *PersistentVolumeClaimVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.PersistentVolumeClaim = value
	return b
}

// RemovePersistentVolumeClaim removes the PersistentVolumeClaim field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemovePersistentVolumeClaim() *VolumeSourceApplyConfiguration {
	b.PersistentVolumeClaim = nil
	return b
}

// GetPersistentVolumeClaim gets the PersistentVolumeClaim field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetPersistentVolumeClaim() (value *PersistentVolumeClaimVolumeSourceApplyConfiguration, ok bool) {
	return b.PersistentVolumeClaim, b.PersistentVolumeClaim != nil
}

// SetRBD sets the RBD field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetRBD(value *RBDVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.RBD = value
	return b
}

// RemoveRBD removes the RBD field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveRBD() *VolumeSourceApplyConfiguration {
	b.RBD = nil
	return b
}

// GetRBD gets the RBD field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetRBD() (value *RBDVolumeSourceApplyConfiguration, ok bool) {
	return b.RBD, b.RBD != nil
}

// SetFlexVolume sets the FlexVolume field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetFlexVolume(value *FlexVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.FlexVolume = value
	return b
}

// RemoveFlexVolume removes the FlexVolume field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveFlexVolume() *VolumeSourceApplyConfiguration {
	b.FlexVolume = nil
	return b
}

// GetFlexVolume gets the FlexVolume field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetFlexVolume() (value *FlexVolumeSourceApplyConfiguration, ok bool) {
	return b.FlexVolume, b.FlexVolume != nil
}

// SetCinder sets the Cinder field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetCinder(value *CinderVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Cinder = value
	return b
}

// RemoveCinder removes the Cinder field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveCinder() *VolumeSourceApplyConfiguration {
	b.Cinder = nil
	return b
}

// GetCinder gets the Cinder field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetCinder() (value *CinderVolumeSourceApplyConfiguration, ok bool) {
	return b.Cinder, b.Cinder != nil
}

// SetCephFS sets the CephFS field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetCephFS(value *CephFSVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.CephFS = value
	return b
}

// RemoveCephFS removes the CephFS field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveCephFS() *VolumeSourceApplyConfiguration {
	b.CephFS = nil
	return b
}

// GetCephFS gets the CephFS field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetCephFS() (value *CephFSVolumeSourceApplyConfiguration, ok bool) {
	return b.CephFS, b.CephFS != nil
}

// SetFlocker sets the Flocker field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetFlocker(value *FlockerVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Flocker = value
	return b
}

// RemoveFlocker removes the Flocker field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveFlocker() *VolumeSourceApplyConfiguration {
	b.Flocker = nil
	return b
}

// GetFlocker gets the Flocker field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetFlocker() (value *FlockerVolumeSourceApplyConfiguration, ok bool) {
	return b.Flocker, b.Flocker != nil
}

// SetDownwardAPI sets the DownwardAPI field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetDownwardAPI(value *DownwardAPIVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.DownwardAPI = value
	return b
}

// RemoveDownwardAPI removes the DownwardAPI field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveDownwardAPI() *VolumeSourceApplyConfiguration {
	b.DownwardAPI = nil
	return b
}

// GetDownwardAPI gets the DownwardAPI field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetDownwardAPI() (value *DownwardAPIVolumeSourceApplyConfiguration, ok bool) {
	return b.DownwardAPI, b.DownwardAPI != nil
}

// SetFC sets the FC field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetFC(value *FCVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.FC = value
	return b
}

// RemoveFC removes the FC field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveFC() *VolumeSourceApplyConfiguration {
	b.FC = nil
	return b
}

// GetFC gets the FC field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetFC() (value *FCVolumeSourceApplyConfiguration, ok bool) {
	return b.FC, b.FC != nil
}

// SetAzureFile sets the AzureFile field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetAzureFile(value *AzureFileVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.AzureFile = value
	return b
}

// RemoveAzureFile removes the AzureFile field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveAzureFile() *VolumeSourceApplyConfiguration {
	b.AzureFile = nil
	return b
}

// GetAzureFile gets the AzureFile field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetAzureFile() (value *AzureFileVolumeSourceApplyConfiguration, ok bool) {
	return b.AzureFile, b.AzureFile != nil
}

// SetConfigMap sets the ConfigMap field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetConfigMap(value *ConfigMapVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.ConfigMap = value
	return b
}

// RemoveConfigMap removes the ConfigMap field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveConfigMap() *VolumeSourceApplyConfiguration {
	b.ConfigMap = nil
	return b
}

// GetConfigMap gets the ConfigMap field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetConfigMap() (value *ConfigMapVolumeSourceApplyConfiguration, ok bool) {
	return b.ConfigMap, b.ConfigMap != nil
}

// SetVsphereVolume sets the VsphereVolume field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetVsphereVolume(value *VsphereVirtualDiskVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.VsphereVolume = value
	return b
}

// RemoveVsphereVolume removes the VsphereVolume field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveVsphereVolume() *VolumeSourceApplyConfiguration {
	b.VsphereVolume = nil
	return b
}

// GetVsphereVolume gets the VsphereVolume field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetVsphereVolume() (value *VsphereVirtualDiskVolumeSourceApplyConfiguration, ok bool) {
	return b.VsphereVolume, b.VsphereVolume != nil
}

// SetQuobyte sets the Quobyte field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetQuobyte(value *QuobyteVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Quobyte = value
	return b
}

// RemoveQuobyte removes the Quobyte field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveQuobyte() *VolumeSourceApplyConfiguration {
	b.Quobyte = nil
	return b
}

// GetQuobyte gets the Quobyte field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetQuobyte() (value *QuobyteVolumeSourceApplyConfiguration, ok bool) {
	return b.Quobyte, b.Quobyte != nil
}

// SetAzureDisk sets the AzureDisk field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetAzureDisk(value *AzureDiskVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.AzureDisk = value
	return b
}

// RemoveAzureDisk removes the AzureDisk field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveAzureDisk() *VolumeSourceApplyConfiguration {
	b.AzureDisk = nil
	return b
}

// GetAzureDisk gets the AzureDisk field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetAzureDisk() (value *AzureDiskVolumeSourceApplyConfiguration, ok bool) {
	return b.AzureDisk, b.AzureDisk != nil
}

// SetPhotonPersistentDisk sets the PhotonPersistentDisk field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetPhotonPersistentDisk(value *PhotonPersistentDiskVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.PhotonPersistentDisk = value
	return b
}

// RemovePhotonPersistentDisk removes the PhotonPersistentDisk field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemovePhotonPersistentDisk() *VolumeSourceApplyConfiguration {
	b.PhotonPersistentDisk = nil
	return b
}

// GetPhotonPersistentDisk gets the PhotonPersistentDisk field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetPhotonPersistentDisk() (value *PhotonPersistentDiskVolumeSourceApplyConfiguration, ok bool) {
	return b.PhotonPersistentDisk, b.PhotonPersistentDisk != nil
}

// SetProjected sets the Projected field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetProjected(value *ProjectedVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Projected = value
	return b
}

// RemoveProjected removes the Projected field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveProjected() *VolumeSourceApplyConfiguration {
	b.Projected = nil
	return b
}

// GetProjected gets the Projected field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetProjected() (value *ProjectedVolumeSourceApplyConfiguration, ok bool) {
	return b.Projected, b.Projected != nil
}

// SetPortworxVolume sets the PortworxVolume field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetPortworxVolume(value *PortworxVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.PortworxVolume = value
	return b
}

// RemovePortworxVolume removes the PortworxVolume field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemovePortworxVolume() *VolumeSourceApplyConfiguration {
	b.PortworxVolume = nil
	return b
}

// GetPortworxVolume gets the PortworxVolume field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetPortworxVolume() (value *PortworxVolumeSourceApplyConfiguration, ok bool) {
	return b.PortworxVolume, b.PortworxVolume != nil
}

// SetScaleIO sets the ScaleIO field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetScaleIO(value *ScaleIOVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.ScaleIO = value
	return b
}

// RemoveScaleIO removes the ScaleIO field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveScaleIO() *VolumeSourceApplyConfiguration {
	b.ScaleIO = nil
	return b
}

// GetScaleIO gets the ScaleIO field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetScaleIO() (value *ScaleIOVolumeSourceApplyConfiguration, ok bool) {
	return b.ScaleIO, b.ScaleIO != nil
}

// SetStorageOS sets the StorageOS field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetStorageOS(value *StorageOSVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.StorageOS = value
	return b
}

// RemoveStorageOS removes the StorageOS field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveStorageOS() *VolumeSourceApplyConfiguration {
	b.StorageOS = nil
	return b
}

// GetStorageOS gets the StorageOS field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetStorageOS() (value *StorageOSVolumeSourceApplyConfiguration, ok bool) {
	return b.StorageOS, b.StorageOS != nil
}

// SetCSI sets the CSI field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetCSI(value *CSIVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.CSI = value
	return b
}

// RemoveCSI removes the CSI field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveCSI() *VolumeSourceApplyConfiguration {
	b.CSI = nil
	return b
}

// GetCSI gets the CSI field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetCSI() (value *CSIVolumeSourceApplyConfiguration, ok bool) {
	return b.CSI, b.CSI != nil
}

// SetEphemeral sets the Ephemeral field in the declarative configuration to the given value.
func (b *VolumeSourceApplyConfiguration) SetEphemeral(value *EphemeralVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Ephemeral = value
	return b
}

// RemoveEphemeral removes the Ephemeral field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) RemoveEphemeral() *VolumeSourceApplyConfiguration {
	b.Ephemeral = nil
	return b
}

// GetEphemeral gets the Ephemeral field from the declarative configuration.
func (b *VolumeSourceApplyConfiguration) GetEphemeral() (value *EphemeralVolumeSourceApplyConfiguration, ok bool) {
	return b.Ephemeral, b.Ephemeral != nil
}

// VolumeSourceList represents a listAlias of VolumeSourceApplyConfiguration.
type VolumeSourceList []*VolumeSourceApplyConfiguration

// VolumeSourceList represents a map of VolumeSourceApplyConfiguration.
type VolumeSourceMap map[string]VolumeSourceApplyConfiguration
