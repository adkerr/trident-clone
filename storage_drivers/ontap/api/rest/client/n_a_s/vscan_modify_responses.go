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

// VscanModifyReader is a Reader for the VscanModify structure.
type VscanModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanModifyOK creates a VscanModifyOK with default headers values
func NewVscanModifyOK() *VscanModifyOK {
	return &VscanModifyOK{}
}

/* VscanModifyOK describes a response with status code 200, with default header values.

OK
*/
type VscanModifyOK struct {
}

func (o *VscanModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}][%d] vscanModifyOK ", 200)
}

func (o *VscanModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanModifyDefault creates a VscanModifyDefault with default headers values
func NewVscanModifyDefault(code int) *VscanModifyDefault {
	return &VscanModifyDefault{
		_statusCode: code,
	}
}

/* VscanModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 10027015   | Attempting to enable a Vscan but no active scanner-pool exists for the specified SVM
| 10027011   | Attempting to enable a Vscan for an SVM for which no CIFS server exists
| 10027023   | Attempting to enable a Vscan for an SVM for which no active Vscan On-Access policy exists

*/
type VscanModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the vscan modify default response
func (o *VscanModifyDefault) Code() int {
	return o._statusCode
}

func (o *VscanModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}][%d] vscan_modify default  %+v", o._statusCode, o.Payload)
}
func (o *VscanModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
