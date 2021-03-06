// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewNfsCollectionPerformanceMetricsGetParams creates a new NfsCollectionPerformanceMetricsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNfsCollectionPerformanceMetricsGetParams() *NfsCollectionPerformanceMetricsGetParams {
	return &NfsCollectionPerformanceMetricsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNfsCollectionPerformanceMetricsGetParamsWithTimeout creates a new NfsCollectionPerformanceMetricsGetParams object
// with the ability to set a timeout on a request.
func NewNfsCollectionPerformanceMetricsGetParamsWithTimeout(timeout time.Duration) *NfsCollectionPerformanceMetricsGetParams {
	return &NfsCollectionPerformanceMetricsGetParams{
		timeout: timeout,
	}
}

// NewNfsCollectionPerformanceMetricsGetParamsWithContext creates a new NfsCollectionPerformanceMetricsGetParams object
// with the ability to set a context for a request.
func NewNfsCollectionPerformanceMetricsGetParamsWithContext(ctx context.Context) *NfsCollectionPerformanceMetricsGetParams {
	return &NfsCollectionPerformanceMetricsGetParams{
		Context: ctx,
	}
}

// NewNfsCollectionPerformanceMetricsGetParamsWithHTTPClient creates a new NfsCollectionPerformanceMetricsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNfsCollectionPerformanceMetricsGetParamsWithHTTPClient(client *http.Client) *NfsCollectionPerformanceMetricsGetParams {
	return &NfsCollectionPerformanceMetricsGetParams{
		HTTPClient: client,
	}
}

/* NfsCollectionPerformanceMetricsGetParams contains all the parameters to send to the API endpoint
   for the nfs collection performance metrics get operation.

   Typically these are written to a http.Request.
*/
type NfsCollectionPerformanceMetricsGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Interval.

	     The time range for the data. Examples can be 1h, 1d, 1m, 1w, 1y.
	The period for each time range is as follows:
	* 1h: Metrics over the most recent hour sampled over 15 seconds.
	* 1d: Metrics over the most recent day sampled over 5 minutes.
	* 1w: Metrics over the most recent week sampled over 30 minutes.
	* 1m: Metrics over the most recent month sampled over 2 hours.
	* 1y: Metrics over the most recent year sampled over a day.


	     Default: "1h"
	*/
	IntervalQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

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

	/* SvmUUID.

	   Unique identifier of the SVM.
	*/
	SVMUUIDPathParameter string

	/* Timestamp.

	   Filter by timestamp
	*/
	TimestampQueryParameter *string

	/* V3Duration.

	   Filter by v3.duration
	*/
	V3DurationQueryParameter *string

	/* V3IopsOther.

	   Filter by v3.iops.other
	*/
	V3IopsOtherQueryParameter *int64

	/* V3IopsRead.

	   Filter by v3.iops.read
	*/
	V3IopsReadQueryParameter *int64

	/* V3IopsTotal.

	   Filter by v3.iops.total
	*/
	V3IopsTotalQueryParameter *int64

	/* V3IopsWrite.

	   Filter by v3.iops.write
	*/
	V3IopsWriteQueryParameter *int64

	/* V3LatencyOther.

	   Filter by v3.latency.other
	*/
	V3LatencyOtherQueryParameter *int64

	/* V3LatencyRead.

	   Filter by v3.latency.read
	*/
	V3LatencyReadQueryParameter *int64

	/* V3LatencyTotal.

	   Filter by v3.latency.total
	*/
	V3LatencyTotalQueryParameter *int64

	/* V3LatencyWrite.

	   Filter by v3.latency.write
	*/
	V3LatencyWriteQueryParameter *int64

	/* V3Status.

	   Filter by v3.status
	*/
	V3StatusQueryParameter *string

	/* V3ThroughputRead.

	   Filter by v3.throughput.read
	*/
	V3ThroughputReadQueryParameter *int64

	/* V3ThroughputTotal.

	   Filter by v3.throughput.total
	*/
	V3ThroughputTotalQueryParameter *int64

	/* V3ThroughputWrite.

	   Filter by v3.throughput.write
	*/
	V3ThroughputWriteQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nfs collection performance metrics get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsCollectionPerformanceMetricsGetParams) WithDefaults() *NfsCollectionPerformanceMetricsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nfs collection performance metrics get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsCollectionPerformanceMetricsGetParams) SetDefaults() {
	var (
		intervalQueryParameterDefault = string("1h")

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NfsCollectionPerformanceMetricsGetParams{
		IntervalQueryParameter: &intervalQueryParameterDefault,
		ReturnRecords:          &returnRecordsDefault,
		ReturnTimeout:          &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithTimeout(timeout time.Duration) *NfsCollectionPerformanceMetricsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithContext(ctx context.Context) *NfsCollectionPerformanceMetricsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithHTTPClient(client *http.Client) *NfsCollectionPerformanceMetricsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithFields(fields []string) *NfsCollectionPerformanceMetricsGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIntervalQueryParameter adds the interval to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithIntervalQueryParameter(interval *string) *NfsCollectionPerformanceMetricsGetParams {
	o.SetIntervalQueryParameter(interval)
	return o
}

// SetIntervalQueryParameter adds the interval to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetIntervalQueryParameter(interval *string) {
	o.IntervalQueryParameter = interval
}

// WithMaxRecords adds the maxRecords to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithMaxRecords(maxRecords *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithOrderBy(orderBy []string) *NfsCollectionPerformanceMetricsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithReturnRecords(returnRecords *bool) *NfsCollectionPerformanceMetricsGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithReturnTimeout(returnTimeout *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSVMUUIDPathParameter adds the svmUUID to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithSVMUUIDPathParameter(svmUUID string) *NfsCollectionPerformanceMetricsGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithTimestampQueryParameter adds the timestamp to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithTimestampQueryParameter(timestamp *string) *NfsCollectionPerformanceMetricsGetParams {
	o.SetTimestampQueryParameter(timestamp)
	return o
}

// SetTimestampQueryParameter adds the timestamp to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetTimestampQueryParameter(timestamp *string) {
	o.TimestampQueryParameter = timestamp
}

