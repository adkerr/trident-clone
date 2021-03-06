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

// ClusterNtpKeysCreateReader is a Reader for the ClusterNtpKeysCreate structure.
type ClusterNtpKeysCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNtpKeysCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewClusterNtpKeysCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNtpKeysCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNtpKeysCreateCreated creates a ClusterNtpKeysCreateCreated with default headers values
func NewClusterNtpKeysCreateCreated() *ClusterNtpKeysCreateCreated {
	return &ClusterNtpKeysCreateCreated{}
}

/* ClusterNtpKeysCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ClusterNtpKeysCreateCreated struct {
}

func (o *ClusterNtpKeysCreateCreated) Error() string {
	return fmt.Sprintf("[POST /cluster/ntp/keys][%d] clusterNtpKeysCreateCreated ", 201)
}

func (o *ClusterNtpKeysCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterNtpKeysCreateDefault creates a ClusterNtpKeysCreateDefault with default headers values
func NewClusterNtpKeysCreateDefault(code int) *ClusterNtpKeysCreateDefault {
	return &ClusterNtpKeysCreateDefault{
		_statusCode: code,
	}
}

/* ClusterNtpKeysCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 2097187 | Invalid value for an NTP symmetric authentication key. A SHA1 key must be exactly 40 hexadecimal digits. |
| 2097189 | Too many NTP keys have been configured. |

*/
type ClusterNtpKeysCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster ntp keys create default response
func (o *ClusterNtpKeysCreateDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNtpKeysCreateDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/ntp/keys][%d] cluster_ntp_keys_create default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterNtpKeysCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNtpKeysCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
