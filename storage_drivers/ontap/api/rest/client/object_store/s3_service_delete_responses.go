// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3ServiceDeleteReader is a Reader for the S3ServiceDelete structure.
type S3ServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3ServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3ServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3ServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3ServiceDeleteOK creates a S3ServiceDeleteOK with default headers values
func NewS3ServiceDeleteOK() *S3ServiceDeleteOK {
	return &S3ServiceDeleteOK{}
}

/* S3ServiceDeleteOK describes a response with status code 200, with default header values.

OK
*/
type S3ServiceDeleteOK struct {
	Payload *models.S3ServiceDeleteResponse
}

func (o *S3ServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/s3/services/{svm.uuid}][%d] s3ServiceDeleteOK  %+v", 200, o.Payload)
}
func (o *S3ServiceDeleteOK) GetPayload() *models.S3ServiceDeleteResponse {
	return o.Payload
}

func (o *S3ServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3ServiceDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3ServiceDeleteDefault creates a S3ServiceDeleteDefault with default headers values
func NewS3ServiceDeleteDefault(code int) *S3ServiceDeleteDefault {
	return &S3ServiceDeleteDefault{
		_statusCode: code,
	}
}

/* S3ServiceDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 92405864   | An error occurs when deleting an S3 user or bucket. The reason for failure is detailed in the error message. Follow the error codes specified for the user or bucket endpoints to see details for the failure. |

*/
type S3ServiceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 service delete default response
func (o *S3ServiceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *S3ServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/s3/services/{svm.uuid}][%d] s3_service_delete default  %+v", o._statusCode, o.Payload)
}
func (o *S3ServiceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3ServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
