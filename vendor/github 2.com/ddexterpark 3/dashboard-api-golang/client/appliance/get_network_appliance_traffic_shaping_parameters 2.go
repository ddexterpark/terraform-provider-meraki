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

// NewGetNetworkApplianceTrafficShapingParams creates a new GetNetworkApplianceTrafficShapingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkApplianceTrafficShapingParams() *GetNetworkApplianceTrafficShapingParams {
	return &GetNetworkApplianceTrafficShapingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkApplianceTrafficShapingParamsWithTimeout creates a new GetNetworkApplianceTrafficShapingParams object
// with the ability to set a timeout on a request.
func NewGetNetworkApplianceTrafficShapingParamsWithTimeout(timeout time.Duration) *GetNetworkApplianceTrafficShapingParams {
	return &GetNetworkApplianceTrafficShapingParams{
		timeout: timeout,
	}
}

// NewGetNetworkApplianceTrafficShapingParamsWithContext creates a new GetNetworkApplianceTrafficShapingParams object
// with the ability to set a context for a request.
func NewGetNetworkApplianceTrafficShapingParamsWithContext(ctx context.Context) *GetNetworkApplianceTrafficShapingParams {
	return &GetNetworkApplianceTrafficShapingParams{
		Context: ctx,
	}
}

// NewGetNetworkApplianceTrafficShapingParamsWithHTTPClient creates a new GetNetworkApplianceTrafficShapingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkApplianceTrafficShapingParamsWithHTTPClient(client *http.Client) *GetNetworkApplianceTrafficShapingParams {
	return &GetNetworkApplianceTrafficShapingParams{
		HTTPClient: client,
	}
}

/* GetNetworkApplianceTrafficShapingParams contains all the parameters to send to the API endpoint
   for the get network appliance traffic shaping operation.

   Typically these are written to a http.Request.
*/
type GetNetworkApplianceTrafficShapingParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network appliance traffic shaping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceTrafficShapingParams) WithDefaults() *GetNetworkApplianceTrafficShapingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network appliance traffic shaping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceTrafficShapingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network appliance traffic shaping params
func (o *GetNetworkApplianceTrafficShapingParams) WithTimeout(timeout time.Duration) *GetNetworkApplianceTrafficShapingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network appliance traffic shaping params
func (o *GetNetworkApplianceTrafficShapingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network appliance traffic shaping params
func (o *GetNetworkApplianceTrafficShapingParams) WithContext(ctx context.Context) *GetNetworkApplianceTrafficShapingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network appliance traffic shaping params
func (o *GetNetworkApplianceTrafficShapingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network appliance traffic shaping params
func (o *GetNetworkApplianceTrafficShapingParams) WithHTTPClient(client *http.Client) *GetNetworkApplianceTrafficShapingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network appliance traffic shaping params
func (o *GetNetworkApplianceTrafficShapingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network appliance traffic shaping params
func (o *GetNetworkApplianceTrafficShapingParams) WithNetworkID(networkID string) *GetNetworkApplianceTrafficShapingParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network appliance traffic shaping params
func (o *GetNetworkApplianceTrafficShapingParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkApplianceTrafficShapingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
