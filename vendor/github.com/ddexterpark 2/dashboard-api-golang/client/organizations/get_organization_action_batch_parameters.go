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

// NewGetOrganizationActionBatchParams creates a new GetOrganizationActionBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationActionBatchParams() *GetOrganizationActionBatchParams {
	return &GetOrganizationActionBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationActionBatchParamsWithTimeout creates a new GetOrganizationActionBatchParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationActionBatchParamsWithTimeout(timeout time.Duration) *GetOrganizationActionBatchParams {
	return &GetOrganizationActionBatchParams{
		timeout: timeout,
	}
}

// NewGetOrganizationActionBatchParamsWithContext creates a new GetOrganizationActionBatchParams object
// with the ability to set a context for a request.
func NewGetOrganizationActionBatchParamsWithContext(ctx context.Context) *GetOrganizationActionBatchParams {
	return &GetOrganizationActionBatchParams{
		Context: ctx,
	}
}

// NewGetOrganizationActionBatchParamsWithHTTPClient creates a new GetOrganizationActionBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationActionBatchParamsWithHTTPClient(client *http.Client) *GetOrganizationActionBatchParams {
	return &GetOrganizationActionBatchParams{
		HTTPClient: client,
	}
}

/* GetOrganizationActionBatchParams contains all the parameters to send to the API endpoint
   for the get organization action batch operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationActionBatchParams struct {

	// ActionBatchID.
	ActionBatchID string

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization action batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationActionBatchParams) WithDefaults() *GetOrganizationActionBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization action batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationActionBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization action batch params
func (o *GetOrganizationActionBatchParams) WithTimeout(timeout time.Duration) *GetOrganizationActionBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization action batch params
func (o *GetOrganizationActionBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization action batch params
func (o *GetOrganizationActionBatchParams) WithContext(ctx context.Context) *GetOrganizationActionBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization action batch params
func (o *GetOrganizationActionBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization action batch params
func (o *GetOrganizationActionBatchParams) WithHTTPClient(client *http.Client) *GetOrganizationActionBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization action batch params
func (o *GetOrganizationActionBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionBatchID adds the actionBatchID to the get organization action batch params
func (o *GetOrganizationActionBatchParams) WithActionBatchID(actionBatchID string) *GetOrganizationActionBatchParams {
	o.SetActionBatchID(actionBatchID)
	return o
}

// SetActionBatchID adds the actionBatchId to the get organization action batch params
func (o *GetOrganizationActionBatchParams) SetActionBatchID(actionBatchID string) {
	o.ActionBatchID = actionBatchID
}

// WithOrganizationID adds the organizationID to the get organization action batch params
func (o *GetOrganizationActionBatchParams) WithOrganizationID(organizationID string) *GetOrganizationActionBatchParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization action batch params
func (o *GetOrganizationActionBatchParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationActionBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionBatchId
	if err := r.SetPathParam("actionBatchId", o.ActionBatchID); err != nil {
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