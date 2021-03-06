// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockRetentionOperationGetReader is a Reader for the SnaplockRetentionOperationGet structure.
type SnaplockRetentionOperationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockRetentionOperationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockRetentionOperationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockRetentionOperationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockRetentionOperationGetOK creates a SnaplockRetentionOperationGetOK with default headers values
func NewSnaplockRetentionOperationGetOK() *SnaplockRetentionOperationGetOK {
	return &SnaplockRetentionOperationGetOK{}
}

/* SnaplockRetentionOperationGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockRetentionOperationGetOK struct {
	Payload *models.EbrOperation
}

func (o *SnaplockRetentionOperationGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/operations/{id}][%d] snaplockRetentionOperationGetOK  %+v", 200, o.Payload)
}
func (o *SnaplockRetentionOperationGetOK) GetPayload() *models.EbrOperation {
	return o.Payload
}

func (o *SnaplockRetentionOperationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EbrOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockRetentionOperationGetDefault creates a SnaplockRetentionOperationGetDefault with default headers values
func NewSnaplockRetentionOperationGetDefault(code int) *SnaplockRetentionOperationGetDefault {
	return &SnaplockRetentionOperationGetDefault{
		_statusCode: code,
	}
}

/* SnaplockRetentionOperationGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnaplockRetentionOperationGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock retention operation get default response
func (o *SnaplockRetentionOperationGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockRetentionOperationGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/operations/{id}][%d] snaplock_retention_operation_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockRetentionOperationGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockRetentionOperationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
