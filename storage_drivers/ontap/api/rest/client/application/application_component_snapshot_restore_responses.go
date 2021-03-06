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

// ApplicationComponentSnapshotRestoreReader is a Reader for the ApplicationComponentSnapshotRestore structure.
type ApplicationComponentSnapshotRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationComponentSnapshotRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewApplicationComponentSnapshotRestoreAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationComponentSnapshotRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationComponentSnapshotRestoreAccepted creates a ApplicationComponentSnapshotRestoreAccepted with default headers values
func NewApplicationComponentSnapshotRestoreAccepted() *ApplicationComponentSnapshotRestoreAccepted {
	return &ApplicationComponentSnapshotRestoreAccepted{}
}

/* ApplicationComponentSnapshotRestoreAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationComponentSnapshotRestoreAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *ApplicationComponentSnapshotRestoreAccepted) Error() string {
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}/restore][%d] applicationComponentSnapshotRestoreAccepted  %+v", 202, o.Payload)
}
func (o *ApplicationComponentSnapshotRestoreAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotRestoreAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationComponentSnapshotRestoreDefault creates a ApplicationComponentSnapshotRestoreDefault with default headers values
func NewApplicationComponentSnapshotRestoreDefault(code int) *ApplicationComponentSnapshotRestoreDefault {
	return &ApplicationComponentSnapshotRestoreDefault{
		_statusCode: code,
	}
}

/* ApplicationComponentSnapshotRestoreDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationComponentSnapshotRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application component snapshot restore default response
func (o *ApplicationComponentSnapshotRestoreDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationComponentSnapshotRestoreDefault) Error() string {
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}/restore][%d] application_component_snapshot_restore default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationComponentSnapshotRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
