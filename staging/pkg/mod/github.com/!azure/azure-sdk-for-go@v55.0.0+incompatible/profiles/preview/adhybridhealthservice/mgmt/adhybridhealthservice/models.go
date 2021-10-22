// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package adhybridhealthservice

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/adhybridhealthservice/mgmt/2014-01-01/adhybridhealthservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AlgorithmStepType = original.AlgorithmStepType

const (
	ConnectorFilter         AlgorithmStepType = original.ConnectorFilter
	Deprovisioning          AlgorithmStepType = original.Deprovisioning
	ExportFlow              AlgorithmStepType = original.ExportFlow
	ImportFlow              AlgorithmStepType = original.ImportFlow
	Join                    AlgorithmStepType = original.Join
	MvDeletion              AlgorithmStepType = original.MvDeletion
	MvObjectTypeChange      AlgorithmStepType = original.MvObjectTypeChange
	Projection              AlgorithmStepType = original.Projection
	Provisioning            AlgorithmStepType = original.Provisioning
	Recall                  AlgorithmStepType = original.Recall
	Staging                 AlgorithmStepType = original.Staging
	Undefined               AlgorithmStepType = original.Undefined
	ValidateConnectorFilter AlgorithmStepType = original.ValidateConnectorFilter
)

type AttributeDeltaOperationType = original.AttributeDeltaOperationType

const (
	AttributeDeltaOperationTypeAdd       AttributeDeltaOperationType = original.AttributeDeltaOperationTypeAdd
	AttributeDeltaOperationTypeDelete    AttributeDeltaOperationType = original.AttributeDeltaOperationTypeDelete
	AttributeDeltaOperationTypeReplace   AttributeDeltaOperationType = original.AttributeDeltaOperationTypeReplace
	AttributeDeltaOperationTypeUndefined AttributeDeltaOperationType = original.AttributeDeltaOperationTypeUndefined
	AttributeDeltaOperationTypeUpdate    AttributeDeltaOperationType = original.AttributeDeltaOperationTypeUpdate
)

type AttributeMappingType = original.AttributeMappingType

const (
	Constant AttributeMappingType = original.Constant
	Direct   AttributeMappingType = original.Direct
	DnPart   AttributeMappingType = original.DnPart
	Script   AttributeMappingType = original.Script
)

type DeltaOperationType = original.DeltaOperationType

const (
	DeltaOperationTypeAdd       DeltaOperationType = original.DeltaOperationTypeAdd
	DeltaOperationTypeDelete    DeltaOperationType = original.DeltaOperationTypeDelete
	DeltaOperationTypeDeleteAdd DeltaOperationType = original.DeltaOperationTypeDeleteAdd
	DeltaOperationTypeNone      DeltaOperationType = original.DeltaOperationTypeNone
	DeltaOperationTypeObsolete  DeltaOperationType = original.DeltaOperationTypeObsolete
	DeltaOperationTypeReplace   DeltaOperationType = original.DeltaOperationTypeReplace
	DeltaOperationTypeUndefined DeltaOperationType = original.DeltaOperationTypeUndefined
	DeltaOperationTypeUpdate    DeltaOperationType = original.DeltaOperationTypeUpdate
)

type HealthStatus = original.HealthStatus

const (
	Error        HealthStatus = original.Error
	Healthy      HealthStatus = original.Healthy
	Missing      HealthStatus = original.Missing
	NotMonitored HealthStatus = original.NotMonitored
	Warning      HealthStatus = original.Warning
)

type Level = original.Level

const (
	LevelError      Level = original.LevelError
	LevelPreWarning Level = original.LevelPreWarning
	LevelWarning    Level = original.LevelWarning
)

type MonitoringLevel = original.MonitoringLevel

const (
	Full    MonitoringLevel = original.Full
	Off     MonitoringLevel = original.Off
	Partial MonitoringLevel = original.Partial
)

type PasswordOperationTypes = original.PasswordOperationTypes

const (
	PasswordOperationTypesChange    PasswordOperationTypes = original.PasswordOperationTypesChange
	PasswordOperationTypesSet       PasswordOperationTypes = original.PasswordOperationTypesSet
	PasswordOperationTypesUndefined PasswordOperationTypes = original.PasswordOperationTypesUndefined
)

type State = original.State

