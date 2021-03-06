// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IscsiSessionGetReader is a Reader for the IscsiSessionGet structure.
type IscsiSessionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiSessionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiSessionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiSessionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiSessionGetOK creates a IscsiSessionGetOK with default headers values
func NewIscsiSessionGetOK() *IscsiSessionGetOK {
	return &IscsiSessionGetOK{}
}

/* IscsiSessionGetOK describes a response with status code 200, with default header values.

OK
*/
type IscsiSessionGetOK struct {
	Payload *models.IscsiSession
}

func (o *IscsiSessionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/iscsi/sessions/{svm.uuid}/{tpgroup}/{tsih}][%d] iscsiSessionGetOK  %+v", 200, o.Payload)
}
func (o *IscsiSessionGetOK) GetPayload() *models.IscsiSession {
	return o.Payload
}

func (o *IscsiSessionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IscsiSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIscsiSessionGetDefault creates a IscsiSessionGetDefault with default headers values
func NewIscsiSessionGetDefault(code int) *IscsiSessionGetDefault {
	return &IscsiSessionGetDefault{
		_statusCode: code,
	}
}

/* IscsiSessionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |

*/
type IscsiSessionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the iscsi session get default response
func (o *IscsiSessionGetDefault) Code() int {
	return o._statusCode
}

func (o *IscsiSessionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/iscsi/sessions/{svm.uuid}/{tpgroup}/{tsih}][%d] iscsi_session_get default  %+v", o._statusCode, o.Payload)
}
func (o *IscsiSessionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiSessionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
