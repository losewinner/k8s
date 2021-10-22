package documentdb

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

// RestorableDatabaseAccountsClient is the client for the RestorableDatabaseAccounts methods of the Documentdb service.
type RestorableDatabaseAccountsClient struct {
	BaseClient
}

// NewRestorableDatabaseAccountsClient creates an instance of the RestorableDatabaseAccountsClient client.
func NewRestorableDatabaseAccountsClient(subscriptionID string) RestorableDatabaseAccountsClient {
	return NewRestorableDatabaseAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestorableDatabaseAccountsClientWithBaseURI creates an instance of the RestorableDatabaseAccountsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewRestorableDatabaseAccountsClientWithBaseURI(baseURI string, subscriptionID string) RestorableDatabaseAccountsClient {
	return RestorableDatabaseAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByLocation retrieves the properties of an existing Azure Cosmos DB restorable database account.
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
// instanceID - the instanceId GUID of a restorable database account.
func (client RestorableDatabaseAccountsClient) GetByLocation(ctx context.Context, location string, instanceID string) (result RestorableDatabaseAccountGetResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableDatabaseAccountsClient.GetByLocation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.RestorableDatabaseAccountsClient", "GetByLocation", err.Error())
	}

	req, err := client.GetByLocationPreparer(ctx, location, instanceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "GetByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "GetByLocation", resp, "Failure sending request")
		return
	}

	result, err = client.GetByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "GetByLocation", resp, "Failure responding to request")
		return
	}

	return
}

// GetByLocationPreparer prepares the GetByLocation request.
func (client RestorableDatabaseAccountsClient) GetByLocationPreparer(ctx context.Context, location string, instanceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":     autorest.Encode("path", instanceID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByLocationSender sends the GetByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableDatabaseAccountsClient) GetByLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByLocationResponder handles the response to the GetByLocation request. The method always
// closes the http.Response Body.
func (client RestorableDatabaseAccountsClient) GetByLocationResponder(resp *http.Response) (result RestorableDatabaseAccountGetResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the restorable Azure Cosmos DB database accounts available under the subscription.
func (client RestorableDatabaseAccountsClient) List(ctx context.Context) (result RestorableDatabaseAccountsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableDatabaseAccountsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.RestorableDatabaseAccountsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RestorableDatabaseAccountsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableDatabaseAccountsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RestorableDatabaseAccountsClient) ListResponder(resp *http.Response) (result RestorableDatabaseAccountsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByLocation lists all the restorable Azure Cosmos DB database accounts available under the subscription and in a
// region.
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
func (client RestorableDatabaseAccountsClient) ListByLocation(ctx context.Context, location string) (result RestorableDatabaseAccountsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableDatabaseAccountsClient.ListByLocation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.RestorableDatabaseAccountsClient", "ListByLocation", err.Error())
	}

	req, err := client.ListByLocationPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableDatabaseAccountsClient", "ListByLocation", resp, "Failure responding to request")
		return
	}

	return
}

// ListByLocationPreparer prepares the ListByLocation request.
func (client RestorableDatabaseAccountsClient) ListByLocationPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableDatabaseAccountsClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client RestorableDatabaseAccountsClient) ListByLocationResponder(resp *http.Response) (result RestorableDatabaseAccountsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
