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

// CifsHomedirSearchPathGetReader is a Reader for the CifsHomedirSearchPathGet structure.
type CifsHomedirSearchPathGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsHomedirSearchPathGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsHomedirSearchPathGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsHomedirSearchPathGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsHomedirSearchPathGetOK creates a CifsHomedirSearchPathGetOK with default headers values
func NewCifsHomedirSearchPathGetOK() *CifsHomedirSearchPathGetOK {
	return &CifsHomedirSearchPathGetOK{}
}

/* CifsHomedirSearchPathGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsHomedirSearchPathGetOK struct {
	Payload *models.CifsSearchPath
}

func (o *CifsHomedirSearchPathGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifsHomedirSearchPathGetOK  %+v", 200, o.Payload)
}
func (o *CifsHomedirSearchPathGetOK) GetPayload() *models.CifsSearchPath {
	return o.Payload
}

func (o *CifsHomedirSearchPathGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsSearchPath)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsHomedirSearchPathGetDefault creates a CifsHomedirSearchPathGetDefault with default headers values
func NewCifsHomedirSearchPathGetDefault(code int) *CifsHomedirSearchPathGetDefault {
	return &CifsHomedirSearchPathGetDefault{
		_statusCode: code,
	}
}

/* CifsHomedirSearchPathGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsHomedirSearchPathGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs homedir search path get default response
func (o *CifsHomedirSearchPathGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsHomedirSearchPathGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifs_homedir_search_path_get default  %+v", o._statusCode, o.Payload)
}
func (o *CifsHomedirSearchPathGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsHomedirSearchPathGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
