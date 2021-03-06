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

// VscanGetReader is a Reader for the VscanGet structure.
type VscanGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanGetOK creates a VscanGetOK with default headers values
func NewVscanGetOK() *VscanGetOK {
	return &VscanGetOK{}
}

/* VscanGetOK describes a response with status code 200, with default header values.

OK
*/
type VscanGetOK struct {
	Payload *models.Vscan
}

func (o *VscanGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}][%d] vscanGetOK  %+v", 200, o.Payload)
}
func (o *VscanGetOK) GetPayload() *models.Vscan {
	return o.Payload
}

func (o *VscanGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vscan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanGetDefault creates a VscanGetDefault with default headers values
func NewVscanGetDefault(code int) *VscanGetDefault {
	return &VscanGetDefault{
		_statusCode: code,
	}
}

/* VscanGetDefault describes a response with status code -1, with default header values.

Error
*/
type VscanGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the vscan get default response
func (o *VscanGetDefault) Code() int {
	return o._statusCode
}

func (o *VscanGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}][%d] vscan_get default  %+v", o._statusCode, o.Payload)
}
func (o *VscanGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
