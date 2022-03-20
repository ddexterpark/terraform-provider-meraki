// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkGroupPolicyReader is a Reader for the DeleteNetworkGroupPolicy structure.
type DeleteNetworkGroupPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkGroupPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkGroupPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkGroupPolicyNoContent creates a DeleteNetworkGroupPolicyNoContent with default headers values
func NewDeleteNetworkGroupPolicyNoContent() *DeleteNetworkGroupPolicyNoContent {
	return &DeleteNetworkGroupPolicyNoContent{}
}

/* DeleteNetworkGroupPolicyNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkGroupPolicyNoContent struct {
}

func (o *DeleteNetworkGroupPolicyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/groupPolicies/{groupPolicyId}][%d] deleteNetworkGroupPolicyNoContent ", 204)
}

func (o *DeleteNetworkGroupPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}