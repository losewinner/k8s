package eventgrid

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TopicTypesClient is the azure EventGrid Management Client
type TopicTypesClient struct {
	BaseClient
}

// NewTopicTypesClient creates an instance of the TopicTypesClient client.
func NewTopicTypesClient(subscriptionID string) TopicTypesClient {
	return NewTopicTypesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTopicTypesClientWithBaseURI creates an instance of the TopicTypesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTopicTypesClientWithBaseURI(baseURI string, subscriptionID string) TopicTypesClient {
	return TopicTypesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get information about a topic type
// Parameters:
// topicTypeName - name of the topic type
func (client TopicTypesClient) Get(ctx context.Context, topicTypeName string) (result TopicTypeInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopicTypesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, topicTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TopicTypesClient) GetPreparer(ctx context.Context, topicTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"topicTypeName": autorest.Encode("path", topicTypeName),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TopicTypesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TopicTypesClient) GetResponder(resp *http.Response) (result TopicTypeInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all registered topic types
func (client TopicTypesClient) List(ctx context.Context) (result TopicTypesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopicTypesClient.List")
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
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client TopicTypesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.EventGrid/topicTypes"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TopicTypesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TopicTypesClient) ListResponder(resp *http.Response) (result TopicTypesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListEventTypes list event types for a topic type
// Parameters:
// topicTypeName - name of the topic type
func (client TopicTypesClient) ListEventTypes(ctx context.Context, topicTypeName string) (result EventTypesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopicTypesClient.ListEventTypes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListEventTypesPreparer(ctx, topicTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "ListEventTypes", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListEventTypesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "ListEventTypes", resp, "Failure sending request")
		return
	}

	result, err = client.ListEventTypesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "ListEventTypes", resp, "Failure responding to request")
	}

	return
}

// ListEventTypesPreparer prepares the ListEventTypes request.
func (client TopicTypesClient) ListEventTypesPreparer(ctx context.Context, topicTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"topicTypeName": autorest.Encode("path", topicTypeName),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListEventTypesSender sends the ListEventTypes request. The method will close the
// http.Response Body if it receives an error.
func (client TopicTypesClient) ListEventTypesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListEventTypesResponder handles the response to the ListEventTypes request. The method always
// closes the http.Response Body.
func (client TopicTypesClient) ListEventTypesResponder(resp *http.Response) (result EventTypesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
