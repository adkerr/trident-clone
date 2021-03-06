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

// NewQuotaReportCollectionGetParams creates a new QuotaReportCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuotaReportCollectionGetParams() *QuotaReportCollectionGetParams {
	return &QuotaReportCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuotaReportCollectionGetParamsWithTimeout creates a new QuotaReportCollectionGetParams object
// with the ability to set a timeout on a request.
func NewQuotaReportCollectionGetParamsWithTimeout(timeout time.Duration) *QuotaReportCollectionGetParams {
	return &QuotaReportCollectionGetParams{
		timeout: timeout,
	}
}

// NewQuotaReportCollectionGetParamsWithContext creates a new QuotaReportCollectionGetParams object
// with the ability to set a context for a request.
func NewQuotaReportCollectionGetParamsWithContext(ctx context.Context) *QuotaReportCollectionGetParams {
	return &QuotaReportCollectionGetParams{
		Context: ctx,
	}
}

// NewQuotaReportCollectionGetParamsWithHTTPClient creates a new QuotaReportCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuotaReportCollectionGetParamsWithHTTPClient(client *http.Client) *QuotaReportCollectionGetParams {
	return &QuotaReportCollectionGetParams{
		HTTPClient: client,
	}
}

/* QuotaReportCollectionGetParams contains all the parameters to send to the API endpoint
   for the quota report collection get operation.

   Typically these are written to a http.Request.
*/
type QuotaReportCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* FilesHardLimit.

	   Filter by files.hard_limit
	*/
	FilesHardLimitQueryParameter *int64

	/* FilesSoftLimit.

	   Filter by files.soft_limit
	*/
	FilesSoftLimitQueryParameter *int64

	/* FilesUsedHardLimitPercent.

	   Filter by files.used.hard_limit_percent
	*/
	FilesUsedHardLimitPercentQueryParameter *int64

	/* FilesUsedSoftLimitPercent.

	   Filter by files.used.soft_limit_percent
	*/
	FilesUsedSoftLimitPercentQueryParameter *int64

	/* FilesUsedTotal.

	   Filter by files.used.total
	*/
	FilesUsedTotalQueryParameter *int64

	/* GroupID.

	   Filter by group.id
	*/
	GroupIDQueryParameter *string

	/* GroupName.

	   Filter by group.name
	*/
	GroupNameQueryParameter *string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* QtreeID.

	   Filter by qtree.id
	*/
	QtreeIDQueryParameter *int64

	/* QtreeName.

	   Filter by qtree.name
	*/
	QtreeNameQueryParameter *string

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

	/* ShowDefaultRecords.

	   The default is true for GET calls. When set to false, the default records are not reported.

	   Default: true
	*/
	ShowDefaultRecordsQueryParameter *bool

	/* SpaceHardLimit.

	   Filter by space.hard_limit
	*/
	SpaceHardLimitQueryParameter *int64

	/* SpaceSoftLimit.

	   Filter by space.soft_limit
	*/
	SpaceSoftLimitQueryParameter *int64

	/* SpaceUsedHardLimitPercent.

	   Filter by space.used.hard_limit_percent
	*/
	SpaceUsedHardLimitPercentQueryParameter *int64

	/* SpaceUsedSoftLimitPercent.

	   Filter by space.used.soft_limit_percent
	*/
	SpaceUsedSoftLimitPercentQueryParameter *int64

	/* SpaceUsedTotal.

	   Filter by space.used.total
	*/
	SpaceUsedTotalQueryParameter *int64

	/* Specifier.

	   Filter by specifier
	*/
	SpecifierQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* UsersID.

	   Filter by users.id
	*/
	UsersIDQueryParameter *string

	/* UsersName.

	   Filter by users.name
	*/
	UsersNameQueryParameter *string

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeNameQueryParameter *string

	/* VolumeUUID.

	   Filter by volume.uuid
	*/
	VolumeUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the quota report collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaReportCollectionGetParams) WithDefaults() *QuotaReportCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quota report collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaReportCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		showDefaultRecordsQueryParameterDefault = bool(true)
	)

	val := QuotaReportCollectionGetParams{
		ReturnRecords:                    &returnRecordsDefault,
		ReturnTimeout:                    &returnTimeoutDefault,
		ShowDefaultRecordsQueryParameter: &showDefaultRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithTimeout(timeout time.Duration) *QuotaReportCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithContext(ctx context.Context) *QuotaReportCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithHTTPClient(client *http.Client) *QuotaReportCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithFields(fields []string) *QuotaReportCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFilesHardLimitQueryParameter adds the filesHardLimit to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithFilesHardLimitQueryParameter(filesHardLimit *int64) *QuotaReportCollectionGetParams {
	o.SetFilesHardLimitQueryParameter(filesHardLimit)
	return o
}

// SetFilesHardLimitQueryParameter adds the filesHardLimit to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetFilesHardLimitQueryParameter(filesHardLimit *int64) {
	o.FilesHardLimitQueryParameter = filesHardLimit
}

