// Code generated by go-swagger; DO NOT EDIT.

package ndmp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NdmpNodeCollectionGetReader is a Reader for the NdmpNodeCollectionGet structure.
type NdmpNodeCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpNodeCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpNodeCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpNodeCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpNodeCollectionGetOK creates a NdmpNodeCollectionGetOK with default headers values
func NewNdmpNodeCollectionGetOK() *NdmpNodeCollectionGetOK {
	return &NdmpNodeCollectionGetOK{}
}

/* NdmpNodeCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NdmpNodeCollectionGetOK struct {
	Payload *models.NdmpNodeResponse
}

func (o *NdmpNodeCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/ndmp/nodes][%d] ndmpNodeCollectionGetOK  %+v", 200, o.Payload)
}
func (o *NdmpNodeCollectionGetOK) GetPayload() *models.NdmpNodeResponse {
	return o.Payload
}

func (o *NdmpNodeCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNdmpNodeCollectionGetDefault creates a NdmpNodeCollectionGetDefault with default headers values
func NewNdmpNodeCollectionGetDefault(code int) *NdmpNodeCollectionGetDefault {
	return &NdmpNodeCollectionGetDefault{
		_statusCode: code,
	}
}

/* NdmpNodeCollectionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 68812801    | Node-scoped operations are not allowed in an SVM-scope.|
| 68812804    | Failed to get the node name from the specified node UUID.|

*/
type NdmpNodeCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ndmp node collection get default response
func (o *NdmpNodeCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NdmpNodeCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/ndmp/nodes][%d] ndmp_node_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *NdmpNodeCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpNodeCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
