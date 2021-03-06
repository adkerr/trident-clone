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

// NetworkIPBgpPeerGroupGetReader is a Reader for the NetworkIPBgpPeerGroupGet structure.
type NetworkIPBgpPeerGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPBgpPeerGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPBgpPeerGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPBgpPeerGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPBgpPeerGroupGetOK creates a NetworkIPBgpPeerGroupGetOK with default headers values
func NewNetworkIPBgpPeerGroupGetOK() *NetworkIPBgpPeerGroupGetOK {
	return &NetworkIPBgpPeerGroupGetOK{}
}

/* NetworkIPBgpPeerGroupGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPBgpPeerGroupGetOK struct {
	Payload *models.BgpPeerGroup
}

func (o *NetworkIPBgpPeerGroupGetOK) Error() string {
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups/{uuid}][%d] networkIpBgpPeerGroupGetOK  %+v", 200, o.Payload)
}
func (o *NetworkIPBgpPeerGroupGetOK) GetPayload() *models.BgpPeerGroup {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BgpPeerGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkIPBgpPeerGroupGetDefault creates a NetworkIPBgpPeerGroupGetDefault with default headers values
func NewNetworkIPBgpPeerGroupGetDefault(code int) *NetworkIPBgpPeerGroupGetDefault {
	return &NetworkIPBgpPeerGroupGetDefault{
		_statusCode: code,
	}
}

/* NetworkIPBgpPeerGroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkIPBgpPeerGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ip bgp peer group get default response
func (o *NetworkIPBgpPeerGroupGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPBgpPeerGroupGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups/{uuid}][%d] network_ip_bgp_peer_group_get default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkIPBgpPeerGroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
