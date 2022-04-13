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

// NewUpdateNetworkMerakiAuthUserParams creates a new UpdateNetworkMerakiAuthUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkMerakiAuthUserParams() *UpdateNetworkMerakiAuthUserParams {
	return &UpdateNetworkMerakiAuthUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkMerakiAuthUserParamsWithTimeout creates a new UpdateNetworkMerakiAuthUserParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkMerakiAuthUserParamsWithTimeout(timeout time.Duration) *UpdateNetworkMerakiAuthUserParams {
	return &UpdateNetworkMerakiAuthUserParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkMerakiAuthUserParamsWithContext creates a new UpdateNetworkMerakiAuthUserParams object
// with the ability to set a context for a request.
func NewUpdateNetworkMerakiAuthUserParamsWithContext(ctx context.Context) *UpdateNetworkMerakiAuthUserParams {
	return &UpdateNetworkMerakiAuthUserParams{
		Context: ctx,
	}
}

// NewUpdateNetworkMerakiAuthUserParamsWithHTTPClient creates a new UpdateNetworkMerakiAuthUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkMerakiAuthUserParamsWithHTTPClient(client *http.Client) *UpdateNetworkMerakiAuthUserParams {
	return &UpdateNetworkMerakiAuthUserParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkMerakiAuthUserParams contains all the parameters to send to the API endpoint
   for the update network meraki auth user operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkMerakiAuthUserParams struct {

	// MerakiAuthUserID.
	MerakiAuthUserID string

	// NetworkID.
	NetworkID string

	// UpdateNetworkMerakiAuthUser.
	UpdateNetworkMerakiAuthUser UpdateNetworkMerakiAuthUserBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network meraki auth user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkMerakiAuthUserParams) WithDefaults() *UpdateNetworkMerakiAuthUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network meraki auth user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkMerakiAuthUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) WithTimeout(timeout time.Duration) *UpdateNetworkMerakiAuthUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) WithContext(ctx context.Context) *UpdateNetworkMerakiAuthUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) WithHTTPClient(client *http.Client) *UpdateNetworkMerakiAuthUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMerakiAuthUserID adds the merakiAuthUserID to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) WithMerakiAuthUserID(merakiAuthUserID string) *UpdateNetworkMerakiAuthUserParams {
	o.SetMerakiAuthUserID(merakiAuthUserID)
	return o
}

// SetMerakiAuthUserID adds the merakiAuthUserId to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) SetMerakiAuthUserID(merakiAuthUserID string) {
	o.MerakiAuthUserID = merakiAuthUserID
}

// WithNetworkID adds the networkID to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) WithNetworkID(networkID string) *UpdateNetworkMerakiAuthUserParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkMerakiAuthUser adds the updateNetworkMerakiAuthUser to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) WithUpdateNetworkMerakiAuthUser(updateNetworkMerakiAuthUser UpdateNetworkMerakiAuthUserBody) *UpdateNetworkMerakiAuthUserParams {
	o.SetUpdateNetworkMerakiAuthUser(updateNetworkMerakiAuthUser)
	return o
}

// SetUpdateNetworkMerakiAuthUser adds the updateNetworkMerakiAuthUser to the update network meraki auth user params
func (o *UpdateNetworkMerakiAuthUserParams) SetUpdateNetworkMerakiAuthUser(updateNetworkMerakiAuthUser UpdateNetworkMerakiAuthUserBody) {
	o.UpdateNetworkMerakiAuthUser = updateNetworkMerakiAuthUser
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkMerakiAuthUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param merakiAuthUserId
	if err := r.SetPathParam("merakiAuthUserId", o.MerakiAuthUserID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkMerakiAuthUser); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
