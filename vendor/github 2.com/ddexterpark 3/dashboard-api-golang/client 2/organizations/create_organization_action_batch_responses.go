// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// CreateOrganizationActionBatchReader is a Reader for the CreateOrganizationActionBatch structure.
type CreateOrganizationActionBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationActionBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationActionBatchCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrganizationActionBatchCreated creates a CreateOrganizationActionBatchCreated with default headers values
func NewCreateOrganizationActionBatchCreated() *CreateOrganizationActionBatchCreated {
	return &CreateOrganizationActionBatchCreated{}
}

/* CreateOrganizationActionBatchCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationActionBatchCreated struct {
	Payload interface{}
}

func (o *CreateOrganizationActionBatchCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/actionBatches][%d] createOrganizationActionBatchCreated  %+v", 201, o.Payload)
}
func (o *CreateOrganizationActionBatchCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationActionBatchCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateOrganizationActionBatchBody create organization action batch body
// Example: {"actions":[{"body":{"enabled":false},"operation":"update","resource":"/devices/QXXX-XXXX-XXXX/switch/ports/3"},{"body":{"name":"Group 1"},"operation":"create","resource":"/networks/L_XXXXX/groupPolicies"}],"confirmed":true,"synchronous":false}
swagger:model CreateOrganizationActionBatchBody
*/
type CreateOrganizationActionBatchBody struct {

	// A set of changes to make as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	// Required: true
	Actions []*CreateOrganizationActionBatchParamsBodyActionsItems0 `json:"actions"`

	// Set to true for immediate execution. Set to false if the action should be previewed before executing. This property cannot be unset once it is true. Defaults to false.
	Confirmed bool `json:"confirmed,omitempty"`

	// Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch. Defaults to false.
	Synchronous bool `json:"synchronous,omitempty"`
}

// Validate validates this create organization action batch body
func (o *CreateOrganizationActionBatchBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchBody) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationActionBatch"+"."+"actions", "body", o.Actions); err != nil {
		return err
	}

	for i := 0; i < len(o.Actions); i++ {
		if swag.IsZero(o.Actions[i]) { // not required
			continue
		}

		if o.Actions[i] != nil {
			if err := o.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationActionBatch" + "." + "actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationActionBatch" + "." + "actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create organization action batch body based on the context it is used
func (o *CreateOrganizationActionBatchBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchBody) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Actions); i++ {

		if o.Actions[i] != nil {
			if err := o.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationActionBatch" + "." + "actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationActionBatch" + "." + "actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationActionBatchBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationActionBatchBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationActionBatchBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateOrganizationActionBatchParamsBodyActionsItems0 create organization action batch params body actions items0
swagger:model CreateOrganizationActionBatchParamsBodyActionsItems0
*/
type CreateOrganizationActionBatchParamsBodyActionsItems0 struct {

	// The body of the action
	Body interface{} `json:"body,omitempty"`

	// The operation to be used
	// Required: true
	Operation *string `json:"operation"`

	// Unique identifier for the resource to be acted on
	// Required: true
	Resource *string `json:"resource"`
}

// Validate validates this create organization action batch params body actions items0
func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", o.Operation); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", o.Resource); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization action batch params body actions items0 based on context it is used
func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationActionBatchParamsBodyActionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
