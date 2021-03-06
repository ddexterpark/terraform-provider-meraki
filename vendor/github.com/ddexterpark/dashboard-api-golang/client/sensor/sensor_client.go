// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sensor API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sensor API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetOrganizationSensorReadingsHistory(params *GetOrganizationSensorReadingsHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationSensorReadingsHistoryOK, error)

	GetOrganizationSensorReadingsLatest(params *GetOrganizationSensorReadingsLatestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationSensorReadingsLatestOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetOrganizationSensorReadingsHistory returns all reported readings from sensors in a given timespan sorted by timestamp

  Return all reported readings from sensors in a given timespan, sorted by timestamp
*/
func (a *Client) GetOrganizationSensorReadingsHistory(params *GetOrganizationSensorReadingsHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationSensorReadingsHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationSensorReadingsHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationSensorReadingsHistory",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationId}/sensor/readings/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationSensorReadingsHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationSensorReadingsHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationSensorReadingsHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationSensorReadingsLatest returns the latest available reading for each metric from each sensor sorted by sensor serial

  Return the latest available reading for each metric from each sensor, sorted by sensor serial
*/
func (a *Client) GetOrganizationSensorReadingsLatest(params *GetOrganizationSensorReadingsLatestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationSensorReadingsLatestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationSensorReadingsLatestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationSensorReadingsLatest",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationId}/sensor/readings/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationSensorReadingsLatestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationSensorReadingsLatestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationSensorReadingsLatest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
