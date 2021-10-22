package computeapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result compute.OperationListResult, err error)
}

var _ OperationsClientAPI = (*compute.OperationsClient)(nil)

// AvailabilitySetsClientAPI contains the set of methods on the AvailabilitySetsClient type.
type AvailabilitySetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters compute.AvailabilitySet) (result compute.AvailabilitySet, err error)
	Delete(ctx context.Context, resourceGroupName string, availabilitySetName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, availabilitySetName string) (result compute.AvailabilitySet, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.AvailabilitySetListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result compute.AvailabilitySetListResultIterator, err error)
	ListAvailableSizes(ctx context.Context, resourceGroupName string, availabilitySetName string) (result compute.VirtualMachineSizeListResult, err error)
	ListBySubscription(ctx context.Context, expand string) (result compute.AvailabilitySetListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, expand string) (result compute.AvailabilitySetListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters compute.AvailabilitySetUpdate) (result compute.AvailabilitySet, err error)
}

var _ AvailabilitySetsClientAPI = (*compute.AvailabilitySetsClient)(nil)

// ProximityPlacementGroupsClientAPI contains the set of methods on the ProximityPlacementGroupsClient type.
type ProximityPlacementGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters compute.ProximityPlacementGroup) (result compute.ProximityPlacementGroup, err error)
	Delete(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, includeColocationStatus string) (result compute.ProximityPlacementGroup, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.ProximityPlacementGroupListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result compute.ProximityPlacementGroupListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result compute.ProximityPlacementGroupListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result compute.ProximityPlacementGroupListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters compute.ProximityPlacementGroupUpdate) (result compute.ProximityPlacementGroup, err error)
}

var _ ProximityPlacementGroupsClientAPI = (*compute.ProximityPlacementGroupsClient)(nil)

// DedicatedHostGroupsClientAPI contains the set of methods on the DedicatedHostGroupsClient type.
type DedicatedHostGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, parameters compute.DedicatedHostGroup) (result compute.DedicatedHostGroup, err error)
	Delete(ctx context.Context, resourceGroupName string, hostGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, hostGroupName string) (result compute.DedicatedHostGroup, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.DedicatedHostGroupListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result compute.DedicatedHostGroupListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result compute.DedicatedHostGroupListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result compute.DedicatedHostGroupListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, hostGroupName string, parameters compute.DedicatedHostGroupUpdate) (result compute.DedicatedHostGroup, err error)
}

var _ DedicatedHostGroupsClientAPI = (*compute.DedicatedHostGroupsClient)(nil)

// DedicatedHostsClientAPI contains the set of methods on the DedicatedHostsClient type.
type DedicatedHostsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters compute.DedicatedHost) (result compute.DedicatedHostsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (result compute.DedicatedHostsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, expand compute.InstanceViewTypes) (result compute.DedicatedHost, err error)
	ListByHostGroup(ctx context.Context, resourceGroupName string, hostGroupName string) (result compute.DedicatedHostListResultPage, err error)
	ListByHostGroupComplete(ctx context.Context, resourceGroupName string, hostGroupName string) (result compute.DedicatedHostListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters compute.DedicatedHostUpdate) (result compute.DedicatedHostsUpdateFuture, err error)
}

var _ DedicatedHostsClientAPI = (*compute.DedicatedHostsClient)(nil)

// VirtualMachineExtensionImagesClientAPI contains the set of methods on the VirtualMachineExtensionImagesClient type.
type VirtualMachineExtensionImagesClientAPI interface {
	Get(ctx context.Context, location string, publisherName string, typeParameter string, version string) (result compute.VirtualMachineExtensionImage, err error)
	ListTypes(ctx context.Context, location string, publisherName string) (result compute.ListVirtualMachineExtensionImage, err error)
	ListVersions(ctx context.Context, location string, publisherName string, typeParameter string, filter string, top *int32, orderby string) (result compute.ListVirtualMachineExtensionImage, err error)
}

