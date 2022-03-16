// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// UpdateOrganizationCameraOnboardingStatusesReader is a Reader for the UpdateOrganizationCameraOnboardingStatuses structure.
type UpdateOrganizationCameraOnboardingStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationCameraOnboardingStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationCameraOnboardingStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationCameraOnboardingStatusesOK creates a UpdateOrganizationCameraOnboardingStatusesOK with default headers values
func NewUpdateOrganizationCameraOnboardingStatusesOK() *UpdateOrganizationCameraOnboardingStatusesOK {
	return &UpdateOrganizationCameraOnboardingStatusesOK{}
}

/* UpdateOrganizationCameraOnboardingStatusesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationCameraOnboardingStatusesOK struct {
	Payload interface{}
}

func (o *UpdateOrganizationCameraOnboardingStatusesOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/camera/onboarding/statuses][%d] updateOrganizationCameraOnboardingStatusesOK  %+v", 200, o.Payload)
}
func (o *UpdateOrganizationCameraOnboardingStatusesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationCameraOnboardingStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateOrganizationCameraOnboardingStatusesBody update organization camera onboarding statuses body
// Example: {}
swagger:model UpdateOrganizationCameraOnboardingStatusesBody
*/
type UpdateOrganizationCameraOnboardingStatusesBody struct {

	// Serial of camera
	Serial string `json:"serial,omitempty"`

	// Note whether credentials were sent successfully
	WirelessCredentialsSent bool `json:"wirelessCredentialsSent,omitempty"`
}

// Validate validates this update organization camera onboarding statuses body
func (o *UpdateOrganizationCameraOnboardingStatusesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization camera onboarding statuses body based on context it is used
func (o *UpdateOrganizationCameraOnboardingStatusesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationCameraOnboardingStatusesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationCameraOnboardingStatusesBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationCameraOnboardingStatusesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
