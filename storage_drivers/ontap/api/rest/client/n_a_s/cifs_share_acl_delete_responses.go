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

// CifsShareACLDeleteReader is a Reader for the CifsShareACLDelete structure.
type CifsShareACLDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareACLDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareACLDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareACLDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareACLDeleteOK creates a CifsShareACLDeleteOK with default headers values
func NewCifsShareACLDeleteOK() *CifsShareACLDeleteOK {
	return &CifsShareACLDeleteOK{}
}

/* CifsShareACLDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareACLDeleteOK struct {
}

func (o *CifsShareACLDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type}][%d] cifsShareAclDeleteOK ", 200)
}

func (o *CifsShareACLDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsShareACLDeleteDefault creates a CifsShareACLDeleteDefault with default headers values
func NewCifsShareACLDeleteDefault(code int) *CifsShareACLDeleteDefault {
	return &CifsShareACLDeleteDefault{
		_statusCode: code,
	}
}

/* CifsShareACLDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type CifsShareACLDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs share acl delete default response
func (o *CifsShareACLDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareACLDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type}][%d] cifs_share_acl_delete default  %+v", o._statusCode, o.Payload)
}
func (o *CifsShareACLDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareACLDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
