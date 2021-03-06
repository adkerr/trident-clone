// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NvmeSubsystemMapGetReader is a Reader for the NvmeSubsystemMapGet structure.
type NvmeSubsystemMapGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemMapGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemMapGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemMapGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemMapGetOK creates a NvmeSubsystemMapGetOK with default headers values
func NewNvmeSubsystemMapGetOK() *NvmeSubsystemMapGetOK {
	return &NvmeSubsystemMapGetOK{}
}

/* NvmeSubsystemMapGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemMapGetOK struct {
	Payload *models.NvmeSubsystemMap
}

func (o *NvmeSubsystemMapGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nvme/subsystem-maps/{subsystem.uuid}/{namespace.uuid}][%d] nvmeSubsystemMapGetOK  %+v", 200, o.Payload)
}
func (o *NvmeSubsystemMapGetOK) GetPayload() *models.NvmeSubsystemMap {
	return o.Payload
}

func (o *NvmeSubsystemMapGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeSubsystemMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeSubsystemMapGetDefault creates a NvmeSubsystemMapGetDefault with default headers values
func NewNvmeSubsystemMapGetDefault(code int) *NvmeSubsystemMapGetDefault {
	return &NvmeSubsystemMapGetDefault{
		_statusCode: code,
	}
}

/* NvmeSubsystemMapGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 72090019 | The specified NVMe namespace is not mapped to the specified NVMe subsystem. |

*/
type NvmeSubsystemMapGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nvme subsystem map get default response
func (o *NvmeSubsystemMapGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemMapGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nvme/subsystem-maps/{subsystem.uuid}/{namespace.uuid}][%d] nvme_subsystem_map_get default  %+v", o._statusCode, o.Payload)
}
func (o *NvmeSubsystemMapGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemMapGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
