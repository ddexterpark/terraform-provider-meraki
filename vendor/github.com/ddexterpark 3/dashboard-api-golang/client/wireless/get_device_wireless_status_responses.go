// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceWirelessStatusReader is a Reader for the GetDeviceWirelessStatus structure.
type GetDeviceWirelessStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceWirelessStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceWirelessStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceWirelessStatusOK creates a GetDeviceWirelessStatusOK with default headers values
func NewGetDeviceWirelessStatusOK() *GetDeviceWirelessStatusOK {
	return &GetDeviceWirelessStatusOK{}
}

/* GetDeviceWirelessStatusOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceWirelessStatusOK struct {
	Payload interface{}
}

func (o *GetDeviceWirelessStatusOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/wireless/status][%d] getDeviceWirelessStatusOK  %+v", 200, o.Payload)
}
func (o *GetDeviceWirelessStatusOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceWirelessStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
