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

// ExportPolicyCollectionGetReader is a Reader for the ExportPolicyCollectionGet structure.
type ExportPolicyCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportPolicyCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportPolicyCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportPolicyCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportPolicyCollectionGetOK creates a ExportPolicyCollectionGetOK with default headers values
func NewExportPolicyCollectionGetOK() *ExportPolicyCollectionGetOK {
	return &ExportPolicyCollectionGetOK{}
}

/* ExportPolicyCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ExportPolicyCollectionGetOK struct {
	Payload *models.ExportPolicyResponse
}

func (o *ExportPolicyCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/export-policies][%d] exportPolicyCollectionGetOK  %+v", 200, o.Payload)
}
func (o *ExportPolicyCollectionGetOK) GetPayload() *models.ExportPolicyResponse {
	return o.Payload
}

func (o *ExportPolicyCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportPolicyCollectionGetDefault creates a ExportPolicyCollectionGetDefault with default headers values
func NewExportPolicyCollectionGetDefault(code int) *ExportPolicyCollectionGetDefault {
	return &ExportPolicyCollectionGetDefault{
		_statusCode: code,
	}
}

/* ExportPolicyCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ExportPolicyCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the export policy collection get default response
func (o *ExportPolicyCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ExportPolicyCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/export-policies][%d] export_policy_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *ExportPolicyCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportPolicyCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
