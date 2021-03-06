// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// EmsDestinationDeleteReader is a Reader for the EmsDestinationDelete structure.
type EmsDestinationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsDestinationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsDestinationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsDestinationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsDestinationDeleteOK creates a EmsDestinationDeleteOK with default headers values
func NewEmsDestinationDeleteOK() *EmsDestinationDeleteOK {
	return &EmsDestinationDeleteOK{}
}

/* EmsDestinationDeleteOK describes a response with status code 200, with default header values.

OK
*/
type EmsDestinationDeleteOK struct {
}

func (o *EmsDestinationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /support/ems/destinations/{name}][%d] emsDestinationDeleteOK ", 200)
}

func (o *EmsDestinationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmsDestinationDeleteDefault creates a EmsDestinationDeleteDefault with default headers values
func NewEmsDestinationDeleteDefault(code int) *EmsDestinationDeleteDefault {
	return &EmsDestinationDeleteDefault{
		_statusCode: code,
	}
}

/* EmsDestinationDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 983152     | Default destinations cannot be modified or removed |

*/
type EmsDestinationDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ems destination delete default response
func (o *EmsDestinationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *EmsDestinationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /support/ems/destinations/{name}][%d] ems_destination_delete default  %+v", o._statusCode, o.Payload)
}
func (o *EmsDestinationDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsDestinationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
