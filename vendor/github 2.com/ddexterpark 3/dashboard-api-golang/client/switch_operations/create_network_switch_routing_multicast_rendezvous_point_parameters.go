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

// NewCreateNetworkSwitchRoutingMulticastRendezvousPointParams creates a new CreateNetworkSwitchRoutingMulticastRendezvousPointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkSwitchRoutingMulticastRendezvousPointParams() *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &CreateNetworkSwitchRoutingMulticastRendezvousPointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkSwitchRoutingMulticastRendezvousPointParamsWithTimeout creates a new CreateNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkSwitchRoutingMulticastRendezvousPointParamsWithTimeout(timeout time.Duration) *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &CreateNetworkSwitchRoutingMulticastRendezvousPointParams{
		timeout: timeout,
	}
}

// NewCreateNetworkSwitchRoutingMulticastRendezvousPointParamsWithContext creates a new CreateNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a context for a request.
func NewCreateNetworkSwitchRoutingMulticastRendezvousPointParamsWithContext(ctx context.Context) *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &CreateNetworkSwitchRoutingMulticastRendezvousPointParams{
		Context: ctx,
	}
}

// NewCreateNetworkSwitchRoutingMulticastRendezvousPointParamsWithHTTPClient creates a new CreateNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkSwitchRoutingMulticastRendezvousPointParamsWithHTTPClient(client *http.Client) *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &CreateNetworkSwitchRoutingMulticastRendezvousPointParams{
		HTTPClient: client,
	}
}

/* CreateNetworkSwitchRoutingMulticastRendezvousPointParams contains all the parameters to send to the API endpoint
   for the create network switch routing multicast rendezvous point operation.

   Typically these are written to a http.Request.
*/
type CreateNetworkSwitchRoutingMulticastRendezvousPointParams struct {

	// CreateNetworkSwitchRoutingMulticastRendezvousPoint.
	CreateNetworkSwitchRoutingMulticastRendezvousPoint CreateNetworkSwitchRoutingMulticastRendezvousPointBody

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network switch routing multicast rendezvous point params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) WithDefaults() *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network switch routing multicast rendezvous point params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) WithTimeout(timeout time.Duration) *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) WithContext(ctx context.Context) *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) WithHTTPClient(client *http.Client) *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkSwitchRoutingMulticastRendezvousPoint adds the createNetworkSwitchRoutingMulticastRendezvousPoint to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) WithCreateNetworkSwitchRoutingMulticastRendezvousPoint(createNetworkSwitchRoutingMulticastRendezvousPoint CreateNetworkSwitchRoutingMulticastRendezvousPointBody) *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetCreateNetworkSwitchRoutingMulticastRendezvousPoint(createNetworkSwitchRoutingMulticastRendezvousPoint)
	return o
}

// SetCreateNetworkSwitchRoutingMulticastRendezvousPoint adds the createNetworkSwitchRoutingMulticastRendezvousPoint to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) SetCreateNetworkSwitchRoutingMulticastRendezvousPoint(createNetworkSwitchRoutingMulticastRendezvousPoint CreateNetworkSwitchRoutingMulticastRendezvousPointBody) {
	o.CreateNetworkSwitchRoutingMulticastRendezvousPoint = createNetworkSwitchRoutingMulticastRendezvousPoint
}

// WithNetworkID adds the networkID to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) WithNetworkID(networkID string) *CreateNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network switch routing multicast rendezvous point params
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkSwitchRoutingMulticastRendezvousPoint); err != nil {
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