var _ VirtualMachineExtensionImagesClientAPI = (*compute.VirtualMachineExtensionImagesClient)(nil)

// VirtualMachineExtensionsClientAPI contains the set of methods on the VirtualMachineExtensionsClient type.
type VirtualMachineExtensionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, extensionParameters compute.VirtualMachineExtension) (result compute.VirtualMachineExtensionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string) (result compute.VirtualMachineExtensionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, expand string) (result compute.VirtualMachineExtension, err error)
	List(ctx context.Context, resourceGroupName string, VMName string, expand string) (result compute.VirtualMachineExtensionsListResult, err error)
	Update(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, extensionParameters compute.VirtualMachineExtensionUpdate) (result compute.VirtualMachineExtensionsUpdateFuture, err error)
}

var _ VirtualMachineExtensionsClientAPI = (*compute.VirtualMachineExtensionsClient)(nil)

// VirtualMachineImagesClientAPI contains the set of methods on the VirtualMachineImagesClient type.
type VirtualMachineImagesClientAPI interface {
	Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string) (result compute.VirtualMachineImage, err error)
	List(ctx context.Context, location string, publisherName string, offer string, skus string, expand string, top *int32, orderby string) (result compute.ListVirtualMachineImageResource, err error)
	ListOffers(ctx context.Context, location string, publisherName string) (result compute.ListVirtualMachineImageResource, err error)
	ListPublishers(ctx context.Context, location string) (result compute.ListVirtualMachineImageResource, err error)
	ListSkus(ctx context.Context, location string, publisherName string, offer string) (result compute.ListVirtualMachineImageResource, err error)
}

var _ VirtualMachineImagesClientAPI = (*compute.VirtualMachineImagesClient)(nil)

// UsageClientAPI contains the set of methods on the UsageClient type.
type UsageClientAPI interface {
	List(ctx context.Context, location string) (result compute.ListUsagesResultPage, err error)
	ListComplete(ctx context.Context, location string) (result compute.ListUsagesResultIterator, err error)
}

var _ UsageClientAPI = (*compute.UsageClient)(nil)

// VirtualMachinesClientAPI contains the set of methods on the VirtualMachinesClient type.
type VirtualMachinesClientAPI interface {
	Capture(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachineCaptureParameters) (result compute.VirtualMachinesCaptureFuture, err error)
	ConvertToManagedDisks(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesConvertToManagedDisksFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachine) (result compute.VirtualMachinesCreateOrUpdateFuture, err error)
	Deallocate(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesDeleteFuture, err error)
	Generalize(ctx context.Context, resourceGroupName string, VMName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, VMName string, expand compute.InstanceViewTypes) (result compute.VirtualMachine, err error)
	InstanceView(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachineInstanceView, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineListResultIterator, err error)
	ListAll(ctx context.Context, statusOnly string) (result compute.VirtualMachineListResultPage, err error)
	ListAllComplete(ctx context.Context, statusOnly string) (result compute.VirtualMachineListResultIterator, err error)
	ListAvailableSizes(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachineSizeListResult, err error)
	ListByLocation(ctx context.Context, location string) (result compute.VirtualMachineListResultPage, err error)
	ListByLocationComplete(ctx context.Context, location string) (result compute.VirtualMachineListResultIterator, err error)
	PerformMaintenance(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesPerformMaintenanceFuture, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMName string, skipShutdown *bool) (result compute.VirtualMachinesPowerOffFuture, err error)
	Reapply(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesReapplyFuture, err error)
	Redeploy(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesRedeployFuture, err error)
	Reimage(ctx context.Context, resourceGroupName string, VMName string, parameters *compute.VirtualMachineReimageParameters) (result compute.VirtualMachinesReimageFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesRestartFuture, err error)
	RunCommand(ctx context.Context, resourceGroupName string, VMName string, parameters compute.RunCommandInput) (result compute.VirtualMachinesRunCommandFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesStartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachineUpdate) (result compute.VirtualMachinesUpdateFuture, err error)
}