// WithFilesSoftLimitQueryParameter adds the filesSoftLimit to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithFilesSoftLimitQueryParameter(filesSoftLimit *int64) *QuotaReportCollectionGetParams {
	o.SetFilesSoftLimitQueryParameter(filesSoftLimit)
	return o
}

// SetFilesSoftLimitQueryParameter adds the filesSoftLimit to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetFilesSoftLimitQueryParameter(filesSoftLimit *int64) {
	o.FilesSoftLimitQueryParameter = filesSoftLimit
}

// WithFilesUsedHardLimitPercentQueryParameter adds the filesUsedHardLimitPercent to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithFilesUsedHardLimitPercentQueryParameter(filesUsedHardLimitPercent *int64) *QuotaReportCollectionGetParams {
	o.SetFilesUsedHardLimitPercentQueryParameter(filesUsedHardLimitPercent)
	return o
}

// SetFilesUsedHardLimitPercentQueryParameter adds the filesUsedHardLimitPercent to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetFilesUsedHardLimitPercentQueryParameter(filesUsedHardLimitPercent *int64) {
	o.FilesUsedHardLimitPercentQueryParameter = filesUsedHardLimitPercent
}

// WithFilesUsedSoftLimitPercentQueryParameter adds the filesUsedSoftLimitPercent to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithFilesUsedSoftLimitPercentQueryParameter(filesUsedSoftLimitPercent *int64) *QuotaReportCollectionGetParams {
	o.SetFilesUsedSoftLimitPercentQueryParameter(filesUsedSoftLimitPercent)
	return o
}

// SetFilesUsedSoftLimitPercentQueryParameter adds the filesUsedSoftLimitPercent to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetFilesUsedSoftLimitPercentQueryParameter(filesUsedSoftLimitPercent *int64) {
	o.FilesUsedSoftLimitPercentQueryParameter = filesUsedSoftLimitPercent
}

// WithFilesUsedTotalQueryParameter adds the filesUsedTotal to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithFilesUsedTotalQueryParameter(filesUsedTotal *int64) *QuotaReportCollectionGetParams {
	o.SetFilesUsedTotalQueryParameter(filesUsedTotal)
	return o
}

// SetFilesUsedTotalQueryParameter adds the filesUsedTotal to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetFilesUsedTotalQueryParameter(filesUsedTotal *int64) {
	o.FilesUsedTotalQueryParameter = filesUsedTotal
}

// WithGroupIDQueryParameter adds the groupID to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithGroupIDQueryParameter(groupID *string) *QuotaReportCollectionGetParams {
	o.SetGroupIDQueryParameter(groupID)
	return o
}

// SetGroupIDQueryParameter adds the groupId to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetGroupIDQueryParameter(groupID *string) {
	o.GroupIDQueryParameter = groupID
}

// WithGroupNameQueryParameter adds the groupName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithGroupNameQueryParameter(groupName *string) *QuotaReportCollectionGetParams {
	o.SetGroupNameQueryParameter(groupName)
	return o
}

// SetGroupNameQueryParameter adds the groupName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetGroupNameQueryParameter(groupName *string) {
	o.GroupNameQueryParameter = groupName
}

// WithIndexQueryParameter adds the index to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithIndexQueryParameter(index *int64) *QuotaReportCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithMaxRecords adds the maxRecords to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithMaxRecords(maxRecords *int64) *QuotaReportCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithOrderBy(orderBy []string) *QuotaReportCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithQtreeIDQueryParameter adds the qtreeID to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithQtreeIDQueryParameter(qtreeID *int64) *QuotaReportCollectionGetParams {
	o.SetQtreeIDQueryParameter(qtreeID)
	return o
}

