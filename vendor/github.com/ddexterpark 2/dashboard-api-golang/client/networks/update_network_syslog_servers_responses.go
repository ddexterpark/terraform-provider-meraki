// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// UpdateNetworkSyslogServersReader is a Reader for the UpdateNetworkSyslogServers structure.
type UpdateNetworkSyslogServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSyslogServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSyslogServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSyslogServersOK creates a UpdateNetworkSyslogServersOK with default headers values
func NewUpdateNetworkSyslogServersOK() *UpdateNetworkSyslogServersOK {
	return &UpdateNetworkSyslogServersOK{}
}

/* UpdateNetworkSyslogServersOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSyslogServersOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSyslogServersOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/syslogServers][%d] updateNetworkSyslogServersOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkSyslogServersOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSyslogServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkSyslogServersBody update network syslog servers body
// Example: {"servers":[{"host":"1.2.3.4","port":443,"roles":["Wireless event log","URLs"]}]}
swagger:model UpdateNetworkSyslogServersBody
*/
type UpdateNetworkSyslogServersBody struct {

	// A list of the syslog servers for this network
	// Required: true
	Servers []*UpdateNetworkSyslogServersParamsBodyServersItems0 `json:"servers"`
}

// Validate validates this update network syslog servers body
func (o *UpdateNetworkSyslogServersBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSyslogServersBody) validateServers(formats strfmt.Registry) error {

	if err := validate.Required("updateNetworkSyslogServers"+"."+"servers", "body", o.Servers); err != nil {
		return err
	}

	for i := 0; i < len(o.Servers); i++ {
		if swag.IsZero(o.Servers[i]) { // not required
			continue
		}

		if o.Servers[i] != nil {
			if err := o.Servers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSyslogServers" + "." + "servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSyslogServers" + "." + "servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network syslog servers body based on the context it is used
func (o *UpdateNetworkSyslogServersBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSyslogServersBody) contextValidateServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Servers); i++ {

		if o.Servers[i] != nil {
			if err := o.Servers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSyslogServers" + "." + "servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSyslogServers" + "." + "servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSyslogServersBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSyslogServersBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSyslogServersBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkSyslogServersParamsBodyServersItems0 update network syslog servers params body servers items0
swagger:model UpdateNetworkSyslogServersParamsBodyServersItems0
*/
type UpdateNetworkSyslogServersParamsBodyServersItems0 struct {

	// The IP address of the syslog server
	// Required: true
	Host *string `json:"host"`

	// The port of the syslog server
	// Required: true
	Port *int64 `json:"port"`

	// A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'
	// Required: true
	Roles []string `json:"roles"`
}

// Validate validates this update network syslog servers params body servers items0
func (o *UpdateNetworkSyslogServersParamsBodyServersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSyslogServersParamsBodyServersItems0) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", o.Host); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSyslogServersParamsBodyServersItems0) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", o.Port); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSyslogServersParamsBodyServersItems0) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", o.Roles); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network syslog servers params body servers items0 based on context it is used
func (o *UpdateNetworkSyslogServersParamsBodyServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSyslogServersParamsBodyServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSyslogServersParamsBodyServersItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSyslogServersParamsBodyServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}