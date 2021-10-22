package servicemap

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SummariesClient is the service Map API Reference
type SummariesClient struct {
	BaseClient
}

// NewSummariesClient creates an instance of the SummariesClient client.
func NewSummariesClient(subscriptionID string) SummariesClient {
	return NewSummariesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSummariesClientWithBaseURI creates an instance of the SummariesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSummariesClientWithBaseURI(baseURI string, subscriptionID string) SummariesClient {
	return SummariesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetMachines returns summary information about the machines in the workspace.
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// startTime - UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m
// endTime - UTC date and time specifying the end time of an interval. When not specified the service uses
// DateTime.UtcNow
func (client SummariesClient) GetMachines(ctx context.Context, resourceGroupName string, workspaceName string, startTime *date.Time, endTime *date.Time) (result MachinesSummary, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SummariesClient.GetMachines")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.SummariesClient", "GetMachines", err.Error())
	}

	req, err := client.GetMachinesPreparer(ctx, resourceGroupName, workspaceName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.SummariesClient", "GetMachines", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMachinesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.SummariesClient", "GetMachines", resp, "Failure sending request")
		return
	}

	result, err = client.GetMachinesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.SummariesClient", "GetMachines", resp, "Failure responding to request")
		return
	}

	return
}

// GetMachinesPreparer prepares the GetMachines request.
func (client SummariesClient) GetMachinesPreparer(ctx context.Context, resourceGroupName string, workspaceName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/summaries/machines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMachinesSender sends the GetMachines request. The method will close the
// http.Response Body if it receives an error.
func (client SummariesClient) GetMachinesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetMachinesResponder handles the response to the GetMachines request. The method always
// closes the http.Response Body.
func (client SummariesClient) GetMachinesResponder(resp *http.Response) (result MachinesSummary, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
