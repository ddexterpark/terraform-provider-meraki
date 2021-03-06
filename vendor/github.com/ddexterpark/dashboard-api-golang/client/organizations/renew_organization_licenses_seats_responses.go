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

// RenewOrganizationLicensesSeatsReader is a Reader for the RenewOrganizationLicensesSeats structure.
type RenewOrganizationLicensesSeatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenewOrganizationLicensesSeatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenewOrganizationLicensesSeatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRenewOrganizationLicensesSeatsOK creates a RenewOrganizationLicensesSeatsOK with default headers values
func NewRenewOrganizationLicensesSeatsOK() *RenewOrganizationLicensesSeatsOK {
	return &RenewOrganizationLicensesSeatsOK{}
}

/* RenewOrganizationLicensesSeatsOK describes a response with status code 200, with default header values.

Successful operation
*/
type RenewOrganizationLicensesSeatsOK struct {
	Payload interface{}
}

func (o *RenewOrganizationLicensesSeatsOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/licenses/renewSeats][%d] renewOrganizationLicensesSeatsOK  %+v", 200, o.Payload)
}
func (o *RenewOrganizationLicensesSeatsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RenewOrganizationLicensesSeatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RenewOrganizationLicensesSeatsBody renew organization licenses seats body
// Example: {"licenseIdToRenew":"123","unusedLicenseId":"1234"}
swagger:model RenewOrganizationLicensesSeatsBody
*/
type RenewOrganizationLicensesSeatsBody struct {

	// The ID of the SM license to renew. This license must already be assigned to an SM network
	// Required: true
	LicenseIDToRenew *string `json:"licenseIdToRenew"`

	// The SM license to use to renew the seats on 'licenseIdToRenew'. This license must have at least as many seats available as there are seats on 'licenseIdToRenew'
	// Required: true
	UnusedLicenseID *string `json:"unusedLicenseId"`
}

// Validate validates this renew organization licenses seats body
func (o *RenewOrganizationLicensesSeatsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLicenseIDToRenew(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnusedLicenseID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RenewOrganizationLicensesSeatsBody) validateLicenseIDToRenew(formats strfmt.Registry) error {

	if err := validate.Required("renewOrganizationLicensesSeats"+"."+"licenseIdToRenew", "body", o.LicenseIDToRenew); err != nil {
		return err
	}

	return nil
}

func (o *RenewOrganizationLicensesSeatsBody) validateUnusedLicenseID(formats strfmt.Registry) error {

	if err := validate.Required("renewOrganizationLicensesSeats"+"."+"unusedLicenseId", "body", o.UnusedLicenseID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this renew organization licenses seats body based on context it is used
func (o *RenewOrganizationLicensesSeatsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RenewOrganizationLicensesSeatsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RenewOrganizationLicensesSeatsBody) UnmarshalBinary(b []byte) error {
	var res RenewOrganizationLicensesSeatsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
