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

// NewUpdateNetworkGroupPolicyParams creates a new UpdateNetworkGroupPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkGroupPolicyParams() *UpdateNetworkGroupPolicyParams {
	return &UpdateNetworkGroupPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkGroupPolicyParamsWithTimeout creates a new UpdateNetworkGroupPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkGroupPolicyParamsWithTimeout(timeout time.Duration) *UpdateNetworkGroupPolicyParams {
	return &UpdateNetworkGroupPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkGroupPolicyParamsWithContext creates a new UpdateNetworkGroupPolicyParams object
// with the ability to set a context for a request.
func NewUpdateNetworkGroupPolicyParamsWithContext(ctx context.Context) *UpdateNetworkGroupPolicyParams {
	return &UpdateNetworkGroupPolicyParams{
		Context: ctx,
	}
}

// NewUpdateNetworkGroupPolicyParamsWithHTTPClient creates a new UpdateNetworkGroupPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkGroupPolicyParamsWithHTTPClient(client *http.Client) *UpdateNetworkGroupPolicyParams {
	return &UpdateNetworkGroupPolicyParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkGroupPolicyParams contains all the parameters to send to the API endpoint
   for the update network group policy operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkGroupPolicyParams struct {

	// GroupPolicyID.
	GroupPolicyID string

	// NetworkID.
	NetworkID string

	// UpdateNetworkGroupPolicy.
	UpdateNetworkGroupPolicy UpdateNetworkGroupPolicyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network group policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkGroupPolicyParams) WithDefaults() *UpdateNetworkGroupPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network group policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkGroupPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) WithTimeout(timeout time.Duration) *UpdateNetworkGroupPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) WithContext(ctx context.Context) *UpdateNetworkGroupPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) WithHTTPClient(client *http.Client) *UpdateNetworkGroupPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupPolicyID adds the groupPolicyID to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) WithGroupPolicyID(groupPolicyID string) *UpdateNetworkGroupPolicyParams {
	o.SetGroupPolicyID(groupPolicyID)
	return o
}

// SetGroupPolicyID adds the groupPolicyId to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) SetGroupPolicyID(groupPolicyID string) {
	o.GroupPolicyID = groupPolicyID
}

// WithNetworkID adds the networkID to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) WithNetworkID(networkID string) *UpdateNetworkGroupPolicyParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkGroupPolicy adds the updateNetworkGroupPolicy to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) WithUpdateNetworkGroupPolicy(updateNetworkGroupPolicy UpdateNetworkGroupPolicyBody) *UpdateNetworkGroupPolicyParams {
	o.SetUpdateNetworkGroupPolicy(updateNetworkGroupPolicy)
	return o
}

// SetUpdateNetworkGroupPolicy adds the updateNetworkGroupPolicy to the update network group policy params
func (o *UpdateNetworkGroupPolicyParams) SetUpdateNetworkGroupPolicy(updateNetworkGroupPolicy UpdateNetworkGroupPolicyBody) {
	o.UpdateNetworkGroupPolicy = updateNetworkGroupPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkGroupPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupPolicyId
	if err := r.SetPathParam("groupPolicyId", o.GroupPolicyID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkGroupPolicy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}