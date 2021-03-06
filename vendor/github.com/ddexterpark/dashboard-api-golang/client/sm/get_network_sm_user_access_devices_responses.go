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

// GetNetworkSmUserAccessDevicesReader is a Reader for the GetNetworkSmUserAccessDevices structure.
type GetNetworkSmUserAccessDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmUserAccessDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmUserAccessDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmUserAccessDevicesOK creates a GetNetworkSmUserAccessDevicesOK with default headers values
func NewGetNetworkSmUserAccessDevicesOK() *GetNetworkSmUserAccessDevicesOK {
	return &GetNetworkSmUserAccessDevicesOK{}
}

/* GetNetworkSmUserAccessDevicesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmUserAccessDevicesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

func (o *GetNetworkSmUserAccessDevicesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/userAccessDevices][%d] getNetworkSmUserAccessDevicesOK  %+v", 200, o.Payload)
}
func (o *GetNetworkSmUserAccessDevicesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkSmUserAccessDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
