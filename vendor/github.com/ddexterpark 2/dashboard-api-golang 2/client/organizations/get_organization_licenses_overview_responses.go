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

// GetOrganizationLicensesOverviewReader is a Reader for the GetOrganizationLicensesOverview structure.
type GetOrganizationLicensesOverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationLicensesOverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationLicensesOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationLicensesOverviewOK creates a GetOrganizationLicensesOverviewOK with default headers values
func NewGetOrganizationLicensesOverviewOK() *GetOrganizationLicensesOverviewOK {
	return &GetOrganizationLicensesOverviewOK{}
}

/* GetOrganizationLicensesOverviewOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationLicensesOverviewOK struct {
	Payload interface{}
}

func (o *GetOrganizationLicensesOverviewOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/licenses/overview][%d] getOrganizationLicensesOverviewOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationLicensesOverviewOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationLicensesOverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}