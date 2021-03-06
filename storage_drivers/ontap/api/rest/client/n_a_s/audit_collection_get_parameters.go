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

// NewAuditCollectionGetParams creates a new AuditCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditCollectionGetParams() *AuditCollectionGetParams {
	return &AuditCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditCollectionGetParamsWithTimeout creates a new AuditCollectionGetParams object
// with the ability to set a timeout on a request.
func NewAuditCollectionGetParamsWithTimeout(timeout time.Duration) *AuditCollectionGetParams {
	return &AuditCollectionGetParams{
		timeout: timeout,
	}
}

// NewAuditCollectionGetParamsWithContext creates a new AuditCollectionGetParams object
// with the ability to set a context for a request.
func NewAuditCollectionGetParamsWithContext(ctx context.Context) *AuditCollectionGetParams {
	return &AuditCollectionGetParams{
		Context: ctx,
	}
}

// NewAuditCollectionGetParamsWithHTTPClient creates a new AuditCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditCollectionGetParamsWithHTTPClient(client *http.Client) *AuditCollectionGetParams {
	return &AuditCollectionGetParams{
		HTTPClient: client,
	}
}

/* AuditCollectionGetParams contains all the parameters to send to the API endpoint
   for the audit collection get operation.

   Typically these are written to a http.Request.
*/
type AuditCollectionGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* EventsAuthorizationPolicy.

	   Filter by events.authorization_policy
	*/
	EventsAuthorizationPolicyQueryParameter *bool

	/* EventsCapStaging.

	   Filter by events.cap_staging
	*/
	EventsCapStagingQueryParameter *bool

	/* EventsCifsLogonLogoff.

	   Filter by events.cifs_logon_logoff
	*/
	EventsCifsLogonLogoffQueryParameter *bool

	/* EventsFileOperations.

	   Filter by events.file_operations
	*/
	EventsFileOperationsQueryParameter *bool

	/* EventsFileShare.

	   Filter by events.file_share
	*/
	EventsFileShareQueryParameter *bool

	/* EventsSecurityGroup.

	   Filter by events.security_group
	*/
	EventsSecurityGroupQueryParameter *bool

	/* EventsUserAccount.

	   Filter by events.user_account
	*/
	EventsUserAccountQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* LogFormat.

	   Filter by log.format
	*/
	LogFormatQueryParameter *string

	/* LogRetentionCount.

	   Filter by log.retention.count
	*/
	LogRetentionCountQueryParameter *int64

	/* LogRetentionDuration.

	   Filter by log.retention.duration
	*/
	LogRetentionDurationQueryParameter *string

	/* LogRotationScheduleDays.

	   Filter by log.rotation.schedule.days
	*/
	LogRotationScheduleDaysQueryParameter *int64

	/* LogRotationScheduleHours.

	   Filter by log.rotation.schedule.hours
	*/
	LogRotationScheduleHoursQueryParameter *int64

	/* LogRotationScheduleMinutes.

	   Filter by log.rotation.schedule.minutes
	*/
	LogRotationScheduleMinutesQueryParameter *int64

	/* LogRotationScheduleMonths.

	   Filter by log.rotation.schedule.months
	*/
	LogRotationScheduleMonthsQueryParameter *int64

	/* LogRotationScheduleWeekdays.

	   Filter by log.rotation.schedule.weekdays
	*/
	LogRotationScheduleWeekdaysQueryParameter *int64

	/* LogRotationSize.

	   Filter by log.rotation.size
	*/
	LogRotationSizeQueryParameter *int64

	/* LogPath.

	   Filter by log_path
	*/
	LogPathQueryParameter *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditCollectionGetParams) WithDefaults() *AuditCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := AuditCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the audit collection get params
