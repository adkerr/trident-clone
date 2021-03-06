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

// KerberosInterfaceCollectionGetReader is a Reader for the KerberosInterfaceCollectionGet structure.
type KerberosInterfaceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KerberosInterfaceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKerberosInterfaceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKerberosInterfaceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKerberosInterfaceCollectionGetOK creates a KerberosInterfaceCollectionGetOK with default headers values
func NewKerberosInterfaceCollectionGetOK() *KerberosInterfaceCollectionGetOK {
	return &KerberosInterfaceCollectionGetOK{}
}

/* KerberosInterfaceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type KerberosInterfaceCollectionGetOK struct {
	Payload *models.KerberosInterfaceResponse
}

func (o *KerberosInterfaceCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/interfaces][%d] kerberosInterfaceCollectionGetOK  %+v", 200, o.Payload)
}
func (o *KerberosInterfaceCollectionGetOK) GetPayload() *models.KerberosInterfaceResponse {
	return o.Payload
}

func (o *KerberosInterfaceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KerberosInterfaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKerberosInterfaceCollectionGetDefault creates a KerberosInterfaceCollectionGetDefault with default headers values
func NewKerberosInterfaceCollectionGetDefault(code int) *KerberosInterfaceCollectionGetDefault {
	return &KerberosInterfaceCollectionGetDefault{
		_statusCode: code,
	}
}

/* KerberosInterfaceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type KerberosInterfaceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the kerberos interface collection get default response
func (o *KerberosInterfaceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *KerberosInterfaceCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/interfaces][%d] kerberos_interface_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *KerberosInterfaceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KerberosInterfaceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
