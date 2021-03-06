// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// QtreeCollectionGetReader is a Reader for the QtreeCollectionGet structure.
type QtreeCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QtreeCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQtreeCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQtreeCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQtreeCollectionGetOK creates a QtreeCollectionGetOK with default headers values
func NewQtreeCollectionGetOK() *QtreeCollectionGetOK {
	return &QtreeCollectionGetOK{}
}

/* QtreeCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type QtreeCollectionGetOK struct {
	Payload *models.QtreeResponse
}

func (o *QtreeCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/qtrees][%d] qtreeCollectionGetOK  %+v", 200, o.Payload)
}
func (o *QtreeCollectionGetOK) GetPayload() *models.QtreeResponse {
	return o.Payload
}

func (o *QtreeCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QtreeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQtreeCollectionGetDefault creates a QtreeCollectionGetDefault with default headers values
func NewQtreeCollectionGetDefault(code int) *QtreeCollectionGetDefault {
	return &QtreeCollectionGetDefault{
		_statusCode: code,
	}
}

/* QtreeCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type QtreeCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the qtree collection get default response
func (o *QtreeCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *QtreeCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/qtrees][%d] qtree_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *QtreeCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QtreeCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
