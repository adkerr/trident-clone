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

// CifsSearchPathDeleteReader is a Reader for the CifsSearchPathDelete structure.
type CifsSearchPathDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSearchPathDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSearchPathDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSearchPathDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSearchPathDeleteOK creates a CifsSearchPathDeleteOK with default headers values
func NewCifsSearchPathDeleteOK() *CifsSearchPathDeleteOK {
	return &CifsSearchPathDeleteOK{}
}

/* CifsSearchPathDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsSearchPathDeleteOK struct {
}

func (o *CifsSearchPathDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifsSearchPathDeleteOK ", 200)
}

func (o *CifsSearchPathDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsSearchPathDeleteDefault creates a CifsSearchPathDeleteDefault with default headers values
func NewCifsSearchPathDeleteDefault(code int) *CifsSearchPathDeleteDefault {
	return &CifsSearchPathDeleteDefault{
		_statusCode: code,
	}
}

/* CifsSearchPathDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type CifsSearchPathDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs search path delete default response
func (o *CifsSearchPathDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsSearchPathDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifs_search_path_delete default  %+v", o._statusCode, o.Payload)
}
func (o *CifsSearchPathDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSearchPathDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
