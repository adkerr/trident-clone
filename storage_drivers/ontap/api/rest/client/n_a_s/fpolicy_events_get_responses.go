// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FpolicyEventsGetReader is a Reader for the FpolicyEventsGet structure.
type FpolicyEventsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEventsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyEventsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEventsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEventsGetOK creates a FpolicyEventsGetOK with default headers values
func NewFpolicyEventsGetOK() *FpolicyEventsGetOK {
	return &FpolicyEventsGetOK{}
}

/* FpolicyEventsGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyEventsGetOK struct {
	Payload *models.FpolicyEvent
}

func (o *FpolicyEventsGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events/{name}][%d] fpolicyEventsGetOK  %+v", 200, o.Payload)
}
func (o *FpolicyEventsGetOK) GetPayload() *models.FpolicyEvent {
	return o.Payload
}

func (o *FpolicyEventsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyEventsGetDefault creates a FpolicyEventsGetDefault with default headers values
func NewFpolicyEventsGetDefault(code int) *FpolicyEventsGetDefault {
	return &FpolicyEventsGetDefault{
		_statusCode: code,
	}
}

/* FpolicyEventsGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyEventsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy events get default response
func (o *FpolicyEventsGetDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEventsGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events/{name}][%d] fpolicy_events_get default  %+v", o._statusCode, o.Payload)
}
func (o *FpolicyEventsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEventsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
