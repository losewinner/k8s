package alertsmanagementapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/alertsmanagement/mgmt/2018-05-05-preview/alertsmanagement"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result alertsmanagement.OperationsListPage, err error)
	ListComplete(ctx context.Context) (result alertsmanagement.OperationsListIterator, err error)
}

var _ OperationsClientAPI = (*alertsmanagement.OperationsClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	ChangeState(ctx context.Context, alertID string, newState alertsmanagement.AlertState) (result alertsmanagement.Alert, err error)
	GetAll(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorCondition alertsmanagement.MonitorCondition, severity alertsmanagement.Severity, alertState alertsmanagement.AlertState, smartGroupID string, includePayload *bool, pageCount *int32, sortBy alertsmanagement.AlertsSortByFields, sortOrder string, timeRange alertsmanagement.TimeRange) (result alertsmanagement.AlertsListPage, err error)
	GetAllComplete(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorCondition alertsmanagement.MonitorCondition, severity alertsmanagement.Severity, alertState alertsmanagement.AlertState, smartGroupID string, includePayload *bool, pageCount *int32, sortBy alertsmanagement.AlertsSortByFields, sortOrder string, timeRange alertsmanagement.TimeRange) (result alertsmanagement.AlertsListIterator, err error)
	GetByID(ctx context.Context, alertID string) (result alertsmanagement.Alert, err error)
	GetHistory(ctx context.Context, alertID string) (result alertsmanagement.AlertModification, err error)
	GetSummary(ctx context.Context, targetResourceGroup string, timeRange alertsmanagement.TimeRange) (result alertsmanagement.AlertsSummary, err error)
}

var _ AlertsClientAPI = (*alertsmanagement.AlertsClient)(nil)

// SmartGroupsClientAPI contains the set of methods on the SmartGroupsClient type.
type SmartGroupsClientAPI interface {
	ChangeState(ctx context.Context, smartGroupID string, newState alertsmanagement.AlertState) (result alertsmanagement.SmartGroup, err error)
	GetAll(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorCondition alertsmanagement.MonitorCondition, severity alertsmanagement.Severity, smartGroupState alertsmanagement.AlertState, timeRange alertsmanagement.TimeRange, pageCount *int32, sortBy alertsmanagement.SmartGroupsSortByFields, sortOrder string) (result alertsmanagement.SmartGroupsList, err error)
	GetByID(ctx context.Context, smartGroupID string) (result alertsmanagement.SmartGroup, err error)
	GetHistory(ctx context.Context, smartGroupID string) (result alertsmanagement.SmartGroupModification, err error)
}

var _ SmartGroupsClientAPI = (*alertsmanagement.SmartGroupsClient)(nil)
