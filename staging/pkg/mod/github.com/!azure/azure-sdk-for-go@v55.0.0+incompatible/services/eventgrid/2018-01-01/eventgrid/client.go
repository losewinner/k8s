// Package eventgrid implements the Azure ARM Eventgrid service API version 2018-01-01.
//
// EventGrid Client
package eventgrid

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

// BaseClient is the base client for Eventgrid.
type BaseClient struct {
	autorest.Client
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithoutDefaults()
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults() BaseClient {
	return BaseClient{
		Client: autorest.NewClientWithUserAgent(UserAgent()),
	}
}

// PublishCloudEventEvents publishes a batch of events to an Azure Event Grid topic.
// Parameters:
// topicHostname - the host name of the topic, e.g. topic1.westus2-1.eventgrid.azure.net
// events - an array of events to be published to Event Grid.
func (client BaseClient) PublishCloudEventEvents(ctx context.Context, topicHostname string, events []CloudEventEvent) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PublishCloudEventEvents")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: events,
			Constraints: []validation.Constraint{{Target: "events", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventgrid.BaseClient", "PublishCloudEventEvents", err.Error())
	}

	req, err := client.PublishCloudEventEventsPreparer(ctx, topicHostname, events)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishCloudEventEvents", nil, "Failure preparing request")
		return
	}

	resp, err := client.PublishCloudEventEventsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishCloudEventEvents", resp, "Failure sending request")
		return
	}

	result, err = client.PublishCloudEventEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishCloudEventEvents", resp, "Failure responding to request")
		return
	}

	return
}

// PublishCloudEventEventsPreparer prepares the PublishCloudEventEvents request.
func (client BaseClient) PublishCloudEventEventsPreparer(ctx context.Context, topicHostname string, events []CloudEventEvent) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"topicHostname": topicHostname,
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/cloudevents-batch+json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{topicHostname}", urlParameters),
		autorest.WithPath("/api/events"),
		autorest.WithJSON(events),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PublishCloudEventEventsSender sends the PublishCloudEventEvents request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PublishCloudEventEventsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PublishCloudEventEventsResponder handles the response to the PublishCloudEventEvents request. The method always
// closes the http.Response Body.
func (client BaseClient) PublishCloudEventEventsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PublishCustomEventEvents publishes a batch of events to an Azure Event Grid topic.
// Parameters:
// topicHostname - the host name of the topic, e.g. topic1.westus2-1.eventgrid.azure.net
// events - an array of events to be published to Event Grid.
func (client BaseClient) PublishCustomEventEvents(ctx context.Context, topicHostname string, events []interface{}) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PublishCustomEventEvents")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: events,
			Constraints: []validation.Constraint{{Target: "events", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventgrid.BaseClient", "PublishCustomEventEvents", err.Error())
	}

	req, err := client.PublishCustomEventEventsPreparer(ctx, topicHostname, events)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishCustomEventEvents", nil, "Failure preparing request")
		return
	}

	resp, err := client.PublishCustomEventEventsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishCustomEventEvents", resp, "Failure sending request")
		return
	}

	result, err = client.PublishCustomEventEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishCustomEventEvents", resp, "Failure responding to request")
		return
	}

	return
}

// PublishCustomEventEventsPreparer prepares the PublishCustomEventEvents request.
func (client BaseClient) PublishCustomEventEventsPreparer(ctx context.Context, topicHostname string, events []interface{}) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"topicHostname": topicHostname,
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{topicHostname}", urlParameters),
		autorest.WithPath("/api/events"),
		autorest.WithJSON(events),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PublishCustomEventEventsSender sends the PublishCustomEventEvents request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PublishCustomEventEventsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PublishCustomEventEventsResponder handles the response to the PublishCustomEventEvents request. The method always
// closes the http.Response Body.
func (client BaseClient) PublishCustomEventEventsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PublishEvents publishes a batch of events to an Azure Event Grid topic.
// Parameters:
// topicHostname - the host name of the topic, e.g. topic1.westus2-1.eventgrid.azure.net
// events - an array of events to be published to Event Grid.
func (client BaseClient) PublishEvents(ctx context.Context, topicHostname string, events []Event) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PublishEvents")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: events,
			Constraints: []validation.Constraint{{Target: "events", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventgrid.BaseClient", "PublishEvents", err.Error())
	}

	req, err := client.PublishEventsPreparer(ctx, topicHostname, events)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishEvents", nil, "Failure preparing request")
		return
	}

	resp, err := client.PublishEventsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishEvents", resp, "Failure sending request")
		return
	}

	result, err = client.PublishEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishEvents", resp, "Failure responding to request")
		return
	}

	return
}

// PublishEventsPreparer prepares the PublishEvents request.
func (client BaseClient) PublishEventsPreparer(ctx context.Context, topicHostname string, events []Event) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"topicHostname": topicHostname,
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{topicHostname}", urlParameters),
		autorest.WithPath("/api/events"),
		autorest.WithJSON(events),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PublishEventsSender sends the PublishEvents request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PublishEventsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PublishEventsResponder handles the response to the PublishEvents request. The method always
// closes the http.Response Body.
func (client BaseClient) PublishEventsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
