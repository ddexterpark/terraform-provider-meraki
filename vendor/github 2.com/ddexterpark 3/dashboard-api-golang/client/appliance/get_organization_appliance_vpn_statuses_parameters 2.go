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

// NewGetOrganizationApplianceVpnStatusesParams creates a new GetOrganizationApplianceVpnStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationApplianceVpnStatusesParams() *GetOrganizationApplianceVpnStatusesParams {
	return &GetOrganizationApplianceVpnStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationApplianceVpnStatusesParamsWithTimeout creates a new GetOrganizationApplianceVpnStatusesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationApplianceVpnStatusesParamsWithTimeout(timeout time.Duration) *GetOrganizationApplianceVpnStatusesParams {
	return &GetOrganizationApplianceVpnStatusesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationApplianceVpnStatusesParamsWithContext creates a new GetOrganizationApplianceVpnStatusesParams object
// with the ability to set a context for a request.
func NewGetOrganizationApplianceVpnStatusesParamsWithContext(ctx context.Context) *GetOrganizationApplianceVpnStatusesParams {
	return &GetOrganizationApplianceVpnStatusesParams{
		Context: ctx,
	}
}

// NewGetOrganizationApplianceVpnStatusesParamsWithHTTPClient creates a new GetOrganizationApplianceVpnStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationApplianceVpnStatusesParamsWithHTTPClient(client *http.Client) *GetOrganizationApplianceVpnStatusesParams {
	return &GetOrganizationApplianceVpnStatusesParams{
		HTTPClient: client,
	}
}

/* GetOrganizationApplianceVpnStatusesParams contains all the parameters to send to the API endpoint
   for the get organization appliance vpn statuses operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationApplianceVpnStatusesParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* NetworkIds.

	   A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456
	*/
	NetworkIds []string

	// OrganizationID.
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 300. Default is 300.
	*/
	PerPage *int64

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization appliance vpn statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationApplianceVpnStatusesParams) WithDefaults() *GetOrganizationApplianceVpnStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization appliance vpn statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationApplianceVpnStatusesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) WithTimeout(timeout time.Duration) *GetOrganizationApplianceVpnStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) WithContext(ctx context.Context) *GetOrganizationApplianceVpnStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) WithHTTPClient(client *http.Client) *GetOrganizationApplianceVpnStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) WithEndingBefore(endingBefore *string) *GetOrganizationApplianceVpnStatusesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithNetworkIds adds the networkIds to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) WithNetworkIds(networkIds []string) *GetOrganizationApplianceVpnStatusesParams {
	o.SetNetworkIds(networkIds)
	return o
}

// SetNetworkIds adds the networkIds to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) SetNetworkIds(networkIds []string) {
	o.NetworkIds = networkIds
}

// WithOrganizationID adds the organizationID to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) WithOrganizationID(organizationID string) *GetOrganizationApplianceVpnStatusesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) WithPerPage(perPage *int64) *GetOrganizationApplianceVpnStatusesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) WithStartingAfter(startingAfter *string) *GetOrganizationApplianceVpnStatusesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization appliance vpn statuses params
func (o *GetOrganizationApplianceVpnStatusesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationApplianceVpnStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.NetworkIds != nil {

		// binding items for networkIds
		joinedNetworkIds := o.bindParamNetworkIds(reg)

		// query array param networkIds
		if err := r.SetQueryParam("networkIds", joinedNetworkIds...); err != nil {
			return err
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrganizationApplianceVpnStatuses binds the parameter networkIds
func (o *GetOrganizationApplianceVpnStatusesParams) bindParamNetworkIds(formats strfmt.Registry) []string {
	networkIdsIR := o.NetworkIds

	var networkIdsIC []string
	for _, networkIdsIIR := range networkIdsIR { // explode []string

		networkIdsIIV := networkIdsIIR // string as string
		networkIdsIC = append(networkIdsIC, networkIdsIIV)
	}

	// items.CollectionFormat: ""
	networkIdsIS := swag.JoinByFormat(networkIdsIC, "")

	return networkIdsIS
}
