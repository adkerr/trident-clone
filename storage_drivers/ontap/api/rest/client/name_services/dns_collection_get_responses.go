// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// DNSCollectionGetReader is a Reader for the DNSCollectionGet structure.
type DNSCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDNSCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDNSCollectionGetOK creates a DNSCollectionGetOK with default headers values
func NewDNSCollectionGetOK() *DNSCollectionGetOK {
	return &DNSCollectionGetOK{}
}

/* DNSCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type DNSCollectionGetOK struct {
	Payload *models.DNSResponse
}

func (o *DNSCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/dns][%d] dnsCollectionGetOK  %+v", 200, o.Payload)
}
func (o *DNSCollectionGetOK) GetPayload() *models.DNSResponse {
	return o.Payload
}

func (o *DNSCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DNSResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSCollectionGetDefault creates a DNSCollectionGetDefault with default headers values
func NewDNSCollectionGetDefault(code int) *DNSCollectionGetDefault {
	return &DNSCollectionGetDefault{
		_statusCode: code,
	}
}

/* DNSCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type DNSCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the dns collection get default response
func (o *DNSCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *DNSCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/dns][%d] dns_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *DNSCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DNSCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
