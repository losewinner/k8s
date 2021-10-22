package labservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LabAccountsClient is the the Managed Labs Client.
type LabAccountsClient struct {
	BaseClient
}

// NewLabAccountsClient creates an instance of the LabAccountsClient client.
func NewLabAccountsClient(subscriptionID string) LabAccountsClient {
	return NewLabAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLabAccountsClientWithBaseURI creates an instance of the LabAccountsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLabAccountsClientWithBaseURI(baseURI string, subscriptionID string) LabAccountsClient {
	return LabAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateLab create a lab in a lab account.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// createLabProperties - properties for creating a managed lab and a default environment setting
func (client LabAccountsClient) CreateLab(ctx context.Context, resourceGroupName string, labAccountName string, createLabProperties CreateLabProperties) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.CreateLab")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createLabProperties,
			Constraints: []validation.Constraint{{Target: "createLabProperties.EnvironmentSettingCreationParameters", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "createLabProperties.EnvironmentSettingCreationParameters.ResourceSettingCreationParameters", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "createLabProperties.EnvironmentSettingCreationParameters.ResourceSettingCreationParameters.GalleryImageResourceID", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "createLabProperties.EnvironmentSettingCreationParameters.ResourceSettingCreationParameters.ReferenceVMCreationParameters", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "createLabProperties.EnvironmentSettingCreationParameters.ResourceSettingCreationParameters.ReferenceVMCreationParameters.UserName", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "createLabProperties.EnvironmentSettingCreationParameters.ResourceSettingCreationParameters.ReferenceVMCreationParameters.Password", Name: validation.Null, Rule: true, Chain: nil},
							}},
					}},
				}},
				{Target: "createLabProperties.LabCreationParameters", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "createLabProperties.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabAccountsClient", "CreateLab", err.Error())
	}

	req, err := client.CreateLabPreparer(ctx, resourceGroupName, labAccountName, createLabProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "CreateLab", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateLabSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "CreateLab", resp, "Failure sending request")
		return
	}

	result, err = client.CreateLabResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "CreateLab", resp, "Failure responding to request")
		return
	}

	return
}

// CreateLabPreparer prepares the CreateLab request.
func (client LabAccountsClient) CreateLabPreparer(ctx context.Context, resourceGroupName string, labAccountName string, createLabProperties CreateLabProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/createLab", pathParameters),
		autorest.WithJSON(createLabProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateLabSender sends the CreateLab request. The method will close the
// http.Response Body if it receives an error.
func (client LabAccountsClient) CreateLabSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateLabResponder handles the response to the CreateLab request. The method always
// closes the http.Response Body.
func (client LabAccountsClient) CreateLabResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create or replace an existing Lab Account.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labAccount - represents a lab account.
func (client LabAccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labAccount LabAccount) (result LabAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, labAccountName, labAccount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LabAccountsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labAccount LabAccount) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", pathParameters),
		autorest.WithJSON(labAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LabAccountsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LabAccountsClient) CreateOrUpdateResponder(resp *http.Response) (result LabAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete lab account. This operation can take a while to complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
func (client LabAccountsClient) Delete(ctx context.Context, resourceGroupName string, labAccountName string) (result LabAccountsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, labAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LabAccountsClient) DeletePreparer(ctx context.Context, resourceGroupName string, labAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LabAccountsClient) DeleteSender(req *http.Request) (future LabAccountsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LabAccountsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get lab account
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// expand - specify the $expand query. Example: 'properties($expand=sizeConfiguration)'
func (client LabAccountsClient) Get(ctx context.Context, resourceGroupName string, labAccountName string, expand string) (result LabAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labAccountName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LabAccountsClient) GetPreparer(ctx context.Context, resourceGroupName string, labAccountName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LabAccountsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LabAccountsClient) GetResponder(resp *http.Response) (result LabAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRegionalAvailability get regional availability information for each size category configured under a lab account
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
func (client LabAccountsClient) GetRegionalAvailability(ctx context.Context, resourceGroupName string, labAccountName string) (result GetRegionalAvailabilityResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.GetRegionalAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetRegionalAvailabilityPreparer(ctx, resourceGroupName, labAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "GetRegionalAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRegionalAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "GetRegionalAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.GetRegionalAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "GetRegionalAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// GetRegionalAvailabilityPreparer prepares the GetRegionalAvailability request.
func (client LabAccountsClient) GetRegionalAvailabilityPreparer(ctx context.Context, resourceGroupName string, labAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/getRegionalAvailability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRegionalAvailabilitySender sends the GetRegionalAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client LabAccountsClient) GetRegionalAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetRegionalAvailabilityResponder handles the response to the GetRegionalAvailability request. The method always
// closes the http.Response Body.
func (client LabAccountsClient) GetRegionalAvailabilityResponder(resp *http.Response) (result GetRegionalAvailabilityResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list lab accounts in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// expand - specify the $expand query. Example: 'properties($expand=sizeConfiguration)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client LabAccountsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationLabAccountPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.rwcla.Response.Response != nil {
				sc = result.rwcla.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.rwcla.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.rwcla, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.rwcla.hasNextLink() && result.rwcla.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client LabAccountsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client LabAccountsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client LabAccountsClient) ListByResourceGroupResponder(resp *http.Response) (result ResponseWithContinuationLabAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client LabAccountsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ResponseWithContinuationLabAccount) (result ResponseWithContinuationLabAccount, err error) {
	req, err := lastResults.responseWithContinuationLabAccountPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client LabAccountsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationLabAccountIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, expand, filter, top, orderby)
	return
}

// ListBySubscription list lab accounts in a subscription.
// Parameters:
// expand - specify the $expand query. Example: 'properties($expand=sizeConfiguration)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client LabAccountsClient) ListBySubscription(ctx context.Context, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationLabAccountPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.rwcla.Response.Response != nil {
				sc = result.rwcla.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.rwcla.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.rwcla, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.rwcla.hasNextLink() && result.rwcla.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client LabAccountsClient) ListBySubscriptionPreparer(ctx context.Context, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/labaccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client LabAccountsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client LabAccountsClient) ListBySubscriptionResponder(resp *http.Response) (result ResponseWithContinuationLabAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client LabAccountsClient) listBySubscriptionNextResults(ctx context.Context, lastResults ResponseWithContinuationLabAccount) (result ResponseWithContinuationLabAccount, err error) {
	req, err := lastResults.responseWithContinuationLabAccountPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client LabAccountsClient) ListBySubscriptionComplete(ctx context.Context, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationLabAccountIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, expand, filter, top, orderby)
	return
}

// Update modify properties of lab accounts.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labAccount - represents a lab account.
func (client LabAccountsClient) Update(ctx context.Context, resourceGroupName string, labAccountName string, labAccount LabAccountFragment) (result LabAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabAccountsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, labAccountName, labAccount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabAccountsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client LabAccountsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labAccount LabAccountFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", pathParameters),
		autorest.WithJSON(labAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client LabAccountsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LabAccountsClient) UpdateResponder(resp *http.Response) (result LabAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
