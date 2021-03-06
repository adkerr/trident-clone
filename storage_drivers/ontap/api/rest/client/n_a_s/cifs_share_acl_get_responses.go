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

// CifsShareACLGetReader is a Reader for the CifsShareACLGet structure.
type CifsShareACLGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareACLGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareACLGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareACLGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareACLGetOK creates a CifsShareACLGetOK with default headers values
func NewCifsShareACLGetOK() *CifsShareACLGetOK {
	return &CifsShareACLGetOK{}
}

/* CifsShareACLGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareACLGetOK struct {
	Payload *models.CifsShareACL
}

func (o *CifsShareACLGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type}][%d] cifsShareAclGetOK  %+v", 200, o.Payload)
}
func (o *CifsShareACLGetOK) GetPayload() *models.CifsShareACL {
	return o.Payload
}

func (o *CifsShareACLGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsShareACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsShareACLGetDefault creates a CifsShareACLGetDefault with default headers values
func NewCifsShareACLGetDefault(code int) *CifsShareACLGetDefault {
	return &CifsShareACLGetDefault{
		_statusCode: code,
	}
}

/* CifsShareACLGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsShareACLGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs share acl get default response
func (o *CifsShareACLGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareACLGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type}][%d] cifs_share_acl_get default  %+v", o._statusCode, o.Payload)
}
func (o *CifsShareACLGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareACLGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
