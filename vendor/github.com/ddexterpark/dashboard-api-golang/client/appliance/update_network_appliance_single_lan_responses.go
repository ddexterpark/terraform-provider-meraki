// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// UpdateNetworkApplianceSingleLanReader is a Reader for the UpdateNetworkApplianceSingleLan structure.
type UpdateNetworkApplianceSingleLanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceSingleLanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceSingleLanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkApplianceSingleLanOK creates a UpdateNetworkApplianceSingleLanOK with default headers values
func NewUpdateNetworkApplianceSingleLanOK() *UpdateNetworkApplianceSingleLanOK {
	return &UpdateNetworkApplianceSingleLanOK{}
}

/* UpdateNetworkApplianceSingleLanOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceSingleLanOK struct {
	Payload interface{}
}

func (o *UpdateNetworkApplianceSingleLanOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/singleLan][%d] updateNetworkApplianceSingleLanOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkApplianceSingleLanOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceSingleLanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkApplianceSingleLanBody update network appliance single lan body
// Example: {}
swagger:model UpdateNetworkApplianceSingleLanBody
*/
type UpdateNetworkApplianceSingleLanBody struct {

	// The appliance IP address of the single LAN
	ApplianceIP string `json:"applianceIp,omitempty"`

	// The subnet of the single LAN configuration
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this update network appliance single lan body
func (o *UpdateNetworkApplianceSingleLanBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance single lan body based on context it is used
func (o *UpdateNetworkApplianceSingleLanBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}