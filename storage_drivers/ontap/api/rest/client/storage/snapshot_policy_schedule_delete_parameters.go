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
)

// NewSnapshotPolicyScheduleDeleteParams creates a new SnapshotPolicyScheduleDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotPolicyScheduleDeleteParams() *SnapshotPolicyScheduleDeleteParams {
	return &SnapshotPolicyScheduleDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotPolicyScheduleDeleteParamsWithTimeout creates a new SnapshotPolicyScheduleDeleteParams object
// with the ability to set a timeout on a request.
func NewSnapshotPolicyScheduleDeleteParamsWithTimeout(timeout time.Duration) *SnapshotPolicyScheduleDeleteParams {
	return &SnapshotPolicyScheduleDeleteParams{
		timeout: timeout,
	}
}

// NewSnapshotPolicyScheduleDeleteParamsWithContext creates a new SnapshotPolicyScheduleDeleteParams object
// with the ability to set a context for a request.
func NewSnapshotPolicyScheduleDeleteParamsWithContext(ctx context.Context) *SnapshotPolicyScheduleDeleteParams {
	return &SnapshotPolicyScheduleDeleteParams{
		Context: ctx,
	}
}

// NewSnapshotPolicyScheduleDeleteParamsWithHTTPClient creates a new SnapshotPolicyScheduleDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotPolicyScheduleDeleteParamsWithHTTPClient(client *http.Client) *SnapshotPolicyScheduleDeleteParams {
	return &SnapshotPolicyScheduleDeleteParams{
		HTTPClient: client,
	}
}

/* SnapshotPolicyScheduleDeleteParams contains all the parameters to send to the API endpoint
   for the snapshot policy schedule delete operation.

   Typically these are written to a http.Request.
*/
type SnapshotPolicyScheduleDeleteParams struct {

	/* SnapshotPolicyUUID.

	   Snapshot copy policy UUID
	*/
	SnapshotPolicyUUIDPathParameter string

	/* UUID.

	   Snapshot copy policy schedule UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot policy schedule delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleDeleteParams) WithDefaults() *SnapshotPolicyScheduleDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot policy schedule delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) WithTimeout(timeout time.Duration) *SnapshotPolicyScheduleDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) WithContext(ctx context.Context) *SnapshotPolicyScheduleDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) WithHTTPClient(client *http.Client) *SnapshotPolicyScheduleDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotPolicyUUIDPathParameter adds the snapshotPolicyUUID to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) WithSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID string) *SnapshotPolicyScheduleDeleteParams {
	o.SetSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID)
	return o
}

// SetSnapshotPolicyUUIDPathParameter adds the snapshotPolicyUuid to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) SetSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID string) {
	o.SnapshotPolicyUUIDPathParameter = snapshotPolicyUUID
}

// WithUUIDPathParameter adds the uuid to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) WithUUIDPathParameter(uuid string) *SnapshotPolicyScheduleDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the snapshot policy schedule delete params
func (o *SnapshotPolicyScheduleDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotPolicyScheduleDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param snapshot-policy.uuid
	if err := r.SetPathParam("snapshot-policy.uuid", o.SnapshotPolicyUUIDPathParameter); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