// SetQtreeIDQueryParameter adds the qtreeId to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetQtreeIDQueryParameter(qtreeID *int64) {
	o.QtreeIDQueryParameter = qtreeID
}

// WithQtreeNameQueryParameter adds the qtreeName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithQtreeNameQueryParameter(qtreeName *string) *QuotaReportCollectionGetParams {
	o.SetQtreeNameQueryParameter(qtreeName)
	return o
}

// SetQtreeNameQueryParameter adds the qtreeName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetQtreeNameQueryParameter(qtreeName *string) {
	o.QtreeNameQueryParameter = qtreeName
}

// WithReturnRecords adds the returnRecords to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithReturnRecords(returnRecords *bool) *QuotaReportCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *QuotaReportCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithShowDefaultRecordsQueryParameter adds the showDefaultRecords to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithShowDefaultRecordsQueryParameter(showDefaultRecords *bool) *QuotaReportCollectionGetParams {
	o.SetShowDefaultRecordsQueryParameter(showDefaultRecords)
	return o
}

// SetShowDefaultRecordsQueryParameter adds the showDefaultRecords to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetShowDefaultRecordsQueryParameter(showDefaultRecords *bool) {
	o.ShowDefaultRecordsQueryParameter = showDefaultRecords
}

// WithSpaceHardLimitQueryParameter adds the spaceHardLimit to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithSpaceHardLimitQueryParameter(spaceHardLimit *int64) *QuotaReportCollectionGetParams {
	o.SetSpaceHardLimitQueryParameter(spaceHardLimit)
	return o
}

// SetSpaceHardLimitQueryParameter adds the spaceHardLimit to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetSpaceHardLimitQueryParameter(spaceHardLimit *int64) {
	o.SpaceHardLimitQueryParameter = spaceHardLimit
}

// WithSpaceSoftLimitQueryParameter adds the spaceSoftLimit to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithSpaceSoftLimitQueryParameter(spaceSoftLimit *int64) *QuotaReportCollectionGetParams {
	o.SetSpaceSoftLimitQueryParameter(spaceSoftLimit)
	return o
}

// SetSpaceSoftLimitQueryParameter adds the spaceSoftLimit to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetSpaceSoftLimitQueryParameter(spaceSoftLimit *int64) {
	o.SpaceSoftLimitQueryParameter = spaceSoftLimit
}

// WithSpaceUsedHardLimitPercentQueryParameter adds the spaceUsedHardLimitPercent to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithSpaceUsedHardLimitPercentQueryParameter(spaceUsedHardLimitPercent *int64) *QuotaReportCollectionGetParams {
	o.SetSpaceUsedHardLimitPercentQueryParameter(spaceUsedHardLimitPercent)
	return o
}

// SetSpaceUsedHardLimitPercentQueryParameter adds the spaceUsedHardLimitPercent to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetSpaceUsedHardLimitPercentQueryParameter(spaceUsedHardLimitPercent *int64) {
	o.SpaceUsedHardLimitPercentQueryParameter = spaceUsedHardLimitPercent
}

// WithSpaceUsedSoftLimitPercentQueryParameter adds the spaceUsedSoftLimitPercent to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithSpaceUsedSoftLimitPercentQueryParameter(spaceUsedSoftLimitPercent *int64) *QuotaReportCollectionGetParams {
	o.SetSpaceUsedSoftLimitPercentQueryParameter(spaceUsedSoftLimitPercent)
	return o
}

// SetSpaceUsedSoftLimitPercentQueryParameter adds the spaceUsedSoftLimitPercent to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetSpaceUsedSoftLimitPercentQueryParameter(spaceUsedSoftLimitPercent *int64) {
	o.SpaceUsedSoftLimitPercentQueryParameter = spaceUsedSoftLimitPercent
}

// WithSpaceUsedTotalQueryParameter adds the spaceUsedTotal to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithSpaceUsedTotalQueryParameter(spaceUsedTotal *int64) *QuotaReportCollectionGetParams {
	o.SetSpaceUsedTotalQueryParameter(spaceUsedTotal)
	return o
}

