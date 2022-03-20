// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewUpdateNetworkMqttBrokerParams creates a new UpdateNetworkMqttBrokerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkMqttBrokerParams() *UpdateNetworkMqttBrokerParams {
	return &UpdateNetworkMqttBrokerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkMqttBrokerParamsWithTimeout creates a new UpdateNetworkMqttBrokerParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkMqttBrokerParamsWithTimeout(timeout time.Duration) *UpdateNetworkMqttBrokerParams {
	return &UpdateNetworkMqttBrokerParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkMqttBrokerParamsWithContext creates a new UpdateNetworkMqttBrokerParams object
// with the ability to set a context for a request.
func NewUpdateNetworkMqttBrokerParamsWithContext(ctx context.Context) *UpdateNetworkMqttBrokerParams {
	return &UpdateNetworkMqttBrokerParams{
		Context: ctx,
	}
}

// NewUpdateNetworkMqttBrokerParamsWithHTTPClient creates a new UpdateNetworkMqttBrokerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkMqttBrokerParamsWithHTTPClient(client *http.Client) *UpdateNetworkMqttBrokerParams {
	return &UpdateNetworkMqttBrokerParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkMqttBrokerParams contains all the parameters to send to the API endpoint
   for the update network mqtt broker operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkMqttBrokerParams struct {

	// MqttBrokerID.
	MqttBrokerID string

	// NetworkID.
	NetworkID string

	// UpdateNetworkMqttBroker.
	UpdateNetworkMqttBroker UpdateNetworkMqttBrokerBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network mqtt broker params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkMqttBrokerParams) WithDefaults() *UpdateNetworkMqttBrokerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network mqtt broker params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkMqttBrokerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) WithTimeout(timeout time.Duration) *UpdateNetworkMqttBrokerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) WithContext(ctx context.Context) *UpdateNetworkMqttBrokerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) WithHTTPClient(client *http.Client) *UpdateNetworkMqttBrokerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMqttBrokerID adds the mqttBrokerID to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) WithMqttBrokerID(mqttBrokerID string) *UpdateNetworkMqttBrokerParams {
	o.SetMqttBrokerID(mqttBrokerID)
	return o
}

// SetMqttBrokerID adds the mqttBrokerId to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) SetMqttBrokerID(mqttBrokerID string) {
	o.MqttBrokerID = mqttBrokerID
}

// WithNetworkID adds the networkID to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) WithNetworkID(networkID string) *UpdateNetworkMqttBrokerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkMqttBroker adds the updateNetworkMqttBroker to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) WithUpdateNetworkMqttBroker(updateNetworkMqttBroker UpdateNetworkMqttBrokerBody) *UpdateNetworkMqttBrokerParams {
	o.SetUpdateNetworkMqttBroker(updateNetworkMqttBroker)
	return o
}

// SetUpdateNetworkMqttBroker adds the updateNetworkMqttBroker to the update network mqtt broker params
func (o *UpdateNetworkMqttBrokerParams) SetUpdateNetworkMqttBroker(updateNetworkMqttBroker UpdateNetworkMqttBrokerBody) {
	o.UpdateNetworkMqttBroker = updateNetworkMqttBroker
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkMqttBrokerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mqttBrokerId
	if err := r.SetPathParam("mqttBrokerId", o.MqttBrokerID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkMqttBroker); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
