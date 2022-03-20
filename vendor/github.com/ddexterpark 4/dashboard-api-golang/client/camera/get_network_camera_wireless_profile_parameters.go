// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// NewGetNetworkCameraWirelessProfileParams creates a new GetNetworkCameraWirelessProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkCameraWirelessProfileParams() *GetNetworkCameraWirelessProfileParams {
	return &GetNetworkCameraWirelessProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkCameraWirelessProfileParamsWithTimeout creates a new GetNetworkCameraWirelessProfileParams object
// with the ability to set a timeout on a request.
func NewGetNetworkCameraWirelessProfileParamsWithTimeout(timeout time.Duration) *GetNetworkCameraWirelessProfileParams {
	return &GetNetworkCameraWirelessProfileParams{
		timeout: timeout,
	}
}

// NewGetNetworkCameraWirelessProfileParamsWithContext creates a new GetNetworkCameraWirelessProfileParams object
// with the ability to set a context for a request.
func NewGetNetworkCameraWirelessProfileParamsWithContext(ctx context.Context) *GetNetworkCameraWirelessProfileParams {
	return &GetNetworkCameraWirelessProfileParams{
		Context: ctx,
	}
}

// NewGetNetworkCameraWirelessProfileParamsWithHTTPClient creates a new GetNetworkCameraWirelessProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkCameraWirelessProfileParamsWithHTTPClient(client *http.Client) *GetNetworkCameraWirelessProfileParams {
	return &GetNetworkCameraWirelessProfileParams{
		HTTPClient: client,
	}
}

/* GetNetworkCameraWirelessProfileParams contains all the parameters to send to the API endpoint
   for the get network camera wireless profile operation.

   Typically these are written to a http.Request.
*/
type GetNetworkCameraWirelessProfileParams struct {

	// NetworkID.
	NetworkID string

	// WirelessProfileID.
	WirelessProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network camera wireless profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCameraWirelessProfileParams) WithDefaults() *GetNetworkCameraWirelessProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network camera wireless profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCameraWirelessProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) WithTimeout(timeout time.Duration) *GetNetworkCameraWirelessProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) WithContext(ctx context.Context) *GetNetworkCameraWirelessProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) WithHTTPClient(client *http.Client) *GetNetworkCameraWirelessProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) WithNetworkID(networkID string) *GetNetworkCameraWirelessProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithWirelessProfileID adds the wirelessProfileID to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) WithWirelessProfileID(wirelessProfileID string) *GetNetworkCameraWirelessProfileParams {
	o.SetWirelessProfileID(wirelessProfileID)
	return o
}

// SetWirelessProfileID adds the wirelessProfileId to the get network camera wireless profile params
func (o *GetNetworkCameraWirelessProfileParams) SetWirelessProfileID(wirelessProfileID string) {
	o.WirelessProfileID = wirelessProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkCameraWirelessProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param wirelessProfileId
	if err := r.SetPathParam("wirelessProfileId", o.WirelessProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
