// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// UpdateNetworkSnmpReader is a Reader for the UpdateNetworkSnmp structure.
type UpdateNetworkSnmpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSnmpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSnmpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSnmpOK creates a UpdateNetworkSnmpOK with default headers values
func NewUpdateNetworkSnmpOK() *UpdateNetworkSnmpOK {
	return &UpdateNetworkSnmpOK{}
}

/* UpdateNetworkSnmpOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSnmpOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSnmpOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/snmp][%d] updateNetworkSnmpOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkSnmpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSnmpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkSnmpBody update network snmp body
// Example: {"access":"users","users":[{"passphrase":"hunter2","username":"AzureDiamond"}]}
swagger:model UpdateNetworkSnmpBody
*/
type UpdateNetworkSnmpBody struct {

	// The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
	// Enum: [none community users]
	Access string `json:"access,omitempty"`

	// The SNMP community string. Only relevant if 'access' is set to 'community'.
	CommunityString string `json:"communityString,omitempty"`

	// The list of SNMP users. Only relevant if 'access' is set to 'users'.
	Users []*UpdateNetworkSnmpParamsBodyUsersItems0 `json:"users"`
}

// Validate validates this update network snmp body
func (o *UpdateNetworkSnmpBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkSnmpBodyTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","community","users"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSnmpBodyTypeAccessPropEnum = append(updateNetworkSnmpBodyTypeAccessPropEnum, v)
	}
}

const (

	// UpdateNetworkSnmpBodyAccessNone captures enum value "none"
	UpdateNetworkSnmpBodyAccessNone string = "none"

	// UpdateNetworkSnmpBodyAccessCommunity captures enum value "community"
	UpdateNetworkSnmpBodyAccessCommunity string = "community"

	// UpdateNetworkSnmpBodyAccessUsers captures enum value "users"
	UpdateNetworkSnmpBodyAccessUsers string = "users"
)

// prop value enum
func (o *UpdateNetworkSnmpBody) validateAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSnmpBodyTypeAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSnmpBody) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(o.Access) { // not required
		return nil
	}

	// value enum
	if err := o.validateAccessEnum("updateNetworkSnmp"+"."+"access", "body", o.Access); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSnmpBody) validateUsers(formats strfmt.Registry) error {
	if swag.IsZero(o.Users) { // not required
		return nil
	}

	for i := 0; i < len(o.Users); i++ {
		if swag.IsZero(o.Users[i]) { // not required
			continue
		}

		if o.Users[i] != nil {
			if err := o.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSnmp" + "." + "users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSnmp" + "." + "users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network snmp body based on the context it is used
func (o *UpdateNetworkSnmpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSnmpBody) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Users); i++ {

		if o.Users[i] != nil {
			if err := o.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSnmp" + "." + "users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSnmp" + "." + "users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSnmpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSnmpBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSnmpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSnmpParamsBodyUsersItems0 update network snmp params body users items0
swagger:model UpdateNetworkSnmpParamsBodyUsersItems0
*/
type UpdateNetworkSnmpParamsBodyUsersItems0 struct {

	// The passphrase for the SNMP user. Required.
	// Required: true
	Passphrase *string `json:"passphrase"`

	// The username for the SNMP user. Required.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this update network snmp params body users items0
func (o *UpdateNetworkSnmpParamsBodyUsersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePassphrase(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSnmpParamsBodyUsersItems0) validatePassphrase(formats strfmt.Registry) error {

	if err := validate.Required("passphrase", "body", o.Passphrase); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSnmpParamsBodyUsersItems0) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network snmp params body users items0 based on context it is used
func (o *UpdateNetworkSnmpParamsBodyUsersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSnmpParamsBodyUsersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSnmpParamsBodyUsersItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSnmpParamsBodyUsersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
