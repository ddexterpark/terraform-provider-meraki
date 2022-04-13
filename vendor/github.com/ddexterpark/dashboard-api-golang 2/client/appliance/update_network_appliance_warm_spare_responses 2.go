// Code generated by go-swagger; DO NOT EDIT.

package appliance

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
	"github.com/go-openapi/validate"
)

// UpdateNetworkApplianceWarmSpareReader is a Reader for the UpdateNetworkApplianceWarmSpare structure.
type UpdateNetworkApplianceWarmSpareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceWarmSpareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceWarmSpareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkApplianceWarmSpareOK creates a UpdateNetworkApplianceWarmSpareOK with default headers values
func NewUpdateNetworkApplianceWarmSpareOK() *UpdateNetworkApplianceWarmSpareOK {
	return &UpdateNetworkApplianceWarmSpareOK{}
}

/* UpdateNetworkApplianceWarmSpareOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceWarmSpareOK struct {
	Payload interface{}
}

func (o *UpdateNetworkApplianceWarmSpareOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/warmSpare][%d] updateNetworkApplianceWarmSpareOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkApplianceWarmSpareOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceWarmSpareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkApplianceWarmSpareBody update network appliance warm spare body
// Example: {"enabled":true,"spareSerial":"Q234-ABCD-5678","uplinkMode":"virtual","virtualIp1":"1.2.3.4","virtualIp2":"1.2.3.4"}
swagger:model UpdateNetworkApplianceWarmSpareBody
*/
type UpdateNetworkApplianceWarmSpareBody struct {

	// Enable warm spare
	// Required: true
	Enabled *bool `json:"enabled"`

	// Serial number of the warm spare appliance
	SpareSerial string `json:"spareSerial,omitempty"`

	// Uplink mode, either virtual or public
	UplinkMode string `json:"uplinkMode,omitempty"`

	// The WAN 1 shared IP
	VirtualIp1 string `json:"virtualIp1,omitempty"`

	// The WAN 2 shared IP
	VirtualIp2 string `json:"virtualIp2,omitempty"`
}

// Validate validates this update network appliance warm spare body
func (o *UpdateNetworkApplianceWarmSpareBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceWarmSpareBody) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("updateNetworkApplianceWarmSpare"+"."+"enabled", "body", o.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance warm spare body based on context it is used
func (o *UpdateNetworkApplianceWarmSpareBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceWarmSpareBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceWarmSpareBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceWarmSpareBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
