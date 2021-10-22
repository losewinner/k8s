package servicefabric

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MeshVolumeClient is the service Fabric REST Client APIs allows management of Service Fabric clusters, applications
// and services.
type MeshVolumeClient struct {
	BaseClient
}

// NewMeshVolumeClient creates an instance of the MeshVolumeClient client.
func NewMeshVolumeClient() MeshVolumeClient {
	return NewMeshVolumeClientWithBaseURI(DefaultBaseURI)
}

// NewMeshVolumeClientWithBaseURI creates an instance of the MeshVolumeClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMeshVolumeClientWithBaseURI(baseURI string) MeshVolumeClient {
	return MeshVolumeClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate creates a Volume resource with the specified name, description and properties. If Volume resource
// with the same name exists, then it is updated with the specified description and properties.
// Parameters:
// volumeResourceName - the identity of the volume.
// volumeResourceDescription - description for creating a Volume resource.
func (client MeshVolumeClient) CreateOrUpdate(ctx context.Context, volumeResourceName string, volumeResourceDescription VolumeResourceDescription) (result VolumeResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshVolumeClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: volumeResourceDescription,
			Constraints: []validation.Constraint{{Target: "volumeResourceDescription.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "volumeResourceDescription.VolumeProperties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "volumeResourceDescription.VolumeProperties.Provider", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "volumeResourceDescription.VolumeProperties.AzureFileParameters", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "volumeResourceDescription.VolumeProperties.AzureFileParameters.AccountName", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "volumeResourceDescription.VolumeProperties.AzureFileParameters.ShareName", Name: validation.Null, Rule: true, Chain: nil},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("servicefabric.MeshVolumeClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, volumeResourceName, volumeResourceDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client MeshVolumeClient) CreateOrUpdatePreparer(ctx context.Context, volumeResourceName string, volumeResourceDescription VolumeResourceDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"volumeResourceName": volumeResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Volumes/{volumeResourceName}", pathParameters),
		autorest.WithJSON(volumeResourceDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client MeshVolumeClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client MeshVolumeClient) CreateOrUpdateResponder(resp *http.Response) (result VolumeResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the Volume resource identified by the name.
// Parameters:
// volumeResourceName - the identity of the volume.
func (client MeshVolumeClient) Delete(ctx context.Context, volumeResourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshVolumeClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, volumeResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MeshVolumeClient) DeletePreparer(ctx context.Context, volumeResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"volumeResourceName": volumeResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Volumes/{volumeResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MeshVolumeClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MeshVolumeClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the information about the Volume resource with the given name. The information include the description and
// other properties of the Volume.
// Parameters:
// volumeResourceName - the identity of the volume.
func (client MeshVolumeClient) Get(ctx context.Context, volumeResourceName string) (result VolumeResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshVolumeClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, volumeResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MeshVolumeClient) GetPreparer(ctx context.Context, volumeResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"volumeResourceName": volumeResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Volumes/{volumeResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MeshVolumeClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MeshVolumeClient) GetResponder(resp *http.Response) (result VolumeResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the information about all volume resources in a given resource group. The information include the
// description and other properties of the Volume.
func (client MeshVolumeClient) List(ctx context.Context) (result PagedVolumeResourceDescriptionList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshVolumeClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshVolumeClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MeshVolumeClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/Resources/Volumes"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MeshVolumeClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MeshVolumeClient) ListResponder(resp *http.Response) (result PagedVolumeResourceDescriptionList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
