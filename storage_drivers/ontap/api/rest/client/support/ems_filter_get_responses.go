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

// EmsFilterGetReader is a Reader for the EmsFilterGet structure.
type EmsFilterGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsFilterGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsFilterGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsFilterGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsFilterGetOK creates a EmsFilterGetOK with default headers values
func NewEmsFilterGetOK() *EmsFilterGetOK {
	return &EmsFilterGetOK{}
}

/* EmsFilterGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsFilterGetOK struct {
	Payload *models.EmsFilter
}

func (o *EmsFilterGetOK) Error() string {
	return fmt.Sprintf("[GET /support/ems/filters/{name}][%d] emsFilterGetOK  %+v", 200, o.Payload)
}
func (o *EmsFilterGetOK) GetPayload() *models.EmsFilter {
	return o.Payload
}

func (o *EmsFilterGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsFilterGetDefault creates a EmsFilterGetDefault with default headers values
func NewEmsFilterGetDefault(code int) *EmsFilterGetDefault {
	return &EmsFilterGetDefault{
		_statusCode: code,
	}
}

/* EmsFilterGetDefault describes a response with status code -1, with default header values.

Error
*/
type EmsFilterGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ems filter get default response
func (o *EmsFilterGetDefault) Code() int {
	return o._statusCode
}

func (o *EmsFilterGetDefault) Error() string {
	return fmt.Sprintf("[GET /support/ems/filters/{name}][%d] ems_filter_get default  %+v", o._statusCode, o.Payload)
}
func (o *EmsFilterGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsFilterGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
