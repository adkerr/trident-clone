// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewEmsEventCollectionGetParams creates a new EmsEventCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsEventCollectionGetParams() *EmsEventCollectionGetParams {
	return &EmsEventCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsEventCollectionGetParamsWithTimeout creates a new EmsEventCollectionGetParams object
// with the ability to set a timeout on a request.
func NewEmsEventCollectionGetParamsWithTimeout(timeout time.Duration) *EmsEventCollectionGetParams {
	return &EmsEventCollectionGetParams{
		timeout: timeout,
	}
}

// NewEmsEventCollectionGetParamsWithContext creates a new EmsEventCollectionGetParams object
// with the ability to set a context for a request.
func NewEmsEventCollectionGetParamsWithContext(ctx context.Context) *EmsEventCollectionGetParams {
	return &EmsEventCollectionGetParams{
		Context: ctx,
	}
}

// NewEmsEventCollectionGetParamsWithHTTPClient creates a new EmsEventCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsEventCollectionGetParamsWithHTTPClient(client *http.Client) *EmsEventCollectionGetParams {
	return &EmsEventCollectionGetParams{
		HTTPClient: client,
	}
}

/* EmsEventCollectionGetParams contains all the parameters to send to the API endpoint
   for the ems event collection get operation.

   Typically these are written to a http.Request.
*/
type EmsEventCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* LogMessage.

	   Filter by log_message
	*/
	LogMessageQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* MessageName.

	   Filter by message.name
	*/
	MessageNameQueryParameter *string

	/* MessageSeverity.

	   Filter by message.severity
	*/
	MessageSeverityQueryParameter *string

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

	/* ParametersName.

	   Filter by parameters.name
	*/
	ParametersNameQueryParameter *string

	/* ParametersValue.

	   Filter by parameters.value
	*/
	ParametersValueQueryParameter *string

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

	/* Source.

	   Filter by source
	*/
	SourceQueryParameter *string

	/* Time.

	   Filter by time
	*/
	TimeQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems event collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsEventCollectionGetParams) WithDefaults() *EmsEventCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems event collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsEventCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := EmsEventCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithTimeout(timeout time.Duration) *EmsEventCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithContext(ctx context.Context) *EmsEventCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithHTTPClient(client *http.Client) *EmsEventCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithFields(fields []string) *EmsEventCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIndexQueryParameter adds the index to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithIndexQueryParameter(index *int64) *EmsEventCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithLogMessageQueryParameter adds the logMessage to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithLogMessageQueryParameter(logMessage *string) *EmsEventCollectionGetParams {
	o.SetLogMessageQueryParameter(logMessage)
	return o
}

// SetLogMessageQueryParameter adds the logMessage to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetLogMessageQueryParameter(logMessage *string) {
	o.LogMessageQueryParameter = logMessage
}

// WithMaxRecords adds the maxRecords to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithMaxRecords(maxRecords *int64) *EmsEventCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithMessageNameQueryParameter adds the messageName to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithMessageNameQueryParameter(messageName *string) *EmsEventCollectionGetParams {
	o.SetMessageNameQueryParameter(messageName)
	return o
}

// SetMessageNameQueryParameter adds the messageName to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetMessageNameQueryParameter(messageName *string) {
	o.MessageNameQueryParameter = messageName
}

// WithMessageSeverityQueryParameter adds the messageSeverity to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithMessageSeverityQueryParameter(messageSeverity *string) *EmsEventCollectionGetParams {
	o.SetMessageSeverityQueryParameter(messageSeverity)
	return o
}

// SetMessageSeverityQueryParameter adds the messageSeverity to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetMessageSeverityQueryParameter(messageSeverity *string) {
	o.MessageSeverityQueryParameter = messageSeverity
}

// WithNodeNameQueryParameter adds the nodeName to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *EmsEventCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *EmsEventCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderBy adds the orderBy to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithOrderBy(orderBy []string) *EmsEventCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithParametersNameQueryParameter adds the parametersName to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithParametersNameQueryParameter(parametersName *string) *EmsEventCollectionGetParams {
	o.SetParametersNameQueryParameter(parametersName)
	return o
}

// SetParametersNameQueryParameter adds the parametersName to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetParametersNameQueryParameter(parametersName *string) {
	o.ParametersNameQueryParameter = parametersName
}

