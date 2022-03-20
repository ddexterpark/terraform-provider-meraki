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

// NewGetNetworkWirelessSsidFirewallL3FirewallRulesParams creates a new GetNetworkWirelessSsidFirewallL3FirewallRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessSsidFirewallL3FirewallRulesParams() *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	return &GetNetworkWirelessSsidFirewallL3FirewallRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessSsidFirewallL3FirewallRulesParamsWithTimeout creates a new GetNetworkWirelessSsidFirewallL3FirewallRulesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessSsidFirewallL3FirewallRulesParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	return &GetNetworkWirelessSsidFirewallL3FirewallRulesParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessSsidFirewallL3FirewallRulesParamsWithContext creates a new GetNetworkWirelessSsidFirewallL3FirewallRulesParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessSsidFirewallL3FirewallRulesParamsWithContext(ctx context.Context) *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	return &GetNetworkWirelessSsidFirewallL3FirewallRulesParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessSsidFirewallL3FirewallRulesParamsWithHTTPClient creates a new GetNetworkWirelessSsidFirewallL3FirewallRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessSsidFirewallL3FirewallRulesParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	return &GetNetworkWirelessSsidFirewallL3FirewallRulesParams{
		HTTPClient: client,
	}
}

/* GetNetworkWirelessSsidFirewallL3FirewallRulesParams contains all the parameters to send to the API endpoint
   for the get network wireless ssid firewall l3 firewall rules operation.

   Typically these are written to a http.Request.
*/
type GetNetworkWirelessSsidFirewallL3FirewallRulesParams struct {

	// NetworkID.
	NetworkID string

	// Number.
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless ssid firewall l3 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) WithDefaults() *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless ssid firewall l3 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) WithContext(ctx context.Context) *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) WithNetworkID(networkID string) *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) WithNumber(number string) *GetNetworkWirelessSsidFirewallL3FirewallRulesParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get network wireless ssid firewall l3 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessSsidFirewallL3FirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
