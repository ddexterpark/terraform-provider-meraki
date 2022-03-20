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

// NewReleaseFromOrganizationInventoryParams creates a new ReleaseFromOrganizationInventoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReleaseFromOrganizationInventoryParams() *ReleaseFromOrganizationInventoryParams {
	return &ReleaseFromOrganizationInventoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReleaseFromOrganizationInventoryParamsWithTimeout creates a new ReleaseFromOrganizationInventoryParams object
// with the ability to set a timeout on a request.
func NewReleaseFromOrganizationInventoryParamsWithTimeout(timeout time.Duration) *ReleaseFromOrganizationInventoryParams {
	return &ReleaseFromOrganizationInventoryParams{
		timeout: timeout,
	}
}

// NewReleaseFromOrganizationInventoryParamsWithContext creates a new ReleaseFromOrganizationInventoryParams object
// with the ability to set a context for a request.
func NewReleaseFromOrganizationInventoryParamsWithContext(ctx context.Context) *ReleaseFromOrganizationInventoryParams {
	return &ReleaseFromOrganizationInventoryParams{
		Context: ctx,
	}
}

// NewReleaseFromOrganizationInventoryParamsWithHTTPClient creates a new ReleaseFromOrganizationInventoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewReleaseFromOrganizationInventoryParamsWithHTTPClient(client *http.Client) *ReleaseFromOrganizationInventoryParams {
	return &ReleaseFromOrganizationInventoryParams{
		HTTPClient: client,
	}
}

/* ReleaseFromOrganizationInventoryParams contains all the parameters to send to the API endpoint
   for the release from organization inventory operation.

   Typically these are written to a http.Request.
*/
type ReleaseFromOrganizationInventoryParams struct {

	// OrganizationID.
	OrganizationID string

	// ReleaseFromOrganizationInventory.
	ReleaseFromOrganizationInventory ReleaseFromOrganizationInventoryBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the release from organization inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseFromOrganizationInventoryParams) WithDefaults() *ReleaseFromOrganizationInventoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the release from organization inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseFromOrganizationInventoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) WithTimeout(timeout time.Duration) *ReleaseFromOrganizationInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) WithContext(ctx context.Context) *ReleaseFromOrganizationInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) WithHTTPClient(client *http.Client) *ReleaseFromOrganizationInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) WithOrganizationID(organizationID string) *ReleaseFromOrganizationInventoryParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithReleaseFromOrganizationInventory adds the releaseFromOrganizationInventory to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) WithReleaseFromOrganizationInventory(releaseFromOrganizationInventory ReleaseFromOrganizationInventoryBody) *ReleaseFromOrganizationInventoryParams {
	o.SetReleaseFromOrganizationInventory(releaseFromOrganizationInventory)
	return o
}

// SetReleaseFromOrganizationInventory adds the releaseFromOrganizationInventory to the release from organization inventory params
func (o *ReleaseFromOrganizationInventoryParams) SetReleaseFromOrganizationInventory(releaseFromOrganizationInventory ReleaseFromOrganizationInventoryBody) {
	o.ReleaseFromOrganizationInventory = releaseFromOrganizationInventory
}

// WriteToRequest writes these params to a swagger request
func (o *ReleaseFromOrganizationInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.ReleaseFromOrganizationInventory); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}