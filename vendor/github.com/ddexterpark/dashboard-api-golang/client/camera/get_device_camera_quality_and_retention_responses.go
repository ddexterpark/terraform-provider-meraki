// Code generated by go-swagger; DO NOT EDIT.

package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceCameraQualityAndRetentionReader is a Reader for the GetDeviceCameraQualityAndRetention structure.
type GetDeviceCameraQualityAndRetentionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCameraQualityAndRetentionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCameraQualityAndRetentionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceCameraQualityAndRetentionOK creates a GetDeviceCameraQualityAndRetentionOK with default headers values
func NewGetDeviceCameraQualityAndRetentionOK() *GetDeviceCameraQualityAndRetentionOK {
	return &GetDeviceCameraQualityAndRetentionOK{}
}

/* GetDeviceCameraQualityAndRetentionOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceCameraQualityAndRetentionOK struct {
	Payload interface{}
}

func (o *GetDeviceCameraQualityAndRetentionOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/qualityAndRetention][%d] getDeviceCameraQualityAndRetentionOK  %+v", 200, o.Payload)
}
func (o *GetDeviceCameraQualityAndRetentionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceCameraQualityAndRetentionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
