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

// NewUpdateNetworkSwitchStpParams creates a new UpdateNetworkSwitchStpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkSwitchStpParams() *UpdateNetworkSwitchStpParams {
	return &UpdateNetworkSwitchStpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchStpParamsWithTimeout creates a new UpdateNetworkSwitchStpParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkSwitchStpParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchStpParams {
	return &UpdateNetworkSwitchStpParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchStpParamsWithContext creates a new UpdateNetworkSwitchStpParams object
// with the ability to set a context for a request.
func NewUpdateNetworkSwitchStpParamsWithContext(ctx context.Context) *UpdateNetworkSwitchStpParams {
	return &UpdateNetworkSwitchStpParams{
		Context: ctx,
	}
}

// NewUpdateNetworkSwitchStpParamsWithHTTPClient creates a new UpdateNetworkSwitchStpParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkSwitchStpParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchStpParams {
	return &UpdateNetworkSwitchStpParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkSwitchStpParams contains all the parameters to send to the API endpoint
   for the update network switch stp operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkSwitchStpParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkSwitchStp.
	UpdateNetworkSwitchStp UpdateNetworkSwitchStpBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network switch stp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchStpParams) WithDefaults() *UpdateNetworkSwitchStpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network switch stp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchStpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchStpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) WithContext(ctx context.Context) *UpdateNetworkSwitchStpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchStpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) WithNetworkID(networkID string) *UpdateNetworkSwitchStpParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkSwitchStp adds the updateNetworkSwitchStp to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) WithUpdateNetworkSwitchStp(updateNetworkSwitchStp UpdateNetworkSwitchStpBody) *UpdateNetworkSwitchStpParams {
	o.SetUpdateNetworkSwitchStp(updateNetworkSwitchStp)
	return o
}

// SetUpdateNetworkSwitchStp adds the updateNetworkSwitchStp to the update network switch stp params
func (o *UpdateNetworkSwitchStpParams) SetUpdateNetworkSwitchStp(updateNetworkSwitchStp UpdateNetworkSwitchStpBody) {
	o.UpdateNetworkSwitchStp = updateNetworkSwitchStp
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchStpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkSwitchStp); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}