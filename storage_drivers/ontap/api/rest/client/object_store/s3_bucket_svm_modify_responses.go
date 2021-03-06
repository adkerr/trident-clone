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

// S3BucketSvmModifyReader is a Reader for the S3BucketSvmModify structure.
type S3BucketSvmModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketSvmModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewS3BucketSvmModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketSvmModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketSvmModifyAccepted creates a S3BucketSvmModifyAccepted with default headers values
func NewS3BucketSvmModifyAccepted() *S3BucketSvmModifyAccepted {
	return &S3BucketSvmModifyAccepted{}
}

/* S3BucketSvmModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3BucketSvmModifyAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *S3BucketSvmModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3BucketSvmModifyAccepted  %+v", 202, o.Payload)
}
func (o *S3BucketSvmModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketSvmModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketSvmModifyDefault creates a S3BucketSvmModifyDefault with default headers values
func NewS3BucketSvmModifyDefault(code int) *S3BucketSvmModifyDefault {
	return &S3BucketSvmModifyDefault{
		_statusCode: code,
	}
}

/* S3BucketSvmModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error code | Message |
| ---------- | ------- |
| 92405778   | "Failed to modify bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Reason: {Reason for failure}. ";
| 92405846   | "Failed to modify the object store volume. Reason: {Reason for failure}.";
| 92405811   | "Failed to modify bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Wait a few minutes and try the operation again.";
| 92405858   | "Failed to \\\"modify\\\" the \\\"bucket\\\" because the operation is only supported on data SVMs.";
| 92405861   | "The specified SVM UUID or bucket UUID does not exist.";

*/
type S3BucketSvmModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 bucket svm modify default response
func (o *S3BucketSvmModifyDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketSvmModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3_bucket_svm_modify default  %+v", o._statusCode, o.Payload)
}
func (o *S3BucketSvmModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketSvmModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
