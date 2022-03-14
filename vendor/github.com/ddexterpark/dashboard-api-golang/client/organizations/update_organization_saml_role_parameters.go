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

// NewUpdateOrganizationSamlRoleParams creates a new UpdateOrganizationSamlRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationSamlRoleParams() *UpdateOrganizationSamlRoleParams {
	return &UpdateOrganizationSamlRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationSamlRoleParamsWithTimeout creates a new UpdateOrganizationSamlRoleParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationSamlRoleParamsWithTimeout(timeout time.Duration) *UpdateOrganizationSamlRoleParams {
	return &UpdateOrganizationSamlRoleParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationSamlRoleParamsWithContext creates a new UpdateOrganizationSamlRoleParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationSamlRoleParamsWithContext(ctx context.Context) *UpdateOrganizationSamlRoleParams {
	return &UpdateOrganizationSamlRoleParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationSamlRoleParamsWithHTTPClient creates a new UpdateOrganizationSamlRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationSamlRoleParamsWithHTTPClient(client *http.Client) *UpdateOrganizationSamlRoleParams {
	return &UpdateOrganizationSamlRoleParams{
		HTTPClient: client,
	}
}

/* UpdateOrganizationSamlRoleParams contains all the parameters to send to the API endpoint
   for the update organization saml role operation.

   Typically these are written to a http.Request.
*/
type UpdateOrganizationSamlRoleParams struct {

	// OrganizationID.
	OrganizationID string

	// SamlRoleID.
	SamlRoleID string

	// UpdateOrganizationSamlRole.
	UpdateOrganizationSamlRole UpdateOrganizationSamlRoleBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization saml role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationSamlRoleParams) WithDefaults() *UpdateOrganizationSamlRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization saml role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationSamlRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) WithTimeout(timeout time.Duration) *UpdateOrganizationSamlRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) WithContext(ctx context.Context) *UpdateOrganizationSamlRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) WithHTTPClient(client *http.Client) *UpdateOrganizationSamlRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) WithOrganizationID(organizationID string) *UpdateOrganizationSamlRoleParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithSamlRoleID adds the samlRoleID to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) WithSamlRoleID(samlRoleID string) *UpdateOrganizationSamlRoleParams {
	o.SetSamlRoleID(samlRoleID)
	return o
}

// SetSamlRoleID adds the samlRoleId to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) SetSamlRoleID(samlRoleID string) {
	o.SamlRoleID = samlRoleID
}

// WithUpdateOrganizationSamlRole adds the updateOrganizationSamlRole to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) WithUpdateOrganizationSamlRole(updateOrganizationSamlRole UpdateOrganizationSamlRoleBody) *UpdateOrganizationSamlRoleParams {
	o.SetUpdateOrganizationSamlRole(updateOrganizationSamlRole)
	return o
}

// SetUpdateOrganizationSamlRole adds the updateOrganizationSamlRole to the update organization saml role params
func (o *UpdateOrganizationSamlRoleParams) SetUpdateOrganizationSamlRole(updateOrganizationSamlRole UpdateOrganizationSamlRoleBody) {
	o.UpdateOrganizationSamlRole = updateOrganizationSamlRole
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationSamlRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	// path param samlRoleId
	if err := r.SetPathParam("samlRoleId", o.SamlRoleID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationSamlRole); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
