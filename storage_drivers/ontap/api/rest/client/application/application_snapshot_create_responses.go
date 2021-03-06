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

// ApplicationSnapshotCreateReader is a Reader for the ApplicationSnapshotCreate structure.
type ApplicationSnapshotCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationSnapshotCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewApplicationSnapshotCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationSnapshotCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationSnapshotCreateAccepted creates a ApplicationSnapshotCreateAccepted with default headers values
func NewApplicationSnapshotCreateAccepted() *ApplicationSnapshotCreateAccepted {
	return &ApplicationSnapshotCreateAccepted{}
}

/* ApplicationSnapshotCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationSnapshotCreateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *ApplicationSnapshotCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots][%d] applicationSnapshotCreateAccepted  %+v", 202, o.Payload)
}
func (o *ApplicationSnapshotCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationSnapshotCreateDefault creates a ApplicationSnapshotCreateDefault with default headers values
func NewApplicationSnapshotCreateDefault(code int) *ApplicationSnapshotCreateDefault {
	return &ApplicationSnapshotCreateDefault{
		_statusCode: code,
	}
}

/* ApplicationSnapshotCreateDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationSnapshotCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application snapshot create default response
func (o *ApplicationSnapshotCreateDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationSnapshotCreateDefault) Error() string {
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots][%d] application_snapshot_create default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationSnapshotCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
