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

// GetNetworkWirelessDevicesConnectionStatsReader is a Reader for the GetNetworkWirelessDevicesConnectionStats structure.
type GetNetworkWirelessDevicesConnectionStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessDevicesConnectionStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessDevicesConnectionStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessDevicesConnectionStatsOK creates a GetNetworkWirelessDevicesConnectionStatsOK with default headers values
func NewGetNetworkWirelessDevicesConnectionStatsOK() *GetNetworkWirelessDevicesConnectionStatsOK {
	return &GetNetworkWirelessDevicesConnectionStatsOK{}
}

/* GetNetworkWirelessDevicesConnectionStatsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessDevicesConnectionStatsOK struct {
	Payload []interface{}
}

func (o *GetNetworkWirelessDevicesConnectionStatsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/devices/connectionStats][%d] getNetworkWirelessDevicesConnectionStatsOK  %+v", 200, o.Payload)
}
func (o *GetNetworkWirelessDevicesConnectionStatsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkWirelessDevicesConnectionStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}