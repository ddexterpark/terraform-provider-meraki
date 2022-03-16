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

// NewGetNetworkSmTargetGroupParams creates a new GetNetworkSmTargetGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSmTargetGroupParams() *GetNetworkSmTargetGroupParams {
	return &GetNetworkSmTargetGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmTargetGroupParamsWithTimeout creates a new GetNetworkSmTargetGroupParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSmTargetGroupParamsWithTimeout(timeout time.Duration) *GetNetworkSmTargetGroupParams {
	return &GetNetworkSmTargetGroupParams{
		timeout: timeout,
	}
}

// NewGetNetworkSmTargetGroupParamsWithContext creates a new GetNetworkSmTargetGroupParams object
// with the ability to set a context for a request.
func NewGetNetworkSmTargetGroupParamsWithContext(ctx context.Context) *GetNetworkSmTargetGroupParams {
	return &GetNetworkSmTargetGroupParams{
		Context: ctx,
	}
}

// NewGetNetworkSmTargetGroupParamsWithHTTPClient creates a new GetNetworkSmTargetGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSmTargetGroupParamsWithHTTPClient(client *http.Client) *GetNetworkSmTargetGroupParams {
	return &GetNetworkSmTargetGroupParams{
		HTTPClient: client,
	}
}

/* GetNetworkSmTargetGroupParams contains all the parameters to send to the API endpoint
   for the get network sm target group operation.

   Typically these are written to a http.Request.
*/
type GetNetworkSmTargetGroupParams struct {

	// NetworkID.
	NetworkID string

	// TargetGroupID.
	TargetGroupID string

	/* WithDetails.

	   Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response
	*/
	WithDetails *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network sm target group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmTargetGroupParams) WithDefaults() *GetNetworkSmTargetGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sm target group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmTargetGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) WithTimeout(timeout time.Duration) *GetNetworkSmTargetGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) WithContext(ctx context.Context) *GetNetworkSmTargetGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) WithHTTPClient(client *http.Client) *GetNetworkSmTargetGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) WithNetworkID(networkID string) *GetNetworkSmTargetGroupParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTargetGroupID adds the targetGroupID to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) WithTargetGroupID(targetGroupID string) *GetNetworkSmTargetGroupParams {
	o.SetTargetGroupID(targetGroupID)
	return o
}

// SetTargetGroupID adds the targetGroupId to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) SetTargetGroupID(targetGroupID string) {
	o.TargetGroupID = targetGroupID
}

// WithWithDetails adds the withDetails to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) WithWithDetails(withDetails *bool) *GetNetworkSmTargetGroupParams {
	o.SetWithDetails(withDetails)
	return o
}

// SetWithDetails adds the withDetails to the get network sm target group params
func (o *GetNetworkSmTargetGroupParams) SetWithDetails(withDetails *bool) {
	o.WithDetails = withDetails
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmTargetGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param targetGroupId
	if err := r.SetPathParam("targetGroupId", o.TargetGroupID); err != nil {
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