var _ VirtualMachinesClientAPI = (*compute.VirtualMachinesClient)(nil)

// VirtualMachineSizesClientAPI contains the set of methods on the VirtualMachineSizesClient type.
type VirtualMachineSizesClientAPI interface {
	List(ctx context.Context, location string) (result compute.VirtualMachineSizeListResult, err error)
}

var _ VirtualMachineSizesClientAPI = (*compute.VirtualMachineSizesClient)(nil)

// ImagesClientAPI contains the set of methods on the ImagesClient type.
type ImagesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters compute.Image) (result compute.ImagesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, imageName string) (result compute.ImagesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, imageName string, expand string) (result compute.Image, err error)
	List(ctx context.Context) (result compute.ImageListResultPage, err error)
	ListComplete(ctx context.Context) (result compute.ImageListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.ImageListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result compute.ImageListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, imageName string, parameters compute.ImageUpdate) (result compute.ImagesUpdateFuture, err error)
}

var _ ImagesClientAPI = (*compute.ImagesClient)(nil)

// VirtualMachineScaleSetsClientAPI contains the set of methods on the VirtualMachineScaleSetsClient type.
type VirtualMachineScaleSetsClientAPI interface {
	ConvertToSinglePlacementGroup(ctx context.Context, resourceGroupName string, VMScaleSetName string, parameters compute.VMScaleSetConvertToSinglePlacementGroupInput) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMScaleSetName string, parameters compute.VirtualMachineScaleSet) (result compute.VirtualMachineScaleSetsCreateOrUpdateFuture, err error)
	Deallocate(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetsDeleteFuture, err error)
	DeleteInstances(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (result compute.VirtualMachineScaleSetsDeleteInstancesFuture, err error)
	ForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx context.Context, resourceGroupName string, VMScaleSetName string, platformUpdateDomain int32) (result compute.RecoveryWalkResponse, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSet, err error)
	GetInstanceView(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetInstanceView, err error)
	GetOSUpgradeHistory(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetListOSUpgradeHistoryPage, err error)
	GetOSUpgradeHistoryComplete(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetListOSUpgradeHistoryIterator, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineScaleSetListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineScaleSetListResultIterator, err error)
	ListAll(ctx context.Context) (result compute.VirtualMachineScaleSetListWithLinkResultPage, err error)
	ListAllComplete(ctx context.Context) (result compute.VirtualMachineScaleSetListWithLinkResultIterator, err error)
	ListSkus(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetListSkusResultPage, err error)
	ListSkusComplete(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetListSkusResultIterator, err error)
	PerformMaintenance(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsPerformMaintenanceFuture, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs, skipShutdown *bool) (result compute.VirtualMachineScaleSetsPowerOffFuture, err error)
	Redeploy(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsRedeployFuture, err error)
	Reimage(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMScaleSetReimageInput *compute.VirtualMachineScaleSetReimageParameters) (result compute.VirtualMachineScaleSetsReimageFuture, err error)
	ReimageAll(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsReimageAllFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsStartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, parameters compute.VirtualMachineScaleSetUpdate) (result compute.VirtualMachineScaleSetsUpdateFuture, err error)
	UpdateInstances(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (result compute.VirtualMachineScaleSetsUpdateInstancesFuture, err error)
}

var _ VirtualMachineScaleSetsClientAPI = (*compute.VirtualMachineScaleSetsClient)(nil)

// VirtualMachineScaleSetExtensionsClientAPI contains the set of methods on the VirtualMachineScaleSetExtensionsClient type.
type VirtualMachineScaleSetExtensionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMScaleSetName string, vmssExtensionName string, extensionParameters compute.VirtualMachineScaleSetExtension) (result compute.VirtualMachineScaleSetExtensionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string, vmssExtensionName string) (result compute.VirtualMachineScaleSetExtensionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, vmssExtensionName string, expand string) (result compute.VirtualMachineScaleSetExtension, err error)
	List(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetExtensionListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetExtensionListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, vmssExtensionName string, extensionParameters compute.VirtualMachineScaleSetExtensionUpdate) (result compute.VirtualMachineScaleSetExtensionsUpdateFuture, err error)
}

