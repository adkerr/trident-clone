// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// HTTPProxyCollectionGetReader is a Reader for the HTTPProxyCollectionGet structure.
type HTTPProxyCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPProxyCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHTTPProxyCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHTTPProxyCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHTTPProxyCollectionGetOK creates a HTTPProxyCollectionGetOK with default headers values
func NewHTTPProxyCollectionGetOK() *HTTPProxyCollectionGetOK {
	return &HTTPProxyCollectionGetOK{}
}

/* HTTPProxyCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type HTTPProxyCollectionGetOK struct {
	Payload *models.NetworkHTTPProxyResponse
}

func (o *HTTPProxyCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /network/http-proxy][%d] httpProxyCollectionGetOK  %+v", 200, o.Payload)
}
func (o *HTTPProxyCollectionGetOK) GetPayload() *models.NetworkHTTPProxyResponse {
	return o.Payload
}

func (o *HTTPProxyCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkHTTPProxyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHTTPProxyCollectionGetDefault creates a HTTPProxyCollectionGetDefault with default headers values
func NewHTTPProxyCollectionGetDefault(code int) *HTTPProxyCollectionGetDefault {
	return &HTTPProxyCollectionGetDefault{
		_statusCode: code,
	}
}

/* HTTPProxyCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type HTTPProxyCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the http proxy collection get default response
func (o *HTTPProxyCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *HTTPProxyCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/http-proxy][%d] http_proxy_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *HTTPProxyCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *HTTPProxyCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
