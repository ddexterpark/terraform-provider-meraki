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

// UpdateNetworkSwitchAccessPolicyReader is a Reader for the UpdateNetworkSwitchAccessPolicy structure.
type UpdateNetworkSwitchAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchAccessPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSwitchAccessPolicyOK creates a UpdateNetworkSwitchAccessPolicyOK with default headers values
func NewUpdateNetworkSwitchAccessPolicyOK() *UpdateNetworkSwitchAccessPolicyOK {
	return &UpdateNetworkSwitchAccessPolicyOK{}
}

/* UpdateNetworkSwitchAccessPolicyOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchAccessPolicyOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSwitchAccessPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}][%d] updateNetworkSwitchAccessPolicyOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkSwitchAccessPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchAccessPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkSwitchAccessPolicyBody update network switch access policy body
// Example: {"accessPolicyType":"Hybrid authentication","guestVlanId":100,"hostMode":"Single-Host","increaseAccessSpeed":false,"name":"Access policy #1","radius":{"criticalAuth":{"dataVlanId":100,"suspendPortBounce":true,"voiceVlanId":100},"failedAuthVlanId":100,"reAuthenticationInterval":120,"suspendReAuthentication":true},"radiusAccountingEnabled":true,"radiusAccountingServers":[{"host":"1.2.3.4","port":22,"secret":"secret"}],"radiusCoaSupportEnabled":false,"radiusGroupAttribute":"11","radiusServers":[{"host":"1.2.3.4","port":22,"secret":"secret"}],"radiusTestingEnabled":false,"urlRedirectWalledGardenEnabled":true,"urlRedirectWalledGardenRanges":["192.168.1.0/24"],"voiceVlanClients":true}
swagger:model UpdateNetworkSwitchAccessPolicyBody
*/
type UpdateNetworkSwitchAccessPolicyBody struct {

	// Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	// Enum: [802.1x MAC authentication bypass Hybrid authentication]
	AccessPolicyType string `json:"accessPolicyType,omitempty"`

	// ID for the guest VLAN allow unauthorized devices access to limited network resources
	GuestVlanID int64 `json:"guestVlanId,omitempty"`

	// Choose the Host Mode for the access policy.
	// Enum: [Single-Host Multi-Domain Multi-Host Multi-Auth]
	HostMode string `json:"hostMode,omitempty"`

	// Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	IncreaseAccessSpeed bool `json:"increaseAccessSpeed,omitempty"`

	// Name of the access policy
	Name string `json:"name,omitempty"`

	// radius
	Radius *UpdateNetworkSwitchAccessPolicyParamsBodyRadius `json:"radius,omitempty"`

	// Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingEnabled bool `json:"radiusAccountingEnabled,omitempty"`

	// List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusAccountingServers []*UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0 `json:"radiusAccountingServers"`

	// Change of authentication for RADIUS re-authentication and disconnection
	RadiusCoaSupportEnabled bool `json:"radiusCoaSupportEnabled,omitempty"`

	// Can be either `""`, which means `None` on Dashboard, or `"11"`, which means `Filter-Id` on Dashboard and will use Group Policy ACLs when supported (firmware 14+)
	RadiusGroupAttribute string `json:"radiusGroupAttribute,omitempty"`

	// List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusServers []*UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0 `json:"radiusServers"`

	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	RadiusTestingEnabled bool `json:"radiusTestingEnabled,omitempty"`

	// Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenEnabled bool `json:"urlRedirectWalledGardenEnabled,omitempty"`

	// IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges []string `json:"urlRedirectWalledGardenRanges"`

	// CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
	VoiceVlanClients bool `json:"voiceVlanClients,omitempty"`
}