// WithV3DurationQueryParameter adds the v3Duration to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3DurationQueryParameter(v3Duration *string) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3DurationQueryParameter(v3Duration)
	return o
}

// SetV3DurationQueryParameter adds the v3Duration to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3DurationQueryParameter(v3Duration *string) {
	o.V3DurationQueryParameter = v3Duration
}

// WithV3IopsOtherQueryParameter adds the v3IopsOther to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3IopsOtherQueryParameter(v3IopsOther *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3IopsOtherQueryParameter(v3IopsOther)
	return o
}

// SetV3IopsOtherQueryParameter adds the v3IopsOther to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3IopsOtherQueryParameter(v3IopsOther *int64) {
	o.V3IopsOtherQueryParameter = v3IopsOther
}

// WithV3IopsReadQueryParameter adds the v3IopsRead to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3IopsReadQueryParameter(v3IopsRead *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3IopsReadQueryParameter(v3IopsRead)
	return o
}

// SetV3IopsReadQueryParameter adds the v3IopsRead to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3IopsReadQueryParameter(v3IopsRead *int64) {
	o.V3IopsReadQueryParameter = v3IopsRead
}

// WithV3IopsTotalQueryParameter adds the v3IopsTotal to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3IopsTotalQueryParameter(v3IopsTotal *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3IopsTotalQueryParameter(v3IopsTotal)
	return o
}

// SetV3IopsTotalQueryParameter adds the v3IopsTotal to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3IopsTotalQueryParameter(v3IopsTotal *int64) {
	o.V3IopsTotalQueryParameter = v3IopsTotal
}

