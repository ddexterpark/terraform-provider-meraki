// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// NewCreateNetworkApplianceStaticRouteParams creates a new CreateNetworkApplianceStaticRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkApplianceStaticRouteParams() *CreateNetworkApplianceStaticRouteParams {
	return &CreateNetworkApplianceStaticRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkApplianceStaticRouteParamsWithTimeout creates a new CreateNetworkApplianceStaticRouteParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkApplianceStaticRouteParamsWithTimeout(timeout time.Duration) *CreateNetworkApplianceStaticRouteParams {
	return &CreateNetworkApplianceStaticRouteParams{
		timeout: timeout,
	}
}

// NewCreateNetworkApplianceStaticRouteParamsWithContext creates a new CreateNetworkApplianceStaticRouteParams object
// with the ability to set a context for a request.
func NewCreateNetworkApplianceStaticRouteParamsWithContext(ctx context.Context) *CreateNetworkApplianceStaticRouteParams {
	return &CreateNetworkApplianceStaticRouteParams{
		Context: ctx,
	}
}

// NewCreateNetworkApplianceStaticRouteParamsWithHTTPClient creates a new CreateNetworkApplianceStaticRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkApplianceStaticRouteParamsWithHTTPClient(client *http.Client) *CreateNetworkApplianceStaticRouteParams {
	return &CreateNetworkApplianceStaticRouteParams{
		HTTPClient: client,
	}
}

/* CreateNetworkApplianceStaticRouteParams contains all the parameters to send to the API endpoint
   for the create network appliance static route operation.

   Typically these are written to a http.Request.
*/
type CreateNetworkApplianceStaticRouteParams struct {

	// CreateNetworkApplianceStaticRoute.
	CreateNetworkApplianceStaticRoute CreateNetworkApplianceStaticRouteBody

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network appliance static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkApplianceStaticRouteParams) WithDefaults() *CreateNetworkApplianceStaticRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network appliance static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkApplianceStaticRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) WithTimeout(timeout time.Duration) *CreateNetworkApplianceStaticRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) WithContext(ctx context.Context) *CreateNetworkApplianceStaticRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) WithHTTPClient(client *http.Client) *CreateNetworkApplianceStaticRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkApplianceStaticRoute adds the createNetworkApplianceStaticRoute to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) WithCreateNetworkApplianceStaticRoute(createNetworkApplianceStaticRoute CreateNetworkApplianceStaticRouteBody) *CreateNetworkApplianceStaticRouteParams {
	o.SetCreateNetworkApplianceStaticRoute(createNetworkApplianceStaticRoute)
	return o
}

// SetCreateNetworkApplianceStaticRoute adds the createNetworkApplianceStaticRoute to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) SetCreateNetworkApplianceStaticRoute(createNetworkApplianceStaticRoute CreateNetworkApplianceStaticRouteBody) {
	o.CreateNetworkApplianceStaticRoute = createNetworkApplianceStaticRoute
}

// WithNetworkID adds the networkID to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) WithNetworkID(networkID string) *CreateNetworkApplianceStaticRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network appliance static route params
func (o *CreateNetworkApplianceStaticRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkApplianceStaticRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkApplianceStaticRoute); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
