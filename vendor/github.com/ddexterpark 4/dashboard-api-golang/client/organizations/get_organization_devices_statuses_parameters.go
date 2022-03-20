// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrganizationDevicesStatusesParams creates a new GetOrganizationDevicesStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationDevicesStatusesParams() *GetOrganizationDevicesStatusesParams {
	return &GetOrganizationDevicesStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationDevicesStatusesParamsWithTimeout creates a new GetOrganizationDevicesStatusesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationDevicesStatusesParamsWithTimeout(timeout time.Duration) *GetOrganizationDevicesStatusesParams {
	return &GetOrganizationDevicesStatusesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationDevicesStatusesParamsWithContext creates a new GetOrganizationDevicesStatusesParams object
// with the ability to set a context for a request.
func NewGetOrganizationDevicesStatusesParamsWithContext(ctx context.Context) *GetOrganizationDevicesStatusesParams {
	return &GetOrganizationDevicesStatusesParams{
		Context: ctx,
	}
}

// NewGetOrganizationDevicesStatusesParamsWithHTTPClient creates a new GetOrganizationDevicesStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationDevicesStatusesParamsWithHTTPClient(client *http.Client) *GetOrganizationDevicesStatusesParams {
	return &GetOrganizationDevicesStatusesParams{
		HTTPClient: client,
	}
}

/* GetOrganizationDevicesStatusesParams contains all the parameters to send to the API endpoint
   for the get organization devices statuses operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationDevicesStatusesParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* Models.

	   Optional parameter to filter devices by models.
	*/
	Models []string

	/* NetworkIds.

	   Optional parameter to filter devices by network ids.
	*/
	NetworkIds []string

	// OrganizationID.
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	*/
	PerPage *int64

	/* ProductTypes.

	   An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and sensor.
	*/
	ProductTypes []string

	/* Serials.

	   Optional parameter to filter devices by serials.
	*/
	Serials []string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* Statuses.

	   Optional parameter to filter devices by statuses. Valid statuses are ["online", "alerting", "offline", "dormant"].
	*/
	Statuses []string

	/* Tags.

	   An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).
	*/
	Tags []string

	/* TagsFilterType.

	   An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
	*/
	TagsFilterType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization devices statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationDevicesStatusesParams) WithDefaults() *GetOrganizationDevicesStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization devices statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationDevicesStatusesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithTimeout(timeout time.Duration) *GetOrganizationDevicesStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithContext(ctx context.Context) *GetOrganizationDevicesStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithHTTPClient(client *http.Client) *GetOrganizationDevicesStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithEndingBefore(endingBefore *string) *GetOrganizationDevicesStatusesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithModels adds the models to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithModels(models []string) *GetOrganizationDevicesStatusesParams {
	o.SetModels(models)
	return o
}

// SetModels adds the models to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetModels(models []string) {
	o.Models = models
}

// WithNetworkIds adds the networkIds to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithNetworkIds(networkIds []string) *GetOrganizationDevicesStatusesParams {
	o.SetNetworkIds(networkIds)
	return o
}

// SetNetworkIds adds the networkIds to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetNetworkIds(networkIds []string) {
	o.NetworkIds = networkIds
}

// WithOrganizationID adds the organizationID to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithOrganizationID(organizationID string) *GetOrganizationDevicesStatusesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithPerPage(perPage *int64) *GetOrganizationDevicesStatusesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithProductTypes adds the productTypes to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithProductTypes(productTypes []string) *GetOrganizationDevicesStatusesParams {
	o.SetProductTypes(productTypes)
	return o
}

// SetProductTypes adds the productTypes to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetProductTypes(productTypes []string) {
	o.ProductTypes = productTypes
}

// WithSerials adds the serials to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithSerials(serials []string) *GetOrganizationDevicesStatusesParams {
	o.SetSerials(serials)
	return o
}

// SetSerials adds the serials to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetSerials(serials []string) {
	o.Serials = serials
}

// WithStartingAfter adds the startingAfter to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithStartingAfter(startingAfter *string) *GetOrganizationDevicesStatusesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithStatuses adds the statuses to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithStatuses(statuses []string) *GetOrganizationDevicesStatusesParams {
	o.SetStatuses(statuses)
	return o
}

// SetStatuses adds the statuses to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetStatuses(statuses []string) {
	o.Statuses = statuses
}

