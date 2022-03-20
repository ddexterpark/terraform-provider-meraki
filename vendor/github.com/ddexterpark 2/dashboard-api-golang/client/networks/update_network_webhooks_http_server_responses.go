// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkWebhooksHTTPServerReader is a Reader for the UpdateNetworkWebhooksHTTPServer structure.
type UpdateNetworkWebhooksHTTPServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkWebhooksHTTPServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkWebhooksHTTPServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkWebhooksHTTPServerOK creates a UpdateNetworkWebhooksHTTPServerOK with default headers values
func NewUpdateNetworkWebhooksHTTPServerOK() *UpdateNetworkWebhooksHTTPServerOK {
	return &UpdateNetworkWebhooksHTTPServerOK{}
}

/* UpdateNetworkWebhooksHTTPServerOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkWebhooksHTTPServerOK struct {
	Payload interface{}
}

func (o *UpdateNetworkWebhooksHTTPServerOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/webhooks/httpServers/{httpServerId}][%d] updateNetworkWebhooksHttpServerOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkWebhooksHTTPServerOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkWebhooksHTTPServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkWebhooksHTTPServerBody update network webhooks HTTP server body
// Example: {"name":"Example Webhook Server","payloadTemplate":{"name":"Meraki (included)","payloadTemplateId":"wpt_00001"},"sharedSecret":"shhh","url":"https://example.com"}
swagger:model UpdateNetworkWebhooksHTTPServerBody
*/
type UpdateNetworkWebhooksHTTPServerBody struct {

	// A name for easy reference to the HTTP server
	Name string `json:"name,omitempty"`

	// payload template
	PayloadTemplate *UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate `json:"payloadTemplate,omitempty"`

	// A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
	SharedSecret string `json:"sharedSecret,omitempty"`
}

// Validate validates this update network webhooks HTTP server body
func (o *UpdateNetworkWebhooksHTTPServerBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePayloadTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWebhooksHTTPServerBody) validatePayloadTemplate(formats strfmt.Registry) error {
	if swag.IsZero(o.PayloadTemplate) { // not required
		return nil
	}

	if o.PayloadTemplate != nil {
		if err := o.PayloadTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWebhooksHttpServer" + "." + "payloadTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWebhooksHttpServer" + "." + "payloadTemplate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network webhooks HTTP server body based on the context it is used
func (o *UpdateNetworkWebhooksHTTPServerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePayloadTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWebhooksHTTPServerBody) contextValidatePayloadTemplate(ctx context.Context, formats strfmt.Registry) error {

	if o.PayloadTemplate != nil {
		if err := o.PayloadTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWebhooksHttpServer" + "." + "payloadTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWebhooksHttpServer" + "." + "payloadTemplate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWebhooksHTTPServerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWebhooksHTTPServerBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWebhooksHTTPServerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate The payload template to use when posting data to the HTTP server.
swagger:model UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate
*/
type UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate struct {

	// The ID of the payload template. Defaults to 'wpt_00001' for the Meraki template. For Webex, use 'wpt_00002'; for Slack, use 'wpt_00003'; for Microsoft Teams, use 'wpt_00004'.
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"`
}

// Validate validates this update network webhooks HTTP server params body payload template
func (o *UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network webhooks HTTP server params body payload template based on context it is used
func (o *UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}