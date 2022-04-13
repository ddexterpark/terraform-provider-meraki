// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateOrganizationAlertsProfileReader is a Reader for the CreateOrganizationAlertsProfile structure.
type CreateOrganizationAlertsProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationAlertsProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationAlertsProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrganizationAlertsProfileCreated creates a CreateOrganizationAlertsProfileCreated with default headers values
func NewCreateOrganizationAlertsProfileCreated() *CreateOrganizationAlertsProfileCreated {
	return &CreateOrganizationAlertsProfileCreated{}
}

/* CreateOrganizationAlertsProfileCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationAlertsProfileCreated struct {
	Payload interface{}
}

func (o *CreateOrganizationAlertsProfileCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/alerts/profiles][%d] createOrganizationAlertsProfileCreated  %+v", 201, o.Payload)
}
func (o *CreateOrganizationAlertsProfileCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationAlertsProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateOrganizationAlertsProfileBody create organization alerts profile body
// Example: {"alertCondition":{"bit_rate_bps":10000,"duration":60,"interface":"wan1","window":600},"description":"WAN 1 high utilization","enabled":true,"networkTags":["tag1","tag2"],"recipients":{"emails":["admin@example.org"],"httpServerIds":["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vcGF0aA=="]},"type":"wanUtilization"}
swagger:model CreateOrganizationAlertsProfileBody
*/
type CreateOrganizationAlertsProfileBody struct {

	// alert condition
	// Required: true
	AlertCondition *CreateOrganizationAlertsProfileParamsBodyAlertCondition `json:"alertCondition"`

	// User supplied description of the alert
	Description string `json:"description,omitempty"`

	// Networks with these tags will be monitored for the alert
	// Required: true
	NetworkTags []string `json:"networkTags"`

	// recipients
	// Required: true
	Recipients *CreateOrganizationAlertsProfileParamsBodyRecipients `json:"recipients"`

	// The alert type
	// Required: true
	// Enum: [voipJitter voipPacketLoss voipMos wanLatency wanPacketLoss wanUtilization wanStatus appOutage]
	Type *string `json:"type"`
}

// Validate validates this create organization alerts profile body
func (o *CreateOrganizationAlertsProfileBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlertCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworkTags(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationAlertsProfileBody) validateAlertCondition(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationAlertsProfile"+"."+"alertCondition", "body", o.AlertCondition); err != nil {
		return err
	}

	if o.AlertCondition != nil {
		if err := o.AlertCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrganizationAlertsProfile" + "." + "alertCondition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrganizationAlertsProfile" + "." + "alertCondition")
			}
			return err
		}
	}

	return nil
}

func (o *CreateOrganizationAlertsProfileBody) validateNetworkTags(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationAlertsProfile"+"."+"networkTags", "body", o.NetworkTags); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAlertsProfileBody) validateRecipients(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationAlertsProfile"+"."+"recipients", "body", o.Recipients); err != nil {
		return err
	}

	if o.Recipients != nil {
		if err := o.Recipients.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrganizationAlertsProfile" + "." + "recipients")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrganizationAlertsProfile" + "." + "recipients")
			}
			return err
		}
	}

	return nil
}

var createOrganizationAlertsProfileBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["voipJitter","voipPacketLoss","voipMos","wanLatency","wanPacketLoss","wanUtilization","wanStatus","appOutage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationAlertsProfileBodyTypeTypePropEnum = append(createOrganizationAlertsProfileBodyTypeTypePropEnum, v)
	}
}

const (

	// CreateOrganizationAlertsProfileBodyTypeVoipJitter captures enum value "voipJitter"
	CreateOrganizationAlertsProfileBodyTypeVoipJitter string = "voipJitter"

	// CreateOrganizationAlertsProfileBodyTypeVoipPacketLoss captures enum value "voipPacketLoss"
	CreateOrganizationAlertsProfileBodyTypeVoipPacketLoss string = "voipPacketLoss"

	// CreateOrganizationAlertsProfileBodyTypeVoipMos captures enum value "voipMos"
	CreateOrganizationAlertsProfileBodyTypeVoipMos string = "voipMos"

	// CreateOrganizationAlertsProfileBodyTypeWanLatency captures enum value "wanLatency"
	CreateOrganizationAlertsProfileBodyTypeWanLatency string = "wanLatency"

	// CreateOrganizationAlertsProfileBodyTypeWanPacketLoss captures enum value "wanPacketLoss"
	CreateOrganizationAlertsProfileBodyTypeWanPacketLoss string = "wanPacketLoss"

	// CreateOrganizationAlertsProfileBodyTypeWanUtilization captures enum value "wanUtilization"
	CreateOrganizationAlertsProfileBodyTypeWanUtilization string = "wanUtilization"

	// CreateOrganizationAlertsProfileBodyTypeWanStatus captures enum value "wanStatus"
	CreateOrganizationAlertsProfileBodyTypeWanStatus string = "wanStatus"

	// CreateOrganizationAlertsProfileBodyTypeAppOutage captures enum value "appOutage"
	CreateOrganizationAlertsProfileBodyTypeAppOutage string = "appOutage"
)

