// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewUnbindNetworkParams creates a new UnbindNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnbindNetworkParams() *UnbindNetworkParams {
	return &UnbindNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnbindNetworkParamsWithTimeout creates a new UnbindNetworkParams object
// with the ability to set a timeout on a request.
func NewUnbindNetworkParamsWithTimeout(timeout time.Duration) *UnbindNetworkParams {
	return &UnbindNetworkParams{
		timeout: timeout,
	}
}

// NewUnbindNetworkParamsWithContext creates a new UnbindNetworkParams object
// with the ability to set a context for a request.
func NewUnbindNetworkParamsWithContext(ctx context.Context) *UnbindNetworkParams {
	return &UnbindNetworkParams{
		Context: ctx,
	}
}

// NewUnbindNetworkParamsWithHTTPClient creates a new UnbindNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnbindNetworkParamsWithHTTPClient(client *http.Client) *UnbindNetworkParams {
	return &UnbindNetworkParams{
		HTTPClient: client,
	}
}

/* UnbindNetworkParams contains all the parameters to send to the API endpoint
   for the unbind network operation.

   Typically these are written to a http.Request.
*/
type UnbindNetworkParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unbind network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnbindNetworkParams) WithDefaults() *UnbindNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unbind network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnbindNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unbind network params
func (o *UnbindNetworkParams) WithTimeout(timeout time.Duration) *UnbindNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unbind network params
func (o *UnbindNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unbind network params
func (o *UnbindNetworkParams) WithContext(ctx context.Context) *UnbindNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unbind network params
func (o *UnbindNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unbind network params
func (o *UnbindNetworkParams) WithHTTPClient(client *http.Client) *UnbindNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unbind network params
func (o *UnbindNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the unbind network params
func (o *UnbindNetworkParams) WithNetworkID(networkID string) *UnbindNetworkParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the unbind network params
func (o *UnbindNetworkParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *UnbindNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
