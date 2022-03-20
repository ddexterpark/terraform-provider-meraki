// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkSwitchRoutingMulticastReader is a Reader for the UpdateNetworkSwitchRoutingMulticast structure.
type UpdateNetworkSwitchRoutingMulticastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchRoutingMulticastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchRoutingMulticastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSwitchRoutingMulticastOK creates a UpdateNetworkSwitchRoutingMulticastOK with default headers values
func NewUpdateNetworkSwitchRoutingMulticastOK() *UpdateNetworkSwitchRoutingMulticastOK {
	return &UpdateNetworkSwitchRoutingMulticastOK{}
}

/* UpdateNetworkSwitchRoutingMulticastOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchRoutingMulticastOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSwitchRoutingMulticastOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/routing/multicast][%d] updateNetworkSwitchRoutingMulticastOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkSwitchRoutingMulticastOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchRoutingMulticastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkSwitchRoutingMulticastBody update network switch routing multicast body
// Example: {"defaultSettings":{"floodUnknownMulticastTrafficEnabled":true,"igmpSnoopingEnabled":true},"overrides":[{"floodUnknownMulticastTrafficEnabled":true,"igmpSnoopingEnabled":true,"switches":["Q234-ABCD-0001","Q234-ABCD-0002","Q234-ABCD-0003"]},{"floodUnknownMulticastTrafficEnabled":true,"igmpSnoopingEnabled":true,"stacks":["789102","123456","129102"]}]}
swagger:model UpdateNetworkSwitchRoutingMulticastBody
*/
type UpdateNetworkSwitchRoutingMulticastBody struct {

	// default settings
	DefaultSettings *UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings `json:"defaultSettings,omitempty"`

	// Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
	Overrides []*UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0 `json:"overrides"`
}

// Validate validates this update network switch routing multicast body
func (o *UpdateNetworkSwitchRoutingMulticastBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDefaultSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchRoutingMulticastBody) validateDefaultSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.DefaultSettings) { // not required
		return nil
	}

	if o.DefaultSettings != nil {
		if err := o.DefaultSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchRoutingMulticast" + "." + "defaultSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchRoutingMulticast" + "." + "defaultSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchRoutingMulticastBody) validateOverrides(formats strfmt.Registry) error {
	if swag.IsZero(o.Overrides) { // not required
		return nil
	}

	for i := 0; i < len(o.Overrides); i++ {
		if swag.IsZero(o.Overrides[i]) { // not required
			continue
		}

		if o.Overrides[i] != nil {
			if err := o.Overrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchRoutingMulticast" + "." + "overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchRoutingMulticast" + "." + "overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network switch routing multicast body based on the context it is used
func (o *UpdateNetworkSwitchRoutingMulticastBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDefaultSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchRoutingMulticastBody) contextValidateDefaultSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.DefaultSettings != nil {
		if err := o.DefaultSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchRoutingMulticast" + "." + "defaultSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchRoutingMulticast" + "." + "defaultSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchRoutingMulticastBody) contextValidateOverrides(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Overrides); i++ {

		if o.Overrides[i] != nil {
			if err := o.Overrides[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchRoutingMulticast" + "." + "overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchRoutingMulticast" + "." + "overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchRoutingMulticastBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchRoutingMulticastBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchRoutingMulticastBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings Default multicast setting for entire network. IGMP snooping and Flood unknown multicast traffic settings are enabled by default.
swagger:model UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings
*/
type UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings struct {

	// Flood unknown multicast traffic setting for entire network
	FloodUnknownMulticastTrafficEnabled bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"`

	// IGMP snooping setting for entire network
	IgmpSnoopingEnabled bool `json:"igmpSnoopingEnabled,omitempty"`
}

// Validate validates this update network switch routing multicast params body default settings
func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch routing multicast params body default settings based on context it is used
func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchRoutingMulticastParamsBodyDefaultSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0 update network switch routing multicast params body overrides items0
swagger:model UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0
*/
type UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0 struct {

	// Flood unknown multicast traffic setting for switches, switch stacks or switch profiles
	// Required: true
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled"`

	// IGMP snooping setting for switches, switch stacks or switch profiles
	// Required: true
	IgmpSnoopingEnabled *bool `json:"igmpSnoopingEnabled"`

	// List of switch stack ids for non-template network
	Stacks []string `json:"stacks"`

	// List of switch profiles ids for template network
	SwitchProfiles []string `json:"switchProfiles"`

	// List of switch serials for non-template network
	Switches []string `json:"switches"`
}

// Validate validates this update network switch routing multicast params body overrides items0
func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFloodUnknownMulticastTrafficEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIgmpSnoopingEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0) validateFloodUnknownMulticastTrafficEnabled(formats strfmt.Registry) error {

	if err := validate.Required("floodUnknownMulticastTrafficEnabled", "body", o.FloodUnknownMulticastTrafficEnabled); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0) validateIgmpSnoopingEnabled(formats strfmt.Registry) error {

	if err := validate.Required("igmpSnoopingEnabled", "body", o.IgmpSnoopingEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch routing multicast params body overrides items0 based on context it is used
func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchRoutingMulticastParamsBodyOverridesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}