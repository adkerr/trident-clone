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

// FpolicyEngineGetReader is a Reader for the FpolicyEngineGet structure.
type FpolicyEngineGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEngineGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyEngineGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEngineGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEngineGetOK creates a FpolicyEngineGetOK with default headers values
func NewFpolicyEngineGetOK() *FpolicyEngineGetOK {
	return &FpolicyEngineGetOK{}
}

/* FpolicyEngineGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyEngineGetOK struct {
	Payload *models.FpolicyEngine
}

func (o *FpolicyEngineGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/engines/{name}][%d] fpolicyEngineGetOK  %+v", 200, o.Payload)
}
func (o *FpolicyEngineGetOK) GetPayload() *models.FpolicyEngine {
	return o.Payload
}

func (o *FpolicyEngineGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyEngine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyEngineGetDefault creates a FpolicyEngineGetDefault with default headers values
func NewFpolicyEngineGetDefault(code int) *FpolicyEngineGetDefault {
	return &FpolicyEngineGetDefault{
		_statusCode: code,
	}
}

/* FpolicyEngineGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyEngineGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy engine get default response
func (o *FpolicyEngineGetDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEngineGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/engines/{name}][%d] fpolicy_engine_get default  %+v", o._statusCode, o.Payload)
}
func (o *FpolicyEngineGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEngineGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