// prop value enum
func (o *CreateOrganizationAlertsProfileBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationAlertsProfileBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationAlertsProfileBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationAlertsProfile"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("createOrganizationAlertsProfile"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create organization alerts profile body based on the context it is used
func (o *CreateOrganizationAlertsProfileBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAlertCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRecipients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationAlertsProfileBody) contextValidateAlertCondition(ctx context.Context, formats strfmt.Registry) error {

	if o.AlertCondition != nil {
		if err := o.AlertCondition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrganizationAlertsProfile" + "." + "alertCondition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrganizationAlertsProfile" + "." + "alertCondition")
			}
			return err
		}
	}

	return nil
}

func (o *CreateOrganizationAlertsProfileBody) contextValidateRecipients(ctx context.Context, formats strfmt.Registry) error {

	if o.Recipients != nil {
		if err := o.Recipients.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrganizationAlertsProfile" + "." + "recipients")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrganizationAlertsProfile" + "." + "recipients")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAlertsProfileBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAlertsProfileBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAlertsProfileBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateOrganizationAlertsProfileParamsBodyAlertCondition The conditions that determine if the alert triggers
swagger:model CreateOrganizationAlertsProfileParamsBodyAlertCondition
*/
type CreateOrganizationAlertsProfileParamsBodyAlertCondition struct {

	// The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.
	BitRateBps int64 `json:"bit_rate_bps,omitempty"`

	// The total duration in seconds that the threshold should be crossed before alerting
	Duration int64 `json:"duration,omitempty"`

	// The uplink observed for the alert.  interface must be one of the following: wan1, wan2, cellular
	// Enum: [wan1 wan2 cellular]
	Interface string `json:"interface,omitempty"`

	// The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.
	JitterMs int64 `json:"jitter_ms,omitempty"`

	// The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.
	LatencyMs int64 `json:"latency_ms,omitempty"`

	// The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.
	LossRatio float32 `json:"loss_ratio,omitempty"`

	// The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.
	Mos float32 `json:"mos,omitempty"`

	// The look back period in seconds for sensing the alert
	Window int64 `json:"window,omitempty"`
}

// Validate validates this create organization alerts profile params body alert condition
func (o *CreateOrganizationAlertsProfileParamsBodyAlertCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createOrganizationAlertsProfileParamsBodyAlertConditionTypeInterfacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wan1","wan2","cellular"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationAlertsProfileParamsBodyAlertConditionTypeInterfacePropEnum = append(createOrganizationAlertsProfileParamsBodyAlertConditionTypeInterfacePropEnum, v)
	}
}

const (

	// CreateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan1 captures enum value "wan1"
	CreateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan1 string = "wan1"

	// CreateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan2 captures enum value "wan2"
	CreateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan2 string = "wan2"

	// CreateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceCellular captures enum value "cellular"
	CreateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceCellular string = "cellular"
)

// prop value enum
func (o *CreateOrganizationAlertsProfileParamsBodyAlertCondition) validateInterfaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationAlertsProfileParamsBodyAlertConditionTypeInterfacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationAlertsProfileParamsBodyAlertCondition) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(o.Interface) { // not required
		return nil
	}

	// value enum
	if err := o.validateInterfaceEnum("createOrganizationAlertsProfile"+"."+"alertCondition"+"."+"interface", "body", o.Interface); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization alerts profile params body alert condition based on context it is used
func (o *CreateOrganizationAlertsProfileParamsBodyAlertCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAlertsProfileParamsBodyAlertCondition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAlertsProfileParamsBodyAlertCondition) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAlertsProfileParamsBodyAlertCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateOrganizationAlertsProfileParamsBodyRecipients List of recipients that will recieve the alert.
swagger:model CreateOrganizationAlertsProfileParamsBodyRecipients
*/
type CreateOrganizationAlertsProfileParamsBodyRecipients struct {

	// A list of emails that will receive information about the alert
	Emails []string `json:"emails"`

	// A list base64 encoded urls of webhook endpoints that will receive information about the alert
	HTTPServerIds []string `json:"httpServerIds"`
}

// Validate validates this create organization alerts profile params body recipients
func (o *CreateOrganizationAlertsProfileParamsBodyRecipients) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create organization alerts profile params body recipients based on context it is used
func (o *CreateOrganizationAlertsProfileParamsBodyRecipients) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAlertsProfileParamsBodyRecipients) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAlertsProfileParamsBodyRecipients) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAlertsProfileParamsBodyRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
