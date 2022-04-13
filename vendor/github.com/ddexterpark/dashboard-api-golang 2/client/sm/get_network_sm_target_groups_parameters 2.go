// Code generated by go-swagger; DO NOT EDIT.

package sm

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
	"github.com/go-openapi/swag"
)

// NewGetNetworkSmTargetGroupsParams creates a new GetNetworkSmTargetGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSmTargetGroupsParams() *GetNetworkSmTargetGroupsParams {
	return &GetNetworkSmTargetGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmTargetGroupsParamsWithTimeout creates a new GetNetworkSmTargetGroupsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSmTargetGroupsParamsWithTimeout(timeout time.Duration) *GetNetworkSmTargetGroupsParams {
	return &GetNetworkSmTargetGroupsParams{
		timeout: timeout,
	}
}

// NewGetNetworkSmTargetGroupsParamsWithContext creates a new GetNetworkSmTargetGroupsParams object
// with the ability to set a context for a request.
func NewGetNetworkSmTargetGroupsParamsWithContext(ctx context.Context) *GetNetworkSmTargetGroupsParams {
	return &GetNetworkSmTargetGroupsParams{
		Context: ctx,
	}
}

// NewGetNetworkSmTargetGroupsParamsWithHTTPClient creates a new GetNetworkSmTargetGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSmTargetGroupsParamsWithHTTPClient(client *http.Client) *GetNetworkSmTargetGroupsParams {
	return &GetNetworkSmTargetGroupsParams{
		HTTPClient: client,
	}
}

/* GetNetworkSmTargetGroupsParams contains all the parameters to send to the API endpoint
   for the get network sm target groups operation.

   Typically these are written to a http.Request.
*/
type GetNetworkSmTargetGroupsParams struct {

	// NetworkID.
	NetworkID string

	/* WithDetails.

	   Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response
	*/
	WithDetails *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network sm target groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmTargetGroupsParams) WithDefaults() *GetNetworkSmTargetGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sm target groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmTargetGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) WithTimeout(timeout time.Duration) *GetNetworkSmTargetGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) WithContext(ctx context.Context) *GetNetworkSmTargetGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) WithHTTPClient(client *http.Client) *GetNetworkSmTargetGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) WithNetworkID(networkID string) *GetNetworkSmTargetGroupsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithWithDetails adds the withDetails to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) WithWithDetails(withDetails *bool) *GetNetworkSmTargetGroupsParams {
	o.SetWithDetails(withDetails)
	return o
}

// SetWithDetails adds the withDetails to the get network sm target groups params
func (o *GetNetworkSmTargetGroupsParams) SetWithDetails(withDetails *bool) {
	o.WithDetails = withDetails
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmTargetGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.WithDetails != nil {

		// query param withDetails
		var qrWithDetails bool

		if o.WithDetails != nil {
			qrWithDetails = *o.WithDetails
		}
		qWithDetails := swag.FormatBool(qrWithDetails)
		if qWithDetails != "" {

			if err := r.SetQueryParam("withDetails", qWithDetails); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
