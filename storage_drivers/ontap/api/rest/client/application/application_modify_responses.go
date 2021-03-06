// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ApplicationModifyReader is a Reader for the ApplicationModify structure.
type ApplicationModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewApplicationModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationModifyAccepted creates a ApplicationModifyAccepted with default headers values
func NewApplicationModifyAccepted() *ApplicationModifyAccepted {
	return &ApplicationModifyAccepted{}
}

/* ApplicationModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationModifyAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *ApplicationModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /application/applications/{uuid}][%d] applicationModifyAccepted  %+v", 202, o.Payload)
}
func (o *ApplicationModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationModifyDefault creates a ApplicationModifyDefault with default headers values
func NewApplicationModifyDefault(code int) *ApplicationModifyDefault {
	return &ApplicationModifyDefault{
		_statusCode: code,
	}
}

/* ApplicationModifyDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application modify default response
func (o *ApplicationModifyDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /application/applications/{uuid}][%d] application_modify default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
