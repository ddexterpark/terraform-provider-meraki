// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateOrganizationSamlIdpReader is a Reader for the UpdateOrganizationSamlIdp structure.
type UpdateOrganizationSamlIdpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationSamlIdpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationSamlIdpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationSamlIdpOK creates a UpdateOrganizationSamlIdpOK with default headers values
func NewUpdateOrganizationSamlIdpOK() *UpdateOrganizationSamlIdpOK {
	return &UpdateOrganizationSamlIdpOK{}
}

/* UpdateOrganizationSamlIdpOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationSamlIdpOK struct {
	Payload interface{}
}

func (o *UpdateOrganizationSamlIdpOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/saml/idps/{idpId}][%d] updateOrganizationSamlIdpOK  %+v", 200, o.Payload)
}
func (o *UpdateOrganizationSamlIdpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationSamlIdpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateOrganizationSamlIdpBody update organization saml idp body
// Example: {"idpId":"ab0c1de23Fg","sloLogoutUrl":"https://somewhere.com","x509certSha1Fingerprint":"00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:99"}
swagger:model UpdateOrganizationSamlIdpBody
*/
type UpdateOrganizationSamlIdpBody struct {

	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutURL string `json:"sloLogoutUrl,omitempty"`

	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint string `json:"x509certSha1Fingerprint,omitempty"`
}

// Validate validates this update organization saml idp body
func (o *UpdateOrganizationSamlIdpBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization saml idp body based on context it is used
func (o *UpdateOrganizationSamlIdpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationSamlIdpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationSamlIdpBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationSamlIdpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}