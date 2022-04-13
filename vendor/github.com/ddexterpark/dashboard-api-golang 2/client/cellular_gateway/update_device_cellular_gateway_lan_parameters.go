// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

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

// NewUpdateDeviceCellularGatewayLanParams creates a new UpdateDeviceCellularGatewayLanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceCellularGatewayLanParams() *UpdateDeviceCellularGatewayLanParams {
	return &UpdateDeviceCellularGatewayLanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceCellularGatewayLanParamsWithTimeout creates a new UpdateDeviceCellularGatewayLanParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceCellularGatewayLanParamsWithTimeout(timeout time.Duration) *UpdateDeviceCellularGatewayLanParams {
	return &UpdateDeviceCellularGatewayLanParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceCellularGatewayLanParamsWithContext creates a new UpdateDeviceCellularGatewayLanParams object
// with the ability to set a context for a request.
func NewUpdateDeviceCellularGatewayLanParamsWithContext(ctx context.Context) *UpdateDeviceCellularGatewayLanParams {
	return &UpdateDeviceCellularGatewayLanParams{
		Context: ctx,
	}
}

// NewUpdateDeviceCellularGatewayLanParamsWithHTTPClient creates a new UpdateDeviceCellularGatewayLanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceCellularGatewayLanParamsWithHTTPClient(client *http.Client) *UpdateDeviceCellularGatewayLanParams {
	return &UpdateDeviceCellularGatewayLanParams{
		HTTPClient: client,
	}
}

/* UpdateDeviceCellularGatewayLanParams contains all the parameters to send to the API endpoint
   for the update device cellular gateway lan operation.

   Typically these are written to a http.Request.
*/
type UpdateDeviceCellularGatewayLanParams struct {

	// Serial.
	Serial string

	// UpdateDeviceCellularGatewayLan.
	UpdateDeviceCellularGatewayLan UpdateDeviceCellularGatewayLanBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device cellular gateway lan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceCellularGatewayLanParams) WithDefaults() *UpdateDeviceCellularGatewayLanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device cellular gateway lan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceCellularGatewayLanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) WithTimeout(timeout time.Duration) *UpdateDeviceCellularGatewayLanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) WithContext(ctx context.Context) *UpdateDeviceCellularGatewayLanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) WithHTTPClient(client *http.Client) *UpdateDeviceCellularGatewayLanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) WithSerial(serial string) *UpdateDeviceCellularGatewayLanParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceCellularGatewayLan adds the updateDeviceCellularGatewayLan to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) WithUpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan UpdateDeviceCellularGatewayLanBody) *UpdateDeviceCellularGatewayLanParams {
	o.SetUpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan)
	return o
}

// SetUpdateDeviceCellularGatewayLan adds the updateDeviceCellularGatewayLan to the update device cellular gateway lan params
func (o *UpdateDeviceCellularGatewayLanParams) SetUpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan UpdateDeviceCellularGatewayLanBody) {
	o.UpdateDeviceCellularGatewayLan = updateDeviceCellularGatewayLan
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceCellularGatewayLanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateDeviceCellularGatewayLan); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
