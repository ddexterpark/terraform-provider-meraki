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

// NewGetNetworkCameraSchedulesParams creates a new GetNetworkCameraSchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkCameraSchedulesParams() *GetNetworkCameraSchedulesParams {
	return &GetNetworkCameraSchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkCameraSchedulesParamsWithTimeout creates a new GetNetworkCameraSchedulesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkCameraSchedulesParamsWithTimeout(timeout time.Duration) *GetNetworkCameraSchedulesParams {
	return &GetNetworkCameraSchedulesParams{
		timeout: timeout,
	}
}

// NewGetNetworkCameraSchedulesParamsWithContext creates a new GetNetworkCameraSchedulesParams object
// with the ability to set a context for a request.
func NewGetNetworkCameraSchedulesParamsWithContext(ctx context.Context) *GetNetworkCameraSchedulesParams {
	return &GetNetworkCameraSchedulesParams{
		Context: ctx,
	}
}

// NewGetNetworkCameraSchedulesParamsWithHTTPClient creates a new GetNetworkCameraSchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkCameraSchedulesParamsWithHTTPClient(client *http.Client) *GetNetworkCameraSchedulesParams {
	return &GetNetworkCameraSchedulesParams{
		HTTPClient: client,
	}
}

/* GetNetworkCameraSchedulesParams contains all the parameters to send to the API endpoint
   for the get network camera schedules operation.

   Typically these are written to a http.Request.
*/
type GetNetworkCameraSchedulesParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network camera schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCameraSchedulesParams) WithDefaults() *GetNetworkCameraSchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network camera schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCameraSchedulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network camera schedules params
func (o *GetNetworkCameraSchedulesParams) WithTimeout(timeout time.Duration) *GetNetworkCameraSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network camera schedules params
func (o *GetNetworkCameraSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network camera schedules params
func (o *GetNetworkCameraSchedulesParams) WithContext(ctx context.Context) *GetNetworkCameraSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network camera schedules params
func (o *GetNetworkCameraSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network camera schedules params
func (o *GetNetworkCameraSchedulesParams) WithHTTPClient(client *http.Client) *GetNetworkCameraSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network camera schedules params
func (o *GetNetworkCameraSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network camera schedules params
func (o *GetNetworkCameraSchedulesParams) WithNetworkID(networkID string) *GetNetworkCameraSchedulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network camera schedules params
func (o *GetNetworkCameraSchedulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkCameraSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
