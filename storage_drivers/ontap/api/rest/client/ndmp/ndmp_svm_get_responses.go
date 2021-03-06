// Code generated by go-swagger; DO NOT EDIT.

package ndmp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NdmpSvmGetReader is a Reader for the NdmpSvmGet structure.
type NdmpSvmGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpSvmGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpSvmGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpSvmGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpSvmGetOK creates a NdmpSvmGetOK with default headers values
func NewNdmpSvmGetOK() *NdmpSvmGetOK {
	return &NdmpSvmGetOK{}
}

/* NdmpSvmGetOK describes a response with status code 200, with default header values.

OK
*/
type NdmpSvmGetOK struct {
	Payload *models.NdmpSvm
}

func (o *NdmpSvmGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/ndmp/svms/{svm.uuid}][%d] ndmpSvmGetOK  %+v", 200, o.Payload)
}
func (o *NdmpSvmGetOK) GetPayload() *models.NdmpSvm {
	return o.Payload
}

func (o *NdmpSvmGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpSvm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNdmpSvmGetDefault creates a NdmpSvmGetDefault with default headers values
func NewNdmpSvmGetDefault(code int) *NdmpSvmGetDefault {
	return &NdmpSvmGetDefault{
		_statusCode: code,
	}
}

/* NdmpSvmGetDefault describes a response with status code -1, with default header values.

Error
*/
type NdmpSvmGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ndmp svm get default response
func (o *NdmpSvmGetDefault) Code() int {
	return o._statusCode
}

func (o *NdmpSvmGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/ndmp/svms/{svm.uuid}][%d] ndmp_svm_get default  %+v", o._statusCode, o.Payload)
}
func (o *NdmpSvmGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpSvmGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
