// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewCreateDeviceLiveToolsPingParams creates a new CreateDeviceLiveToolsPingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDeviceLiveToolsPingParams() *CreateDeviceLiveToolsPingParams {
	return &CreateDeviceLiveToolsPingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeviceLiveToolsPingParamsWithTimeout creates a new CreateDeviceLiveToolsPingParams object
// with the ability to set a timeout on a request.
func NewCreateDeviceLiveToolsPingParamsWithTimeout(timeout time.Duration) *CreateDeviceLiveToolsPingParams {
	return &CreateDeviceLiveToolsPingParams{
		timeout: timeout,
	}
}

// NewCreateDeviceLiveToolsPingParamsWithContext creates a new CreateDeviceLiveToolsPingParams object
// with the ability to set a context for a request.
func NewCreateDeviceLiveToolsPingParamsWithContext(ctx context.Context) *CreateDeviceLiveToolsPingParams {
	return &CreateDeviceLiveToolsPingParams{
		Context: ctx,
	}
}

// NewCreateDeviceLiveToolsPingParamsWithHTTPClient creates a new CreateDeviceLiveToolsPingParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDeviceLiveToolsPingParamsWithHTTPClient(client *http.Client) *CreateDeviceLiveToolsPingParams {
	return &CreateDeviceLiveToolsPingParams{
		HTTPClient: client,
	}
}

/* CreateDeviceLiveToolsPingParams contains all the parameters to send to the API endpoint
   for the create device live tools ping operation.

   Typically these are written to a http.Request.
*/
type CreateDeviceLiveToolsPingParams struct {

	// CreateDeviceLiveToolsPing.
	CreateDeviceLiveToolsPing CreateDeviceLiveToolsPingBody

	// Serial.
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create device live tools ping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeviceLiveToolsPingParams) WithDefaults() *CreateDeviceLiveToolsPingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create device live tools ping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeviceLiveToolsPingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) WithTimeout(timeout time.Duration) *CreateDeviceLiveToolsPingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) WithContext(ctx context.Context) *CreateDeviceLiveToolsPingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) WithHTTPClient(client *http.Client) *CreateDeviceLiveToolsPingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateDeviceLiveToolsPing adds the createDeviceLiveToolsPing to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) WithCreateDeviceLiveToolsPing(createDeviceLiveToolsPing CreateDeviceLiveToolsPingBody) *CreateDeviceLiveToolsPingParams {
	o.SetCreateDeviceLiveToolsPing(createDeviceLiveToolsPing)
	return o
}

// SetCreateDeviceLiveToolsPing adds the createDeviceLiveToolsPing to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) SetCreateDeviceLiveToolsPing(createDeviceLiveToolsPing CreateDeviceLiveToolsPingBody) {
	o.CreateDeviceLiveToolsPing = createDeviceLiveToolsPing
}

// WithSerial adds the serial to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) WithSerial(serial string) *CreateDeviceLiveToolsPingParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the create device live tools ping params
func (o *CreateDeviceLiveToolsPingParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeviceLiveToolsPingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateDeviceLiveToolsPing); err != nil {
		return err
	}

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
