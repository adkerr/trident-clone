// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewDiskModifyParams creates a new DiskModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiskModifyParams() *DiskModifyParams {
	return &DiskModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiskModifyParamsWithTimeout creates a new DiskModifyParams object
// with the ability to set a timeout on a request.
func NewDiskModifyParamsWithTimeout(timeout time.Duration) *DiskModifyParams {
	return &DiskModifyParams{
		timeout: timeout,
	}
}

// NewDiskModifyParamsWithContext creates a new DiskModifyParams object
// with the ability to set a context for a request.
func NewDiskModifyParamsWithContext(ctx context.Context) *DiskModifyParams {
	return &DiskModifyParams{
		Context: ctx,
	}
}

// NewDiskModifyParamsWithHTTPClient creates a new DiskModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiskModifyParamsWithHTTPClient(client *http.Client) *DiskModifyParams {
	return &DiskModifyParams{
		HTTPClient: client,
	}
}

/* DiskModifyParams contains all the parameters to send to the API endpoint
   for the disk modify operation.

   Typically these are written to a http.Request.
*/
type DiskModifyParams struct {

	/* EncryptionOperation.

	   Name of the operation to apply to encrypting disks.
	*/
	EncryptionOperationQueryParameter string

	/* Name.

	   Disk name
	*/
	NameQueryParameter string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disk modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiskModifyParams) WithDefaults() *DiskModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disk modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiskModifyParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := DiskModifyParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the disk modify params
func (o *DiskModifyParams) WithTimeout(timeout time.Duration) *DiskModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disk modify params
func (o *DiskModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disk modify params
func (o *DiskModifyParams) WithContext(ctx context.Context) *DiskModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disk modify params
func (o *DiskModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disk modify params
func (o *DiskModifyParams) WithHTTPClient(client *http.Client) *DiskModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disk modify params
func (o *DiskModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEncryptionOperationQueryParameter adds the encryptionOperation to the disk modify params
func (o *DiskModifyParams) WithEncryptionOperationQueryParameter(encryptionOperation string) *DiskModifyParams {
	o.SetEncryptionOperationQueryParameter(encryptionOperation)
	return o
}

// SetEncryptionOperationQueryParameter adds the encryptionOperation to the disk modify params
func (o *DiskModifyParams) SetEncryptionOperationQueryParameter(encryptionOperation string) {
	o.EncryptionOperationQueryParameter = encryptionOperation
}

// WithNameQueryParameter adds the name to the disk modify params
func (o *DiskModifyParams) WithNameQueryParameter(name string) *DiskModifyParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the disk modify params
func (o *DiskModifyParams) SetNameQueryParameter(name string) {
	o.NameQueryParameter = name
}

// WithReturnRecords adds the returnRecords to the disk modify params
func (o *DiskModifyParams) WithReturnRecords(returnRecords *bool) *DiskModifyParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the disk modify params
func (o *DiskModifyParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *DiskModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param encryption_operation
	qrEncryptionOperation := o.EncryptionOperationQueryParameter
	qEncryptionOperation := qrEncryptionOperation
	if qEncryptionOperation != "" {

		if err := r.SetQueryParam("encryption_operation", qEncryptionOperation); err != nil {
			return err
		}
	}

	// query param name
	qrName := o.NameQueryParameter
	qName := qrName
	if qName != "" {

		if err := r.SetQueryParam("name", qName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
