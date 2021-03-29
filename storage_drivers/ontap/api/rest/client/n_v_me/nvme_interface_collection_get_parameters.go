// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewNvmeInterfaceCollectionGetParams creates a new NvmeInterfaceCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeInterfaceCollectionGetParams() *NvmeInterfaceCollectionGetParams {
	return &NvmeInterfaceCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeInterfaceCollectionGetParamsWithTimeout creates a new NvmeInterfaceCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNvmeInterfaceCollectionGetParamsWithTimeout(timeout time.Duration) *NvmeInterfaceCollectionGetParams {
	return &NvmeInterfaceCollectionGetParams{
		timeout: timeout,
	}
}

// NewNvmeInterfaceCollectionGetParamsWithContext creates a new NvmeInterfaceCollectionGetParams object
// with the ability to set a context for a request.
func NewNvmeInterfaceCollectionGetParamsWithContext(ctx context.Context) *NvmeInterfaceCollectionGetParams {
	return &NvmeInterfaceCollectionGetParams{
		Context: ctx,
	}
}

// NewNvmeInterfaceCollectionGetParamsWithHTTPClient creates a new NvmeInterfaceCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeInterfaceCollectionGetParamsWithHTTPClient(client *http.Client) *NvmeInterfaceCollectionGetParams {
	return &NvmeInterfaceCollectionGetParams{
		HTTPClient: client,
	}
}

/* NvmeInterfaceCollectionGetParams contains all the parameters to send to the API endpoint
   for the nvme interface collection get operation.

   Typically these are written to a http.Request.
*/
type NvmeInterfaceCollectionGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* FcInterfacePortName.

	   Filter by fc_interface.port.name
	*/
	FcInterfacePortNameQueryParameter *string

	/* FcInterfacePortNodeName.

	   Filter by fc_interface.port.node.name
	*/
	FcInterfacePortNodeNameQueryParameter *string

	/* FcInterfacePortUUID.

	   Filter by fc_interface.port.uuid
	*/
	FcInterfacePortUUIDQueryParameter *string

	/* FcInterfaceWwnn.

	   Filter by fc_interface.wwnn
	*/
	FcInterfaceWwnnQueryParameter *string

	/* FcInterfaceWwpn.

	   Filter by fc_interface.wwpn
	*/
	FcInterfaceWwpnQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* TransportAddress.

	   Filter by transport_address
	*/
	TransportAddressQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme interface collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeInterfaceCollectionGetParams) WithDefaults() *NvmeInterfaceCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme interface collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeInterfaceCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NvmeInterfaceCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithTimeout(timeout time.Duration) *NvmeInterfaceCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithContext(ctx context.Context) *NvmeInterfaceCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithHTTPClient(client *http.Client) *NvmeInterfaceCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabledQueryParameter adds the enabled to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithEnabledQueryParameter(enabled *bool) *NvmeInterfaceCollectionGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithFcInterfacePortNameQueryParameter adds the fcInterfacePortName to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithFcInterfacePortNameQueryParameter(fcInterfacePortName *string) *NvmeInterfaceCollectionGetParams {
	o.SetFcInterfacePortNameQueryParameter(fcInterfacePortName)
	return o
}

// SetFcInterfacePortNameQueryParameter adds the fcInterfacePortName to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetFcInterfacePortNameQueryParameter(fcInterfacePortName *string) {
	o.FcInterfacePortNameQueryParameter = fcInterfacePortName
}

// WithFcInterfacePortNodeNameQueryParameter adds the fcInterfacePortNodeName to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithFcInterfacePortNodeNameQueryParameter(fcInterfacePortNodeName *string) *NvmeInterfaceCollectionGetParams {
	o.SetFcInterfacePortNodeNameQueryParameter(fcInterfacePortNodeName)
	return o
}

// SetFcInterfacePortNodeNameQueryParameter adds the fcInterfacePortNodeName to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetFcInterfacePortNodeNameQueryParameter(fcInterfacePortNodeName *string) {
	o.FcInterfacePortNodeNameQueryParameter = fcInterfacePortNodeName
}

// WithFcInterfacePortUUIDQueryParameter adds the fcInterfacePortUUID to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithFcInterfacePortUUIDQueryParameter(fcInterfacePortUUID *string) *NvmeInterfaceCollectionGetParams {
	o.SetFcInterfacePortUUIDQueryParameter(fcInterfacePortUUID)
	return o
}

// SetFcInterfacePortUUIDQueryParameter adds the fcInterfacePortUuid to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetFcInterfacePortUUIDQueryParameter(fcInterfacePortUUID *string) {
	o.FcInterfacePortUUIDQueryParameter = fcInterfacePortUUID
}

// WithFcInterfaceWwnnQueryParameter adds the fcInterfaceWwnn to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithFcInterfaceWwnnQueryParameter(fcInterfaceWwnn *string) *NvmeInterfaceCollectionGetParams {
	o.SetFcInterfaceWwnnQueryParameter(fcInterfaceWwnn)
	return o
}

// SetFcInterfaceWwnnQueryParameter adds the fcInterfaceWwnn to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetFcInterfaceWwnnQueryParameter(fcInterfaceWwnn *string) {
	o.FcInterfaceWwnnQueryParameter = fcInterfaceWwnn
}

