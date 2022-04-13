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
	"github.com/go-openapi/swag"
)

// NewGetNetworkApplianceSecurityEventsParams creates a new GetNetworkApplianceSecurityEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkApplianceSecurityEventsParams() *GetNetworkApplianceSecurityEventsParams {
	return &GetNetworkApplianceSecurityEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkApplianceSecurityEventsParamsWithTimeout creates a new GetNetworkApplianceSecurityEventsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkApplianceSecurityEventsParamsWithTimeout(timeout time.Duration) *GetNetworkApplianceSecurityEventsParams {
	return &GetNetworkApplianceSecurityEventsParams{
		timeout: timeout,
	}
}

// NewGetNetworkApplianceSecurityEventsParamsWithContext creates a new GetNetworkApplianceSecurityEventsParams object
// with the ability to set a context for a request.
func NewGetNetworkApplianceSecurityEventsParamsWithContext(ctx context.Context) *GetNetworkApplianceSecurityEventsParams {
	return &GetNetworkApplianceSecurityEventsParams{
		Context: ctx,
	}
}

// NewGetNetworkApplianceSecurityEventsParamsWithHTTPClient creates a new GetNetworkApplianceSecurityEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkApplianceSecurityEventsParamsWithHTTPClient(client *http.Client) *GetNetworkApplianceSecurityEventsParams {
	return &GetNetworkApplianceSecurityEventsParams{
		HTTPClient: client,
	}
}

/* GetNetworkApplianceSecurityEventsParams contains all the parameters to send to the API endpoint
   for the get network appliance security events operation.

   Typically these are written to a http.Request.
*/
type GetNetworkApplianceSecurityEventsParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	// NetworkID.
	NetworkID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	*/
	PerPage *int64

	/* SortOrder.

	   Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order.
	*/
	SortOrder *string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* T0.

	   The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network appliance security events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceSecurityEventsParams) WithDefaults() *GetNetworkApplianceSecurityEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network appliance security events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceSecurityEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithTimeout(timeout time.Duration) *GetNetworkApplianceSecurityEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithContext(ctx context.Context) *GetNetworkApplianceSecurityEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithHTTPClient(client *http.Client) *GetNetworkApplianceSecurityEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithEndingBefore(endingBefore *string) *GetNetworkApplianceSecurityEventsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithNetworkID adds the networkID to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithNetworkID(networkID string) *GetNetworkApplianceSecurityEventsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPerPage adds the perPage to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithPerPage(perPage *int64) *GetNetworkApplianceSecurityEventsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSortOrder adds the sortOrder to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithSortOrder(sortOrder *string) *GetNetworkApplianceSecurityEventsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStartingAfter adds the startingAfter to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithStartingAfter(startingAfter *string) *GetNetworkApplianceSecurityEventsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithT0 adds the t0 to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithT0(t0 *string) *GetNetworkApplianceSecurityEventsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithT1(t1 *string) *GetNetworkApplianceSecurityEventsParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) WithTimespan(timespan *float32) *GetNetworkApplianceSecurityEventsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get network appliance security events params
func (o *GetNetworkApplianceSecurityEventsParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkApplianceSecurityEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingBefore != nil {

		// query param endingBefore
		var qrEndingBefore string

		if o.EndingBefore != nil {
			qrEndingBefore = *o.EndingBefore
		}
		qEndingBefore := qrEndingBefore
		if qEndingBefore != "" {

			if err := r.SetQueryParam("endingBefore", qEndingBefore); err != nil {
				return err
			}
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}
	}

	if o.StartingAfter != nil {

		// query param startingAfter
		var qrStartingAfter string

		if o.StartingAfter != nil {
			qrStartingAfter = *o.StartingAfter
		}
		qStartingAfter := qrStartingAfter
		if qStartingAfter != "" {

			if err := r.SetQueryParam("startingAfter", qStartingAfter); err != nil {
				return err
			}
		}
	}

	if o.T0 != nil {

		// query param t0
		var qrT0 string

		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {

			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}
	}

	if o.T1 != nil {

		// query param t1
		var qrT1 string

		if o.T1 != nil {
			qrT1 = *o.T1
		}
		qT1 := qrT1
		if qT1 != "" {

			if err := r.SetQueryParam("t1", qT1); err != nil {
				return err
			}
		}
	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float32

		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat32(qrTimespan)
		if qTimespan != "" {

			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
