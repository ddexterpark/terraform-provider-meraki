// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSmDeviceSecurityCentersReader is a Reader for the GetNetworkSmDeviceSecurityCenters structure.
type GetNetworkSmDeviceSecurityCentersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceSecurityCentersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceSecurityCentersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmDeviceSecurityCentersOK creates a GetNetworkSmDeviceSecurityCentersOK with default headers values
func NewGetNetworkSmDeviceSecurityCentersOK() *GetNetworkSmDeviceSecurityCentersOK {
	return &GetNetworkSmDeviceSecurityCentersOK{}
}

/* GetNetworkSmDeviceSecurityCentersOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmDeviceSecurityCentersOK struct {
	Payload []interface{}
}

func (o *GetNetworkSmDeviceSecurityCentersOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/securityCenters][%d] getNetworkSmDeviceSecurityCentersOK  %+v", 200, o.Payload)
}
func (o *GetNetworkSmDeviceSecurityCentersOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkSmDeviceSecurityCentersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
