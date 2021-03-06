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

// ExportRuleCollectionGetReader is a Reader for the ExportRuleCollectionGet structure.
type ExportRuleCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportRuleCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleCollectionGetOK creates a ExportRuleCollectionGetOK with default headers values
func NewExportRuleCollectionGetOK() *ExportRuleCollectionGetOK {
	return &ExportRuleCollectionGetOK{}
}

/* ExportRuleCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ExportRuleCollectionGetOK struct {
	Payload *models.ExportRuleResponse
}

func (o *ExportRuleCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules][%d] exportRuleCollectionGetOK  %+v", 200, o.Payload)
}
func (o *ExportRuleCollectionGetOK) GetPayload() *models.ExportRuleResponse {
	return o.Payload
}

func (o *ExportRuleCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportRuleCollectionGetDefault creates a ExportRuleCollectionGetDefault with default headers values
func NewExportRuleCollectionGetDefault(code int) *ExportRuleCollectionGetDefault {
	return &ExportRuleCollectionGetDefault{
		_statusCode: code,
	}
}

/* ExportRuleCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ExportRuleCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the export rule collection get default response
func (o *ExportRuleCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules][%d] export_rule_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *ExportRuleCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
