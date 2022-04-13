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

// NewUpdateNetworkSwitchStormControlParams creates a new UpdateNetworkSwitchStormControlParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkSwitchStormControlParams() *UpdateNetworkSwitchStormControlParams {
	return &UpdateNetworkSwitchStormControlParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchStormControlParamsWithTimeout creates a new UpdateNetworkSwitchStormControlParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkSwitchStormControlParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchStormControlParams {
	return &UpdateNetworkSwitchStormControlParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchStormControlParamsWithContext creates a new UpdateNetworkSwitchStormControlParams object
// with the ability to set a context for a request.
func NewUpdateNetworkSwitchStormControlParamsWithContext(ctx context.Context) *UpdateNetworkSwitchStormControlParams {
	return &UpdateNetworkSwitchStormControlParams{
		Context: ctx,
	}
}

// NewUpdateNetworkSwitchStormControlParamsWithHTTPClient creates a new UpdateNetworkSwitchStormControlParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkSwitchStormControlParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchStormControlParams {
	return &UpdateNetworkSwitchStormControlParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkSwitchStormControlParams contains all the parameters to send to the API endpoint
   for the update network switch storm control operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkSwitchStormControlParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkSwitchStormControl.
	UpdateNetworkSwitchStormControl UpdateNetworkSwitchStormControlBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network switch storm control params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchStormControlParams) WithDefaults() *UpdateNetworkSwitchStormControlParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network switch storm control params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchStormControlParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchStormControlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) WithContext(ctx context.Context) *UpdateNetworkSwitchStormControlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchStormControlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) WithNetworkID(networkID string) *UpdateNetworkSwitchStormControlParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkSwitchStormControl adds the updateNetworkSwitchStormControl to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) WithUpdateNetworkSwitchStormControl(updateNetworkSwitchStormControl UpdateNetworkSwitchStormControlBody) *UpdateNetworkSwitchStormControlParams {
	o.SetUpdateNetworkSwitchStormControl(updateNetworkSwitchStormControl)
	return o
}

// SetUpdateNetworkSwitchStormControl adds the updateNetworkSwitchStormControl to the update network switch storm control params
func (o *UpdateNetworkSwitchStormControlParams) SetUpdateNetworkSwitchStormControl(updateNetworkSwitchStormControl UpdateNetworkSwitchStormControlBody) {
	o.UpdateNetworkSwitchStormControl = updateNetworkSwitchStormControl
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchStormControlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkSwitchStormControl); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
