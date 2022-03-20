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

// NewGetOrganizationConfigTemplateSwitchProfilePortsParams creates a new GetOrganizationConfigTemplateSwitchProfilePortsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationConfigTemplateSwitchProfilePortsParams() *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	return &GetOrganizationConfigTemplateSwitchProfilePortsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationConfigTemplateSwitchProfilePortsParamsWithTimeout creates a new GetOrganizationConfigTemplateSwitchProfilePortsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationConfigTemplateSwitchProfilePortsParamsWithTimeout(timeout time.Duration) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	return &GetOrganizationConfigTemplateSwitchProfilePortsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationConfigTemplateSwitchProfilePortsParamsWithContext creates a new GetOrganizationConfigTemplateSwitchProfilePortsParams object
// with the ability to set a context for a request.
func NewGetOrganizationConfigTemplateSwitchProfilePortsParamsWithContext(ctx context.Context) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	return &GetOrganizationConfigTemplateSwitchProfilePortsParams{
		Context: ctx,
	}
}

// NewGetOrganizationConfigTemplateSwitchProfilePortsParamsWithHTTPClient creates a new GetOrganizationConfigTemplateSwitchProfilePortsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationConfigTemplateSwitchProfilePortsParamsWithHTTPClient(client *http.Client) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	return &GetOrganizationConfigTemplateSwitchProfilePortsParams{
		HTTPClient: client,
	}
}

/* GetOrganizationConfigTemplateSwitchProfilePortsParams contains all the parameters to send to the API endpoint
   for the get organization config template switch profile ports operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationConfigTemplateSwitchProfilePortsParams struct {

	// ConfigTemplateID.
	ConfigTemplateID string

	// OrganizationID.
	OrganizationID string

	// ProfileID.
	ProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization config template switch profile ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) WithDefaults() *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization config template switch profile ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) WithTimeout(timeout time.Duration) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) WithContext(ctx context.Context) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) WithHTTPClient(client *http.Client) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigTemplateID adds the configTemplateID to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) WithConfigTemplateID(configTemplateID string) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	o.SetConfigTemplateID(configTemplateID)
	return o
}

// SetConfigTemplateID adds the configTemplateId to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) SetConfigTemplateID(configTemplateID string) {
	o.ConfigTemplateID = configTemplateID
}

// WithOrganizationID adds the organizationID to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) WithOrganizationID(organizationID string) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProfileID adds the profileID to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) WithProfileID(profileID string) *GetOrganizationConfigTemplateSwitchProfilePortsParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the get organization config template switch profile ports params
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationConfigTemplateSwitchProfilePortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param configTemplateId
	if err := r.SetPathParam("configTemplateId", o.ConfigTemplateID); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	// path param profileId
	if err := r.SetPathParam("profileId", o.ProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}