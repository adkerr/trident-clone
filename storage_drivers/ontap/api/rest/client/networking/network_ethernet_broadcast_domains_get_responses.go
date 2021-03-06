// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetworkEthernetBroadcastDomainsGetReader is a Reader for the NetworkEthernetBroadcastDomainsGet structure.
type NetworkEthernetBroadcastDomainsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetBroadcastDomainsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetBroadcastDomainsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetBroadcastDomainsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetBroadcastDomainsGetOK creates a NetworkEthernetBroadcastDomainsGetOK with default headers values
func NewNetworkEthernetBroadcastDomainsGetOK() *NetworkEthernetBroadcastDomainsGetOK {
	return &NetworkEthernetBroadcastDomainsGetOK{}
}

/* NetworkEthernetBroadcastDomainsGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetBroadcastDomainsGetOK struct {
	Payload *models.BroadcastDomainResponse
}

func (o *NetworkEthernetBroadcastDomainsGetOK) Error() string {
	return fmt.Sprintf("[GET /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainsGetOK  %+v", 200, o.Payload)
}
func (o *NetworkEthernetBroadcastDomainsGetOK) GetPayload() *models.BroadcastDomainResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastDomainResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkEthernetBroadcastDomainsGetDefault creates a NetworkEthernetBroadcastDomainsGetDefault with default headers values
func NewNetworkEthernetBroadcastDomainsGetDefault(code int) *NetworkEthernetBroadcastDomainsGetDefault {
	return &NetworkEthernetBroadcastDomainsGetDefault{
		_statusCode: code,
	}
}

/* NetworkEthernetBroadcastDomainsGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkEthernetBroadcastDomainsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ethernet broadcast domains get default response
func (o *NetworkEthernetBroadcastDomainsGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetBroadcastDomainsGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domains_get default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkEthernetBroadcastDomainsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
