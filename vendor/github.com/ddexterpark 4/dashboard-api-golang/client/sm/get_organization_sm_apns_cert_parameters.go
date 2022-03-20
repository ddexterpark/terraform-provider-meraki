// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// NewGetOrganizationSmApnsCertParams creates a new GetOrganizationSmApnsCertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationSmApnsCertParams() *GetOrganizationSmApnsCertParams {
	return &GetOrganizationSmApnsCertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationSmApnsCertParamsWithTimeout creates a new GetOrganizationSmApnsCertParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationSmApnsCertParamsWithTimeout(timeout time.Duration) *GetOrganizationSmApnsCertParams {
	return &GetOrganizationSmApnsCertParams{
		timeout: timeout,
	}
}

// NewGetOrganizationSmApnsCertParamsWithContext creates a new GetOrganizationSmApnsCertParams object
// with the ability to set a context for a request.
func NewGetOrganizationSmApnsCertParamsWithContext(ctx context.Context) *GetOrganizationSmApnsCertParams {
	return &GetOrganizationSmApnsCertParams{
		Context: ctx,
	}
}

// NewGetOrganizationSmApnsCertParamsWithHTTPClient creates a new GetOrganizationSmApnsCertParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationSmApnsCertParamsWithHTTPClient(client *http.Client) *GetOrganizationSmApnsCertParams {
	return &GetOrganizationSmApnsCertParams{
		HTTPClient: client,
	}
}

/* GetOrganizationSmApnsCertParams contains all the parameters to send to the API endpoint
   for the get organization sm apns cert operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationSmApnsCertParams struct {

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization sm apns cert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSmApnsCertParams) WithDefaults() *GetOrganizationSmApnsCertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization sm apns cert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSmApnsCertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization sm apns cert params
func (o *GetOrganizationSmApnsCertParams) WithTimeout(timeout time.Duration) *GetOrganizationSmApnsCertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization sm apns cert params
func (o *GetOrganizationSmApnsCertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization sm apns cert params
func (o *GetOrganizationSmApnsCertParams) WithContext(ctx context.Context) *GetOrganizationSmApnsCertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization sm apns cert params
func (o *GetOrganizationSmApnsCertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization sm apns cert params
func (o *GetOrganizationSmApnsCertParams) WithHTTPClient(client *http.Client) *GetOrganizationSmApnsCertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization sm apns cert params
func (o *GetOrganizationSmApnsCertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization sm apns cert params
func (o *GetOrganizationSmApnsCertParams) WithOrganizationID(organizationID string) *GetOrganizationSmApnsCertParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization sm apns cert params
func (o *GetOrganizationSmApnsCertParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationSmApnsCertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}