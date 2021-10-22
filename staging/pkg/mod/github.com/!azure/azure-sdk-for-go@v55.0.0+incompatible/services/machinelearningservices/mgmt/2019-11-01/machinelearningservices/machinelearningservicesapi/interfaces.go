package machinelearningservicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/machinelearningservices/mgmt/2019-11-01/machinelearningservices"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	ListSkus(ctx context.Context) (result machinelearningservices.SkuListResultPage, err error)
	ListSkusComplete(ctx context.Context) (result machinelearningservices.SkuListResultIterator, err error)
}

var _ BaseClientAPI = (*machinelearningservices.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result machinelearningservices.OperationListResult, err error)
}

var _ OperationsClientAPI = (*machinelearningservices.OperationsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters machinelearningservices.Workspace) (result machinelearningservices.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skiptoken string) (result machinelearningservices.WorkspaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skiptoken string) (result machinelearningservices.WorkspaceListResultIterator, err error)
	ListBySubscription(ctx context.Context, skiptoken string) (result machinelearningservices.WorkspaceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, skiptoken string) (result machinelearningservices.WorkspaceListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListWorkspaceKeysResult, err error)
	ResyncKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters machinelearningservices.WorkspaceUpdateParameters) (result machinelearningservices.Workspace, err error)
}

var _ WorkspacesClientAPI = (*machinelearningservices.WorkspacesClient)(nil)

// WorkspaceFeaturesClientAPI contains the set of methods on the WorkspaceFeaturesClient type.
type WorkspaceFeaturesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListAmlUserFeatureResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListAmlUserFeatureResultIterator, err error)
}

var _ WorkspaceFeaturesClientAPI = (*machinelearningservices.WorkspaceFeaturesClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result machinelearningservices.ListUsagesResultPage, err error)
	ListComplete(ctx context.Context, location string) (result machinelearningservices.ListUsagesResultIterator, err error)
}

var _ UsagesClientAPI = (*machinelearningservices.UsagesClient)(nil)

// VirtualMachineSizesClientAPI contains the set of methods on the VirtualMachineSizesClient type.
type VirtualMachineSizesClientAPI interface {
	List(ctx context.Context, location string) (result machinelearningservices.VirtualMachineSizeListResult, err error)
}

var _ VirtualMachineSizesClientAPI = (*machinelearningservices.VirtualMachineSizesClient)(nil)

// QuotasClientAPI contains the set of methods on the QuotasClient type.
type QuotasClientAPI interface {
	List(ctx context.Context, location string) (result machinelearningservices.ListWorkspaceQuotasPage, err error)
	ListComplete(ctx context.Context, location string) (result machinelearningservices.ListWorkspaceQuotasIterator, err error)
	Update(ctx context.Context, location string, parameters machinelearningservices.QuotaUpdateParameters) (result machinelearningservices.UpdateWorkspaceQuotasResult, err error)
}

var _ QuotasClientAPI = (*machinelearningservices.QuotasClient)(nil)

// MachineLearningComputeClientAPI contains the set of methods on the MachineLearningComputeClient type.
type MachineLearningComputeClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters machinelearningservices.ComputeResource) (result machinelearningservices.MachineLearningComputeCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction machinelearningservices.UnderlyingResourceAction) (result machinelearningservices.MachineLearningComputeDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeResource, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skiptoken string) (result machinelearningservices.PaginatedComputeResourcesListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skiptoken string) (result machinelearningservices.PaginatedComputeResourcesListIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeSecretsModel, err error)
	ListNodes(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.AmlComputeNodesInformation, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters machinelearningservices.ClusterUpdateParameters) (result machinelearningservices.MachineLearningComputeUpdateFuture, err error)
}

var _ MachineLearningComputeClientAPI = (*machinelearningservices.MachineLearningComputeClient)(nil)
