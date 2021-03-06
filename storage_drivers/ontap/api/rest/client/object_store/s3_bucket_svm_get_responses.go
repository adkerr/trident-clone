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

// S3BucketSvmGetReader is a Reader for the S3BucketSvmGet structure.
type S3BucketSvmGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketSvmGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketSvmGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketSvmGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketSvmGetOK creates a S3BucketSvmGetOK with default headers values
func NewS3BucketSvmGetOK() *S3BucketSvmGetOK {
	return &S3BucketSvmGetOK{}
}

/* S3BucketSvmGetOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketSvmGetOK struct {
	Payload *models.S3BucketSvm
}

func (o *S3BucketSvmGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3BucketSvmGetOK  %+v", 200, o.Payload)
}
func (o *S3BucketSvmGetOK) GetPayload() *models.S3BucketSvm {
	return o.Payload
}

func (o *S3BucketSvmGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BucketSvm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketSvmGetDefault creates a S3BucketSvmGetDefault with default headers values
func NewS3BucketSvmGetDefault(code int) *S3BucketSvmGetDefault {
	return &S3BucketSvmGetDefault{
		_statusCode: code,
	}
}

/* S3BucketSvmGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3BucketSvmGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 bucket svm get default response
func (o *S3BucketSvmGetDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketSvmGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3_bucket_svm_get default  %+v", o._statusCode, o.Payload)
}
func (o *S3BucketSvmGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketSvmGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
