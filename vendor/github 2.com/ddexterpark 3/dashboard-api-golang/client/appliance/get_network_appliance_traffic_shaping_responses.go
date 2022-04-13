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

// GetNetworkApplianceTrafficShapingReader is a Reader for the GetNetworkApplianceTrafficShaping structure.
type GetNetworkApplianceTrafficShapingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceTrafficShapingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceTrafficShapingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceTrafficShapingOK creates a GetNetworkApplianceTrafficShapingOK with default headers values
func NewGetNetworkApplianceTrafficShapingOK() *GetNetworkApplianceTrafficShapingOK {
	return &GetNetworkApplianceTrafficShapingOK{}
}

/* GetNetworkApplianceTrafficShapingOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceTrafficShapingOK struct {
	Payload interface{}
}

func (o *GetNetworkApplianceTrafficShapingOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/trafficShaping][%d] getNetworkApplianceTrafficShapingOK  %+v", 200, o.Payload)
}
func (o *GetNetworkApplianceTrafficShapingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceTrafficShapingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
