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

// NewUpdateNetworkSwitchStackRoutingStaticRouteParams creates a new UpdateNetworkSwitchStackRoutingStaticRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkSwitchStackRoutingStaticRouteParams() *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	return &UpdateNetworkSwitchStackRoutingStaticRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchStackRoutingStaticRouteParamsWithTimeout creates a new UpdateNetworkSwitchStackRoutingStaticRouteParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkSwitchStackRoutingStaticRouteParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	return &UpdateNetworkSwitchStackRoutingStaticRouteParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchStackRoutingStaticRouteParamsWithContext creates a new UpdateNetworkSwitchStackRoutingStaticRouteParams object
// with the ability to set a context for a request.
func NewUpdateNetworkSwitchStackRoutingStaticRouteParamsWithContext(ctx context.Context) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	return &UpdateNetworkSwitchStackRoutingStaticRouteParams{
		Context: ctx,
	}
}

// NewUpdateNetworkSwitchStackRoutingStaticRouteParamsWithHTTPClient creates a new UpdateNetworkSwitchStackRoutingStaticRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkSwitchStackRoutingStaticRouteParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	return &UpdateNetworkSwitchStackRoutingStaticRouteParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkSwitchStackRoutingStaticRouteParams contains all the parameters to send to the API endpoint
   for the update network switch stack routing static route operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkSwitchStackRoutingStaticRouteParams struct {

	// NetworkID.
	NetworkID string

	// StaticRouteID.
	StaticRouteID string

	// SwitchStackID.
	SwitchStackID string

	// UpdateNetworkSwitchStackRoutingStaticRoute.
	UpdateNetworkSwitchStackRoutingStaticRoute UpdateNetworkSwitchStackRoutingStaticRouteBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network switch stack routing static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WithDefaults() *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network switch stack routing static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WithContext(ctx context.Context) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WithNetworkID(networkID string) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithStaticRouteID adds the staticRouteID to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WithStaticRouteID(staticRouteID string) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	o.SetStaticRouteID(staticRouteID)
	return o
}

// SetStaticRouteID adds the staticRouteId to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) SetStaticRouteID(staticRouteID string) {
	o.StaticRouteID = staticRouteID
}

// WithSwitchStackID adds the switchStackID to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WithSwitchStackID(switchStackID string) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	o.SetSwitchStackID(switchStackID)
	return o
}

// SetSwitchStackID adds the switchStackId to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) SetSwitchStackID(switchStackID string) {
	o.SwitchStackID = switchStackID
}

// WithUpdateNetworkSwitchStackRoutingStaticRoute adds the updateNetworkSwitchStackRoutingStaticRoute to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WithUpdateNetworkSwitchStackRoutingStaticRoute(updateNetworkSwitchStackRoutingStaticRoute UpdateNetworkSwitchStackRoutingStaticRouteBody) *UpdateNetworkSwitchStackRoutingStaticRouteParams {
	o.SetUpdateNetworkSwitchStackRoutingStaticRoute(updateNetworkSwitchStackRoutingStaticRoute)
	return o
}

// SetUpdateNetworkSwitchStackRoutingStaticRoute adds the updateNetworkSwitchStackRoutingStaticRoute to the update network switch stack routing static route params
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) SetUpdateNetworkSwitchStackRoutingStaticRoute(updateNetworkSwitchStackRoutingStaticRoute UpdateNetworkSwitchStackRoutingStaticRouteBody) {
	o.UpdateNetworkSwitchStackRoutingStaticRoute = updateNetworkSwitchStackRoutingStaticRoute
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchStackRoutingStaticRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param staticRouteId
	if err := r.SetPathParam("staticRouteId", o.StaticRouteID); err != nil {
		return err
	}

	// path param switchStackId
	if err := r.SetPathParam("switchStackId", o.SwitchStackID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkSwitchStackRoutingStaticRoute); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
