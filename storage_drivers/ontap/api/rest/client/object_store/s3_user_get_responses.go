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

// S3UserGetReader is a Reader for the S3UserGet structure.
type S3UserGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3UserGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3UserGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3UserGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3UserGetOK creates a S3UserGetOK with default headers values
func NewS3UserGetOK() *S3UserGetOK {
	return &S3UserGetOK{}
}

/* S3UserGetOK describes a response with status code 200, with default header values.

OK
*/
type S3UserGetOK struct {
	Payload *models.S3User
}

func (o *S3UserGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/users/{name}][%d] s3UserGetOK  %+v", 200, o.Payload)
}
func (o *S3UserGetOK) GetPayload() *models.S3User {
	return o.Payload
}

func (o *S3UserGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3UserGetDefault creates a S3UserGetDefault with default headers values
func NewS3UserGetDefault(code int) *S3UserGetDefault {
	return &S3UserGetDefault{
		_statusCode: code,
	}
}

/* S3UserGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3UserGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 user get default response
func (o *S3UserGetDefault) Code() int {
	return o._statusCode
}

func (o *S3UserGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/users/{name}][%d] s3_user_get default  %+v", o._statusCode, o.Payload)
}
func (o *S3UserGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3UserGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
