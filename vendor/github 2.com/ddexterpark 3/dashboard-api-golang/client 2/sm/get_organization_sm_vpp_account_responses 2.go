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

// GetOrganizationSmVppAccountReader is a Reader for the GetOrganizationSmVppAccount structure.
type GetOrganizationSmVppAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSmVppAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSmVppAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationSmVppAccountOK creates a GetOrganizationSmVppAccountOK with default headers values
func NewGetOrganizationSmVppAccountOK() *GetOrganizationSmVppAccountOK {
	return &GetOrganizationSmVppAccountOK{}
}

/* GetOrganizationSmVppAccountOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationSmVppAccountOK struct {
	Payload interface{}
}

func (o *GetOrganizationSmVppAccountOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/sm/vppAccounts/{vppAccountId}][%d] getOrganizationSmVppAccountOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationSmVppAccountOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationSmVppAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
