// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworkFloorPlanParams creates a new GetNetworkFloorPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkFloorPlanParams() *GetNetworkFloorPlanParams {
	return &GetNetworkFloorPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkFloorPlanParamsWithTimeout creates a new GetNetworkFloorPlanParams object
// with the ability to set a timeout on a request.
func NewGetNetworkFloorPlanParamsWithTimeout(timeout time.Duration) *GetNetworkFloorPlanParams {
	return &GetNetworkFloorPlanParams{
		timeout: timeout,
	}
}

// NewGetNetworkFloorPlanParamsWithContext creates a new GetNetworkFloorPlanParams object
// with the ability to set a context for a request.
func NewGetNetworkFloorPlanParamsWithContext(ctx context.Context) *GetNetworkFloorPlanParams {
	return &GetNetworkFloorPlanParams{
		Context: ctx,
	}
}

// NewGetNetworkFloorPlanParamsWithHTTPClient creates a new GetNetworkFloorPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkFloorPlanParamsWithHTTPClient(client *http.Client) *GetNetworkFloorPlanParams {
	return &GetNetworkFloorPlanParams{
		HTTPClient: client,
	}
}

/* GetNetworkFloorPlanParams contains all the parameters to send to the API endpoint
   for the get network floor plan operation.

   Typically these are written to a http.Request.
*/
type GetNetworkFloorPlanParams struct {

	// FloorPlanID.
	FloorPlanID string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network floor plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkFloorPlanParams) WithDefaults() *GetNetworkFloorPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network floor plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkFloorPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network floor plan params
func (o *GetNetworkFloorPlanParams) WithTimeout(timeout time.Duration) *GetNetworkFloorPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network floor plan params
func (o *GetNetworkFloorPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network floor plan params
func (o *GetNetworkFloorPlanParams) WithContext(ctx context.Context) *GetNetworkFloorPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network floor plan params
func (o *GetNetworkFloorPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network floor plan params
func (o *GetNetworkFloorPlanParams) WithHTTPClient(client *http.Client) *GetNetworkFloorPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network floor plan params
func (o *GetNetworkFloorPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFloorPlanID adds the floorPlanID to the get network floor plan params
func (o *GetNetworkFloorPlanParams) WithFloorPlanID(floorPlanID string) *GetNetworkFloorPlanParams {
	o.SetFloorPlanID(floorPlanID)
	return o
}

// SetFloorPlanID adds the floorPlanId to the get network floor plan params
func (o *GetNetworkFloorPlanParams) SetFloorPlanID(floorPlanID string) {
	o.FloorPlanID = floorPlanID
}

// WithNetworkID adds the networkID to the get network floor plan params
func (o *GetNetworkFloorPlanParams) WithNetworkID(networkID string) *GetNetworkFloorPlanParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network floor plan params
func (o *GetNetworkFloorPlanParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkFloorPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param floorPlanId
	if err := r.SetPathParam("floorPlanId", o.FloorPlanID); err != nil {
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