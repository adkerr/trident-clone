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

// CifsShareCollectionGetReader is a Reader for the CifsShareCollectionGet structure.
type CifsShareCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareCollectionGetOK creates a CifsShareCollectionGetOK with default headers values
func NewCifsShareCollectionGetOK() *CifsShareCollectionGetOK {
	return &CifsShareCollectionGetOK{}
}

/* CifsShareCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareCollectionGetOK struct {
	Payload *models.CifsShareResponse
}

func (o *CifsShareCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/shares][%d] cifsShareCollectionGetOK  %+v", 200, o.Payload)
}
func (o *CifsShareCollectionGetOK) GetPayload() *models.CifsShareResponse {
	return o.Payload
}

func (o *CifsShareCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsShareResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsShareCollectionGetDefault creates a CifsShareCollectionGetDefault with default headers values
func NewCifsShareCollectionGetDefault(code int) *CifsShareCollectionGetDefault {
	return &CifsShareCollectionGetDefault{
		_statusCode: code,
	}
}

/* CifsShareCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsShareCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs share collection get default response
func (o *CifsShareCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/shares][%d] cifs_share_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *CifsShareCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
