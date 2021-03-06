// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnapshotPolicyCollectionGetReader is a Reader for the SnapshotPolicyCollectionGet structure.
type SnapshotPolicyCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotPolicyCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyCollectionGetOK creates a SnapshotPolicyCollectionGetOK with default headers values
func NewSnapshotPolicyCollectionGetOK() *SnapshotPolicyCollectionGetOK {
	return &SnapshotPolicyCollectionGetOK{}
}

/* SnapshotPolicyCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotPolicyCollectionGetOK struct {
	Payload *models.SnapshotPolicyResponse
}

func (o *SnapshotPolicyCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snapshot-policies][%d] snapshotPolicyCollectionGetOK  %+v", 200, o.Payload)
}
func (o *SnapshotPolicyCollectionGetOK) GetPayload() *models.SnapshotPolicyResponse {
	return o.Payload
}

func (o *SnapshotPolicyCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotPolicyCollectionGetDefault creates a SnapshotPolicyCollectionGetDefault with default headers values
func NewSnapshotPolicyCollectionGetDefault(code int) *SnapshotPolicyCollectionGetDefault {
	return &SnapshotPolicyCollectionGetDefault{
		_statusCode: code,
	}
}

/* SnapshotPolicyCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnapshotPolicyCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot policy collection get default response
func (o *SnapshotPolicyCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotPolicyCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snapshot-policies][%d] snapshot_policy_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnapshotPolicyCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
