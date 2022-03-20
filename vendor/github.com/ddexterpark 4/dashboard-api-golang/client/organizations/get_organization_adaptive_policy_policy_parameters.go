// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrganizationAdaptivePolicyPolicyParams creates a new GetOrganizationAdaptivePolicyPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationAdaptivePolicyPolicyParams() *GetOrganizationAdaptivePolicyPolicyParams {
	return &GetOrganizationAdaptivePolicyPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationAdaptivePolicyPolicyParamsWithTimeout creates a new GetOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationAdaptivePolicyPolicyParamsWithTimeout(timeout time.Duration) *GetOrganizationAdaptivePolicyPolicyParams {
	return &GetOrganizationAdaptivePolicyPolicyParams{
		timeout: timeout,
	}
}

// NewGetOrganizationAdaptivePolicyPolicyParamsWithContext creates a new GetOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a context for a request.
func NewGetOrganizationAdaptivePolicyPolicyParamsWithContext(ctx context.Context) *GetOrganizationAdaptivePolicyPolicyParams {
	return &GetOrganizationAdaptivePolicyPolicyParams{
		Context: ctx,
	}
}

// NewGetOrganizationAdaptivePolicyPolicyParamsWithHTTPClient creates a new GetOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationAdaptivePolicyPolicyParamsWithHTTPClient(client *http.Client) *GetOrganizationAdaptivePolicyPolicyParams {
	return &GetOrganizationAdaptivePolicyPolicyParams{
		HTTPClient: client,
	}
}

/* GetOrganizationAdaptivePolicyPolicyParams contains all the parameters to send to the API endpoint
   for the get organization adaptive policy policy operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationAdaptivePolicyPolicyParams struct {

	// AdaptivePolicyID.
	AdaptivePolicyID string

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization adaptive policy policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAdaptivePolicyPolicyParams) WithDefaults() *GetOrganizationAdaptivePolicyPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization adaptive policy policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAdaptivePolicyPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) WithTimeout(timeout time.Duration) *GetOrganizationAdaptivePolicyPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) WithContext(ctx context.Context) *GetOrganizationAdaptivePolicyPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) WithHTTPClient(client *http.Client) *GetOrganizationAdaptivePolicyPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdaptivePolicyID adds the adaptivePolicyID to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) WithAdaptivePolicyID(adaptivePolicyID string) *GetOrganizationAdaptivePolicyPolicyParams {
	o.SetAdaptivePolicyID(adaptivePolicyID)
	return o
}

// SetAdaptivePolicyID adds the adaptivePolicyId to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) SetAdaptivePolicyID(adaptivePolicyID string) {
	o.AdaptivePolicyID = adaptivePolicyID
}

// WithOrganizationID adds the organizationID to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) WithOrganizationID(organizationID string) *GetOrganizationAdaptivePolicyPolicyParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization adaptive policy policy params
func (o *GetOrganizationAdaptivePolicyPolicyParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationAdaptivePolicyPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adaptivePolicyId
	if err := r.SetPathParam("adaptivePolicyId", o.AdaptivePolicyID); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}