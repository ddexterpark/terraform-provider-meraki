// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// UpdateNetworkSwitchAccessControlListsReader is a Reader for the UpdateNetworkSwitchAccessControlLists structure.
type UpdateNetworkSwitchAccessControlListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchAccessControlListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchAccessControlListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSwitchAccessControlListsOK creates a UpdateNetworkSwitchAccessControlListsOK with default headers values
func NewUpdateNetworkSwitchAccessControlListsOK() *UpdateNetworkSwitchAccessControlListsOK {
	return &UpdateNetworkSwitchAccessControlListsOK{}
}

/* UpdateNetworkSwitchAccessControlListsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchAccessControlListsOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSwitchAccessControlListsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/accessControlLists][%d] updateNetworkSwitchAccessControlListsOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkSwitchAccessControlListsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchAccessControlListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkSwitchAccessControlListsBody update network switch access control lists body
// Example: {"rules":[{"comment":"Deny SSH","dstCidr":"172.16.30/24","dstPort":"22","ipVersion":"ipv4","policy":"deny","protocol":"tcp","srcCidr":"10.1.10.0/24","srcPort":"any","vlan":"10"}]}
swagger:model UpdateNetworkSwitchAccessControlListsBody
*/
type UpdateNetworkSwitchAccessControlListsBody struct {

	// An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
	// Required: true
	Rules []*UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0 `json:"rules"`
}

// Validate validates this update network switch access control lists body
func (o *UpdateNetworkSwitchAccessControlListsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessControlListsBody) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("updateNetworkSwitchAccessControlLists"+"."+"rules", "body", o.Rules); err != nil {
		return err
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchAccessControlLists" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchAccessControlLists" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network switch access control lists body based on the context it is used
func (o *UpdateNetworkSwitchAccessControlListsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessControlListsBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {
			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchAccessControlLists" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchAccessControlLists" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessControlListsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessControlListsBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAccessControlListsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0 update network switch access control lists params body rules items0
swagger:model UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0
*/
type UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0 struct {

	// Description of the rule (optional).
	Comment string `json:"comment,omitempty"`

	// Destination IP address (in IP or CIDR notation) or 'any'.
	// Required: true
	DstCidr *string `json:"dstCidr"`

	// Destination port. Must be in the range of 1-65535 or 'any'. Default is 'any'.
	DstPort string `json:"dstPort,omitempty"`

	// IP address version (must be 'any', 'ipv4' or 'ipv6'). Applicable only if network supports IPv6. Default value is 'ipv4'.
	// Enum: [any ipv4 ipv6]
	IPVersion string `json:"ipVersion,omitempty"`

	// 'allow' or 'deny' traffic specified by this rule.
	// Required: true
	// Enum: [allow deny]
	Policy *string `json:"policy"`

	// The type of protocol (must be 'tcp', 'udp', or 'any').
	// Required: true
	// Enum: [tcp udp any]
	Protocol *string `json:"protocol"`

	// Source IP address (in IP or CIDR notation) or 'any'.
	// Required: true
	SrcCidr *string `json:"srcCidr"`

	// Source port. Must be in the range  of 1-65535 or 'any'. Default is 'any'.
	SrcPort string `json:"srcPort,omitempty"`

	// Incoming traffic VLAN. Must be in the range of 1-4095 or 'any'. Default is 'any'.
	Vlan string `json:"vlan,omitempty"`
}

// Validate validates this update network switch access control lists params body rules items0
func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDstCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSrcCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) validateDstCidr(formats strfmt.Registry) error {

	if err := validate.Required("dstCidr", "body", o.DstCidr); err != nil {
		return err
	}

	return nil
}

var updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypeIPVersionPropEnum = append(updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypeIPVersionPropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0IPVersionAny captures enum value "any"
	UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0IPVersionAny string = "any"

	// UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0IPVersionIPV4 captures enum value "ipv4"
	UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0IPVersionIPV4 string = "ipv4"

	// UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0IPVersionIPV6 captures enum value "ipv6"
	UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0IPVersionIPV6 string = "ipv6"
)

// prop value enum
func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypeIPVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) validateIPVersion(formats strfmt.Registry) error {
	if swag.IsZero(o.IPVersion) { // not required
		return nil
	}

	// value enum
	if err := o.validateIPVersionEnum("ipVersion", "body", o.IPVersion); err != nil {
		return err
	}

	return nil
}

var updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypePolicyPropEnum = append(updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypePolicyPropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0PolicyAllow captures enum value "allow"
	UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0PolicyAllow string = "allow"

	// UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0PolicyDeny captures enum value "deny"
	UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0PolicyDeny string = "deny"
)

// prop value enum
func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) validatePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", o.Policy); err != nil {
		return err
	}

	// value enum
	if err := o.validatePolicyEnum("policy", "body", *o.Policy); err != nil {
		return err
	}

	return nil
}

var updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypeProtocolPropEnum = append(updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypeProtocolPropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0ProtocolTCP captures enum value "tcp"
	UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0ProtocolTCP string = "tcp"

	// UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0ProtocolUDP captures enum value "udp"
	UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0ProtocolUDP string = "udp"

	// UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0ProtocolAny captures enum value "any"
	UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0ProtocolAny string = "any"
)

// prop value enum
func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchAccessControlListsParamsBodyRulesItems0TypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", o.Protocol); err != nil {
		return err
	}

	// value enum
	if err := o.validateProtocolEnum("protocol", "body", *o.Protocol); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) validateSrcCidr(formats strfmt.Registry) error {

	if err := validate.Required("srcCidr", "body", o.SrcCidr); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch access control lists params body rules items0 based on context it is used
func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAccessControlListsParamsBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}