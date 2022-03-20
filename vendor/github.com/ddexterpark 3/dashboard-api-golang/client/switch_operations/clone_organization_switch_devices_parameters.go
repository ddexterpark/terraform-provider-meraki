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

// NewCloneOrganizationSwitchDevicesParams creates a new CloneOrganizationSwitchDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloneOrganizationSwitchDevicesParams() *CloneOrganizationSwitchDevicesParams {
	return &CloneOrganizationSwitchDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloneOrganizationSwitchDevicesParamsWithTimeout creates a new CloneOrganizationSwitchDevicesParams object
// with the ability to set a timeout on a request.
func NewCloneOrganizationSwitchDevicesParamsWithTimeout(timeout time.Duration) *CloneOrganizationSwitchDevicesParams {
	return &CloneOrganizationSwitchDevicesParams{
		timeout: timeout,
	}
}

// NewCloneOrganizationSwitchDevicesParamsWithContext creates a new CloneOrganizationSwitchDevicesParams object
// with the ability to set a context for a request.
func NewCloneOrganizationSwitchDevicesParamsWithContext(ctx context.Context) *CloneOrganizationSwitchDevicesParams {
	return &CloneOrganizationSwitchDevicesParams{
		Context: ctx,
	}
}

// NewCloneOrganizationSwitchDevicesParamsWithHTTPClient creates a new CloneOrganizationSwitchDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloneOrganizationSwitchDevicesParamsWithHTTPClient(client *http.Client) *CloneOrganizationSwitchDevicesParams {
	return &CloneOrganizationSwitchDevicesParams{
		HTTPClient: client,
	}
}

/* CloneOrganizationSwitchDevicesParams contains all the parameters to send to the API endpoint
   for the clone organization switch devices operation.

   Typically these are written to a http.Request.
*/
type CloneOrganizationSwitchDevicesParams struct {

	// CloneOrganizationSwitchDevices.
	CloneOrganizationSwitchDevices CloneOrganizationSwitchDevicesBody

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clone organization switch devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneOrganizationSwitchDevicesParams) WithDefaults() *CloneOrganizationSwitchDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clone organization switch devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneOrganizationSwitchDevicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) WithTimeout(timeout time.Duration) *CloneOrganizationSwitchDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) WithContext(ctx context.Context) *CloneOrganizationSwitchDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) WithHTTPClient(client *http.Client) *CloneOrganizationSwitchDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloneOrganizationSwitchDevices adds the cloneOrganizationSwitchDevices to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) WithCloneOrganizationSwitchDevices(cloneOrganizationSwitchDevices CloneOrganizationSwitchDevicesBody) *CloneOrganizationSwitchDevicesParams {
	o.SetCloneOrganizationSwitchDevices(cloneOrganizationSwitchDevices)
	return o
}

// SetCloneOrganizationSwitchDevices adds the cloneOrganizationSwitchDevices to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) SetCloneOrganizationSwitchDevices(cloneOrganizationSwitchDevices CloneOrganizationSwitchDevicesBody) {
	o.CloneOrganizationSwitchDevices = cloneOrganizationSwitchDevices
}

// WithOrganizationID adds the organizationID to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) WithOrganizationID(organizationID string) *CloneOrganizationSwitchDevicesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the clone organization switch devices params
func (o *CloneOrganizationSwitchDevicesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CloneOrganizationSwitchDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CloneOrganizationSwitchDevices); err != nil {
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