var _ VirtualMachineScaleSetExtensionsClientAPI = (*compute.VirtualMachineScaleSetExtensionsClient)(nil)

// VirtualMachineScaleSetRollingUpgradesClientAPI contains the set of methods on the VirtualMachineScaleSetRollingUpgradesClient type.
type VirtualMachineScaleSetRollingUpgradesClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetRollingUpgradesCancelFuture, err error)
	GetLatest(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.RollingUpgradeStatusInfo, err error)
	StartExtensionUpgrade(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetRollingUpgradesStartExtensionUpgradeFuture, err error)
	StartOSUpgrade(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetRollingUpgradesStartOSUpgradeFuture, err error)
}

var _ VirtualMachineScaleSetRollingUpgradesClientAPI = (*compute.VirtualMachineScaleSetRollingUpgradesClient)(nil)

// VirtualMachineScaleSetVMExtensionsClientAPI contains the set of methods on the VirtualMachineScaleSetVMExtensionsClient type.
type VirtualMachineScaleSetVMExtensionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, extensionParameters compute.VirtualMachineExtension) (result compute.VirtualMachineScaleSetVMExtensionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string) (result compute.VirtualMachineScaleSetVMExtensionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, expand string) (result compute.VirtualMachineExtension, err error)
	List(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, expand string) (result compute.VirtualMachineExtensionsListResult, err error)
	Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, extensionParameters compute.VirtualMachineExtensionUpdate) (result compute.VirtualMachineScaleSetVMExtensionsUpdateFuture, err error)
}

var _ VirtualMachineScaleSetVMExtensionsClientAPI = (*compute.VirtualMachineScaleSetVMExtensionsClient)(nil)

// VirtualMachineScaleSetVMsClientAPI contains the set of methods on the VirtualMachineScaleSetVMsClient type.
type VirtualMachineScaleSetVMsClientAPI interface {
	Deallocate(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, expand compute.InstanceViewTypes) (result compute.VirtualMachineScaleSetVM, err error)
	GetInstanceView(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMInstanceView, err error)
	List(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (result compute.VirtualMachineScaleSetVMListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (result compute.VirtualMachineScaleSetVMListResultIterator, err error)
	PerformMaintenance(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsPerformMaintenanceFuture, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, skipShutdown *bool) (result compute.VirtualMachineScaleSetVMsPowerOffFuture, err error)
	Redeploy(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsRedeployFuture, err error)
	Reimage(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMScaleSetVMReimageInput *compute.VirtualMachineScaleSetVMReimageParameters) (result compute.VirtualMachineScaleSetVMsReimageFuture, err error)
	ReimageAll(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsReimageAllFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsRestartFuture, err error)
	RunCommand(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, parameters compute.RunCommandInput) (result compute.VirtualMachineScaleSetVMsRunCommandFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsStartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, parameters compute.VirtualMachineScaleSetVM) (result compute.VirtualMachineScaleSetVMsUpdateFuture, err error)
}

var _ VirtualMachineScaleSetVMsClientAPI = (*compute.VirtualMachineScaleSetVMsClient)(nil)

// LogAnalyticsClientAPI contains the set of methods on the LogAnalyticsClient type.
type LogAnalyticsClientAPI interface {
	ExportRequestRateByInterval(ctx context.Context, parameters compute.RequestRateByIntervalInput, location string) (result compute.LogAnalyticsExportRequestRateByIntervalFuture, err error)
	ExportThrottledRequests(ctx context.Context, parameters compute.ThrottledRequestsInput, location string) (result compute.LogAnalyticsExportThrottledRequestsFuture, err error)
}

var _ LogAnalyticsClientAPI = (*compute.LogAnalyticsClient)(nil)

// VirtualMachineRunCommandsClientAPI contains the set of methods on the VirtualMachineRunCommandsClient type.
type VirtualMachineRunCommandsClientAPI interface {
	Get(ctx context.Context, location string, commandID string) (result compute.RunCommandDocument, err error)
	List(ctx context.Context, location string) (result compute.RunCommandListResultPage, err error)
	ListComplete(ctx context.Context, location string) (result compute.RunCommandListResultIterator, err error)
}

var _ VirtualMachineRunCommandsClientAPI = (*compute.VirtualMachineRunCommandsClient)(nil)

// ResourceSkusClientAPI contains the set of methods on the ResourceSkusClient type.
type ResourceSkusClientAPI interface {
	List(ctx context.Context, filter string) (result compute.ResourceSkusResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result compute.ResourceSkusResultIterator, err error)
}

var _ ResourceSkusClientAPI = (*compute.ResourceSkusClient)(nil)

// DisksClientAPI contains the set of methods on the DisksClient type.
type DisksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk compute.Disk) (result compute.DisksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, diskName string) (result compute.DisksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, diskName string) (result compute.Disk, err error)
	GrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData compute.GrantAccessData) (result compute.DisksGrantAccessFuture, err error)
	List(ctx context.Context) (result compute.DiskListPage, err error)
	ListComplete(ctx context.Context) (result compute.DiskListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.DiskListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result compute.DiskListIterator, err error)
	RevokeAccess(ctx context.Context, resourceGroupName string, diskName string) (result compute.DisksRevokeAccessFuture, err error)
	Update(ctx context.Context, resourceGroupName string, diskName string, disk compute.DiskUpdate) (result compute.DisksUpdateFuture, err error)
}

