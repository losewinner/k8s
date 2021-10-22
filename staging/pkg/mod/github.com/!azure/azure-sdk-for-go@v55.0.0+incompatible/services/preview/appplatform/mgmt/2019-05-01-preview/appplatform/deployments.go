package appplatform

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

// DeploymentsClient is the REST API for Azure Spring Cloud
type DeploymentsClient struct {
	BaseClient
}

// NewDeploymentsClient creates an instance of the DeploymentsClient client.
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return NewDeploymentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeploymentsClientWithBaseURI creates an instance of the DeploymentsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return DeploymentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a new Deployment or update an exiting Deployment.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// deploymentName - the name of the Deployment resource.
// deploymentResource - parameters for the create or update operation
func (client DeploymentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource DeploymentResource) (result DeploymentsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, appName, deploymentName, deploymentResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DeploymentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource DeploymentResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}", pathParameters),
		autorest.WithJSON(deploymentResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) CreateOrUpdateSender(req *http.Request) (future DeploymentsCreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) CreateOrUpdateResponder(resp *http.Response) (result DeploymentResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete operation to delete a Deployment.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// deploymentName - the name of the Deployment resource.
func (client DeploymentsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, appName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DeploymentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a Deployment and its properties.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// deploymentName - the name of the Deployment resource.
func (client DeploymentsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result DeploymentResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, appName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeploymentsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) GetResponder(resp *http.Response) (result DeploymentResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLogFileURL get deployment log file URL
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// deploymentName - the name of the Deployment resource.
func (client DeploymentsClient) GetLogFileURL(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result LogFileURLResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.GetLogFileURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetLogFileURLPreparer(ctx, resourceGroupName, serviceName, appName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "GetLogFileURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLogFileURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "GetLogFileURL", resp, "Failure sending request")
		return
	}

	result, err = client.GetLogFileURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "GetLogFileURL", resp, "Failure responding to request")
		return
	}

	return
}

// GetLogFileURLPreparer prepares the GetLogFileURL request.
func (client DeploymentsClient) GetLogFileURLPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}/getLogFileUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLogFileURLSender sends the GetLogFileURL request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) GetLogFileURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetLogFileURLResponder handles the response to the GetLogFileURL request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) GetLogFileURLResponder(resp *http.Response) (result LogFileURLResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List handles requests to list all resources in an App.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// version - version of the deployments to be listed
func (client DeploymentsClient) List(ctx context.Context, resourceGroupName string, serviceName string, appName string, version []string) (result DeploymentResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.List")
		defer func() {
			sc := -1
			if result.drc.Response.Response != nil {
				sc = result.drc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, appName, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.drc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "List", resp, "Failure sending request")
		return
	}

	result.drc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.drc.hasNextLink() && result.drc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DeploymentsClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, version []string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if version != nil && len(version) > 0 {
		queryParameters["version"] = version
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) ListResponder(resp *http.Response) (result DeploymentResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DeploymentsClient) listNextResults(ctx context.Context, lastResults DeploymentResourceCollection) (result DeploymentResourceCollection, err error) {
	req, err := lastResults.deploymentResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeploymentsClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, appName string, version []string) (result DeploymentResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, serviceName, appName, version)
	return
}

// ListClusterAllDeployments list deployments for a certain service
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// version - version of the deployments to be listed
func (client DeploymentsClient) ListClusterAllDeployments(ctx context.Context, resourceGroupName string, serviceName string, version []string) (result DeploymentResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.ListClusterAllDeployments")
		defer func() {
			sc := -1
			if result.drc.Response.Response != nil {
				sc = result.drc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listClusterAllDeploymentsNextResults
	req, err := client.ListClusterAllDeploymentsPreparer(ctx, resourceGroupName, serviceName, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "ListClusterAllDeployments", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListClusterAllDeploymentsSender(req)
	if err != nil {
		result.drc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "ListClusterAllDeployments", resp, "Failure sending request")
		return
	}

	result.drc, err = client.ListClusterAllDeploymentsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "ListClusterAllDeployments", resp, "Failure responding to request")
		return
	}
	if result.drc.hasNextLink() && result.drc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListClusterAllDeploymentsPreparer prepares the ListClusterAllDeployments request.
func (client DeploymentsClient) ListClusterAllDeploymentsPreparer(ctx context.Context, resourceGroupName string, serviceName string, version []string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if version != nil && len(version) > 0 {
		queryParameters["version"] = version
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/deployments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListClusterAllDeploymentsSender sends the ListClusterAllDeployments request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) ListClusterAllDeploymentsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListClusterAllDeploymentsResponder handles the response to the ListClusterAllDeployments request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) ListClusterAllDeploymentsResponder(resp *http.Response) (result DeploymentResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listClusterAllDeploymentsNextResults retrieves the next set of results, if any.
func (client DeploymentsClient) listClusterAllDeploymentsNextResults(ctx context.Context, lastResults DeploymentResourceCollection) (result DeploymentResourceCollection, err error) {
	req, err := lastResults.deploymentResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "listClusterAllDeploymentsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListClusterAllDeploymentsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "listClusterAllDeploymentsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListClusterAllDeploymentsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "listClusterAllDeploymentsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListClusterAllDeploymentsComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeploymentsClient) ListClusterAllDeploymentsComplete(ctx context.Context, resourceGroupName string, serviceName string, version []string) (result DeploymentResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.ListClusterAllDeployments")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListClusterAllDeployments(ctx, resourceGroupName, serviceName, version)
	return
}

// Restart restart the deployment.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// deploymentName - the name of the Deployment resource.
func (client DeploymentsClient) Restart(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result DeploymentsRestartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.Restart")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RestartPreparer(ctx, resourceGroupName, serviceName, appName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Restart", nil, "Failure preparing request")
		return
	}

	result, err = client.RestartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Restart", nil, "Failure sending request")
		return
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client DeploymentsClient) RestartPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) RestartSender(req *http.Request) (future DeploymentsRestartFuture, err error) {
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

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start start the deployment.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// deploymentName - the name of the Deployment resource.
func (client DeploymentsClient) Start(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result DeploymentsStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.Start")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, serviceName, appName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Start", nil, "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client DeploymentsClient) StartPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) StartSender(req *http.Request) (future DeploymentsStartFuture, err error) {
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

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stop the deployment.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// deploymentName - the name of the Deployment resource.
func (client DeploymentsClient) Stop(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result DeploymentsStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.Stop")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, serviceName, appName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Stop", nil, "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client DeploymentsClient) StopPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) StopSender(req *http.Request) (future DeploymentsStopFuture, err error) {
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

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update operation to update an exiting Deployment.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// deploymentName - the name of the Deployment resource.
// deploymentResource - parameters for the update operation
func (client DeploymentsClient) Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource DeploymentResource) (result DeploymentsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, appName, deploymentName, deploymentResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.DeploymentsClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DeploymentsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource DeploymentResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}", pathParameters),
		autorest.WithJSON(deploymentResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) UpdateSender(req *http.Request) (future DeploymentsUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) UpdateResponder(resp *http.Response) (result DeploymentResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