// Validate validates this update network switch access policy body
func (o *UpdateNetworkSwitchAccessPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessPolicyType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHostMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadius(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusAccountingServers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkSwitchAccessPolicyBodyTypeAccessPolicyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["802.1x","MAC authentication bypass","Hybrid authentication"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchAccessPolicyBodyTypeAccessPolicyTypePropEnum = append(updateNetworkSwitchAccessPolicyBodyTypeAccessPolicyTypePropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchAccessPolicyBodyAccessPolicyTypeNr802Dot1x captures enum value "802.1x"
	UpdateNetworkSwitchAccessPolicyBodyAccessPolicyTypeNr802Dot1x string = "802.1x"

	// UpdateNetworkSwitchAccessPolicyBodyAccessPolicyTypeMACAuthenticationBypass captures enum value "MAC authentication bypass"
	UpdateNetworkSwitchAccessPolicyBodyAccessPolicyTypeMACAuthenticationBypass string = "MAC authentication bypass"

	// UpdateNetworkSwitchAccessPolicyBodyAccessPolicyTypeHybridAuthentication captures enum value "Hybrid authentication"
	UpdateNetworkSwitchAccessPolicyBodyAccessPolicyTypeHybridAuthentication string = "Hybrid authentication"
)

// prop value enum
func (o *UpdateNetworkSwitchAccessPolicyBody) validateAccessPolicyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchAccessPolicyBodyTypeAccessPolicyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyBody) validateAccessPolicyType(formats strfmt.Registry) error {
	if swag.IsZero(o.AccessPolicyType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAccessPolicyTypeEnum("updateNetworkSwitchAccessPolicy"+"."+"accessPolicyType", "body", o.AccessPolicyType); err != nil {
		return err
	}

	return nil
}

var updateNetworkSwitchAccessPolicyBodyTypeHostModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Single-Host","Multi-Domain","Multi-Host","Multi-Auth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchAccessPolicyBodyTypeHostModePropEnum = append(updateNetworkSwitchAccessPolicyBodyTypeHostModePropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchAccessPolicyBodyHostModeSingleDashHost captures enum value "Single-Host"
	UpdateNetworkSwitchAccessPolicyBodyHostModeSingleDashHost string = "Single-Host"

	// UpdateNetworkSwitchAccessPolicyBodyHostModeMultiDashDomain captures enum value "Multi-Domain"
	UpdateNetworkSwitchAccessPolicyBodyHostModeMultiDashDomain string = "Multi-Domain"

	// UpdateNetworkSwitchAccessPolicyBodyHostModeMultiDashHost captures enum value "Multi-Host"
	UpdateNetworkSwitchAccessPolicyBodyHostModeMultiDashHost string = "Multi-Host"

	// UpdateNetworkSwitchAccessPolicyBodyHostModeMultiDashAuth captures enum value "Multi-Auth"
	UpdateNetworkSwitchAccessPolicyBodyHostModeMultiDashAuth string = "Multi-Auth"
)

// prop value enum
func (o *UpdateNetworkSwitchAccessPolicyBody) validateHostModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchAccessPolicyBodyTypeHostModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyBody) validateHostMode(formats strfmt.Registry) error {
	if swag.IsZero(o.HostMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateHostModeEnum("updateNetworkSwitchAccessPolicy"+"."+"hostMode", "body", o.HostMode); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyBody) validateRadius(formats strfmt.Registry) error {
	if swag.IsZero(o.Radius) { // not required
		return nil
	}

	if o.Radius != nil {
		if err := o.Radius.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radius")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radius")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyBody) validateRadiusAccountingServers(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusAccountingServers) { // not required
		return nil
	}

	for i := 0; i < len(o.RadiusAccountingServers); i++ {
		if swag.IsZero(o.RadiusAccountingServers[i]) { // not required
			continue
		}

		if o.RadiusAccountingServers[i] != nil {
			if err := o.RadiusAccountingServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyBody) validateRadiusServers(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusServers) { // not required
		return nil
	}

	for i := 0; i < len(o.RadiusServers); i++ {
		if swag.IsZero(o.RadiusServers[i]) { // not required
			continue
		}

		if o.RadiusServers[i] != nil {
			if err := o.RadiusServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network switch access policy body based on the context it is used
func (o *UpdateNetworkSwitchAccessPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRadius(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRadiusAccountingServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRadiusServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyBody) contextValidateRadius(ctx context.Context, formats strfmt.Registry) error {

	if o.Radius != nil {
		if err := o.Radius.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radius")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radius")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyBody) contextValidateRadiusAccountingServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusAccountingServers); i++ {

		if o.RadiusAccountingServers[i] != nil {
			if err := o.RadiusAccountingServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyBody) contextValidateRadiusServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusServers); i++ {

		if o.RadiusServers[i] != nil {
			if err := o.RadiusServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAccessPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSwitchAccessPolicyParamsBodyRadius Object for RADIUS Settings
swagger:model UpdateNetworkSwitchAccessPolicyParamsBodyRadius
*/
type UpdateNetworkSwitchAccessPolicyParamsBodyRadius struct {

	// critical auth
	CriticalAuth *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth `json:"criticalAuth,omitempty"`
}

// Validate validates this update network switch access policy params body radius
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadius) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCriticalAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadius) validateCriticalAuth(formats strfmt.Registry) error {
	if swag.IsZero(o.CriticalAuth) { // not required
		return nil
	}

	if o.CriticalAuth != nil {
		if err := o.CriticalAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radius" + "." + "criticalAuth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radius" + "." + "criticalAuth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network switch access policy params body radius based on the context it is used
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadius) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCriticalAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadius) contextValidateCriticalAuth(ctx context.Context, formats strfmt.Registry) error {

	if o.CriticalAuth != nil {
		if err := o.CriticalAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radius" + "." + "criticalAuth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchAccessPolicy" + "." + "radius" + "." + "criticalAuth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadius) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadius) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAccessPolicyParamsBodyRadius
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0 update network switch access policy params body radius accounting servers items0
swagger:model UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0
*/
type UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0 struct {

	// Public IP address of the RADIUS accounting server
	// Required: true
	Host *string `json:"host"`

	// UDP port that the RADIUS Accounting server listens on for access requests
	// Required: true
	Port *int64 `json:"port"`

	// RADIUS client shared secret
	// Required: true
	Secret *string `json:"secret"`
}

// Validate validates this update network switch access policy params body radius accounting servers items0
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", o.Host); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", o.Port); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", o.Secret); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch access policy params body radius accounting servers items0 based on context it is used
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth Critical auth settings for when authentication is rejected by the RADIUS server
swagger:model UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth
*/
type UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth struct {

	// VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	DataVlanID int64 `json:"dataVlanId,omitempty"`

	// Enable to suspend port bounce when RADIUS servers are unreachable
	SuspendPortBounce bool `json:"suspendPortBounce,omitempty"`

	// VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	VoiceVlanID int64 `json:"voiceVlanId,omitempty"`
}

// Validate validates this update network switch access policy params body radius critical auth
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch access policy params body radius critical auth based on context it is used
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0 update network switch access policy params body radius servers items0
swagger:model UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0
*/
type UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0 struct {

	// Public IP address of the RADIUS server
	// Required: true
	Host *string `json:"host"`

	// UDP port that the RADIUS server listens on for access requests
	// Required: true
	Port *int64 `json:"port"`

	// RADIUS client shared secret
	// Required: true
	Secret *string `json:"secret"`
}

// Validate validates this update network switch access policy params body radius servers items0
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", o.Host); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", o.Port); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", o.Secret); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch access policy params body radius servers items0 based on context it is used
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}