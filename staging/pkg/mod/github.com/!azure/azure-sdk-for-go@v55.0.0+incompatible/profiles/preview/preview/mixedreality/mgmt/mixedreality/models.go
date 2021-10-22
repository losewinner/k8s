// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package mixedreality

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/mixedreality/mgmt/2021-03-01-preview/mixedreality"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type NameUnavailableReason = original.NameUnavailableReason

const (
	AlreadyExists NameUnavailableReason = original.AlreadyExists
	Invalid       NameUnavailableReason = original.Invalid
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type SkuTier = original.SkuTier

const (
	Basic    SkuTier = original.Basic
	Free     SkuTier = original.Free
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type AccountKeyRegenerateRequest = original.AccountKeyRegenerateRequest
type AccountKeys = original.AccountKeys
type AccountProperties = original.AccountProperties
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Identity = original.Identity
type LogSpecification = original.LogSpecification
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type ObjectAnchorsAccount = original.ObjectAnchorsAccount
type ObjectAnchorsAccountIdentity = original.ObjectAnchorsAccountIdentity
type ObjectAnchorsAccountPage = original.ObjectAnchorsAccountPage
type ObjectAnchorsAccountPageIterator = original.ObjectAnchorsAccountPageIterator
type ObjectAnchorsAccountPagePage = original.ObjectAnchorsAccountPagePage
type ObjectAnchorsAccountsClient = original.ObjectAnchorsAccountsClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationPage = original.OperationPage
type OperationPageIterator = original.OperationPageIterator
type OperationPagePage = original.OperationPagePage
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type Plan = original.Plan
type ProxyResource = original.ProxyResource
type RemoteRenderingAccount = original.RemoteRenderingAccount
type RemoteRenderingAccountPage = original.RemoteRenderingAccountPage
type RemoteRenderingAccountPageIterator = original.RemoteRenderingAccountPageIterator
type RemoteRenderingAccountPagePage = original.RemoteRenderingAccountPagePage
type RemoteRenderingAccountsClient = original.RemoteRenderingAccountsClient
type Resource = original.Resource
type ResourceModelWithAllowedPropertySet = original.ResourceModelWithAllowedPropertySet
type ResourceModelWithAllowedPropertySetIdentity = original.ResourceModelWithAllowedPropertySetIdentity
type ResourceModelWithAllowedPropertySetPlan = original.ResourceModelWithAllowedPropertySetPlan
type ResourceModelWithAllowedPropertySetSku = original.ResourceModelWithAllowedPropertySetSku
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type SpatialAnchorsAccount = original.SpatialAnchorsAccount
type SpatialAnchorsAccountPage = original.SpatialAnchorsAccountPage
type SpatialAnchorsAccountPageIterator = original.SpatialAnchorsAccountPageIterator
type SpatialAnchorsAccountPagePage = original.SpatialAnchorsAccountPagePage
type SpatialAnchorsAccountsClient = original.SpatialAnchorsAccountsClient
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewObjectAnchorsAccountPageIterator(page ObjectAnchorsAccountPagePage) ObjectAnchorsAccountPageIterator {
	return original.NewObjectAnchorsAccountPageIterator(page)
}
func NewObjectAnchorsAccountPagePage(cur ObjectAnchorsAccountPage, getNextPage func(context.Context, ObjectAnchorsAccountPage) (ObjectAnchorsAccountPage, error)) ObjectAnchorsAccountPagePage {
	return original.NewObjectAnchorsAccountPagePage(cur, getNextPage)
}
func NewObjectAnchorsAccountsClient(subscriptionID string) ObjectAnchorsAccountsClient {
	return original.NewObjectAnchorsAccountsClient(subscriptionID)
}
func NewObjectAnchorsAccountsClientWithBaseURI(baseURI string, subscriptionID string) ObjectAnchorsAccountsClient {
	return original.NewObjectAnchorsAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationPageIterator(page OperationPagePage) OperationPageIterator {
	return original.NewOperationPageIterator(page)
}
func NewOperationPagePage(cur OperationPage, getNextPage func(context.Context, OperationPage) (OperationPage, error)) OperationPagePage {
	return original.NewOperationPagePage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRemoteRenderingAccountPageIterator(page RemoteRenderingAccountPagePage) RemoteRenderingAccountPageIterator {
	return original.NewRemoteRenderingAccountPageIterator(page)
}
func NewRemoteRenderingAccountPagePage(cur RemoteRenderingAccountPage, getNextPage func(context.Context, RemoteRenderingAccountPage) (RemoteRenderingAccountPage, error)) RemoteRenderingAccountPagePage {
	return original.NewRemoteRenderingAccountPagePage(cur, getNextPage)
}
func NewRemoteRenderingAccountsClient(subscriptionID string) RemoteRenderingAccountsClient {
	return original.NewRemoteRenderingAccountsClient(subscriptionID)
}
func NewRemoteRenderingAccountsClientWithBaseURI(baseURI string, subscriptionID string) RemoteRenderingAccountsClient {
	return original.NewRemoteRenderingAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSpatialAnchorsAccountPageIterator(page SpatialAnchorsAccountPagePage) SpatialAnchorsAccountPageIterator {
	return original.NewSpatialAnchorsAccountPageIterator(page)
}
func NewSpatialAnchorsAccountPagePage(cur SpatialAnchorsAccountPage, getNextPage func(context.Context, SpatialAnchorsAccountPage) (SpatialAnchorsAccountPage, error)) SpatialAnchorsAccountPagePage {
	return original.NewSpatialAnchorsAccountPagePage(cur, getNextPage)
}
func NewSpatialAnchorsAccountsClient(subscriptionID string) SpatialAnchorsAccountsClient {
	return original.NewSpatialAnchorsAccountsClient(subscriptionID)
}
func NewSpatialAnchorsAccountsClientWithBaseURI(baseURI string, subscriptionID string) SpatialAnchorsAccountsClient {
	return original.NewSpatialAnchorsAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleNameUnavailableReasonValues() []NameUnavailableReason {
	return original.PossibleNameUnavailableReasonValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
