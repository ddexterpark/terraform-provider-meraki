// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// NewGetDeviceApplianceDhcpSubnetsParams creates a new GetDeviceApplianceDhcpSubnetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceApplianceDhcpSubnetsParams() *GetDeviceApplianceDhcpSubnetsParams {
	return &GetDeviceApplianceDhcpSubnetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceApplianceDhcpSubnetsParamsWithTimeout creates a new GetDeviceApplianceDhcpSubnetsParams object
// with the ability to set a timeout on a request.
func NewGetDeviceApplianceDhcpSubnetsParamsWithTimeout(timeout time.Duration) *GetDeviceApplianceDhcpSubnetsParams {
	return &GetDeviceApplianceDhcpSubnetsParams{
		timeout: timeout,
	}
}

// NewGetDeviceApplianceDhcpSubnetsParamsWithContext creates a new GetDeviceApplianceDhcpSubnetsParams object
// with the ability to set a context for a request.
func NewGetDeviceApplianceDhcpSubnetsParamsWithContext(ctx context.Context) *GetDeviceApplianceDhcpSubnetsParams {
	return &GetDeviceApplianceDhcpSubnetsParams{
		Context: ctx,
	}
}

// NewGetDeviceApplianceDhcpSubnetsParamsWithHTTPClient creates a new GetDeviceApplianceDhcpSubnetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceApplianceDhcpSubnetsParamsWithHTTPClient(client *http.Client) *GetDeviceApplianceDhcpSubnetsParams {
	return &GetDeviceApplianceDhcpSubnetsParams{
		HTTPClient: client,
	}
}

/* GetDeviceApplianceDhcpSubnetsParams contains all the parameters to send to the API endpoint
   for the get device appliance dhcp subnets operation.

   Typically these are written to a http.Request.
*/
type GetDeviceApplianceDhcpSubnetsParams struct {

	// Serial.
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device appliance dhcp subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceApplianceDhcpSubnetsParams) WithDefaults() *GetDeviceApplianceDhcpSubnetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device appliance dhcp subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceApplianceDhcpSubnetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device appliance dhcp subnets params
func (o *GetDeviceApplianceDhcpSubnetsParams) WithTimeout(timeout time.Duration) *GetDeviceApplianceDhcpSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device appliance dhcp subnets params
func (o *GetDeviceApplianceDhcpSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device appliance dhcp subnets params
func (o *GetDeviceApplianceDhcpSubnetsParams) WithContext(ctx context.Context) *GetDeviceApplianceDhcpSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device appliance dhcp subnets params
func (o *GetDeviceApplianceDhcpSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device appliance dhcp subnets params
func (o *GetDeviceApplianceDhcpSubnetsParams) WithHTTPClient(client *http.Client) *GetDeviceApplianceDhcpSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device appliance dhcp subnets params
func (o *GetDeviceApplianceDhcpSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device appliance dhcp subnets params
func (o *GetDeviceApplianceDhcpSubnetsParams) WithSerial(serial string) *GetDeviceApplianceDhcpSubnetsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device appliance dhcp subnets params
func (o *GetDeviceApplianceDhcpSubnetsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceApplianceDhcpSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
