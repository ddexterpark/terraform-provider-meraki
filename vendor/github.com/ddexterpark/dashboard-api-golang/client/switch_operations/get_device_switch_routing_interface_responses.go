// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceSwitchRoutingInterfaceReader is a Reader for the GetDeviceSwitchRoutingInterface structure.
type GetDeviceSwitchRoutingInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceSwitchRoutingInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceSwitchRoutingInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceSwitchRoutingInterfaceOK creates a GetDeviceSwitchRoutingInterfaceOK with default headers values
func NewGetDeviceSwitchRoutingInterfaceOK() *GetDeviceSwitchRoutingInterfaceOK {
	return &GetDeviceSwitchRoutingInterfaceOK{}
}

/* GetDeviceSwitchRoutingInterfaceOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceSwitchRoutingInterfaceOK struct {
	Payload interface{}
}

func (o *GetDeviceSwitchRoutingInterfaceOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/routing/interfaces/{interfaceId}][%d] getDeviceSwitchRoutingInterfaceOK  %+v", 200, o.Payload)
}
func (o *GetDeviceSwitchRoutingInterfaceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceSwitchRoutingInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
