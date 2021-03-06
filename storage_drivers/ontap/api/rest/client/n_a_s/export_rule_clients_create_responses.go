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

// ExportRuleClientsCreateReader is a Reader for the ExportRuleClientsCreate structure.
type ExportRuleClientsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleClientsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExportRuleClientsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleClientsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleClientsCreateCreated creates a ExportRuleClientsCreateCreated with default headers values
func NewExportRuleClientsCreateCreated() *ExportRuleClientsCreateCreated {
	return &ExportRuleClientsCreateCreated{}
}

/* ExportRuleClientsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ExportRuleClientsCreateCreated struct {
	Payload *models.ExportClientResponse
}

func (o *ExportRuleClientsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients][%d] exportRuleClientsCreateCreated  %+v", 201, o.Payload)
}
func (o *ExportRuleClientsCreateCreated) GetPayload() *models.ExportClientResponse {
	return o.Payload
}

func (o *ExportRuleClientsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportClientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportRuleClientsCreateDefault creates a ExportRuleClientsCreateDefault with default headers values
func NewExportRuleClientsCreateDefault(code int) *ExportRuleClientsCreateDefault {
	return &ExportRuleClientsCreateDefault{
		_statusCode: code,
	}
}

/* ExportRuleClientsCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1703954    | Export policy does not exist |
| 1704036    | Invalid clientmatch:  missing domain name |
| 1704037    | Invalid clientmatch:  missing network name |
| 1704038    | Invalid clientmatch:  missing netgroup name |
| 1704039    | Invalid clientmatch |
| 1704040    | Invalid clientmatch: address bytes masked out by netmask are non-zero |
| 1704041    | Invalid clientmatch: address bytes masked to zero by netmask |
| 1704042    | Invalid clientmatch: too many bits in netmask |
| 1704043    | Invalid clientmatch: invalid netmask |
| 1704044    | Invalid clientmatch: invalid characters in host name |
| 1704045    | Invalid clientmatch: invalid characters in domain name |
| 1704050    | Invalid clientmatch: the clientmatch list contains a duplicate string. Duplicate strings in a clientmatch list are not supported |
| 1704051    | Warning: Not adding any new strings to the clientmatch field for ruleindex. All of the match strings are already in the clientmatch list |
| 1704064    | Clientmatch host name too long |
| 1704065    | Clientmatch domain name too long |

*/
type ExportRuleClientsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the export rule clients create default response
func (o *ExportRuleClientsCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleClientsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients][%d] export_rule_clients_create default  %+v", o._statusCode, o.Payload)
}
func (o *ExportRuleClientsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleClientsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
