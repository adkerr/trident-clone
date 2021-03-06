// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ScheduleCreateReader is a Reader for the ScheduleCreate structure.
type ScheduleCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewScheduleCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleCreateCreated creates a ScheduleCreateCreated with default headers values
func NewScheduleCreateCreated() *ScheduleCreateCreated {
	return &ScheduleCreateCreated{}
}

/* ScheduleCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ScheduleCreateCreated struct {
}

func (o *ScheduleCreateCreated) Error() string {
	return fmt.Sprintf("[POST /cluster/schedules][%d] scheduleCreateCreated ", 201)
}

func (o *ScheduleCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleCreateDefault creates a ScheduleCreateDefault with default headers values
func NewScheduleCreateDefault(code int) *ScheduleCreateDefault {
	return &ScheduleCreateDefault{
		_statusCode: code,
	}
}

/* ScheduleCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 458788 | The schedule specified is not a valid schedule. |
| 459760 | The schedule specified is not a valid schedule. |
| 459763 | Schedule cannot be created locally using the remote cluster name as the owner. |
| 459764 | Cannot create a schedule with the same name as an existing schedule from the MetroCluster partner cluster but of a different schedule type. |
| 460783 | As this is a MetroCluster configuration and the local cluster is waiting for switchback, changes to non-system schedules are not allowed. |
| 460784 | An error occurred creating the remote cluster version of this schedule. |

*/
type ScheduleCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the schedule create default response
func (o *ScheduleCreateDefault) Code() int {
	return o._statusCode
}

func (o *ScheduleCreateDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/schedules][%d] schedule_create default  %+v", o._statusCode, o.Payload)
}
func (o *ScheduleCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ScheduleCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