func (o *AuditCollectionGetParams) WithTimeout(timeout time.Duration) *AuditCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit collection get params
func (o *AuditCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit collection get params
func (o *AuditCollectionGetParams) WithContext(ctx context.Context) *AuditCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit collection get params
func (o *AuditCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit collection get params
func (o *AuditCollectionGetParams) WithHTTPClient(client *http.Client) *AuditCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit collection get params
func (o *AuditCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabledQueryParameter adds the enabled to the audit collection get params
func (o *AuditCollectionGetParams) WithEnabledQueryParameter(enabled *bool) *AuditCollectionGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the audit collection get params
func (o *AuditCollectionGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithEventsAuthorizationPolicyQueryParameter adds the eventsAuthorizationPolicy to the audit collection get params
func (o *AuditCollectionGetParams) WithEventsAuthorizationPolicyQueryParameter(eventsAuthorizationPolicy *bool) *AuditCollectionGetParams {
	o.SetEventsAuthorizationPolicyQueryParameter(eventsAuthorizationPolicy)
	return o
}

// SetEventsAuthorizationPolicyQueryParameter adds the eventsAuthorizationPolicy to the audit collection get params
func (o *AuditCollectionGetParams) SetEventsAuthorizationPolicyQueryParameter(eventsAuthorizationPolicy *bool) {
	o.EventsAuthorizationPolicyQueryParameter = eventsAuthorizationPolicy
}

// WithEventsCapStagingQueryParameter adds the eventsCapStaging to the audit collection get params
func (o *AuditCollectionGetParams) WithEventsCapStagingQueryParameter(eventsCapStaging *bool) *AuditCollectionGetParams {
	o.SetEventsCapStagingQueryParameter(eventsCapStaging)
	return o
}

// SetEventsCapStagingQueryParameter adds the eventsCapStaging to the audit collection get params
func (o *AuditCollectionGetParams) SetEventsCapStagingQueryParameter(eventsCapStaging *bool) {
	o.EventsCapStagingQueryParameter = eventsCapStaging
}

// WithEventsCifsLogonLogoffQueryParameter adds the eventsCifsLogonLogoff to the audit collection get params
func (o *AuditCollectionGetParams) WithEventsCifsLogonLogoffQueryParameter(eventsCifsLogonLogoff *bool) *AuditCollectionGetParams {
	o.SetEventsCifsLogonLogoffQueryParameter(eventsCifsLogonLogoff)
	return o
}

// SetEventsCifsLogonLogoffQueryParameter adds the eventsCifsLogonLogoff to the audit collection get params
func (o *AuditCollectionGetParams) SetEventsCifsLogonLogoffQueryParameter(eventsCifsLogonLogoff *bool) {
	o.EventsCifsLogonLogoffQueryParameter = eventsCifsLogonLogoff
}

// WithEventsFileOperationsQueryParameter adds the eventsFileOperations to the audit collection get params
func (o *AuditCollectionGetParams) WithEventsFileOperationsQueryParameter(eventsFileOperations *bool) *AuditCollectionGetParams {
	o.SetEventsFileOperationsQueryParameter(eventsFileOperations)
	return o
}

// SetEventsFileOperationsQueryParameter adds the eventsFileOperations to the audit collection get params
func (o *AuditCollectionGetParams) SetEventsFileOperationsQueryParameter(eventsFileOperations *bool) {
	o.EventsFileOperationsQueryParameter = eventsFileOperations
}

// WithEventsFileShareQueryParameter adds the eventsFileShare to the audit collection get params
func (o *AuditCollectionGetParams) WithEventsFileShareQueryParameter(eventsFileShare *bool) *AuditCollectionGetParams {
	o.SetEventsFileShareQueryParameter(eventsFileShare)
	return o
}

// SetEventsFileShareQueryParameter adds the eventsFileShare to the audit collection get params
func (o *AuditCollectionGetParams) SetEventsFileShareQueryParameter(eventsFileShare *bool) {
	o.EventsFileShareQueryParameter = eventsFileShare
}

// WithEventsSecurityGroupQueryParameter adds the eventsSecurityGroup to the audit collection get params
func (o *AuditCollectionGetParams) WithEventsSecurityGroupQueryParameter(eventsSecurityGroup *bool) *AuditCollectionGetParams {
	o.SetEventsSecurityGroupQueryParameter(eventsSecurityGroup)
	return o
}

// SetEventsSecurityGroupQueryParameter adds the eventsSecurityGroup to the audit collection get params
func (o *AuditCollectionGetParams) SetEventsSecurityGroupQueryParameter(eventsSecurityGroup *bool) {
	o.EventsSecurityGroupQueryParameter = eventsSecurityGroup
}

// WithEventsUserAccountQueryParameter adds the eventsUserAccount to the audit collection get params
func (o *AuditCollectionGetParams) WithEventsUserAccountQueryParameter(eventsUserAccount *bool) *AuditCollectionGetParams {
	o.SetEventsUserAccountQueryParameter(eventsUserAccount)
	return o
}

// SetEventsUserAccountQueryParameter adds the eventsUserAccount to the audit collection get params
func (o *AuditCollectionGetParams) SetEventsUserAccountQueryParameter(eventsUserAccount *bool) {
	o.EventsUserAccountQueryParameter = eventsUserAccount
}

// WithFields adds the fields to the audit collection get params
func (o *AuditCollectionGetParams) WithFields(fields []string) *AuditCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the audit collection get params
func (o *AuditCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLogFormatQueryParameter adds the logFormat to the audit collection get params
func (o *AuditCollectionGetParams) WithLogFormatQueryParameter(logFormat *string) *AuditCollectionGetParams {
	o.SetLogFormatQueryParameter(logFormat)
	return o
}

// SetLogFormatQueryParameter adds the logFormat to the audit collection get params
func (o *AuditCollectionGetParams) SetLogFormatQueryParameter(logFormat *string) {
	o.LogFormatQueryParameter = logFormat
}

// WithLogRetentionCountQueryParameter adds the logRetentionCount to the audit collection get params
func (o *AuditCollectionGetParams) WithLogRetentionCountQueryParameter(logRetentionCount *int64) *AuditCollectionGetParams {
	o.SetLogRetentionCountQueryParameter(logRetentionCount)
	return o
}

// SetLogRetentionCountQueryParameter adds the logRetentionCount to the audit collection get params
func (o *AuditCollectionGetParams) SetLogRetentionCountQueryParameter(logRetentionCount *int64) {
	o.LogRetentionCountQueryParameter = logRetentionCount
}

// WithLogRetentionDurationQueryParameter adds the logRetentionDuration to the audit collection get params
func (o *AuditCollectionGetParams) WithLogRetentionDurationQueryParameter(logRetentionDuration *string) *AuditCollectionGetParams {
	o.SetLogRetentionDurationQueryParameter(logRetentionDuration)
	return o
}

// SetLogRetentionDurationQueryParameter adds the logRetentionDuration to the audit collection get params
func (o *AuditCollectionGetParams) SetLogRetentionDurationQueryParameter(logRetentionDuration *string) {
	o.LogRetentionDurationQueryParameter = logRetentionDuration
}

// WithLogRotationScheduleDaysQueryParameter adds the logRotationScheduleDays to the audit collection get params
func (o *AuditCollectionGetParams) WithLogRotationScheduleDaysQueryParameter(logRotationScheduleDays *int64) *AuditCollectionGetParams {
	o.SetLogRotationScheduleDaysQueryParameter(logRotationScheduleDays)
	return o
}

// SetLogRotationScheduleDaysQueryParameter adds the logRotationScheduleDays to the audit collection get params
func (o *AuditCollectionGetParams) SetLogRotationScheduleDaysQueryParameter(logRotationScheduleDays *int64) {
	o.LogRotationScheduleDaysQueryParameter = logRotationScheduleDays
}

// WithLogRotationScheduleHoursQueryParameter adds the logRotationScheduleHours to the audit collection get params
func (o *AuditCollectionGetParams) WithLogRotationScheduleHoursQueryParameter(logRotationScheduleHours *int64) *AuditCollectionGetParams {
	o.SetLogRotationScheduleHoursQueryParameter(logRotationScheduleHours)
	return o
}

// SetLogRotationScheduleHoursQueryParameter adds the logRotationScheduleHours to the audit collection get params
func (o *AuditCollectionGetParams) SetLogRotationScheduleHoursQueryParameter(logRotationScheduleHours *int64) {
	o.LogRotationScheduleHoursQueryParameter = logRotationScheduleHours
}

// WithLogRotationScheduleMinutesQueryParameter adds the logRotationScheduleMinutes to the audit collection get params
func (o *AuditCollectionGetParams) WithLogRotationScheduleMinutesQueryParameter(logRotationScheduleMinutes *int64) *AuditCollectionGetParams {
	o.SetLogRotationScheduleMinutesQueryParameter(logRotationScheduleMinutes)
	return o
}

// SetLogRotationScheduleMinutesQueryParameter adds the logRotationScheduleMinutes to the audit collection get params
func (o *AuditCollectionGetParams) SetLogRotationScheduleMinutesQueryParameter(logRotationScheduleMinutes *int64) {
	o.LogRotationScheduleMinutesQueryParameter = logRotationScheduleMinutes
}

// WithLogRotationScheduleMonthsQueryParameter adds the logRotationScheduleMonths to the audit collection get params
func (o *AuditCollectionGetParams) WithLogRotationScheduleMonthsQueryParameter(logRotationScheduleMonths *int64) *AuditCollectionGetParams {
	o.SetLogRotationScheduleMonthsQueryParameter(logRotationScheduleMonths)
	return o
}

// SetLogRotationScheduleMonthsQueryParameter adds the logRotationScheduleMonths to the audit collection get params
func (o *AuditCollectionGetParams) SetLogRotationScheduleMonthsQueryParameter(logRotationScheduleMonths *int64) {
	o.LogRotationScheduleMonthsQueryParameter = logRotationScheduleMonths
}

// WithLogRotationScheduleWeekdaysQueryParameter adds the logRotationScheduleWeekdays to the audit collection get params
func (o *AuditCollectionGetParams) WithLogRotationScheduleWeekdaysQueryParameter(logRotationScheduleWeekdays *int64) *AuditCollectionGetParams {
	o.SetLogRotationScheduleWeekdaysQueryParameter(logRotationScheduleWeekdays)
	return o
}

// SetLogRotationScheduleWeekdaysQueryParameter adds the logRotationScheduleWeekdays to the audit collection get params
func (o *AuditCollectionGetParams) SetLogRotationScheduleWeekdaysQueryParameter(logRotationScheduleWeekdays *int64) {
	o.LogRotationScheduleWeekdaysQueryParameter = logRotationScheduleWeekdays
}

// WithLogRotationSizeQueryParameter adds the logRotationSize to the audit collection get params
func (o *AuditCollectionGetParams) WithLogRotationSizeQueryParameter(logRotationSize *int64) *AuditCollectionGetParams {
	o.SetLogRotationSizeQueryParameter(logRotationSize)
	return o
}

// SetLogRotationSizeQueryParameter adds the logRotationSize to the audit collection get params
func (o *AuditCollectionGetParams) SetLogRotationSizeQueryParameter(logRotationSize *int64) {
	o.LogRotationSizeQueryParameter = logRotationSize
}

// WithLogPathQueryParameter adds the logPath to the audit collection get params
func (o *AuditCollectionGetParams) WithLogPathQueryParameter(logPath *string) *AuditCollectionGetParams {
	o.SetLogPathQueryParameter(logPath)
	return o
}

// SetLogPathQueryParameter adds the logPath to the audit collection get params
func (o *AuditCollectionGetParams) SetLogPathQueryParameter(logPath *string) {
	o.LogPathQueryParameter = logPath
}

// WithMaxRecords adds the maxRecords to the audit collection get params
func (o *AuditCollectionGetParams) WithMaxRecords(maxRecords *int64) *AuditCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the audit collection get params
func (o *AuditCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the audit collection get params
func (o *AuditCollectionGetParams) WithOrderBy(orderBy []string) *AuditCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the audit collection get params
func (o *AuditCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the audit collection get params
func (o *AuditCollectionGetParams) WithReturnRecords(returnRecords *bool) *AuditCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the audit collection get params
func (o *AuditCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the audit collection get params
func (o *AuditCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *AuditCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the audit collection get params
func (o *AuditCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the audit collection get params
func (o *AuditCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *AuditCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the audit collection get params
func (o *AuditCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the audit collection get params
func (o *AuditCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *AuditCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the audit collection get params
func (o *AuditCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *AuditCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EventsAuthorizationPolicyQueryParameter != nil {

		// query param events.authorization_policy
		var qrEventsAuthorizationPolicy bool

		if o.EventsAuthorizationPolicyQueryParameter != nil {
			qrEventsAuthorizationPolicy = *o.EventsAuthorizationPolicyQueryParameter
		}
		qEventsAuthorizationPolicy := swag.FormatBool(qrEventsAuthorizationPolicy)
		if qEventsAuthorizationPolicy != "" {

			if err := r.SetQueryParam("events.authorization_policy", qEventsAuthorizationPolicy); err != nil {
				return err
			}
		}
	}

	if o.EventsCapStagingQueryParameter != nil {

		// query param events.cap_staging
		var qrEventsCapStaging bool

		if o.EventsCapStagingQueryParameter != nil {
			qrEventsCapStaging = *o.EventsCapStagingQueryParameter
		}
		qEventsCapStaging := swag.FormatBool(qrEventsCapStaging)
		if qEventsCapStaging != "" {

			if err := r.SetQueryParam("events.cap_staging", qEventsCapStaging); err != nil {
				return err
			}
		}
	}

	if o.EventsCifsLogonLogoffQueryParameter != nil {

		// query param events.cifs_logon_logoff
		var qrEventsCifsLogonLogoff bool

		if o.EventsCifsLogonLogoffQueryParameter != nil {
			qrEventsCifsLogonLogoff = *o.EventsCifsLogonLogoffQueryParameter
		}
		qEventsCifsLogonLogoff := swag.FormatBool(qrEventsCifsLogonLogoff)
		if qEventsCifsLogonLogoff != "" {

			if err := r.SetQueryParam("events.cifs_logon_logoff", qEventsCifsLogonLogoff); err != nil {
				return err
			}
		}
	}

	if o.EventsFileOperationsQueryParameter != nil {

		// query param events.file_operations
		var qrEventsFileOperations bool

		if o.EventsFileOperationsQueryParameter != nil {
			qrEventsFileOperations = *o.EventsFileOperationsQueryParameter
		}
		qEventsFileOperations := swag.FormatBool(qrEventsFileOperations)
		if qEventsFileOperations != "" {

			if err := r.SetQueryParam("events.file_operations", qEventsFileOperations); err != nil {
				return err
			}
		}
	}

	if o.EventsFileShareQueryParameter != nil {

		// query param events.file_share
		var qrEventsFileShare bool

		if o.EventsFileShareQueryParameter != nil {
			qrEventsFileShare = *o.EventsFileShareQueryParameter
		}
		qEventsFileShare := swag.FormatBool(qrEventsFileShare)
		if qEventsFileShare != "" {

			if err := r.SetQueryParam("events.file_share", qEventsFileShare); err != nil {
				return err
			}
		}
	}

	if o.EventsSecurityGroupQueryParameter != nil {

		// query param events.security_group
		var qrEventsSecurityGroup bool

		if o.EventsSecurityGroupQueryParameter != nil {
			qrEventsSecurityGroup = *o.EventsSecurityGroupQueryParameter
		}
		qEventsSecurityGroup := swag.FormatBool(qrEventsSecurityGroup)
		if qEventsSecurityGroup != "" {

			if err := r.SetQueryParam("events.security_group", qEventsSecurityGroup); err != nil {
				return err
			}
		}
	}

	if o.EventsUserAccountQueryParameter != nil {

		// query param events.user_account
		var qrEventsUserAccount bool

		if o.EventsUserAccountQueryParameter != nil {
			qrEventsUserAccount = *o.EventsUserAccountQueryParameter
		}
		qEventsUserAccount := swag.FormatBool(qrEventsUserAccount)
		if qEventsUserAccount != "" {

			if err := r.SetQueryParam("events.user_account", qEventsUserAccount); err != nil {
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

	if o.LogFormatQueryParameter != nil {

		// query param log.format
		var qrLogFormat string

		if o.LogFormatQueryParameter != nil {
			qrLogFormat = *o.LogFormatQueryParameter
		}
		qLogFormat := qrLogFormat
		if qLogFormat != "" {

			if err := r.SetQueryParam("log.format", qLogFormat); err != nil {
				return err
			}
		}
	}

	if o.LogRetentionCountQueryParameter != nil {

		// query param log.retention.count
		var qrLogRetentionCount int64

		if o.LogRetentionCountQueryParameter != nil {
			qrLogRetentionCount = *o.LogRetentionCountQueryParameter
		}
		qLogRetentionCount := swag.FormatInt64(qrLogRetentionCount)
		if qLogRetentionCount != "" {

			if err := r.SetQueryParam("log.retention.count", qLogRetentionCount); err != nil {
				return err
			}
		}
	}

	if o.LogRetentionDurationQueryParameter != nil {

		// query param log.retention.duration
		var qrLogRetentionDuration string

		if o.LogRetentionDurationQueryParameter != nil {
			qrLogRetentionDuration = *o.LogRetentionDurationQueryParameter
		}
		qLogRetentionDuration := qrLogRetentionDuration
		if qLogRetentionDuration != "" {

			if err := r.SetQueryParam("log.retention.duration", qLogRetentionDuration); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleDaysQueryParameter != nil {

		// query param log.rotation.schedule.days
		var qrLogRotationScheduleDays int64

		if o.LogRotationScheduleDaysQueryParameter != nil {
			qrLogRotationScheduleDays = *o.LogRotationScheduleDaysQueryParameter
		}
		qLogRotationScheduleDays := swag.FormatInt64(qrLogRotationScheduleDays)
		if qLogRotationScheduleDays != "" {

			if err := r.SetQueryParam("log.rotation.schedule.days", qLogRotationScheduleDays); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleHoursQueryParameter != nil {

		// query param log.rotation.schedule.hours
		var qrLogRotationScheduleHours int64

		if o.LogRotationScheduleHoursQueryParameter != nil {
			qrLogRotationScheduleHours = *o.LogRotationScheduleHoursQueryParameter
		}
		qLogRotationScheduleHours := swag.FormatInt64(qrLogRotationScheduleHours)
		if qLogRotationScheduleHours != "" {

			if err := r.SetQueryParam("log.rotation.schedule.hours", qLogRotationScheduleHours); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleMinutesQueryParameter != nil {

		// query param log.rotation.schedule.minutes
		var qrLogRotationScheduleMinutes int64

		if o.LogRotationScheduleMinutesQueryParameter != nil {
			qrLogRotationScheduleMinutes = *o.LogRotationScheduleMinutesQueryParameter
		}
		qLogRotationScheduleMinutes := swag.FormatInt64(qrLogRotationScheduleMinutes)
		if qLogRotationScheduleMinutes != "" {

			if err := r.SetQueryParam("log.rotation.schedule.minutes", qLogRotationScheduleMinutes); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleMonthsQueryParameter != nil {

		// query param log.rotation.schedule.months
		var qrLogRotationScheduleMonths int64

		if o.LogRotationScheduleMonthsQueryParameter != nil {
			qrLogRotationScheduleMonths = *o.LogRotationScheduleMonthsQueryParameter
		}
		qLogRotationScheduleMonths := swag.FormatInt64(qrLogRotationScheduleMonths)
		if qLogRotationScheduleMonths != "" {

			if err := r.SetQueryParam("log.rotation.schedule.months", qLogRotationScheduleMonths); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleWeekdaysQueryParameter != nil {

		// query param log.rotation.schedule.weekdays
		var qrLogRotationScheduleWeekdays int64

		if o.LogRotationScheduleWeekdaysQueryParameter != nil {
			qrLogRotationScheduleWeekdays = *o.LogRotationScheduleWeekdaysQueryParameter
		}
		qLogRotationScheduleWeekdays := swag.FormatInt64(qrLogRotationScheduleWeekdays)
		if qLogRotationScheduleWeekdays != "" {

			if err := r.SetQueryParam("log.rotation.schedule.weekdays", qLogRotationScheduleWeekdays); err != nil {
				return err
			}
		}
	}

	if o.LogRotationSizeQueryParameter != nil {

		// query param log.rotation.size
		var qrLogRotationSize int64

		if o.LogRotationSizeQueryParameter != nil {
			qrLogRotationSize = *o.LogRotationSizeQueryParameter
		}
		qLogRotationSize := swag.FormatInt64(qrLogRotationSize)
		if qLogRotationSize != "" {

			if err := r.SetQueryParam("log.rotation.size", qLogRotationSize); err != nil {
				return err
			}
		}
	}

	if o.LogPathQueryParameter != nil {

		// query param log_path
		var qrLogPath string

		if o.LogPathQueryParameter != nil {
			qrLogPath = *o.LogPathQueryParameter
		}
		qLogPath := qrLogPath
		if qLogPath != "" {

			if err := r.SetQueryParam("log_path", qLogPath); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamAuditCollectionGet binds the parameter fields
func (o *AuditCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamAuditCollectionGet binds the parameter order_by
func (o *AuditCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
