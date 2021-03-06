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

// CifsShareDeleteReader is a Reader for the CifsShareDelete structure.
type CifsShareDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareDeleteOK creates a CifsShareDeleteOK with default headers values
func NewCifsShareDeleteOK() *CifsShareDeleteOK {
	return &CifsShareDeleteOK{}
}

/* CifsShareDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareDeleteOK struct {
}

func (o *CifsShareDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifsShareDeleteOK ", 200)
}

func (o *CifsShareDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsShareDeleteDefault creates a CifsShareDeleteDefault with default headers values
func NewCifsShareDeleteDefault(code int) *CifsShareDeleteDefault {
	return &CifsShareDeleteDefault{
		_statusCode: code,
	}
}

/* CifsShareDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 655393     | Standard admin shares cannot be removed |

*/
type CifsShareDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs share delete default response
func (o *CifsShareDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifs_share_delete default  %+v", o._statusCode, o.Payload)
}
func (o *CifsShareDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