// WithParametersValueQueryParameter adds the parametersValue to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithParametersValueQueryParameter(parametersValue *string) *EmsEventCollectionGetParams {
	o.SetParametersValueQueryParameter(parametersValue)
	return o
}

// SetParametersValueQueryParameter adds the parametersValue to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetParametersValueQueryParameter(parametersValue *string) {
	o.ParametersValueQueryParameter = parametersValue
}

// WithReturnRecords adds the returnRecords to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithReturnRecords(returnRecords *bool) *EmsEventCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *EmsEventCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSourceQueryParameter adds the source to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithSourceQueryParameter(source *string) *EmsEventCollectionGetParams {
	o.SetSourceQueryParameter(source)
	return o
}

// SetSourceQueryParameter adds the source to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetSourceQueryParameter(source *string) {
	o.SourceQueryParameter = source
}

// WithTimeQueryParameter adds the time to the ems event collection get params
func (o *EmsEventCollectionGetParams) WithTimeQueryParameter(time *string) *EmsEventCollectionGetParams {
	o.SetTimeQueryParameter(time)
	return o
}

// SetTimeQueryParameter adds the time to the ems event collection get params
func (o *EmsEventCollectionGetParams) SetTimeQueryParameter(time *string) {
	o.TimeQueryParameter = time
}

// WriteToRequest writes these params to a swagger request
func (o *EmsEventCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IndexQueryParameter != nil {

		// query param index
		var qrIndex int64

		if o.IndexQueryParameter != nil {
			qrIndex = *o.IndexQueryParameter
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
				return err
			}
		}
	}

	if o.LogMessageQueryParameter != nil {

		// query param log_message
		var qrLogMessage string

		if o.LogMessageQueryParameter != nil {
			qrLogMessage = *o.LogMessageQueryParameter
		}
		qLogMessage := qrLogMessage
		if qLogMessage != "" {

			if err := r.SetQueryParam("log_message", qLogMessage); err != nil {
				return err
			}
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

	if o.MessageNameQueryParameter != nil {

		// query param message.name
		var qrMessageName string

		if o.MessageNameQueryParameter != nil {
			qrMessageName = *o.MessageNameQueryParameter
		}
		qMessageName := qrMessageName
		if qMessageName != "" {

			if err := r.SetQueryParam("message.name", qMessageName); err != nil {
				return err
			}
		}
	}

	if o.MessageSeverityQueryParameter != nil {

		// query param message.severity
		var qrMessageSeverity string

		if o.MessageSeverityQueryParameter != nil {
			qrMessageSeverity = *o.MessageSeverityQueryParameter
		}
		qMessageSeverity := qrMessageSeverity
		if qMessageSeverity != "" {

			if err := r.SetQueryParam("message.severity", qMessageSeverity); err != nil {
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

	if o.ParametersNameQueryParameter != nil {

		// query param parameters.name
		var qrParametersName string

		if o.ParametersNameQueryParameter != nil {
			qrParametersName = *o.ParametersNameQueryParameter
		}
		qParametersName := qrParametersName
		if qParametersName != "" {

			if err := r.SetQueryParam("parameters.name", qParametersName); err != nil {
				return err
			}
		}
	}

	if o.ParametersValueQueryParameter != nil {

		// query param parameters.value
		var qrParametersValue string

		if o.ParametersValueQueryParameter != nil {
			qrParametersValue = *o.ParametersValueQueryParameter
		}
		qParametersValue := qrParametersValue
		if qParametersValue != "" {

			if err := r.SetQueryParam("parameters.value", qParametersValue); err != nil {
				return err
			}
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

	if o.SourceQueryParameter != nil {

		// query param source
		var qrSource string

		if o.SourceQueryParameter != nil {
			qrSource = *o.SourceQueryParameter
		}
		qSource := qrSource
		if qSource != "" {

			if err := r.SetQueryParam("source", qSource); err != nil {
				return err
			}
		}
	}

	if o.TimeQueryParameter != nil {

		// query param time
		var qrTime string

		if o.TimeQueryParameter != nil {
			qrTime = *o.TimeQueryParameter
		}
		qTime := qrTime
		if qTime != "" {

			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEmsEventCollectionGet binds the parameter fields
func (o *EmsEventCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamEmsEventCollectionGet binds the parameter order_by
func (o *EmsEventCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