const (
	Active                   State = original.Active
	ResolvedByPositiveResult State = original.ResolvedByPositiveResult
	ResolvedByStateChange    State = original.ResolvedByStateChange
	ResolvedByTimer          State = original.ResolvedByTimer
	ResolvedManually         State = original.ResolvedManually
)

type ValueDeltaOperationType = original.ValueDeltaOperationType

const (
	ValueDeltaOperationTypeAdd       ValueDeltaOperationType = original.ValueDeltaOperationTypeAdd
	ValueDeltaOperationTypeDelete    ValueDeltaOperationType = original.ValueDeltaOperationTypeDelete
	ValueDeltaOperationTypeUndefined ValueDeltaOperationType = original.ValueDeltaOperationTypeUndefined
	ValueDeltaOperationTypeUpdate    ValueDeltaOperationType = original.ValueDeltaOperationTypeUpdate
)

type ValueType = original.ValueType

const (
	ValueTypeBinary    ValueType = original.ValueTypeBinary
	ValueTypeBoolean   ValueType = original.ValueTypeBoolean
	ValueTypeDn        ValueType = original.ValueTypeDn
	ValueTypeInteger   ValueType = original.ValueTypeInteger
	ValueTypeString    ValueType = original.ValueTypeString
	ValueTypeUndefined ValueType = original.ValueTypeUndefined
)