// WithV3IopsWriteQueryParameter adds the v3IopsWrite to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3IopsWriteQueryParameter(v3IopsWrite *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3IopsWriteQueryParameter(v3IopsWrite)
	return o
}

// SetV3IopsWriteQueryParameter adds the v3IopsWrite to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3IopsWriteQueryParameter(v3IopsWrite *int64) {
	o.V3IopsWriteQueryParameter = v3IopsWrite
}

// WithV3LatencyOtherQueryParameter adds the v3LatencyOther to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3LatencyOtherQueryParameter(v3LatencyOther *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3LatencyOtherQueryParameter(v3LatencyOther)
	return o
}

// SetV3LatencyOtherQueryParameter adds the v3LatencyOther to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3LatencyOtherQueryParameter(v3LatencyOther *int64) {
	o.V3LatencyOtherQueryParameter = v3LatencyOther
}

// WithV3LatencyReadQueryParameter adds the v3LatencyRead to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3LatencyReadQueryParameter(v3LatencyRead *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3LatencyReadQueryParameter(v3LatencyRead)
	return o
}

// SetV3LatencyReadQueryParameter adds the v3LatencyRead to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3LatencyReadQueryParameter(v3LatencyRead *int64) {
	o.V3LatencyReadQueryParameter = v3LatencyRead
}

// WithV3LatencyTotalQueryParameter adds the v3LatencyTotal to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3LatencyTotalQueryParameter(v3LatencyTotal *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3LatencyTotalQueryParameter(v3LatencyTotal)
	return o
}

// SetV3LatencyTotalQueryParameter adds the v3LatencyTotal to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3LatencyTotalQueryParameter(v3LatencyTotal *int64) {
	o.V3LatencyTotalQueryParameter = v3LatencyTotal
}

// WithV3LatencyWriteQueryParameter adds the v3LatencyWrite to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3LatencyWriteQueryParameter(v3LatencyWrite *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3LatencyWriteQueryParameter(v3LatencyWrite)
	return o
}

// SetV3LatencyWriteQueryParameter adds the v3LatencyWrite to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3LatencyWriteQueryParameter(v3LatencyWrite *int64) {
	o.V3LatencyWriteQueryParameter = v3LatencyWrite
}

// WithV3StatusQueryParameter adds the v3Status to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3StatusQueryParameter(v3Status *string) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3StatusQueryParameter(v3Status)
	return o
}

// SetV3StatusQueryParameter adds the v3Status to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3StatusQueryParameter(v3Status *string) {
	o.V3StatusQueryParameter = v3Status
}

// WithV3ThroughputReadQueryParameter adds the v3ThroughputRead to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3ThroughputReadQueryParameter(v3ThroughputRead *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3ThroughputReadQueryParameter(v3ThroughputRead)
	return o
}

// SetV3ThroughputReadQueryParameter adds the v3ThroughputRead to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3ThroughputReadQueryParameter(v3ThroughputRead *int64) {
	o.V3ThroughputReadQueryParameter = v3ThroughputRead
}

// WithV3ThroughputTotalQueryParameter adds the v3ThroughputTotal to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3ThroughputTotalQueryParameter(v3ThroughputTotal *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3ThroughputTotalQueryParameter(v3ThroughputTotal)
	return o
}

// SetV3ThroughputTotalQueryParameter adds the v3ThroughputTotal to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3ThroughputTotalQueryParameter(v3ThroughputTotal *int64) {
	o.V3ThroughputTotalQueryParameter = v3ThroughputTotal
}

// WithV3ThroughputWriteQueryParameter adds the v3ThroughputWrite to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) WithV3ThroughputWriteQueryParameter(v3ThroughputWrite *int64) *NfsCollectionPerformanceMetricsGetParams {
	o.SetV3ThroughputWriteQueryParameter(v3ThroughputWrite)
	return o
}

