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

// ClusterNisGetReader is a Reader for the ClusterNisGet structure.
type ClusterNisGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNisGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNisGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNisGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNisGetOK creates a ClusterNisGetOK with default headers values
func NewClusterNisGetOK() *ClusterNisGetOK {
	return &ClusterNisGetOK{}
}

/* ClusterNisGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNisGetOK struct {
	Payload *models.ClusterNisService
}

func (o *ClusterNisGetOK) Error() string {
	return fmt.Sprintf("[GET /security/authentication/cluster/nis][%d] clusterNisGetOK  %+v", 200, o.Payload)
}
func (o *ClusterNisGetOK) GetPayload() *models.ClusterNisService {
	return o.Payload
}

func (o *ClusterNisGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNisService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNisGetDefault creates a ClusterNisGetDefault with default headers values
func NewClusterNisGetDefault(code int) *ClusterNisGetDefault {
	return &ClusterNisGetDefault{
		_statusCode: code,
	}
}

/* ClusterNisGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterNisGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster nis get default response
func (o *ClusterNisGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNisGetDefault) Error() string {
	return fmt.Sprintf("[GET /security/authentication/cluster/nis][%d] cluster_nis_get default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterNisGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNisGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
