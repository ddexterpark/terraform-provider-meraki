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

// NewUpdateNetworkWebhooksPayloadTemplateParams creates a new UpdateNetworkWebhooksPayloadTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkWebhooksPayloadTemplateParams() *UpdateNetworkWebhooksPayloadTemplateParams {
	return &UpdateNetworkWebhooksPayloadTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWebhooksPayloadTemplateParamsWithTimeout creates a new UpdateNetworkWebhooksPayloadTemplateParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkWebhooksPayloadTemplateParamsWithTimeout(timeout time.Duration) *UpdateNetworkWebhooksPayloadTemplateParams {
	return &UpdateNetworkWebhooksPayloadTemplateParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkWebhooksPayloadTemplateParamsWithContext creates a new UpdateNetworkWebhooksPayloadTemplateParams object
// with the ability to set a context for a request.
func NewUpdateNetworkWebhooksPayloadTemplateParamsWithContext(ctx context.Context) *UpdateNetworkWebhooksPayloadTemplateParams {
	return &UpdateNetworkWebhooksPayloadTemplateParams{
		Context: ctx,
	}
}

// NewUpdateNetworkWebhooksPayloadTemplateParamsWithHTTPClient creates a new UpdateNetworkWebhooksPayloadTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkWebhooksPayloadTemplateParamsWithHTTPClient(client *http.Client) *UpdateNetworkWebhooksPayloadTemplateParams {
	return &UpdateNetworkWebhooksPayloadTemplateParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkWebhooksPayloadTemplateParams contains all the parameters to send to the API endpoint
   for the update network webhooks payload template operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkWebhooksPayloadTemplateParams struct {

	// NetworkID.
	NetworkID string

	// PayloadTemplateID.
	PayloadTemplateID string

	// UpdateNetworkWebhooksPayloadTemplate.
	UpdateNetworkWebhooksPayloadTemplate UpdateNetworkWebhooksPayloadTemplateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network webhooks payload template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWebhooksPayloadTemplateParams) WithDefaults() *UpdateNetworkWebhooksPayloadTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network webhooks payload template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWebhooksPayloadTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) WithTimeout(timeout time.Duration) *UpdateNetworkWebhooksPayloadTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) WithContext(ctx context.Context) *UpdateNetworkWebhooksPayloadTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) WithHTTPClient(client *http.Client) *UpdateNetworkWebhooksPayloadTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) WithNetworkID(networkID string) *UpdateNetworkWebhooksPayloadTemplateParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPayloadTemplateID adds the payloadTemplateID to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) WithPayloadTemplateID(payloadTemplateID string) *UpdateNetworkWebhooksPayloadTemplateParams {
	o.SetPayloadTemplateID(payloadTemplateID)
	return o
}

// SetPayloadTemplateID adds the payloadTemplateId to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) SetPayloadTemplateID(payloadTemplateID string) {
	o.PayloadTemplateID = payloadTemplateID
}

// WithUpdateNetworkWebhooksPayloadTemplate adds the updateNetworkWebhooksPayloadTemplate to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) WithUpdateNetworkWebhooksPayloadTemplate(updateNetworkWebhooksPayloadTemplate UpdateNetworkWebhooksPayloadTemplateBody) *UpdateNetworkWebhooksPayloadTemplateParams {
	o.SetUpdateNetworkWebhooksPayloadTemplate(updateNetworkWebhooksPayloadTemplate)
	return o
}

// SetUpdateNetworkWebhooksPayloadTemplate adds the updateNetworkWebhooksPayloadTemplate to the update network webhooks payload template params
func (o *UpdateNetworkWebhooksPayloadTemplateParams) SetUpdateNetworkWebhooksPayloadTemplate(updateNetworkWebhooksPayloadTemplate UpdateNetworkWebhooksPayloadTemplateBody) {
	o.UpdateNetworkWebhooksPayloadTemplate = updateNetworkWebhooksPayloadTemplate
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWebhooksPayloadTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param payloadTemplateId
	if err := r.SetPathParam("payloadTemplateId", o.PayloadTemplateID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkWebhooksPayloadTemplate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