// SetSpaceUsedTotalQueryParameter adds the spaceUsedTotal to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetSpaceUsedTotalQueryParameter(spaceUsedTotal *int64) {
	o.SpaceUsedTotalQueryParameter = spaceUsedTotal
}

// WithSpecifierQueryParameter adds the specifier to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithSpecifierQueryParameter(specifier *string) *QuotaReportCollectionGetParams {
	o.SetSpecifierQueryParameter(specifier)
	return o
}

// SetSpecifierQueryParameter adds the specifier to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetSpecifierQueryParameter(specifier *string) {
	o.SpecifierQueryParameter = specifier
}

// WithSVMNameQueryParameter adds the svmName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *QuotaReportCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *QuotaReportCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTypeQueryParameter adds the typeVar to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithTypeQueryParameter(typeVar *string) *QuotaReportCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithUsersIDQueryParameter adds the usersID to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithUsersIDQueryParameter(usersID *string) *QuotaReportCollectionGetParams {
	o.SetUsersIDQueryParameter(usersID)
	return o
}

// SetUsersIDQueryParameter adds the usersId to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetUsersIDQueryParameter(usersID *string) {
	o.UsersIDQueryParameter = usersID
}

// WithUsersNameQueryParameter adds the usersName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithUsersNameQueryParameter(usersName *string) *QuotaReportCollectionGetParams {
	o.SetUsersNameQueryParameter(usersName)
	return o
}

// SetUsersNameQueryParameter adds the usersName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetUsersNameQueryParameter(usersName *string) {
	o.UsersNameQueryParameter = usersName
}

// WithVolumeNameQueryParameter adds the volumeName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithVolumeNameQueryParameter(volumeName *string) *QuotaReportCollectionGetParams {
	o.SetVolumeNameQueryParameter(volumeName)
	return o
}

// SetVolumeNameQueryParameter adds the volumeName to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetVolumeNameQueryParameter(volumeName *string) {
	o.VolumeNameQueryParameter = volumeName
}

// WithVolumeUUIDQueryParameter adds the volumeUUID to the quota report collection get params
func (o *QuotaReportCollectionGetParams) WithVolumeUUIDQueryParameter(volumeUUID *string) *QuotaReportCollectionGetParams {
	o.SetVolumeUUIDQueryParameter(volumeUUID)
	return o
}