type AdDomainServiceMembersClient = original.AdDomainServiceMembersClient
type AdditionalInformation = original.AdditionalInformation
type AddsConfiguration = original.AddsConfiguration
type AddsConfigurationIterator = original.AddsConfigurationIterator
type AddsConfigurationPage = original.AddsConfigurationPage
type AddsServiceClient = original.AddsServiceClient
type AddsServiceMember = original.AddsServiceMember
type AddsServiceMembers = original.AddsServiceMembers
type AddsServiceMembersClient = original.AddsServiceMembersClient
type AddsServiceMembersIterator = original.AddsServiceMembersIterator
type AddsServiceMembersPage = original.AddsServiceMembersPage
type AddsServicesClient = original.AddsServicesClient
type AddsServicesReplicationStatusClient = original.AddsServicesReplicationStatusClient
type AddsServicesServiceMembersClient = original.AddsServicesServiceMembersClient
type AddsServicesUserPreferenceClient = original.AddsServicesUserPreferenceClient
type Agent = original.Agent
type Alert = original.Alert
type AlertFeedback = original.AlertFeedback
type AlertFeedbacks = original.AlertFeedbacks
type Alerts = original.Alerts
type AlertsClient = original.AlertsClient
type AlertsIterator = original.AlertsIterator
type AlertsPage = original.AlertsPage
type AssociatedObject = original.AssociatedObject
type AttributeDelta = original.AttributeDelta
type AttributeMapping = original.AttributeMapping
type AttributeMppingSource = original.AttributeMppingSource
type BaseClient = original.BaseClient
type ChangeNotReimported = original.ChangeNotReimported
type ChangeNotReimportedDelta = original.ChangeNotReimportedDelta
type ChangeNotReimportedEntry = original.ChangeNotReimportedEntry
type ConfigurationClient = original.ConfigurationClient
type Connector = original.Connector
type ConnectorConnectionError = original.ConnectorConnectionError
type ConnectorConnectionErrors = original.ConnectorConnectionErrors
type ConnectorMetadata = original.ConnectorMetadata
type ConnectorMetadataDetails = original.ConnectorMetadataDetails
type ConnectorObjectError = original.ConnectorObjectError
type ConnectorObjectErrors = original.ConnectorObjectErrors
type Connectors = original.Connectors
type Credential = original.Credential
type Credentials = original.Credentials
type DataFreshnessDetails = original.DataFreshnessDetails
type Dimension = original.Dimension
type Dimensions = original.Dimensions
type DimensionsClient = original.DimensionsClient
type DimensionsIterator = original.DimensionsIterator
type DimensionsPage = original.DimensionsPage
type Display = original.Display
type ErrorCount = original.ErrorCount
type ErrorCounts = original.ErrorCounts
type ErrorDetail = original.ErrorDetail
type ErrorReportUsersEntries = original.ErrorReportUsersEntries
type ErrorReportUsersEntry = original.ErrorReportUsersEntry
type ExportError = original.ExportError
type ExportErrors = original.ExportErrors
type ExportStatus = original.ExportStatus
type ExportStatuses = original.ExportStatuses
type ExportStatusesIterator = original.ExportStatusesIterator
type ExportStatusesPage = original.ExportStatusesPage
type ExtensionErrorInfo = original.ExtensionErrorInfo
type ForestSummary = original.ForestSummary
type GlobalConfiguration = original.GlobalConfiguration
type GlobalConfigurations = original.GlobalConfigurations
type HelpLink = original.HelpLink
type Hotfix = original.Hotfix
type Hotfixes = original.Hotfixes
type IPAddressAggregate = original.IPAddressAggregate
type IPAddressAggregateSetting = original.IPAddressAggregateSetting
type IPAddressAggregates = original.IPAddressAggregates
type IPAddressAggregatesIterator = original.IPAddressAggregatesIterator
type IPAddressAggregatesPage = original.IPAddressAggregatesPage
type ImportError = original.ImportError
type ImportErrors = original.ImportErrors
type InboundReplicationNeighbor = original.InboundReplicationNeighbor
type InboundReplicationNeighbors = original.InboundReplicationNeighbors
type Item = original.Item
type Items = original.Items
type ListClient = original.ListClient
type MergedExportError = original.MergedExportError
type MergedExportErrors = original.MergedExportErrors
type MetricGroup = original.MetricGroup
type MetricMetadata = original.MetricMetadata
type MetricMetadataList = original.MetricMetadataList
type MetricMetadataListIterator = original.MetricMetadataListIterator
type MetricMetadataListPage = original.MetricMetadataListPage
type MetricSet = original.MetricSet
type MetricSets = original.MetricSets
type Metrics = original.Metrics
type MetricsIterator = original.MetricsIterator
type MetricsPage = original.MetricsPage
type ModuleConfiguration = original.ModuleConfiguration
type ModuleConfigurations = original.ModuleConfigurations
type ObjectWithSyncError = original.ObjectWithSyncError
type Operation = original.Operation
type OperationListResponse = original.OperationListResponse
type OperationListResponseIterator = original.OperationListResponseIterator
type OperationListResponsePage = original.OperationListResponsePage
type OperationsClient = original.OperationsClient
type Partition = original.Partition
type PartitionScope = original.PartitionScope
type PasswordHashSyncConfiguration = original.PasswordHashSyncConfiguration
type PasswordManagementSettings = original.PasswordManagementSettings
type ReplicationDetailsList = original.ReplicationDetailsList
type ReplicationStatus = original.ReplicationStatus
type ReplicationSummary = original.ReplicationSummary
type ReplicationSummaryList = original.ReplicationSummaryList
type ReportsClient = original.ReportsClient
type Result = original.Result
type RiskyIPBlobURI = original.RiskyIPBlobURI
type RiskyIPBlobUris = original.RiskyIPBlobUris
type RuleErrorInfo = original.RuleErrorInfo
type RunProfile = original.RunProfile
type RunProfiles = original.RunProfiles
type RunStep = original.RunStep
type ServiceClient = original.ServiceClient
type ServiceConfiguration = original.ServiceConfiguration
type ServiceMember = original.ServiceMember
type ServiceMembers = original.ServiceMembers
type ServiceMembersClient = original.ServiceMembersClient
type ServiceMembersIterator = original.ServiceMembersIterator
type ServiceMembersPage = original.ServiceMembersPage
type ServiceProperties = original.ServiceProperties
type Services = original.Services
type ServicesClient = original.ServicesClient
type ServicesIterator = original.ServicesIterator
type ServicesPage = original.ServicesPage
type TabularExportError = original.TabularExportError
type Tenant = original.Tenant
type TenantOnboardingDetails = original.TenantOnboardingDetails
type UpdateClient = original.UpdateClient
type UserPreference = original.UserPreference
type ValueDelta = original.ValueDelta