// SetV3ThroughputWriteQueryParameter adds the v3ThroughputWrite to the nfs collection performance metrics get params
func (o *NfsCollectionPerformanceMetricsGetParams) SetV3ThroughputWriteQueryParameter(v3ThroughputWrite *int64) {
	o.V3ThroughputWriteQueryParameter = v3ThroughputWrite
}

// WriteToRequest writes these params to a swagger request
func (o *NfsCollectionPerformanceMetricsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IntervalQueryParameter != nil {

		// query param interval
		var qrInterval string

		if o.IntervalQueryParameter != nil {
			qrInterval = *o.IntervalQueryParameter
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if o.TimestampQueryParameter != nil {

		// query param timestamp
		var qrTimestamp string

		if o.TimestampQueryParameter != nil {
			qrTimestamp = *o.TimestampQueryParameter
		}
		qTimestamp := qrTimestamp
		if qTimestamp != "" {

			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
		}
	}

	if o.V3DurationQueryParameter != nil {

		// query param v3.duration
		var qrV3Duration string

		if o.V3DurationQueryParameter != nil {
			qrV3Duration = *o.V3DurationQueryParameter
		}
		qV3Duration := qrV3Duration
		if qV3Duration != "" {

			if err := r.SetQueryParam("v3.duration", qV3Duration); err != nil {
				return err
			}
		}
	}

	if o.V3IopsOtherQueryParameter != nil {

		// query param v3.iops.other
		var qrV3IopsOther int64

		if o.V3IopsOtherQueryParameter != nil {
			qrV3IopsOther = *o.V3IopsOtherQueryParameter
		}
		qV3IopsOther := swag.FormatInt64(qrV3IopsOther)
		if qV3IopsOther != "" {

			if err := r.SetQueryParam("v3.iops.other", qV3IopsOther); err != nil {
				return err
			}
		}
	}

	if o.V3IopsReadQueryParameter != nil {

		// query param v3.iops.read
		var qrV3IopsRead int64

		if o.V3IopsReadQueryParameter != nil {
			qrV3IopsRead = *o.V3IopsReadQueryParameter
		}
		qV3IopsRead := swag.FormatInt64(qrV3IopsRead)
		if qV3IopsRead != "" {

			if err := r.SetQueryParam("v3.iops.read", qV3IopsRead); err != nil {
				return err
			}
		}
	}

	if o.V3IopsTotalQueryParameter != nil {

		// query param v3.iops.total
		var qrV3IopsTotal int64

		if o.V3IopsTotalQueryParameter != nil {
			qrV3IopsTotal = *o.V3IopsTotalQueryParameter
		}
		qV3IopsTotal := swag.FormatInt64(qrV3IopsTotal)
		if qV3IopsTotal != "" {

			if err := r.SetQueryParam("v3.iops.total", qV3IopsTotal); err != nil {
				return err
			}
		}
	}

	if o.V3IopsWriteQueryParameter != nil {

		// query param v3.iops.write
		var qrV3IopsWrite int64

		if o.V3IopsWriteQueryParameter != nil {
			qrV3IopsWrite = *o.V3IopsWriteQueryParameter
		}
		qV3IopsWrite := swag.FormatInt64(qrV3IopsWrite)
		if qV3IopsWrite != "" {

			if err := r.SetQueryParam("v3.iops.write", qV3IopsWrite); err != nil {
				return err
			}
		}
	}

	if o.V3LatencyOtherQueryParameter != nil {

		// query param v3.latency.other
		var qrV3LatencyOther int64

		if o.V3LatencyOtherQueryParameter != nil {
			qrV3LatencyOther = *o.V3LatencyOtherQueryParameter
		}
		qV3LatencyOther := swag.FormatInt64(qrV3LatencyOther)
		if qV3LatencyOther != "" {

			if err := r.SetQueryParam("v3.latency.other", qV3LatencyOther); err != nil {
				return err
			}
		}
	}

	if o.V3LatencyReadQueryParameter != nil {

		// query param v3.latency.read
		var qrV3LatencyRead int64

		if o.V3LatencyReadQueryParameter != nil {
			qrV3LatencyRead = *o.V3LatencyReadQueryParameter
		}
		qV3LatencyRead := swag.FormatInt64(qrV3LatencyRead)
		if qV3LatencyRead != "" {

			if err := r.SetQueryParam("v3.latency.read", qV3LatencyRead); err != nil {
				return err
			}
		}
	}

	if o.V3LatencyTotalQueryParameter != nil {

		// query param v3.latency.total
		var qrV3LatencyTotal int64

		if o.V3LatencyTotalQueryParameter != nil {
			qrV3LatencyTotal = *o.V3LatencyTotalQueryParameter
		}
		qV3LatencyTotal := swag.FormatInt64(qrV3LatencyTotal)
		if qV3LatencyTotal != "" {

			if err := r.SetQueryParam("v3.latency.total", qV3LatencyTotal); err != nil {
				return err
			}
		}
	}

	if o.V3LatencyWriteQueryParameter != nil {

		// query param v3.latency.write
		var qrV3LatencyWrite int64

		if o.V3LatencyWriteQueryParameter != nil {
			qrV3LatencyWrite = *o.V3LatencyWriteQueryParameter
		}
		qV3LatencyWrite := swag.FormatInt64(qrV3LatencyWrite)
		if qV3LatencyWrite != "" {

			if err := r.SetQueryParam("v3.latency.write", qV3LatencyWrite); err != nil {
				return err
			}
		}
	}

	if o.V3StatusQueryParameter != nil {

		// query param v3.status
		var qrV3Status string

		if o.V3StatusQueryParameter != nil {
			qrV3Status = *o.V3StatusQueryParameter
		}
		qV3Status := qrV3Status
		if qV3Status != "" {

			if err := r.SetQueryParam("v3.status", qV3Status); err != nil {
				return err
			}
		}
	}

	if o.V3ThroughputReadQueryParameter != nil {

		// query param v3.throughput.read
		var qrV3ThroughputRead int64

		if o.V3ThroughputReadQueryParameter != nil {
			qrV3ThroughputRead = *o.V3ThroughputReadQueryParameter
		}
		qV3ThroughputRead := swag.FormatInt64(qrV3ThroughputRead)
		if qV3ThroughputRead != "" {

			if err := r.SetQueryParam("v3.throughput.read", qV3ThroughputRead); err != nil {
				return err
			}
		}
	}

	if o.V3ThroughputTotalQueryParameter != nil {

		// query param v3.throughput.total
		var qrV3ThroughputTotal int64

		if o.V3ThroughputTotalQueryParameter != nil {
			qrV3ThroughputTotal = *o.V3ThroughputTotalQueryParameter
		}
		qV3ThroughputTotal := swag.FormatInt64(qrV3ThroughputTotal)
		if qV3ThroughputTotal != "" {

			if err := r.SetQueryParam("v3.throughput.total", qV3ThroughputTotal); err != nil {
				return err
			}
		}
	}

	if o.V3ThroughputWriteQueryParameter != nil {

		// query param v3.throughput.write
		var qrV3ThroughputWrite int64

		if o.V3ThroughputWriteQueryParameter != nil {
			qrV3ThroughputWrite = *o.V3ThroughputWriteQueryParameter
		}
		qV3ThroughputWrite := swag.FormatInt64(qrV3ThroughputWrite)
		if qV3ThroughputWrite != "" {

			if err := r.SetQueryParam("v3.throughput.write", qV3ThroughputWrite); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNfsCollectionPerformanceMetricsGet binds the parameter fields
func (o *NfsCollectionPerformanceMetricsGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNfsCollectionPerformanceMetricsGet binds the parameter order_by
func (o *NfsCollectionPerformanceMetricsGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
