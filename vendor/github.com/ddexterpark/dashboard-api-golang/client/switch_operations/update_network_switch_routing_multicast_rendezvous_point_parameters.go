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

// NewUpdateNetworkSwitchRoutingMulticastRendezvousPointParams creates a new UpdateNetworkSwitchRoutingMulticastRendezvousPointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkSwitchRoutingMulticastRendezvousPointParams() *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &UpdateNetworkSwitchRoutingMulticastRendezvousPointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchRoutingMulticastRendezvousPointParamsWithTimeout creates a new UpdateNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkSwitchRoutingMulticastRendezvousPointParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &UpdateNetworkSwitchRoutingMulticastRendezvousPointParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchRoutingMulticastRendezvousPointParamsWithContext creates a new UpdateNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a context for a request.
func NewUpdateNetworkSwitchRoutingMulticastRendezvousPointParamsWithContext(ctx context.Context) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &UpdateNetworkSwitchRoutingMulticastRendezvousPointParams{
		Context: ctx,
	}
}

// NewUpdateNetworkSwitchRoutingMulticastRendezvousPointParamsWithHTTPClient creates a new UpdateNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkSwitchRoutingMulticastRendezvousPointParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &UpdateNetworkSwitchRoutingMulticastRendezvousPointParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkSwitchRoutingMulticastRendezvousPointParams contains all the parameters to send to the API endpoint
   for the update network switch routing multicast rendezvous point operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkSwitchRoutingMulticastRendezvousPointParams struct {

	// NetworkID.
	NetworkID string

	// RendezvousPointID.
	RendezvousPointID string

	// UpdateNetworkSwitchRoutingMulticastRendezvousPoint.
	UpdateNetworkSwitchRoutingMulticastRendezvousPoint UpdateNetworkSwitchRoutingMulticastRendezvousPointBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network switch routing multicast rendezvous point params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) WithDefaults() *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network switch routing multicast rendezvous point params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) WithContext(ctx context.Context) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) WithNetworkID(networkID string) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRendezvousPointID adds the rendezvousPointID to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) WithRendezvousPointID(rendezvousPointID string) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetRendezvousPointID(rendezvousPointID)
	return o
}

// SetRendezvousPointID adds the rendezvousPointId to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) SetRendezvousPointID(rendezvousPointID string) {
	o.RendezvousPointID = rendezvousPointID
}

// WithUpdateNetworkSwitchRoutingMulticastRendezvousPoint adds the updateNetworkSwitchRoutingMulticastRendezvousPoint to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) WithUpdateNetworkSwitchRoutingMulticastRendezvousPoint(updateNetworkSwitchRoutingMulticastRendezvousPoint UpdateNetworkSwitchRoutingMulticastRendezvousPointBody) *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetUpdateNetworkSwitchRoutingMulticastRendezvousPoint(updateNetworkSwitchRoutingMulticastRendezvousPoint)
	return o
}

// SetUpdateNetworkSwitchRoutingMulticastRendezvousPoint adds the updateNetworkSwitchRoutingMulticastRendezvousPoint to the update network switch routing multicast rendezvous point params
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) SetUpdateNetworkSwitchRoutingMulticastRendezvousPoint(updateNetworkSwitchRoutingMulticastRendezvousPoint UpdateNetworkSwitchRoutingMulticastRendezvousPointBody) {
	o.UpdateNetworkSwitchRoutingMulticastRendezvousPoint = updateNetworkSwitchRoutingMulticastRendezvousPoint
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param rendezvousPointId
	if err := r.SetPathParam("rendezvousPointId", o.RendezvousPointID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkSwitchRoutingMulticastRendezvousPoint); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}