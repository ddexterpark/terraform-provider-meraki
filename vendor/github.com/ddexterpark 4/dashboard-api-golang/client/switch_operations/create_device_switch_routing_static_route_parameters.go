// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// NewCreateDeviceSwitchRoutingStaticRouteParams creates a new CreateDeviceSwitchRoutingStaticRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDeviceSwitchRoutingStaticRouteParams() *CreateDeviceSwitchRoutingStaticRouteParams {
	return &CreateDeviceSwitchRoutingStaticRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeviceSwitchRoutingStaticRouteParamsWithTimeout creates a new CreateDeviceSwitchRoutingStaticRouteParams object
// with the ability to set a timeout on a request.
func NewCreateDeviceSwitchRoutingStaticRouteParamsWithTimeout(timeout time.Duration) *CreateDeviceSwitchRoutingStaticRouteParams {
	return &CreateDeviceSwitchRoutingStaticRouteParams{
		timeout: timeout,
	}
}

// NewCreateDeviceSwitchRoutingStaticRouteParamsWithContext creates a new CreateDeviceSwitchRoutingStaticRouteParams object
// with the ability to set a context for a request.
func NewCreateDeviceSwitchRoutingStaticRouteParamsWithContext(ctx context.Context) *CreateDeviceSwitchRoutingStaticRouteParams {
	return &CreateDeviceSwitchRoutingStaticRouteParams{
		Context: ctx,
	}
}

// NewCreateDeviceSwitchRoutingStaticRouteParamsWithHTTPClient creates a new CreateDeviceSwitchRoutingStaticRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDeviceSwitchRoutingStaticRouteParamsWithHTTPClient(client *http.Client) *CreateDeviceSwitchRoutingStaticRouteParams {
	return &CreateDeviceSwitchRoutingStaticRouteParams{
		HTTPClient: client,
	}
}

/* CreateDeviceSwitchRoutingStaticRouteParams contains all the parameters to send to the API endpoint
   for the create device switch routing static route operation.

   Typically these are written to a http.Request.
*/
type CreateDeviceSwitchRoutingStaticRouteParams struct {

	// CreateDeviceSwitchRoutingStaticRoute.
	CreateDeviceSwitchRoutingStaticRoute CreateDeviceSwitchRoutingStaticRouteBody

	// Serial.
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create device switch routing static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeviceSwitchRoutingStaticRouteParams) WithDefaults() *CreateDeviceSwitchRoutingStaticRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create device switch routing static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeviceSwitchRoutingStaticRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) WithTimeout(timeout time.Duration) *CreateDeviceSwitchRoutingStaticRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) WithContext(ctx context.Context) *CreateDeviceSwitchRoutingStaticRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) WithHTTPClient(client *http.Client) *CreateDeviceSwitchRoutingStaticRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateDeviceSwitchRoutingStaticRoute adds the createDeviceSwitchRoutingStaticRoute to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) WithCreateDeviceSwitchRoutingStaticRoute(createDeviceSwitchRoutingStaticRoute CreateDeviceSwitchRoutingStaticRouteBody) *CreateDeviceSwitchRoutingStaticRouteParams {
	o.SetCreateDeviceSwitchRoutingStaticRoute(createDeviceSwitchRoutingStaticRoute)
	return o
}

// SetCreateDeviceSwitchRoutingStaticRoute adds the createDeviceSwitchRoutingStaticRoute to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) SetCreateDeviceSwitchRoutingStaticRoute(createDeviceSwitchRoutingStaticRoute CreateDeviceSwitchRoutingStaticRouteBody) {
	o.CreateDeviceSwitchRoutingStaticRoute = createDeviceSwitchRoutingStaticRoute
}

// WithSerial adds the serial to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) WithSerial(serial string) *CreateDeviceSwitchRoutingStaticRouteParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the create device switch routing static route params
func (o *CreateDeviceSwitchRoutingStaticRouteParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeviceSwitchRoutingStaticRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateDeviceSwitchRoutingStaticRoute); err != nil {
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
