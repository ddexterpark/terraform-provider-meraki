// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RefreshNetworkSmDeviceDetailsReader is a Reader for the RefreshNetworkSmDeviceDetails structure.
type RefreshNetworkSmDeviceDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshNetworkSmDeviceDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshNetworkSmDeviceDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefreshNetworkSmDeviceDetailsOK creates a RefreshNetworkSmDeviceDetailsOK with default headers values
func NewRefreshNetworkSmDeviceDetailsOK() *RefreshNetworkSmDeviceDetailsOK {
	return &RefreshNetworkSmDeviceDetailsOK{}
}

/* RefreshNetworkSmDeviceDetailsOK describes a response with status code 200, with default header values.

Successful operation
*/
type RefreshNetworkSmDeviceDetailsOK struct {
}

func (o *RefreshNetworkSmDeviceDetailsOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/devices/{deviceId}/refreshDetails][%d] refreshNetworkSmDeviceDetailsOK ", 200)
}

func (o *RefreshNetworkSmDeviceDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
