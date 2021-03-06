// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloneOrganizationReader is a Reader for the CloneOrganization structure.
type CloneOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCloneOrganizationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneOrganizationCreated creates a CloneOrganizationCreated with default headers values
func NewCloneOrganizationCreated() *CloneOrganizationCreated {
	return &CloneOrganizationCreated{}
}

/* CloneOrganizationCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CloneOrganizationCreated struct {
	Payload interface{}
}

func (o *CloneOrganizationCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/clone][%d] cloneOrganizationCreated  %+v", 201, o.Payload)
}
func (o *CloneOrganizationCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CloneOrganizationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CloneOrganizationBody clone organization body
// Example: {"name":"My organization"}
swagger:model CloneOrganizationBody
*/
type CloneOrganizationBody struct {

	// The name of the new organization
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this clone organization body
func (o *CloneOrganizationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CloneOrganizationBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cloneOrganization"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this clone organization body based on context it is used
func (o *CloneOrganizationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloneOrganizationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloneOrganizationBody) UnmarshalBinary(b []byte) error {
	var res CloneOrganizationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
