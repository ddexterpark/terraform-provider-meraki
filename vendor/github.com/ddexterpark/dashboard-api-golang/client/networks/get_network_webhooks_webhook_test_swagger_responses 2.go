// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkWebhooksWebhookTestReader is a Reader for the GetNetworkWebhooksWebhookTest structure.
type GetNetworkWebhooksWebhookTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWebhooksWebhookTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWebhooksWebhookTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWebhooksWebhookTestOK creates a GetNetworkWebhooksWebhookTestOK with default headers values
func NewGetNetworkWebhooksWebhookTestOK() *GetNetworkWebhooksWebhookTestOK {
	return &GetNetworkWebhooksWebhookTestOK{}
}

/* GetNetworkWebhooksWebhookTestOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWebhooksWebhookTestOK struct {
	Payload interface{}
}

func (o *GetNetworkWebhooksWebhookTestOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/webhooks/webhookTests/{webhookTestId}][%d] getNetworkWebhooksWebhookTestOK  %+v", 200, o.Payload)
}
func (o *GetNetworkWebhooksWebhookTestOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkWebhooksWebhookTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
