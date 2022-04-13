// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationAdaptivePolicyAclsReader is a Reader for the GetOrganizationAdaptivePolicyAcls structure.
type GetOrganizationAdaptivePolicyAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationAdaptivePolicyAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationAdaptivePolicyAclsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationAdaptivePolicyAclsOK creates a GetOrganizationAdaptivePolicyAclsOK with default headers values
func NewGetOrganizationAdaptivePolicyAclsOK() *GetOrganizationAdaptivePolicyAclsOK {
	return &GetOrganizationAdaptivePolicyAclsOK{}
}

/* GetOrganizationAdaptivePolicyAclsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationAdaptivePolicyAclsOK struct {
	Payload []interface{}
}

func (o *GetOrganizationAdaptivePolicyAclsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/acls][%d] getOrganizationAdaptivePolicyAclsOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationAdaptivePolicyAclsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationAdaptivePolicyAclsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
