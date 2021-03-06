// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// CreateDeviceLiveToolsPingDeviceReader is a Reader for the CreateDeviceLiveToolsPingDevice structure.
type CreateDeviceLiveToolsPingDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeviceLiveToolsPingDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeviceLiveToolsPingDeviceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDeviceLiveToolsPingDeviceCreated creates a CreateDeviceLiveToolsPingDeviceCreated with default headers values
func NewCreateDeviceLiveToolsPingDeviceCreated() *CreateDeviceLiveToolsPingDeviceCreated {
	return &CreateDeviceLiveToolsPingDeviceCreated{}
}

/* CreateDeviceLiveToolsPingDeviceCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateDeviceLiveToolsPingDeviceCreated struct {
	Payload interface{}
}

func (o *CreateDeviceLiveToolsPingDeviceCreated) Error() string {
	return fmt.Sprintf("[POST /devices/{serial}/liveTools/pingDevice][%d] createDeviceLiveToolsPingDeviceCreated  %+v", 201, o.Payload)
}
func (o *CreateDeviceLiveToolsPingDeviceCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateDeviceLiveToolsPingDeviceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateDeviceLiveToolsPingDeviceBody create device live tools ping device body
// Example: {"pingId":"1","request":{"count":2,"serial":"Q234-ABCD-5678","target":"75.75.75.75"},"results":{"latencies":{"average":15.8,"maximum":15.9,"minimum":15.8},"loss":{"percentage":0},"received":5,"replies":[{"latency":15.7,"sequenceId":1,"size":84},{"latency":15.9,"sequenceId":0,"size":84}],"sent":5},"status":"complete","url":"/devices/SERIAL/liveTools/ping/1284392014819"}
swagger:model CreateDeviceLiveToolsPingDeviceBody
*/
type CreateDeviceLiveToolsPingDeviceBody struct {

	// Count parameter to pass to ping. [1..5], default 5
	Count int64 `json:"count,omitempty"`
}

// Validate validates this create device live tools ping device body
func (o *CreateDeviceLiveToolsPingDeviceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create device live tools ping device body based on context it is used
func (o *CreateDeviceLiveToolsPingDeviceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateDeviceLiveToolsPingDeviceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateDeviceLiveToolsPingDeviceBody) UnmarshalBinary(b []byte) error {
	var res CreateDeviceLiveToolsPingDeviceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