var _ DisksClientAPI = (*compute.DisksClient)(nil)

// SnapshotsClientAPI contains the set of methods on the SnapshotsClient type.
type SnapshotsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot compute.Snapshot) (result compute.SnapshotsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, snapshotName string) (result compute.SnapshotsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, snapshotName string) (result compute.Snapshot, err error)
	GrantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData compute.GrantAccessData) (result compute.SnapshotsGrantAccessFuture, err error)
	List(ctx context.Context) (result compute.SnapshotListPage, err error)
	ListComplete(ctx context.Context) (result compute.SnapshotListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.SnapshotListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result compute.SnapshotListIterator, err error)
	RevokeAccess(ctx context.Context, resourceGroupName string, snapshotName string) (result compute.SnapshotsRevokeAccessFuture, err error)
	Update(ctx context.Context, resourceGroupName string, snapshotName string, snapshot compute.SnapshotUpdate) (result compute.SnapshotsUpdateFuture, err error)
}

var _ SnapshotsClientAPI = (*compute.SnapshotsClient)(nil)

// DiskEncryptionSetsClientAPI contains the set of methods on the DiskEncryptionSetsClient type.
type DiskEncryptionSetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet compute.DiskEncryptionSet) (result compute.DiskEncryptionSetsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (result compute.DiskEncryptionSetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (result compute.DiskEncryptionSet, err error)
	List(ctx context.Context) (result compute.DiskEncryptionSetListPage, err error)
	ListComplete(ctx context.Context) (result compute.DiskEncryptionSetListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.DiskEncryptionSetListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result compute.DiskEncryptionSetListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet compute.DiskEncryptionSetUpdate) (result compute.DiskEncryptionSetsUpdateFuture, err error)
}

var _ DiskEncryptionSetsClientAPI = (*compute.DiskEncryptionSetsClient)(nil)

// GalleriesClientAPI contains the set of methods on the GalleriesClient type.
type GalleriesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery compute.Gallery) (result compute.GalleriesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, galleryName string) (result compute.GalleriesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, galleryName string) (result compute.Gallery, err error)
	List(ctx context.Context) (result compute.GalleryListPage, err error)
	ListComplete(ctx context.Context) (result compute.GalleryListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.GalleryListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result compute.GalleryListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, galleryName string, gallery compute.GalleryUpdate) (result compute.GalleriesUpdateFuture, err error)
}

