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

// NewCifsSymlinkMappingCollectionGetParams creates a new CifsSymlinkMappingCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsSymlinkMappingCollectionGetParams() *CifsSymlinkMappingCollectionGetParams {
	return &CifsSymlinkMappingCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsSymlinkMappingCollectionGetParamsWithTimeout creates a new CifsSymlinkMappingCollectionGetParams object
// with the ability to set a timeout on a request.
func NewCifsSymlinkMappingCollectionGetParamsWithTimeout(timeout time.Duration) *CifsSymlinkMappingCollectionGetParams {
	return &CifsSymlinkMappingCollectionGetParams{
		timeout: timeout,
	}
}

// NewCifsSymlinkMappingCollectionGetParamsWithContext creates a new CifsSymlinkMappingCollectionGetParams object
// with the ability to set a context for a request.
func NewCifsSymlinkMappingCollectionGetParamsWithContext(ctx context.Context) *CifsSymlinkMappingCollectionGetParams {
	return &CifsSymlinkMappingCollectionGetParams{
		Context: ctx,
	}
}

// NewCifsSymlinkMappingCollectionGetParamsWithHTTPClient creates a new CifsSymlinkMappingCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsSymlinkMappingCollectionGetParamsWithHTTPClient(client *http.Client) *CifsSymlinkMappingCollectionGetParams {
	return &CifsSymlinkMappingCollectionGetParams{
		HTTPClient: client,
	}
}

/* CifsSymlinkMappingCollectionGetParams contains all the parameters to send to the API endpoint
   for the cifs symlink mapping collection get operation.

   Typically these are written to a http.Request.
*/
type CifsSymlinkMappingCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

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

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* TargetHomeDirectory.

	   Filter by target.home_directory
	*/
	TargetHomeDirectoryQueryParameter *bool

	/* TargetLocality.

	   Filter by target.locality
	*/
	TargetLocalityQueryParameter *string

	/* TargetPath.

	   Filter by target.path
	*/
	TargetPathQueryParameter *string

	/* TargetServer.

	   Filter by target.server
	*/
	TargetServerQueryParameter *string

	/* TargetShare.

	   Filter by target.share
	*/
	TargetShareQueryParameter *string

	/* UnixPath.

	   Filter by unix_path
	*/
	UnixPathQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs symlink mapping collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSymlinkMappingCollectionGetParams) WithDefaults() *CifsSymlinkMappingCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs symlink mapping collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSymlinkMappingCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := CifsSymlinkMappingCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithTimeout(timeout time.Duration) *CifsSymlinkMappingCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithContext(ctx context.Context) *CifsSymlinkMappingCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithHTTPClient(client *http.Client) *CifsSymlinkMappingCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithFields(fields []string) *CifsSymlinkMappingCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithMaxRecords(maxRecords *int64) *CifsSymlinkMappingCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithOrderBy(orderBy []string) *CifsSymlinkMappingCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithReturnRecords(returnRecords *bool) *CifsSymlinkMappingCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *CifsSymlinkMappingCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *CifsSymlinkMappingCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *CifsSymlinkMappingCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTargetHomeDirectoryQueryParameter adds the targetHomeDirectory to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithTargetHomeDirectoryQueryParameter(targetHomeDirectory *bool) *CifsSymlinkMappingCollectionGetParams {
	o.SetTargetHomeDirectoryQueryParameter(targetHomeDirectory)
	return o
}

// SetTargetHomeDirectoryQueryParameter adds the targetHomeDirectory to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetTargetHomeDirectoryQueryParameter(targetHomeDirectory *bool) {
	o.TargetHomeDirectoryQueryParameter = targetHomeDirectory
}

// WithTargetLocalityQueryParameter adds the targetLocality to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithTargetLocalityQueryParameter(targetLocality *string) *CifsSymlinkMappingCollectionGetParams {
	o.SetTargetLocalityQueryParameter(targetLocality)
	return o
}

// SetTargetLocalityQueryParameter adds the targetLocality to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetTargetLocalityQueryParameter(targetLocality *string) {
	o.TargetLocalityQueryParameter = targetLocality
}

// WithTargetPathQueryParameter adds the targetPath to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithTargetPathQueryParameter(targetPath *string) *CifsSymlinkMappingCollectionGetParams {
	o.SetTargetPathQueryParameter(targetPath)
	return o
}