func New() BaseClient {
	return original.New()
}
func NewAdDomainServiceMembersClient() AdDomainServiceMembersClient {
	return original.NewAdDomainServiceMembersClient()
}
func NewAdDomainServiceMembersClientWithBaseURI(baseURI string) AdDomainServiceMembersClient {
	return original.NewAdDomainServiceMembersClientWithBaseURI(baseURI)
}
func NewAddsConfigurationIterator(page AddsConfigurationPage) AddsConfigurationIterator {
	return original.NewAddsConfigurationIterator(page)
}
func NewAddsConfigurationPage(cur AddsConfiguration, getNextPage func(context.Context, AddsConfiguration) (AddsConfiguration, error)) AddsConfigurationPage {
	return original.NewAddsConfigurationPage(cur, getNextPage)
}
func NewAddsServiceClient() AddsServiceClient {
	return original.NewAddsServiceClient()
}
func NewAddsServiceClientWithBaseURI(baseURI string) AddsServiceClient {
	return original.NewAddsServiceClientWithBaseURI(baseURI)
}
func NewAddsServiceMembersClient() AddsServiceMembersClient {
	return original.NewAddsServiceMembersClient()
}
func NewAddsServiceMembersClientWithBaseURI(baseURI string) AddsServiceMembersClient {
	return original.NewAddsServiceMembersClientWithBaseURI(baseURI)
}
func NewAddsServiceMembersIterator(page AddsServiceMembersPage) AddsServiceMembersIterator {
	return original.NewAddsServiceMembersIterator(page)
}
func NewAddsServiceMembersPage(cur AddsServiceMembers, getNextPage func(context.Context, AddsServiceMembers) (AddsServiceMembers, error)) AddsServiceMembersPage {
	return original.NewAddsServiceMembersPage(cur, getNextPage)
}
func NewAddsServicesClient() AddsServicesClient {
	return original.NewAddsServicesClient()
}
func NewAddsServicesClientWithBaseURI(baseURI string) AddsServicesClient {
	return original.NewAddsServicesClientWithBaseURI(baseURI)
}
func NewAddsServicesReplicationStatusClient() AddsServicesReplicationStatusClient {
	return original.NewAddsServicesReplicationStatusClient()
}
func NewAddsServicesReplicationStatusClientWithBaseURI(baseURI string) AddsServicesReplicationStatusClient {
	return original.NewAddsServicesReplicationStatusClientWithBaseURI(baseURI)
}
func NewAddsServicesServiceMembersClient() AddsServicesServiceMembersClient {
	return original.NewAddsServicesServiceMembersClient()
}
func NewAddsServicesServiceMembersClientWithBaseURI(baseURI string) AddsServicesServiceMembersClient {
	return original.NewAddsServicesServiceMembersClientWithBaseURI(baseURI)
}
func NewAddsServicesUserPreferenceClient() AddsServicesUserPreferenceClient {
	return original.NewAddsServicesUserPreferenceClient()
}
func NewAddsServicesUserPreferenceClientWithBaseURI(baseURI string) AddsServicesUserPreferenceClient {
	return original.NewAddsServicesUserPreferenceClientWithBaseURI(baseURI)
}
func NewAlertsClient() AlertsClient {
	return original.NewAlertsClient()
}
func NewAlertsClientWithBaseURI(baseURI string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI)
}
func NewAlertsIterator(page AlertsPage) AlertsIterator {
	return original.NewAlertsIterator(page)
}
func NewAlertsPage(cur Alerts, getNextPage func(context.Context, Alerts) (Alerts, error)) AlertsPage {
	return original.NewAlertsPage(cur, getNextPage)
}
func NewConfigurationClient() ConfigurationClient {
	return original.NewConfigurationClient()
}
func NewConfigurationClientWithBaseURI(baseURI string) ConfigurationClient {
	return original.NewConfigurationClientWithBaseURI(baseURI)
}
func NewDimensionsClient() DimensionsClient {
	return original.NewDimensionsClient()
}
func NewDimensionsClientWithBaseURI(baseURI string) DimensionsClient {
	return original.NewDimensionsClientWithBaseURI(baseURI)
}
func NewDimensionsIterator(page DimensionsPage) DimensionsIterator {
	return original.NewDimensionsIterator(page)
}
func NewDimensionsPage(cur Dimensions, getNextPage func(context.Context, Dimensions) (Dimensions, error)) DimensionsPage {
	return original.NewDimensionsPage(cur, getNextPage)
}
func NewExportStatusesIterator(page ExportStatusesPage) ExportStatusesIterator {
	return original.NewExportStatusesIterator(page)
}
func NewExportStatusesPage(cur ExportStatuses, getNextPage func(context.Context, ExportStatuses) (ExportStatuses, error)) ExportStatusesPage {
	return original.NewExportStatusesPage(cur, getNextPage)
}
func NewIPAddressAggregatesIterator(page IPAddressAggregatesPage) IPAddressAggregatesIterator {
	return original.NewIPAddressAggregatesIterator(page)
}
func NewIPAddressAggregatesPage(cur IPAddressAggregates, getNextPage func(context.Context, IPAddressAggregates) (IPAddressAggregates, error)) IPAddressAggregatesPage {
	return original.NewIPAddressAggregatesPage(cur, getNextPage)
}
func NewListClient() ListClient {
	return original.NewListClient()
}
func NewListClientWithBaseURI(baseURI string) ListClient {
	return original.NewListClientWithBaseURI(baseURI)
}
func NewMetricMetadataListIterator(page MetricMetadataListPage) MetricMetadataListIterator {
	return original.NewMetricMetadataListIterator(page)
}
func NewMetricMetadataListPage(cur MetricMetadataList, getNextPage func(context.Context, MetricMetadataList) (MetricMetadataList, error)) MetricMetadataListPage {
	return original.NewMetricMetadataListPage(cur, getNextPage)
}
func NewMetricsIterator(page MetricsPage) MetricsIterator {
	return original.NewMetricsIterator(page)
}
func NewMetricsPage(cur Metrics, getNextPage func(context.Context, Metrics) (Metrics, error)) MetricsPage {
	return original.NewMetricsPage(cur, getNextPage)
}
func NewOperationListResponseIterator(page OperationListResponsePage) OperationListResponseIterator {
	return original.NewOperationListResponseIterator(page)
}
func NewOperationListResponsePage(cur OperationListResponse, getNextPage func(context.Context, OperationListResponse) (OperationListResponse, error)) OperationListResponsePage {
	return original.NewOperationListResponsePage(cur, getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewReportsClient() ReportsClient {
	return original.NewReportsClient()
}
func NewReportsClientWithBaseURI(baseURI string) ReportsClient {
	return original.NewReportsClientWithBaseURI(baseURI)
}
func NewServiceClient() ServiceClient {
	return original.NewServiceClient()
}
func NewServiceClientWithBaseURI(baseURI string) ServiceClient {
	return original.NewServiceClientWithBaseURI(baseURI)
}
func NewServiceMembersClient() ServiceMembersClient {
	return original.NewServiceMembersClient()
}
func NewServiceMembersClientWithBaseURI(baseURI string) ServiceMembersClient {
	return original.NewServiceMembersClientWithBaseURI(baseURI)
}
func NewServiceMembersIterator(page ServiceMembersPage) ServiceMembersIterator {
	return original.NewServiceMembersIterator(page)
}
func NewServiceMembersPage(cur ServiceMembers, getNextPage func(context.Context, ServiceMembers) (ServiceMembers, error)) ServiceMembersPage {
	return original.NewServiceMembersPage(cur, getNextPage)
}
func NewServicesClient() ServicesClient {
	return original.NewServicesClient()
}
func NewServicesClientWithBaseURI(baseURI string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI)
}
func NewServicesIterator(page ServicesPage) ServicesIterator {
	return original.NewServicesIterator(page)
}
func NewServicesPage(cur Services, getNextPage func(context.Context, Services) (Services, error)) ServicesPage {
	return original.NewServicesPage(cur, getNextPage)
}
func NewUpdateClient() UpdateClient {
	return original.NewUpdateClient()
}
func NewUpdateClientWithBaseURI(baseURI string) UpdateClient {
	return original.NewUpdateClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleAlgorithmStepTypeValues() []AlgorithmStepType {
	return original.PossibleAlgorithmStepTypeValues()
}
func PossibleAttributeDeltaOperationTypeValues() []AttributeDeltaOperationType {
	return original.PossibleAttributeDeltaOperationTypeValues()
}
func PossibleAttributeMappingTypeValues() []AttributeMappingType {
	return original.PossibleAttributeMappingTypeValues()
}
func PossibleDeltaOperationTypeValues() []DeltaOperationType {
	return original.PossibleDeltaOperationTypeValues()
}
func PossibleHealthStatusValues() []HealthStatus {
	return original.PossibleHealthStatusValues()
}
func PossibleLevelValues() []Level {
	return original.PossibleLevelValues()
}
func PossibleMonitoringLevelValues() []MonitoringLevel {
	return original.PossibleMonitoringLevelValues()
}
func PossiblePasswordOperationTypesValues() []PasswordOperationTypes {
	return original.PossiblePasswordOperationTypesValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleValueDeltaOperationTypeValues() []ValueDeltaOperationType {
	return original.PossibleValueDeltaOperationTypeValues()
}
func PossibleValueTypeValues() []ValueType {
	return original.PossibleValueTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
