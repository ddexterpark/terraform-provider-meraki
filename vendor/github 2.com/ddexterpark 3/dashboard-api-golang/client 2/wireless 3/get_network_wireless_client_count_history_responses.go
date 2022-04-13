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

// GetNetworkWirelessClientCountHistoryReader is a Reader for the GetNetworkWirelessClientCountHistory structure.
type GetNetworkWirelessClientCountHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessClientCountHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessClientCountHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessClientCountHistoryOK creates a GetNetworkWirelessClientCountHistoryOK with default headers values
func NewGetNetworkWirelessClientCountHistoryOK() *GetNetworkWirelessClientCountHistoryOK {
	return &GetNetworkWirelessClientCountHistoryOK{}
}

/* GetNetworkWirelessClientCountHistoryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessClientCountHistoryOK struct {
	Payload []interface{}
}

func (o *GetNetworkWirelessClientCountHistoryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/clientCountHistory][%d] getNetworkWirelessClientCountHistoryOK  %+v", 200, o.Payload)
}
func (o *GetNetworkWirelessClientCountHistoryOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkWirelessClientCountHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
