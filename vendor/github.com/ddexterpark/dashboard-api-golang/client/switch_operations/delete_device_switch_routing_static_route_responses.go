// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDeviceSwitchRoutingStaticRouteReader is a Reader for the DeleteDeviceSwitchRoutingStaticRoute structure.
type DeleteDeviceSwitchRoutingStaticRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceSwitchRoutingStaticRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDeviceSwitchRoutingStaticRouteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDeviceSwitchRoutingStaticRouteNoContent creates a DeleteDeviceSwitchRoutingStaticRouteNoContent with default headers values
func NewDeleteDeviceSwitchRoutingStaticRouteNoContent() *DeleteDeviceSwitchRoutingStaticRouteNoContent {
	return &DeleteDeviceSwitchRoutingStaticRouteNoContent{}
}

/* DeleteDeviceSwitchRoutingStaticRouteNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteDeviceSwitchRoutingStaticRouteNoContent struct {
}

func (o *DeleteDeviceSwitchRoutingStaticRouteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /devices/{serial}/switch/routing/staticRoutes/{staticRouteId}][%d] deleteDeviceSwitchRoutingStaticRouteNoContent ", 204)
}

func (o *DeleteDeviceSwitchRoutingStaticRouteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