// WithFcInterfaceWwpnQueryParameter adds the fcInterfaceWwpn to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithFcInterfaceWwpnQueryParameter(fcInterfaceWwpn *string) *NvmeInterfaceCollectionGetParams {
	o.SetFcInterfaceWwpnQueryParameter(fcInterfaceWwpn)
	return o
}

// SetFcInterfaceWwpnQueryParameter adds the fcInterfaceWwpn to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetFcInterfaceWwpnQueryParameter(fcInterfaceWwpn *string) {
	o.FcInterfaceWwpnQueryParameter = fcInterfaceWwpn
}

// WithFields adds the fields to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithFields(fields []string) *NvmeInterfaceCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithMaxRecords(maxRecords *int64) *NvmeInterfaceCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNameQueryParameter adds the name to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithNameQueryParameter(name *string) *NvmeInterfaceCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithNodeNameQueryParameter adds the nodeName to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *NvmeInterfaceCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *NvmeInterfaceCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderBy adds the orderBy to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithOrderBy(orderBy []string) *NvmeInterfaceCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithReturnRecords(returnRecords *bool) *NvmeInterfaceCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *NvmeInterfaceCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *NvmeInterfaceCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NvmeInterfaceCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTransportAddressQueryParameter adds the transportAddress to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithTransportAddressQueryParameter(transportAddress *string) *NvmeInterfaceCollectionGetParams {
	o.SetTransportAddressQueryParameter(transportAddress)
	return o
}

// SetTransportAddressQueryParameter adds the transportAddress to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetTransportAddressQueryParameter(transportAddress *string) {
	o.TransportAddressQueryParameter = transportAddress
}

// WithUUIDQueryParameter adds the uuid to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) WithUUIDQueryParameter(uuid *string) *NvmeInterfaceCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the nvme interface collection get params
func (o *NvmeInterfaceCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeInterfaceCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnabledQueryParameter != nil {

		// query param enabled
		var qrEnabled bool

		if o.EnabledQueryParameter != nil {
			qrEnabled = *o.EnabledQueryParameter
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.FcInterfacePortNameQueryParameter != nil {

		// query param fc_interface.port.name
		var qrFcInterfacePortName string

		if o.FcInterfacePortNameQueryParameter != nil {
			qrFcInterfacePortName = *o.FcInterfacePortNameQueryParameter
		}
		qFcInterfacePortName := qrFcInterfacePortName
		if qFcInterfacePortName != "" {

			if err := r.SetQueryParam("fc_interface.port.name", qFcInterfacePortName); err != nil {
				return err
			}
		}
	}

	if o.FcInterfacePortNodeNameQueryParameter != nil {

		// query param fc_interface.port.node.name
		var qrFcInterfacePortNodeName string

		if o.FcInterfacePortNodeNameQueryParameter != nil {
			qrFcInterfacePortNodeName = *o.FcInterfacePortNodeNameQueryParameter
		}
		qFcInterfacePortNodeName := qrFcInterfacePortNodeName
		if qFcInterfacePortNodeName != "" {

			if err := r.SetQueryParam("fc_interface.port.node.name", qFcInterfacePortNodeName); err != nil {
				return err
			}
		}
	}

	if o.FcInterfacePortUUIDQueryParameter != nil {

		// query param fc_interface.port.uuid
		var qrFcInterfacePortUUID string

		if o.FcInterfacePortUUIDQueryParameter != nil {
			qrFcInterfacePortUUID = *o.FcInterfacePortUUIDQueryParameter
		}
		qFcInterfacePortUUID := qrFcInterfacePortUUID
		if qFcInterfacePortUUID != "" {

			if err := r.SetQueryParam("fc_interface.port.uuid", qFcInterfacePortUUID); err != nil {
				return err
			}
		}
	}

	if o.FcInterfaceWwnnQueryParameter != nil {

		// query param fc_interface.wwnn
		var qrFcInterfaceWwnn string

		if o.FcInterfaceWwnnQueryParameter != nil {
			qrFcInterfaceWwnn = *o.FcInterfaceWwnnQueryParameter
		}
		qFcInterfaceWwnn := qrFcInterfaceWwnn
		if qFcInterfaceWwnn != "" {

			if err := r.SetQueryParam("fc_interface.wwnn", qFcInterfaceWwnn); err != nil {
				return err
			}
		}
	}

	if o.FcInterfaceWwpnQueryParameter != nil {

		// query param fc_interface.wwpn
		var qrFcInterfaceWwpn string

		if o.FcInterfaceWwpnQueryParameter != nil {
			qrFcInterfaceWwpn = *o.FcInterfaceWwpnQueryParameter
		}
		qFcInterfaceWwpn := qrFcInterfaceWwpn
		if qFcInterfaceWwpn != "" {

			if err := r.SetQueryParam("fc_interface.wwpn", qFcInterfaceWwpn); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.TransportAddressQueryParameter != nil {

		// query param transport_address
		var qrTransportAddress string

		if o.TransportAddressQueryParameter != nil {
			qrTransportAddress = *o.TransportAddressQueryParameter
		}
		qTransportAddress := qrTransportAddress
		if qTransportAddress != "" {

			if err := r.SetQueryParam("transport_address", qTransportAddress); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNvmeInterfaceCollectionGet binds the parameter fields
func (o *NvmeInterfaceCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamNvmeInterfaceCollectionGet binds the parameter order_by
func (o *NvmeInterfaceCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}