// WithTags adds the tags to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithTags(tags []string) *GetOrganizationDevicesStatusesParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithTagsFilterType adds the tagsFilterType to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) WithTagsFilterType(tagsFilterType *string) *GetOrganizationDevicesStatusesParams {
	o.SetTagsFilterType(tagsFilterType)
	return o
}

// SetTagsFilterType adds the tagsFilterType to the get organization devices statuses params
func (o *GetOrganizationDevicesStatusesParams) SetTagsFilterType(tagsFilterType *string) {
	o.TagsFilterType = tagsFilterType
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationDevicesStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Models != nil {

		// binding items for models
		joinedModels := o.bindParamModels(reg)

		// query array param models
		if err := r.SetQueryParam("models", joinedModels...); err != nil {
			return err
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

	if o.ProductTypes != nil {

		// binding items for productTypes
		joinedProductTypes := o.bindParamProductTypes(reg)

		// query array param productTypes
		if err := r.SetQueryParam("productTypes", joinedProductTypes...); err != nil {
			return err
		}
	}

	if o.Serials != nil {

		// binding items for serials
		joinedSerials := o.bindParamSerials(reg)

		// query array param serials
		if err := r.SetQueryParam("serials", joinedSerials...); err != nil {
			return err
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

	if o.Statuses != nil {

		// binding items for statuses
		joinedStatuses := o.bindParamStatuses(reg)

		// query array param statuses
		if err := r.SetQueryParam("statuses", joinedStatuses...); err != nil {
			return err
		}
	}

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// query array param tags
		if err := r.SetQueryParam("tags", joinedTags...); err != nil {
			return err
		}
	}

	if o.TagsFilterType != nil {

		// query param tagsFilterType
		var qrTagsFilterType string

		if o.TagsFilterType != nil {
			qrTagsFilterType = *o.TagsFilterType
		}
		qTagsFilterType := qrTagsFilterType
		if qTagsFilterType != "" {

			if err := r.SetQueryParam("tagsFilterType", qTagsFilterType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrganizationDevicesStatuses binds the parameter models
func (o *GetOrganizationDevicesStatusesParams) bindParamModels(formats strfmt.Registry) []string {
	modelsIR := o.Models

	var modelsIC []string
	for _, modelsIIR := range modelsIR { // explode []string

		modelsIIV := modelsIIR // string as string
		modelsIC = append(modelsIC, modelsIIV)
	}

	// items.CollectionFormat: ""
	modelsIS := swag.JoinByFormat(modelsIC, "")

	return modelsIS
}

// bindParamGetOrganizationDevicesStatuses binds the parameter networkIds
func (o *GetOrganizationDevicesStatusesParams) bindParamNetworkIds(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationDevicesStatuses binds the parameter productTypes
func (o *GetOrganizationDevicesStatusesParams) bindParamProductTypes(formats strfmt.Registry) []string {
	productTypesIR := o.ProductTypes

	var productTypesIC []string
	for _, productTypesIIR := range productTypesIR { // explode []string

		productTypesIIV := productTypesIIR // string as string
		productTypesIC = append(productTypesIC, productTypesIIV)
	}

	// items.CollectionFormat: ""
	productTypesIS := swag.JoinByFormat(productTypesIC, "")

	return productTypesIS
}

// bindParamGetOrganizationDevicesStatuses binds the parameter serials
func (o *GetOrganizationDevicesStatusesParams) bindParamSerials(formats strfmt.Registry) []string {
	serialsIR := o.Serials

	var serialsIC []string
	for _, serialsIIR := range serialsIR { // explode []string

		serialsIIV := serialsIIR // string as string
		serialsIC = append(serialsIC, serialsIIV)
	}

	// items.CollectionFormat: ""
	serialsIS := swag.JoinByFormat(serialsIC, "")

	return serialsIS
}

// bindParamGetOrganizationDevicesStatuses binds the parameter statuses
func (o *GetOrganizationDevicesStatusesParams) bindParamStatuses(formats strfmt.Registry) []string {
	statusesIR := o.Statuses

	var statusesIC []string
	for _, statusesIIR := range statusesIR { // explode []string

		statusesIIV := statusesIIR // string as string
		statusesIC = append(statusesIC, statusesIIV)
	}

	// items.CollectionFormat: ""
	statusesIS := swag.JoinByFormat(statusesIC, "")

	return statusesIS
}

// bindParamGetOrganizationDevicesStatuses binds the parameter tags
func (o *GetOrganizationDevicesStatusesParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: ""
	tagsIS := swag.JoinByFormat(tagsIC, "")

	return tagsIS
}
