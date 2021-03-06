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

// ApplicationSnapshotCollectionGetReader is a Reader for the ApplicationSnapshotCollectionGet structure.
type ApplicationSnapshotCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationSnapshotCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationSnapshotCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationSnapshotCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationSnapshotCollectionGetOK creates a ApplicationSnapshotCollectionGetOK with default headers values
func NewApplicationSnapshotCollectionGetOK() *ApplicationSnapshotCollectionGetOK {
	return &ApplicationSnapshotCollectionGetOK{}
}

/* ApplicationSnapshotCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationSnapshotCollectionGetOK struct {
	Payload *models.ApplicationSnapshotResponse
}

func (o *ApplicationSnapshotCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/snapshots][%d] applicationSnapshotCollectionGetOK  %+v", 200, o.Payload)
}
func (o *ApplicationSnapshotCollectionGetOK) GetPayload() *models.ApplicationSnapshotResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationSnapshotCollectionGetDefault creates a ApplicationSnapshotCollectionGetDefault with default headers values
func NewApplicationSnapshotCollectionGetDefault(code int) *ApplicationSnapshotCollectionGetDefault {
	return &ApplicationSnapshotCollectionGetDefault{
		_statusCode: code,
	}
}

/* ApplicationSnapshotCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationSnapshotCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application snapshot collection get default response
func (o *ApplicationSnapshotCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationSnapshotCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/snapshots][%d] application_snapshot_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationSnapshotCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
