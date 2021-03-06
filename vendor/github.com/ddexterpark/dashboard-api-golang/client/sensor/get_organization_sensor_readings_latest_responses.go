// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationSensorReadingsLatestReader is a Reader for the GetOrganizationSensorReadingsLatest structure.
type GetOrganizationSensorReadingsLatestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSensorReadingsLatestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSensorReadingsLatestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationSensorReadingsLatestOK creates a GetOrganizationSensorReadingsLatestOK with default headers values
func NewGetOrganizationSensorReadingsLatestOK() *GetOrganizationSensorReadingsLatestOK {
	return &GetOrganizationSensorReadingsLatestOK{}
}

/* GetOrganizationSensorReadingsLatestOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationSensorReadingsLatestOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

func (o *GetOrganizationSensorReadingsLatestOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/sensor/readings/latest][%d] getOrganizationSensorReadingsLatestOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationSensorReadingsLatestOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationSensorReadingsLatestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
