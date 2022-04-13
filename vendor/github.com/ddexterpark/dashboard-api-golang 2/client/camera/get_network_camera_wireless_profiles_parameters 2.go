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

// NewGetNetworkCameraWirelessProfilesParams creates a new GetNetworkCameraWirelessProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkCameraWirelessProfilesParams() *GetNetworkCameraWirelessProfilesParams {
	return &GetNetworkCameraWirelessProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkCameraWirelessProfilesParamsWithTimeout creates a new GetNetworkCameraWirelessProfilesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkCameraWirelessProfilesParamsWithTimeout(timeout time.Duration) *GetNetworkCameraWirelessProfilesParams {
	return &GetNetworkCameraWirelessProfilesParams{
		timeout: timeout,
	}
}

// NewGetNetworkCameraWirelessProfilesParamsWithContext creates a new GetNetworkCameraWirelessProfilesParams object
// with the ability to set a context for a request.
func NewGetNetworkCameraWirelessProfilesParamsWithContext(ctx context.Context) *GetNetworkCameraWirelessProfilesParams {
	return &GetNetworkCameraWirelessProfilesParams{
		Context: ctx,
	}
}

// NewGetNetworkCameraWirelessProfilesParamsWithHTTPClient creates a new GetNetworkCameraWirelessProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkCameraWirelessProfilesParamsWithHTTPClient(client *http.Client) *GetNetworkCameraWirelessProfilesParams {
	return &GetNetworkCameraWirelessProfilesParams{
		HTTPClient: client,
	}
}

/* GetNetworkCameraWirelessProfilesParams contains all the parameters to send to the API endpoint
   for the get network camera wireless profiles operation.

   Typically these are written to a http.Request.
*/
type GetNetworkCameraWirelessProfilesParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network camera wireless profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCameraWirelessProfilesParams) WithDefaults() *GetNetworkCameraWirelessProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network camera wireless profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCameraWirelessProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network camera wireless profiles params
func (o *GetNetworkCameraWirelessProfilesParams) WithTimeout(timeout time.Duration) *GetNetworkCameraWirelessProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network camera wireless profiles params
func (o *GetNetworkCameraWirelessProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network camera wireless profiles params
func (o *GetNetworkCameraWirelessProfilesParams) WithContext(ctx context.Context) *GetNetworkCameraWirelessProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network camera wireless profiles params
func (o *GetNetworkCameraWirelessProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network camera wireless profiles params
func (o *GetNetworkCameraWirelessProfilesParams) WithHTTPClient(client *http.Client) *GetNetworkCameraWirelessProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network camera wireless profiles params
func (o *GetNetworkCameraWirelessProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network camera wireless profiles params
func (o *GetNetworkCameraWirelessProfilesParams) WithNetworkID(networkID string) *GetNetworkCameraWirelessProfilesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network camera wireless profiles params
func (o *GetNetworkCameraWirelessProfilesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkCameraWirelessProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
