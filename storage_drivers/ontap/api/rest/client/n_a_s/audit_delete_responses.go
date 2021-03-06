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

// AuditDeleteReader is a Reader for the AuditDelete structure.
type AuditDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuditDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAuditDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuditDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuditDeleteAccepted creates a AuditDeleteAccepted with default headers values
func NewAuditDeleteAccepted() *AuditDeleteAccepted {
	return &AuditDeleteAccepted{}
}

/* AuditDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AuditDeleteAccepted struct {
}

func (o *AuditDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /protocols/audit/{svm.uuid}][%d] auditDeleteAccepted ", 202)
}

func (o *AuditDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuditDeleteDefault creates a AuditDeleteDefault with default headers values
func NewAuditDeleteDefault(code int) *AuditDeleteDefault {
	return &AuditDeleteDefault{
		_statusCode: code,
	}
}

/* AuditDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 9699349    | Auditing should be disabled before deleting the audit configuration           |
| 9699350    | Audit configuration cannot be deleted, final consolidation is in progress     |
| 9699410    | Failed to disable multiproto.audit.evtxlog.support support capability         |
| 9699430    | Failed to disable multiproto.audit.cifslogonlogoff.support support capability |
| 9699433    | Failed to disable multiproto.audit.capstaging.support support capability      |

*/
type AuditDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the audit delete default response
func (o *AuditDeleteDefault) Code() int {
	return o._statusCode
}

func (o *AuditDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/audit/{svm.uuid}][%d] audit_delete default  %+v", o._statusCode, o.Payload)
}
func (o *AuditDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AuditDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
