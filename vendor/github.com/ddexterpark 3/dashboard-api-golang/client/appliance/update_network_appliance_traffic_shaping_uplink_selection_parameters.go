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

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionParams creates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionParams() *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	return &UpdateNetworkApplianceTrafficShapingUplinkSelectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionParamsWithTimeout creates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	return &UpdateNetworkApplianceTrafficShapingUplinkSelectionParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionParamsWithContext creates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionParamsWithContext(ctx context.Context) *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	return &UpdateNetworkApplianceTrafficShapingUplinkSelectionParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionParamsWithHTTPClient creates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	return &UpdateNetworkApplianceTrafficShapingUplinkSelectionParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkApplianceTrafficShapingUplinkSelectionParams contains all the parameters to send to the API endpoint
   for the update network appliance traffic shaping uplink selection operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceTrafficShapingUplinkSelectionParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkApplianceTrafficShapingUplinkSelection.
	UpdateNetworkApplianceTrafficShapingUplinkSelection UpdateNetworkApplianceTrafficShapingUplinkSelectionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance traffic shaping uplink selection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) WithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance traffic shaping uplink selection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) WithContext(ctx context.Context) *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) WithNetworkID(networkID string) *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceTrafficShapingUplinkSelection adds the updateNetworkApplianceTrafficShapingUplinkSelection to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) WithUpdateNetworkApplianceTrafficShapingUplinkSelection(updateNetworkApplianceTrafficShapingUplinkSelection UpdateNetworkApplianceTrafficShapingUplinkSelectionBody) *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams {
	o.SetUpdateNetworkApplianceTrafficShapingUplinkSelection(updateNetworkApplianceTrafficShapingUplinkSelection)
	return o
}

// SetUpdateNetworkApplianceTrafficShapingUplinkSelection adds the updateNetworkApplianceTrafficShapingUplinkSelection to the update network appliance traffic shaping uplink selection params
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) SetUpdateNetworkApplianceTrafficShapingUplinkSelection(updateNetworkApplianceTrafficShapingUplinkSelection UpdateNetworkApplianceTrafficShapingUplinkSelectionBody) {
	o.UpdateNetworkApplianceTrafficShapingUplinkSelection = updateNetworkApplianceTrafficShapingUplinkSelection
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceTrafficShapingUplinkSelection); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
