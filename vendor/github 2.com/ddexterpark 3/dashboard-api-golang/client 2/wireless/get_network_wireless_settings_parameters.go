// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// NewGetNetworkWirelessSettingsParams creates a new GetNetworkWirelessSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessSettingsParams() *GetNetworkWirelessSettingsParams {
	return &GetNetworkWirelessSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessSettingsParamsWithTimeout creates a new GetNetworkWirelessSettingsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessSettingsParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessSettingsParams {
	return &GetNetworkWirelessSettingsParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessSettingsParamsWithContext creates a new GetNetworkWirelessSettingsParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessSettingsParamsWithContext(ctx context.Context) *GetNetworkWirelessSettingsParams {
	return &GetNetworkWirelessSettingsParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessSettingsParamsWithHTTPClient creates a new GetNetworkWirelessSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessSettingsParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessSettingsParams {
	return &GetNetworkWirelessSettingsParams{
		HTTPClient: client,
	}
}

/* GetNetworkWirelessSettingsParams contains all the parameters to send to the API endpoint
   for the get network wireless settings operation.

   Typically these are written to a http.Request.
*/
type GetNetworkWirelessSettingsParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSettingsParams) WithDefaults() *GetNetworkWirelessSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless settings params
func (o *GetNetworkWirelessSettingsParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless settings params
func (o *GetNetworkWirelessSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless settings params
func (o *GetNetworkWirelessSettingsParams) WithContext(ctx context.Context) *GetNetworkWirelessSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless settings params
func (o *GetNetworkWirelessSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless settings params
func (o *GetNetworkWirelessSettingsParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless settings params
func (o *GetNetworkWirelessSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless settings params
func (o *GetNetworkWirelessSettingsParams) WithNetworkID(networkID string) *GetNetworkWirelessSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless settings params
func (o *GetNetworkWirelessSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
