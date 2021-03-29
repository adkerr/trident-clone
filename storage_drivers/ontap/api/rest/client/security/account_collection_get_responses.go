// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AccountCollectionGetReader is a Reader for the AccountCollectionGet structure.
type AccountCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountCollectionGetOK creates a AccountCollectionGetOK with default headers values
func NewAccountCollectionGetOK() *AccountCollectionGetOK {
	return &AccountCollectionGetOK{}
}

/* AccountCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AccountCollectionGetOK struct {
	Payload *models.AccountResponse
}

func (o *AccountCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /security/accounts][%d] accountCollectionGetOK  %+v", 200, o.Payload)
}
func (o *AccountCollectionGetOK) GetPayload() *models.AccountResponse {
	return o.Payload
}

func (o *AccountCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountCollectionGetDefault creates a AccountCollectionGetDefault with default headers values
func NewAccountCollectionGetDefault(code int) *AccountCollectionGetDefault {
	return &AccountCollectionGetDefault{
		_statusCode: code,
	}
}

/* AccountCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type AccountCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the account collection get default response
func (o *AccountCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AccountCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /security/accounts][%d] account_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *AccountCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}