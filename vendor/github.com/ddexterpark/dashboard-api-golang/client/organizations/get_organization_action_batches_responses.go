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

// GetOrganizationActionBatchesReader is a Reader for the GetOrganizationActionBatches structure.
type GetOrganizationActionBatchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationActionBatchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationActionBatchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationActionBatchesOK creates a GetOrganizationActionBatchesOK with default headers values
func NewGetOrganizationActionBatchesOK() *GetOrganizationActionBatchesOK {
	return &GetOrganizationActionBatchesOK{}
}

/* GetOrganizationActionBatchesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationActionBatchesOK struct {
	Payload []interface{}
}

func (o *GetOrganizationActionBatchesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/actionBatches][%d] getOrganizationActionBatchesOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationActionBatchesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationActionBatchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