// SetTargetPathQueryParameter adds the targetPath to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetTargetPathQueryParameter(targetPath *string) {
	o.TargetPathQueryParameter = targetPath
}

// WithTargetServerQueryParameter adds the targetServer to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithTargetServerQueryParameter(targetServer *string) *CifsSymlinkMappingCollectionGetParams {
	o.SetTargetServerQueryParameter(targetServer)
	return o
}

// SetTargetServerQueryParameter adds the targetServer to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetTargetServerQueryParameter(targetServer *string) {
	o.TargetServerQueryParameter = targetServer
}

// WithTargetShareQueryParameter adds the targetShare to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithTargetShareQueryParameter(targetShare *string) *CifsSymlinkMappingCollectionGetParams {
	o.SetTargetShareQueryParameter(targetShare)
	return o
}

// SetTargetShareQueryParameter adds the targetShare to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetTargetShareQueryParameter(targetShare *string) {
	o.TargetShareQueryParameter = targetShare
}

// WithUnixPathQueryParameter adds the unixPath to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) WithUnixPathQueryParameter(unixPath *string) *CifsSymlinkMappingCollectionGetParams {
	o.SetUnixPathQueryParameter(unixPath)
	return o
}

// SetUnixPathQueryParameter adds the unixPath to the cifs symlink mapping collection get params
func (o *CifsSymlinkMappingCollectionGetParams) SetUnixPathQueryParameter(unixPath *string) {
	o.UnixPathQueryParameter = unixPath
}

// WriteToRequest writes these params to a swagger request
func (o *CifsSymlinkMappingCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.TargetHomeDirectoryQueryParameter != nil {

		// query param target.home_directory
		var qrTargetHomeDirectory bool

		if o.TargetHomeDirectoryQueryParameter != nil {
			qrTargetHomeDirectory = *o.TargetHomeDirectoryQueryParameter
		}
		qTargetHomeDirectory := swag.FormatBool(qrTargetHomeDirectory)
		if qTargetHomeDirectory != "" {

			if err := r.SetQueryParam("target.home_directory", qTargetHomeDirectory); err != nil {
				return err
			}
		}
	}

	if o.TargetLocalityQueryParameter != nil {

		// query param target.locality
		var qrTargetLocality string

		if o.TargetLocalityQueryParameter != nil {
			qrTargetLocality = *o.TargetLocalityQueryParameter
		}
		qTargetLocality := qrTargetLocality
		if qTargetLocality != "" {

			if err := r.SetQueryParam("target.locality", qTargetLocality); err != nil {
				return err
			}
		}
	}

	if o.TargetPathQueryParameter != nil {

		// query param target.path
		var qrTargetPath string

		if o.TargetPathQueryParameter != nil {
			qrTargetPath = *o.TargetPathQueryParameter
		}
		qTargetPath := qrTargetPath
		if qTargetPath != "" {

			if err := r.SetQueryParam("target.path", qTargetPath); err != nil {
				return err
			}
		}
	}

	if o.TargetServerQueryParameter != nil {

		// query param target.server
		var qrTargetServer string

		if o.TargetServerQueryParameter != nil {
			qrTargetServer = *o.TargetServerQueryParameter
		}
		qTargetServer := qrTargetServer
		if qTargetServer != "" {

			if err := r.SetQueryParam("target.server", qTargetServer); err != nil {
				return err
			}
		}
	}

	if o.TargetShareQueryParameter != nil {

		// query param target.share
		var qrTargetShare string

		if o.TargetShareQueryParameter != nil {
			qrTargetShare = *o.TargetShareQueryParameter
		}
		qTargetShare := qrTargetShare
		if qTargetShare != "" {

			if err := r.SetQueryParam("target.share", qTargetShare); err != nil {
				return err
			}
		}
	}

	if o.UnixPathQueryParameter != nil {

		// query param unix_path
		var qrUnixPath string

		if o.UnixPathQueryParameter != nil {
			qrUnixPath = *o.UnixPathQueryParameter
		}
		qUnixPath := qrUnixPath
		if qUnixPath != "" {

			if err := r.SetQueryParam("unix_path", qUnixPath); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsSymlinkMappingCollectionGet binds the parameter fields
func (o *CifsSymlinkMappingCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCifsSymlinkMappingCollectionGet binds the parameter order_by
func (o *CifsSymlinkMappingCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
