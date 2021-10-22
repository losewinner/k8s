package backup

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

// ProtectionPolicyOperationStatusesClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionPolicyOperationStatusesClient struct {
	BaseClient
}

// NewProtectionPolicyOperationStatusesClient creates an instance of the ProtectionPolicyOperationStatusesClient
// client.
func NewProtectionPolicyOperationStatusesClient(subscriptionID string) ProtectionPolicyOperationStatusesClient {
	return NewProtectionPolicyOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionPolicyOperationStatusesClientWithBaseURI creates an instance of the
// ProtectionPolicyOperationStatusesClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProtectionPolicyOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) ProtectionPolicyOperationStatusesClient {
	return ProtectionPolicyOperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get provides the status of the asynchronous operations like backup, restore. The status can be in progress,
// completed
// or failed. You can refer to the Operation Status enum for all the possible states of an operation. Some operations
// create jobs. This method returns the list of jobs associated with operation.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// policyName - backup policy name whose operation's status needs to be fetched.
// operationID - operation ID which represents an operation whose status needs to be fetched.
func (client ProtectionPolicyOperationStatusesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectionPolicyOperationStatusesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, policyName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationStatusesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectionPolicyOperationStatusesClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}/operations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionPolicyOperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectionPolicyOperationStatusesClient) GetResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
