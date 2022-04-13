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

// NewGetNetworkWirelessSsidEapOverrideParams creates a new GetNetworkWirelessSsidEapOverrideParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessSsidEapOverrideParams() *GetNetworkWirelessSsidEapOverrideParams {
	return &GetNetworkWirelessSsidEapOverrideParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessSsidEapOverrideParamsWithTimeout creates a new GetNetworkWirelessSsidEapOverrideParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessSsidEapOverrideParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessSsidEapOverrideParams {
	return &GetNetworkWirelessSsidEapOverrideParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessSsidEapOverrideParamsWithContext creates a new GetNetworkWirelessSsidEapOverrideParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessSsidEapOverrideParamsWithContext(ctx context.Context) *GetNetworkWirelessSsidEapOverrideParams {
	return &GetNetworkWirelessSsidEapOverrideParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessSsidEapOverrideParamsWithHTTPClient creates a new GetNetworkWirelessSsidEapOverrideParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessSsidEapOverrideParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessSsidEapOverrideParams {
	return &GetNetworkWirelessSsidEapOverrideParams{
		HTTPClient: client,
	}
}

/* GetNetworkWirelessSsidEapOverrideParams contains all the parameters to send to the API endpoint
   for the get network wireless ssid eap override operation.

   Typically these are written to a http.Request.
*/
type GetNetworkWirelessSsidEapOverrideParams struct {

	// NetworkID.
	NetworkID string

	// Number.
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless ssid eap override params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidEapOverrideParams) WithDefaults() *GetNetworkWirelessSsidEapOverrideParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless ssid eap override params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidEapOverrideParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessSsidEapOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) WithContext(ctx context.Context) *GetNetworkWirelessSsidEapOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessSsidEapOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) WithNetworkID(networkID string) *GetNetworkWirelessSsidEapOverrideParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) WithNumber(number string) *GetNetworkWirelessSsidEapOverrideParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get network wireless ssid eap override params
func (o *GetNetworkWirelessSsidEapOverrideParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessSsidEapOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
