// Code generated by go-swagger; DO NOT EDIT.

package insight

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

// NewDeleteOrganizationInsightMonitoredMediaServerParams creates a new DeleteOrganizationInsightMonitoredMediaServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationInsightMonitoredMediaServerParams() *DeleteOrganizationInsightMonitoredMediaServerParams {
	return &DeleteOrganizationInsightMonitoredMediaServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationInsightMonitoredMediaServerParamsWithTimeout creates a new DeleteOrganizationInsightMonitoredMediaServerParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationInsightMonitoredMediaServerParamsWithTimeout(timeout time.Duration) *DeleteOrganizationInsightMonitoredMediaServerParams {
	return &DeleteOrganizationInsightMonitoredMediaServerParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationInsightMonitoredMediaServerParamsWithContext creates a new DeleteOrganizationInsightMonitoredMediaServerParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationInsightMonitoredMediaServerParamsWithContext(ctx context.Context) *DeleteOrganizationInsightMonitoredMediaServerParams {
	return &DeleteOrganizationInsightMonitoredMediaServerParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationInsightMonitoredMediaServerParamsWithHTTPClient creates a new DeleteOrganizationInsightMonitoredMediaServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationInsightMonitoredMediaServerParamsWithHTTPClient(client *http.Client) *DeleteOrganizationInsightMonitoredMediaServerParams {
	return &DeleteOrganizationInsightMonitoredMediaServerParams{
		HTTPClient: client,
	}
}

/* DeleteOrganizationInsightMonitoredMediaServerParams contains all the parameters to send to the API endpoint
   for the delete organization insight monitored media server operation.

   Typically these are written to a http.Request.
*/
type DeleteOrganizationInsightMonitoredMediaServerParams struct {

	// MonitoredMediaServerID.
	MonitoredMediaServerID string

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization insight monitored media server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) WithDefaults() *DeleteOrganizationInsightMonitoredMediaServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization insight monitored media server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) WithTimeout(timeout time.Duration) *DeleteOrganizationInsightMonitoredMediaServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) WithContext(ctx context.Context) *DeleteOrganizationInsightMonitoredMediaServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) WithHTTPClient(client *http.Client) *DeleteOrganizationInsightMonitoredMediaServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitoredMediaServerID adds the monitoredMediaServerID to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) WithMonitoredMediaServerID(monitoredMediaServerID string) *DeleteOrganizationInsightMonitoredMediaServerParams {
	o.SetMonitoredMediaServerID(monitoredMediaServerID)
	return o
}

// SetMonitoredMediaServerID adds the monitoredMediaServerId to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) SetMonitoredMediaServerID(monitoredMediaServerID string) {
	o.MonitoredMediaServerID = monitoredMediaServerID
}

// WithOrganizationID adds the organizationID to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) WithOrganizationID(organizationID string) *DeleteOrganizationInsightMonitoredMediaServerParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete organization insight monitored media server params
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationInsightMonitoredMediaServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param monitoredMediaServerId
	if err := r.SetPathParam("monitoredMediaServerId", o.MonitoredMediaServerID); err != nil {
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
