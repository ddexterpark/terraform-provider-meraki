// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// ModifyNetworkSmDevicesTagsReader is a Reader for the ModifyNetworkSmDevicesTags structure.
type ModifyNetworkSmDevicesTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyNetworkSmDevicesTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyNetworkSmDevicesTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModifyNetworkSmDevicesTagsOK creates a ModifyNetworkSmDevicesTagsOK with default headers values
func NewModifyNetworkSmDevicesTagsOK() *ModifyNetworkSmDevicesTagsOK {
	return &ModifyNetworkSmDevicesTagsOK{}
}

/* ModifyNetworkSmDevicesTagsOK describes a response with status code 200, with default header values.

Successful operation
*/
type ModifyNetworkSmDevicesTagsOK struct {
	Payload []interface{}
}

func (o *ModifyNetworkSmDevicesTagsOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/devices/modifyTags][%d] modifyNetworkSmDevicesTagsOK  %+v", 200, o.Payload)
}
func (o *ModifyNetworkSmDevicesTagsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *ModifyNetworkSmDevicesTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ModifyNetworkSmDevicesTagsBody modify network sm devices tags body
// Example: {"scope":["withAny, old_tag"],"tags":["tag1","tag2"],"updateAction":"add"}
swagger:model ModifyNetworkSmDevicesTagsBody
*/
type ModifyNetworkSmDevicesTagsBody struct {

	// The ids of the devices to be modified.
	Ids []string `json:"ids"`

	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be modified.
	Scope []string `json:"scope"`

	// The serials of the devices to be modified.
	Serials []string `json:"serials"`

	// The tags to be added, deleted, or updated.
	// Required: true
	Tags []string `json:"tags"`

	// One of add, delete, or update. Only devices that have been modified will be returned.
	// Required: true
	UpdateAction *string `json:"updateAction"`

	// The wifiMacs of the devices to be modified.
	WifiMacs []string `json:"wifiMacs"`
}

// Validate validates this modify network sm devices tags body
func (o *ModifyNetworkSmDevicesTagsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyNetworkSmDevicesTagsBody) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("modifyNetworkSmDevicesTags"+"."+"tags", "body", o.Tags); err != nil {
		return err
	}

	return nil
}

func (o *ModifyNetworkSmDevicesTagsBody) validateUpdateAction(formats strfmt.Registry) error {

	if err := validate.Required("modifyNetworkSmDevicesTags"+"."+"updateAction", "body", o.UpdateAction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this modify network sm devices tags body based on context it is used
func (o *ModifyNetworkSmDevicesTagsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModifyNetworkSmDevicesTagsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyNetworkSmDevicesTagsBody) UnmarshalBinary(b []byte) error {
	var res ModifyNetworkSmDevicesTagsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