// SetVolumeUUIDQueryParameter adds the volumeUuid to the quota report collection get params
func (o *QuotaReportCollectionGetParams) SetVolumeUUIDQueryParameter(volumeUUID *string) {
	o.VolumeUUIDQueryParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *QuotaReportCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilesHardLimitQueryParameter != nil {

		// query param files.hard_limit
		var qrFilesHardLimit int64

		if o.FilesHardLimitQueryParameter != nil {
			qrFilesHardLimit = *o.FilesHardLimitQueryParameter
		}
		qFilesHardLimit := swag.FormatInt64(qrFilesHardLimit)
		if qFilesHardLimit != "" {

			if err := r.SetQueryParam("files.hard_limit", qFilesHardLimit); err != nil {
				return err
			}
		}
	}

	if o.FilesSoftLimitQueryParameter != nil {

		// query param files.soft_limit
		var qrFilesSoftLimit int64

		if o.FilesSoftLimitQueryParameter != nil {
			qrFilesSoftLimit = *o.FilesSoftLimitQueryParameter
		}
		qFilesSoftLimit := swag.FormatInt64(qrFilesSoftLimit)
		if qFilesSoftLimit != "" {

			if err := r.SetQueryParam("files.soft_limit", qFilesSoftLimit); err != nil {
				return err
			}
		}
	}

	if o.FilesUsedHardLimitPercentQueryParameter != nil {

		// query param files.used.hard_limit_percent
		var qrFilesUsedHardLimitPercent int64

		if o.FilesUsedHardLimitPercentQueryParameter != nil {
			qrFilesUsedHardLimitPercent = *o.FilesUsedHardLimitPercentQueryParameter
		}
		qFilesUsedHardLimitPercent := swag.FormatInt64(qrFilesUsedHardLimitPercent)
		if qFilesUsedHardLimitPercent != "" {

			if err := r.SetQueryParam("files.used.hard_limit_percent", qFilesUsedHardLimitPercent); err != nil {
				return err
			}
		}
	}

	if o.FilesUsedSoftLimitPercentQueryParameter != nil {

		// query param files.used.soft_limit_percent
		var qrFilesUsedSoftLimitPercent int64

		if o.FilesUsedSoftLimitPercentQueryParameter != nil {
			qrFilesUsedSoftLimitPercent = *o.FilesUsedSoftLimitPercentQueryParameter
		}
		qFilesUsedSoftLimitPercent := swag.FormatInt64(qrFilesUsedSoftLimitPercent)
		if qFilesUsedSoftLimitPercent != "" {

			if err := r.SetQueryParam("files.used.soft_limit_percent", qFilesUsedSoftLimitPercent); err != nil {
				return err
			}
		}
	}

	if o.FilesUsedTotalQueryParameter != nil {

		// query param files.used.total
		var qrFilesUsedTotal int64

		if o.FilesUsedTotalQueryParameter != nil {
			qrFilesUsedTotal = *o.FilesUsedTotalQueryParameter
		}
		qFilesUsedTotal := swag.FormatInt64(qrFilesUsedTotal)
		if qFilesUsedTotal != "" {

			if err := r.SetQueryParam("files.used.total", qFilesUsedTotal); err != nil {
				return err
			}
		}
	}

	if o.GroupIDQueryParameter != nil {

		// query param group.id
		var qrGroupID string

		if o.GroupIDQueryParameter != nil {
			qrGroupID = *o.GroupIDQueryParameter
		}
		qGroupID := qrGroupID
		if qGroupID != "" {

			if err := r.SetQueryParam("group.id", qGroupID); err != nil {
				return err
			}
		}
	}

	if o.GroupNameQueryParameter != nil {

		// query param group.name
		var qrGroupName string

		if o.GroupNameQueryParameter != nil {
			qrGroupName = *o.GroupNameQueryParameter
		}
		qGroupName := qrGroupName
		if qGroupName != "" {

			if err := r.SetQueryParam("group.name", qGroupName); err != nil {
				return err
			}
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

	if o.QtreeIDQueryParameter != nil {

		// query param qtree.id
		var qrQtreeID int64

		if o.QtreeIDQueryParameter != nil {
			qrQtreeID = *o.QtreeIDQueryParameter
		}
		qQtreeID := swag.FormatInt64(qrQtreeID)
		if qQtreeID != "" {

			if err := r.SetQueryParam("qtree.id", qQtreeID); err != nil {
				return err
			}
		}
	}

	if o.QtreeNameQueryParameter != nil {

		// query param qtree.name
		var qrQtreeName string

		if o.QtreeNameQueryParameter != nil {
			qrQtreeName = *o.QtreeNameQueryParameter
		}
		qQtreeName := qrQtreeName
		if qQtreeName != "" {

			if err := r.SetQueryParam("qtree.name", qQtreeName); err != nil {
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

	if o.ShowDefaultRecordsQueryParameter != nil {

		// query param show_default_records
		var qrShowDefaultRecords bool

		if o.ShowDefaultRecordsQueryParameter != nil {
			qrShowDefaultRecords = *o.ShowDefaultRecordsQueryParameter
		}
		qShowDefaultRecords := swag.FormatBool(qrShowDefaultRecords)
		if qShowDefaultRecords != "" {

			if err := r.SetQueryParam("show_default_records", qShowDefaultRecords); err != nil {
				return err
			}
		}
	}

	if o.SpaceHardLimitQueryParameter != nil {

		// query param space.hard_limit
		var qrSpaceHardLimit int64

		if o.SpaceHardLimitQueryParameter != nil {
			qrSpaceHardLimit = *o.SpaceHardLimitQueryParameter
		}
		qSpaceHardLimit := swag.FormatInt64(qrSpaceHardLimit)
		if qSpaceHardLimit != "" {

			if err := r.SetQueryParam("space.hard_limit", qSpaceHardLimit); err != nil {
				return err
			}
		}
	}

	if o.SpaceSoftLimitQueryParameter != nil {

		// query param space.soft_limit
		var qrSpaceSoftLimit int64

		if o.SpaceSoftLimitQueryParameter != nil {
			qrSpaceSoftLimit = *o.SpaceSoftLimitQueryParameter
		}
		qSpaceSoftLimit := swag.FormatInt64(qrSpaceSoftLimit)
		if qSpaceSoftLimit != "" {

			if err := r.SetQueryParam("space.soft_limit", qSpaceSoftLimit); err != nil {
				return err
			}
		}
	}

	if o.SpaceUsedHardLimitPercentQueryParameter != nil {

		// query param space.used.hard_limit_percent
		var qrSpaceUsedHardLimitPercent int64

		if o.SpaceUsedHardLimitPercentQueryParameter != nil {
			qrSpaceUsedHardLimitPercent = *o.SpaceUsedHardLimitPercentQueryParameter
		}
		qSpaceUsedHardLimitPercent := swag.FormatInt64(qrSpaceUsedHardLimitPercent)
		if qSpaceUsedHardLimitPercent != "" {

			if err := r.SetQueryParam("space.used.hard_limit_percent", qSpaceUsedHardLimitPercent); err != nil {
				return err
			}
		}
	}

	if o.SpaceUsedSoftLimitPercentQueryParameter != nil {

		// query param space.used.soft_limit_percent
		var qrSpaceUsedSoftLimitPercent int64

		if o.SpaceUsedSoftLimitPercentQueryParameter != nil {
			qrSpaceUsedSoftLimitPercent = *o.SpaceUsedSoftLimitPercentQueryParameter
		}
		qSpaceUsedSoftLimitPercent := swag.FormatInt64(qrSpaceUsedSoftLimitPercent)
		if qSpaceUsedSoftLimitPercent != "" {

			if err := r.SetQueryParam("space.used.soft_limit_percent", qSpaceUsedSoftLimitPercent); err != nil {
				return err
			}
		}
	}

	if o.SpaceUsedTotalQueryParameter != nil {

		// query param space.used.total
		var qrSpaceUsedTotal int64

		if o.SpaceUsedTotalQueryParameter != nil {
			qrSpaceUsedTotal = *o.SpaceUsedTotalQueryParameter
		}
		qSpaceUsedTotal := swag.FormatInt64(qrSpaceUsedTotal)
		if qSpaceUsedTotal != "" {

			if err := r.SetQueryParam("space.used.total", qSpaceUsedTotal); err != nil {
				return err
			}
		}
	}

	if o.SpecifierQueryParameter != nil {

		// query param specifier
		var qrSpecifier string

		if o.SpecifierQueryParameter != nil {
			qrSpecifier = *o.SpecifierQueryParameter
		}
		qSpecifier := qrSpecifier
		if qSpecifier != "" {

			if err := r.SetQueryParam("specifier", qSpecifier); err != nil {
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

	if o.TypeQueryParameter != nil {

		// query param type
		var qrType string

		if o.TypeQueryParameter != nil {
			qrType = *o.TypeQueryParameter
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.UsersIDQueryParameter != nil {

		// query param users.id
		var qrUsersID string

		if o.UsersIDQueryParameter != nil {
			qrUsersID = *o.UsersIDQueryParameter
		}
		qUsersID := qrUsersID
		if qUsersID != "" {

			if err := r.SetQueryParam("users.id", qUsersID); err != nil {
				return err
			}
		}
	}

	if o.UsersNameQueryParameter != nil {

		// query param users.name
		var qrUsersName string

		if o.UsersNameQueryParameter != nil {
			qrUsersName = *o.UsersNameQueryParameter
		}
		qUsersName := qrUsersName
		if qUsersName != "" {

			if err := r.SetQueryParam("users.name", qUsersName); err != nil {
				return err
			}
		}
	}

	if o.VolumeNameQueryParameter != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeNameQueryParameter != nil {
			qrVolumeName = *o.VolumeNameQueryParameter
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	if o.VolumeUUIDQueryParameter != nil {

		// query param volume.uuid
		var qrVolumeUUID string

		if o.VolumeUUIDQueryParameter != nil {
			qrVolumeUUID = *o.VolumeUUIDQueryParameter
		}
		qVolumeUUID := qrVolumeUUID
		if qVolumeUUID != "" {

			if err := r.SetQueryParam("volume.uuid", qVolumeUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamQuotaReportCollectionGet binds the parameter fields
func (o *QuotaReportCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamQuotaReportCollectionGet binds the parameter order_by
func (o *QuotaReportCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