var _ GalleriesClientAPI = (*compute.GalleriesClient)(nil)

// GalleryImagesClientAPI contains the set of methods on the GalleryImagesClient type.
type GalleryImagesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage compute.GalleryImage) (result compute.GalleryImagesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string) (result compute.GalleryImagesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string) (result compute.GalleryImage, err error)
	ListByGallery(ctx context.Context, resourceGroupName string, galleryName string) (result compute.GalleryImageListPage, err error)
	ListByGalleryComplete(ctx context.Context, resourceGroupName string, galleryName string) (result compute.GalleryImageListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage compute.GalleryImageUpdate) (result compute.GalleryImagesUpdateFuture, err error)
}

var _ GalleryImagesClientAPI = (*compute.GalleryImagesClient)(nil)

// GalleryImageVersionsClientAPI contains the set of methods on the GalleryImageVersionsClient type.
type GalleryImageVersionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion compute.GalleryImageVersion) (result compute.GalleryImageVersionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (result compute.GalleryImageVersionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, expand compute.ReplicationStatusTypes) (result compute.GalleryImageVersion, err error)
	ListByGalleryImage(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string) (result compute.GalleryImageVersionListPage, err error)
	ListByGalleryImageComplete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string) (result compute.GalleryImageVersionListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion compute.GalleryImageVersionUpdate) (result compute.GalleryImageVersionsUpdateFuture, err error)
}

var _ GalleryImageVersionsClientAPI = (*compute.GalleryImageVersionsClient)(nil)

// GalleryApplicationsClientAPI contains the set of methods on the GalleryApplicationsClient type.
type GalleryApplicationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication compute.GalleryApplication) (result compute.GalleryApplicationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string) (result compute.GalleryApplicationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string) (result compute.GalleryApplication, err error)
	ListByGallery(ctx context.Context, resourceGroupName string, galleryName string) (result compute.GalleryApplicationListPage, err error)
	ListByGalleryComplete(ctx context.Context, resourceGroupName string, galleryName string) (result compute.GalleryApplicationListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication compute.GalleryApplicationUpdate) (result compute.GalleryApplicationsUpdateFuture, err error)
}

var _ GalleryApplicationsClientAPI = (*compute.GalleryApplicationsClient)(nil)

// GalleryApplicationVersionsClientAPI contains the set of methods on the GalleryApplicationVersionsClient type.
type GalleryApplicationVersionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion compute.GalleryApplicationVersion) (result compute.GalleryApplicationVersionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string) (result compute.GalleryApplicationVersionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, expand compute.ReplicationStatusTypes) (result compute.GalleryApplicationVersion, err error)
	ListByGalleryApplication(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string) (result compute.GalleryApplicationVersionListPage, err error)
	ListByGalleryApplicationComplete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string) (result compute.GalleryApplicationVersionListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion compute.GalleryApplicationVersionUpdate) (result compute.GalleryApplicationVersionsUpdateFuture, err error)
}

var _ GalleryApplicationVersionsClientAPI = (*compute.GalleryApplicationVersionsClient)(nil)

// ContainerServicesClientAPI contains the set of methods on the ContainerServicesClient type.
type ContainerServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters compute.ContainerService) (result compute.ContainerServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (result compute.ContainerServicesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, containerServiceName string) (result compute.ContainerService, err error)
	List(ctx context.Context) (result compute.ContainerServiceListResultPage, err error)
	ListComplete(ctx context.Context) (result compute.ContainerServiceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.ContainerServiceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result compute.ContainerServiceListResultIterator, err error)
}

var _ ContainerServicesClientAPI = (*compute.ContainerServicesClient)(nil)
