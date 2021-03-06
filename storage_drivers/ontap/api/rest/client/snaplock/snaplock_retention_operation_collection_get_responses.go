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

// SnaplockRetentionOperationCollectionGetReader is a Reader for the SnaplockRetentionOperationCollectionGet structure.
type SnaplockRetentionOperationCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockRetentionOperationCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockRetentionOperationCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockRetentionOperationCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockRetentionOperationCollectionGetOK creates a SnaplockRetentionOperationCollectionGetOK with default headers values
func NewSnaplockRetentionOperationCollectionGetOK() *SnaplockRetentionOperationCollectionGetOK {
	return &SnaplockRetentionOperationCollectionGetOK{}
}

/* SnaplockRetentionOperationCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockRetentionOperationCollectionGetOK struct {
	Payload *models.EbrOperationResponse
}

func (o *SnaplockRetentionOperationCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/operations][%d] snaplockRetentionOperationCollectionGetOK  %+v", 200, o.Payload)
}
func (o *SnaplockRetentionOperationCollectionGetOK) GetPayload() *models.EbrOperationResponse {
	return o.Payload
}

func (o *SnaplockRetentionOperationCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EbrOperationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockRetentionOperationCollectionGetDefault creates a SnaplockRetentionOperationCollectionGetDefault with default headers values
func NewSnaplockRetentionOperationCollectionGetDefault(code int) *SnaplockRetentionOperationCollectionGetDefault {
	return &SnaplockRetentionOperationCollectionGetDefault{
		_statusCode: code,
	}
}

/* SnaplockRetentionOperationCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnaplockRetentionOperationCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock retention operation collection get default response
func (o *SnaplockRetentionOperationCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockRetentionOperationCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/operations][%d] snaplock_retention_operation_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockRetentionOperationCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockRetentionOperationCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
