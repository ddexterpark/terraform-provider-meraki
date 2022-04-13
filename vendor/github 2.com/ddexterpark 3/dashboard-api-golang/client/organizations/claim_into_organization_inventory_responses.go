// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClaimIntoOrganizationInventoryReader is a Reader for the ClaimIntoOrganizationInventory structure.
type ClaimIntoOrganizationInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClaimIntoOrganizationInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClaimIntoOrganizationInventoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewClaimIntoOrganizationInventoryOK creates a ClaimIntoOrganizationInventoryOK with default headers values
func NewClaimIntoOrganizationInventoryOK() *ClaimIntoOrganizationInventoryOK {
	return &ClaimIntoOrganizationInventoryOK{}
}

/* ClaimIntoOrganizationInventoryOK describes a response with status code 200, with default header values.

Successful operation
*/
type ClaimIntoOrganizationInventoryOK struct {
	Payload interface{}
}

func (o *ClaimIntoOrganizationInventoryOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/inventory/claim][%d] claimIntoOrganizationInventoryOK  %+v", 200, o.Payload)
}
func (o *ClaimIntoOrganizationInventoryOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ClaimIntoOrganizationInventoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ClaimIntoOrganizationInventoryBody claim into organization inventory body
// Example: {"licenses":[{"key":"Z2XXXXXXXXXX","mode":"addDevices"}],"orders":["4CXXXXXXX"],"serials":["Q234-ABCD-5678"]}
swagger:model ClaimIntoOrganizationInventoryBody
*/
type ClaimIntoOrganizationInventoryBody struct {

	// The licenses that should be claimed
	Licenses []*ClaimIntoOrganizationInventoryParamsBodyLicensesItems0 `json:"licenses"`

	// The numbers of the orders that should be claimed
	Orders []string `json:"orders"`

	// The serials of the devices that should be claimed
	Serials []string `json:"serials"`
}

// Validate validates this claim into organization inventory body
func (o *ClaimIntoOrganizationInventoryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLicenses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClaimIntoOrganizationInventoryBody) validateLicenses(formats strfmt.Registry) error {
	if swag.IsZero(o.Licenses) { // not required
		return nil
	}

	for i := 0; i < len(o.Licenses); i++ {
		if swag.IsZero(o.Licenses[i]) { // not required
			continue
		}

		if o.Licenses[i] != nil {
			if err := o.Licenses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claimIntoOrganizationInventory" + "." + "licenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("claimIntoOrganizationInventory" + "." + "licenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this claim into organization inventory body based on the context it is used
func (o *ClaimIntoOrganizationInventoryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClaimIntoOrganizationInventoryBody) contextValidateLicenses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Licenses); i++ {

		if o.Licenses[i] != nil {
			if err := o.Licenses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claimIntoOrganizationInventory" + "." + "licenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("claimIntoOrganizationInventory" + "." + "licenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClaimIntoOrganizationInventoryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClaimIntoOrganizationInventoryBody) UnmarshalBinary(b []byte) error {
	var res ClaimIntoOrganizationInventoryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClaimIntoOrganizationInventoryParamsBodyLicensesItems0 claim into organization inventory params body licenses items0
swagger:model ClaimIntoOrganizationInventoryParamsBodyLicensesItems0
*/
type ClaimIntoOrganizationInventoryParamsBodyLicensesItems0 struct {

	// The key of the license
	// Required: true
	Key *string `json:"key"`

	// Co-term licensing only: either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. Does not apply to organizations using per-device licensing model.
	// Enum: [addDevices renew]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this claim into organization inventory params body licenses items0
func (o *ClaimIntoOrganizationInventoryParamsBodyLicensesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClaimIntoOrganizationInventoryParamsBodyLicensesItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

var claimIntoOrganizationInventoryParamsBodyLicensesItems0TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["addDevices","renew"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		claimIntoOrganizationInventoryParamsBodyLicensesItems0TypeModePropEnum = append(claimIntoOrganizationInventoryParamsBodyLicensesItems0TypeModePropEnum, v)
	}
}

const (

	// ClaimIntoOrganizationInventoryParamsBodyLicensesItems0ModeAddDevices captures enum value "addDevices"
	ClaimIntoOrganizationInventoryParamsBodyLicensesItems0ModeAddDevices string = "addDevices"

	// ClaimIntoOrganizationInventoryParamsBodyLicensesItems0ModeRenew captures enum value "renew"
	ClaimIntoOrganizationInventoryParamsBodyLicensesItems0ModeRenew string = "renew"
)

// prop value enum
func (o *ClaimIntoOrganizationInventoryParamsBodyLicensesItems0) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, claimIntoOrganizationInventoryParamsBodyLicensesItems0TypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ClaimIntoOrganizationInventoryParamsBodyLicensesItems0) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(o.Mode) { // not required
		return nil
	}

	// value enum
	if err := o.validateModeEnum("mode", "body", o.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this claim into organization inventory params body licenses items0 based on context it is used
func (o *ClaimIntoOrganizationInventoryParamsBodyLicensesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClaimIntoOrganizationInventoryParamsBodyLicensesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClaimIntoOrganizationInventoryParamsBodyLicensesItems0) UnmarshalBinary(b []byte) error {
	var res ClaimIntoOrganizationInventoryParamsBodyLicensesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
