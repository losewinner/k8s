package dtl

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PolicySetClient is the azure DevTest Labs REST API version 2015-05-21-preview.
type PolicySetClient struct {
	BaseClient
}

// NewPolicySetClient creates an instance of the PolicySetClient client.
func NewPolicySetClient(subscriptionID string) PolicySetClient {
	return NewPolicySetClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPolicySetClientWithBaseURI creates an instance of the PolicySetClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPolicySetClientWithBaseURI(baseURI string, subscriptionID string) PolicySetClient {
	return PolicySetClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// EvaluatePolicies evaluates Lab Policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the policy set.
func (client PolicySetClient) EvaluatePolicies(ctx context.Context, resourceGroupName string, labName string, name string, evaluatePoliciesRequest EvaluatePoliciesRequest) (result EvaluatePoliciesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicySetClient.EvaluatePolicies")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EvaluatePoliciesPreparer(ctx, resourceGroupName, labName, name, evaluatePoliciesRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.PolicySetClient", "EvaluatePolicies", nil, "Failure preparing request")
		return
	}

	resp, err := client.EvaluatePoliciesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.PolicySetClient", "EvaluatePolicies", resp, "Failure sending request")
		return
	}

	result, err = client.EvaluatePoliciesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.PolicySetClient", "EvaluatePolicies", resp, "Failure responding to request")
		return
	}

	return
}

// EvaluatePoliciesPreparer prepares the EvaluatePolicies request.
func (client PolicySetClient) EvaluatePoliciesPreparer(ctx context.Context, resourceGroupName string, labName string, name string, evaluatePoliciesRequest EvaluatePoliciesRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{name}/evaluatePolicies", pathParameters),
		autorest.WithJSON(evaluatePoliciesRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EvaluatePoliciesSender sends the EvaluatePolicies request. The method will close the
// http.Response Body if it receives an error.
func (client PolicySetClient) EvaluatePoliciesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EvaluatePoliciesResponder handles the response to the EvaluatePolicies request. The method always
// closes the http.Response Body.
func (client PolicySetClient) EvaluatePoliciesResponder(resp *http.Response) (result EvaluatePoliciesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
