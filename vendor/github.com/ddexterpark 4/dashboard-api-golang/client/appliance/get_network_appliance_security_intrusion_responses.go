// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkApplianceSecurityIntrusionReader is a Reader for the GetNetworkApplianceSecurityIntrusion structure.
type GetNetworkApplianceSecurityIntrusionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceSecurityIntrusionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceSecurityIntrusionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceSecurityIntrusionOK creates a GetNetworkApplianceSecurityIntrusionOK with default headers values
func NewGetNetworkApplianceSecurityIntrusionOK() *GetNetworkApplianceSecurityIntrusionOK {
	return &GetNetworkApplianceSecurityIntrusionOK{}
}

/* GetNetworkApplianceSecurityIntrusionOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceSecurityIntrusionOK struct {
	Payload interface{}
}

func (o *GetNetworkApplianceSecurityIntrusionOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/security/intrusion][%d] getNetworkApplianceSecurityIntrusionOK  %+v", 200, o.Payload)
}
func (o *GetNetworkApplianceSecurityIntrusionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceSecurityIntrusionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
