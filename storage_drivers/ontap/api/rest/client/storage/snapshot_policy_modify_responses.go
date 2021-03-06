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

// SnapshotPolicyModifyReader is a Reader for the SnapshotPolicyModify structure.
type SnapshotPolicyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotPolicyModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyModifyOK creates a SnapshotPolicyModifyOK with default headers values
func NewSnapshotPolicyModifyOK() *SnapshotPolicyModifyOK {
	return &SnapshotPolicyModifyOK{}
}

/* SnapshotPolicyModifyOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotPolicyModifyOK struct {
}

func (o *SnapshotPolicyModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/snapshot-policies/{uuid}][%d] snapshotPolicyModifyOK ", 200)
}

func (o *SnapshotPolicyModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnapshotPolicyModifyDefault creates a SnapshotPolicyModifyDefault with default headers values
func NewSnapshotPolicyModifyDefault(code int) *SnapshotPolicyModifyDefault {
	return &SnapshotPolicyModifyDefault{
		_statusCode: code,
	}
}

/* SnapshotPolicyModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Code
| Error Code | Description |
| ---------- | ----------- |
| 1638414    | Cannot enable policy. Reason: Specified schedule not found. |

*/
type SnapshotPolicyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot policy modify default response
func (o *SnapshotPolicyModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotPolicyModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/snapshot-policies/{uuid}][%d] snapshot_policy_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SnapshotPolicyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
