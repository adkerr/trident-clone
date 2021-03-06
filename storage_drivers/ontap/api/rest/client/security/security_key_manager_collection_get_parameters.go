// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityKeyManagerCollectionGetParams creates a new SecurityKeyManagerCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeyManagerCollectionGetParams() *SecurityKeyManagerCollectionGetParams {
	return &SecurityKeyManagerCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeyManagerCollectionGetParamsWithTimeout creates a new SecurityKeyManagerCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSecurityKeyManagerCollectionGetParamsWithTimeout(timeout time.Duration) *SecurityKeyManagerCollectionGetParams {
	return &SecurityKeyManagerCollectionGetParams{
		timeout: timeout,
	}
}

// NewSecurityKeyManagerCollectionGetParamsWithContext creates a new SecurityKeyManagerCollectionGetParams object
// with the ability to set a context for a request.
func NewSecurityKeyManagerCollectionGetParamsWithContext(ctx context.Context) *SecurityKeyManagerCollectionGetParams {
	return &SecurityKeyManagerCollectionGetParams{
		Context: ctx,
	}
}

// NewSecurityKeyManagerCollectionGetParamsWithHTTPClient creates a new SecurityKeyManagerCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeyManagerCollectionGetParamsWithHTTPClient(client *http.Client) *SecurityKeyManagerCollectionGetParams {
	return &SecurityKeyManagerCollectionGetParams{
		HTTPClient: client,
	}
}

/* SecurityKeyManagerCollectionGetParams contains all the parameters to send to the API endpoint
   for the security key manager collection get operation.

   Typically these are written to a http.Request.
*/
type SecurityKeyManagerCollectionGetParams struct {

	/* ExternalClientCertificateUUID.

	   Filter by external.client_certificate.uuid
	*/
	ExternalClientCertificateUUIDQueryParameter *string

	/* ExternalServerCaCertificatesUUID.

	   Filter by external.server_ca_certificates.uuid
	*/
	ExternalServerCaCertificatesUUIDQueryParameter *string

	/* ExternalServersConnectivityClusterAvailability.

	   Filter by external.servers.connectivity.cluster_availability
	*/
	ExternalServersConnectivityClusterAvailabilityQueryParameter *bool

	/* ExternalServersConnectivityRecordsNodeName.

	   Filter by external.servers.connectivity.records.node.name
	*/
	ExternalServersConnectivityRecordsNodeNameQueryParameter *string

	/* ExternalServersConnectivityRecordsNodeUUID.

	   Filter by external.servers.connectivity.records.node.uuid
	*/
	ExternalServersConnectivityRecordsNodeUUIDQueryParameter *string

	/* ExternalServersConnectivityRecordsState.

	   Filter by external.servers.connectivity.records.state
	*/
	ExternalServersConnectivityRecordsStateQueryParameter *string

	/* ExternalServersServer.

	   Filter by external.servers.server
	*/
	ExternalServersServerQueryParameter *string

	/* ExternalServersTimeout.

	   Filter by external.servers.timeout
	*/
	ExternalServersTimeoutQueryParameter *int64

	/* ExternalServersUsername.

	   Filter by external.servers.username
	*/
	ExternalServersUsernameQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IsDefaultDataAtRestEncryptionDisabled.

	   Filter by is_default_data_at_rest_encryption_disabled
	*/
	IsDefaultDataAtRestEncryptionDisabledQueryParameter *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OnboardEnabled.

	   Filter by onboard.enabled
	*/
	OnboardEnabledQueryParameter *bool

	/* OnboardKeyBackup.

	   Filter by onboard.key_backup
	*/
	OnboardKeyBackupQueryParameter *string

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

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	/* StatusCode.

	   Filter by status.code
	*/
	StatusCodeQueryParameter *int64

	/* StatusMessage.

	   Filter by status.message
	*/
	StatusMessageQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	/* VolumeEncryptionCode.

	   Filter by volume_encryption.code
	*/
	VolumeEncryptionCodeQueryParameter *int64

	/* VolumeEncryptionMessage.

	   Filter by volume_encryption.message
	*/
	VolumeEncryptionMessageQueryParameter *string

	/* VolumeEncryptionSupported.

	   Filter by volume_encryption.supported
	*/
	VolumeEncryptionSupportedQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security key manager collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerCollectionGetParams) WithDefaults() *SecurityKeyManagerCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security key manager collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SecurityKeyManagerCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithTimeout(timeout time.Duration) *SecurityKeyManagerCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithContext(ctx context.Context) *SecurityKeyManagerCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithHTTPClient(client *http.Client) *SecurityKeyManagerCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalClientCertificateUUIDQueryParameter adds the externalClientCertificateUUID to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalClientCertificateUUIDQueryParameter(externalClientCertificateUUID *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalClientCertificateUUIDQueryParameter(externalClientCertificateUUID)
	return o
}

// SetExternalClientCertificateUUIDQueryParameter adds the externalClientCertificateUuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalClientCertificateUUIDQueryParameter(externalClientCertificateUUID *string) {
	o.ExternalClientCertificateUUIDQueryParameter = externalClientCertificateUUID
}

// WithExternalServerCaCertificatesUUIDQueryParameter adds the externalServerCaCertificatesUUID to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServerCaCertificatesUUIDQueryParameter(externalServerCaCertificatesUUID *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServerCaCertificatesUUIDQueryParameter(externalServerCaCertificatesUUID)
	return o
}

// SetExternalServerCaCertificatesUUIDQueryParameter adds the externalServerCaCertificatesUuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServerCaCertificatesUUIDQueryParameter(externalServerCaCertificatesUUID *string) {
	o.ExternalServerCaCertificatesUUIDQueryParameter = externalServerCaCertificatesUUID
}

// WithExternalServersConnectivityClusterAvailabilityQueryParameter adds the externalServersConnectivityClusterAvailability to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersConnectivityClusterAvailabilityQueryParameter(externalServersConnectivityClusterAvailability *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersConnectivityClusterAvailabilityQueryParameter(externalServersConnectivityClusterAvailability)
	return o
}

// SetExternalServersConnectivityClusterAvailabilityQueryParameter adds the externalServersConnectivityClusterAvailability to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersConnectivityClusterAvailabilityQueryParameter(externalServersConnectivityClusterAvailability *bool) {
	o.ExternalServersConnectivityClusterAvailabilityQueryParameter = externalServersConnectivityClusterAvailability
}

// WithExternalServersConnectivityRecordsNodeNameQueryParameter adds the externalServersConnectivityRecordsNodeName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersConnectivityRecordsNodeNameQueryParameter(externalServersConnectivityRecordsNodeName *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersConnectivityRecordsNodeNameQueryParameter(externalServersConnectivityRecordsNodeName)
	return o
}

// SetExternalServersConnectivityRecordsNodeNameQueryParameter adds the externalServersConnectivityRecordsNodeName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersConnectivityRecordsNodeNameQueryParameter(externalServersConnectivityRecordsNodeName *string) {
	o.ExternalServersConnectivityRecordsNodeNameQueryParameter = externalServersConnectivityRecordsNodeName
}

// WithExternalServersConnectivityRecordsNodeUUIDQueryParameter adds the externalServersConnectivityRecordsNodeUUID to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersConnectivityRecordsNodeUUIDQueryParameter(externalServersConnectivityRecordsNodeUUID *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersConnectivityRecordsNodeUUIDQueryParameter(externalServersConnectivityRecordsNodeUUID)
	return o
}

// SetExternalServersConnectivityRecordsNodeUUIDQueryParameter adds the externalServersConnectivityRecordsNodeUuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersConnectivityRecordsNodeUUIDQueryParameter(externalServersConnectivityRecordsNodeUUID *string) {
	o.ExternalServersConnectivityRecordsNodeUUIDQueryParameter = externalServersConnectivityRecordsNodeUUID
}

// WithExternalServersConnectivityRecordsStateQueryParameter adds the externalServersConnectivityRecordsState to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersConnectivityRecordsStateQueryParameter(externalServersConnectivityRecordsState *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersConnectivityRecordsStateQueryParameter(externalServersConnectivityRecordsState)
	return o
}

// SetExternalServersConnectivityRecordsStateQueryParameter adds the externalServersConnectivityRecordsState to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersConnectivityRecordsStateQueryParameter(externalServersConnectivityRecordsState *string) {
	o.ExternalServersConnectivityRecordsStateQueryParameter = externalServersConnectivityRecordsState
}

// WithExternalServersServerQueryParameter adds the externalServersServer to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersServerQueryParameter(externalServersServer *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersServerQueryParameter(externalServersServer)
	return o
}

// SetExternalServersServerQueryParameter adds the externalServersServer to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersServerQueryParameter(externalServersServer *string) {
	o.ExternalServersServerQueryParameter = externalServersServer
}

// WithExternalServersTimeoutQueryParameter adds the externalServersTimeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersTimeoutQueryParameter(externalServersTimeout *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersTimeoutQueryParameter(externalServersTimeout)
	return o
}

// SetExternalServersTimeoutQueryParameter adds the externalServersTimeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersTimeoutQueryParameter(externalServersTimeout *int64) {
	o.ExternalServersTimeoutQueryParameter = externalServersTimeout
}

// WithExternalServersUsernameQueryParameter adds the externalServersUsername to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersUsernameQueryParameter(externalServersUsername *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersUsernameQueryParameter(externalServersUsername)
	return o
}

// SetExternalServersUsernameQueryParameter adds the externalServersUsername to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersUsernameQueryParameter(externalServersUsername *string) {
	o.ExternalServersUsernameQueryParameter = externalServersUsername
}

// WithFields adds the fields to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithFields(fields []string) *SecurityKeyManagerCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIsDefaultDataAtRestEncryptionDisabledQueryParameter adds the isDefaultDataAtRestEncryptionDisabled to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithIsDefaultDataAtRestEncryptionDisabledQueryParameter(isDefaultDataAtRestEncryptionDisabled *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetIsDefaultDataAtRestEncryptionDisabledQueryParameter(isDefaultDataAtRestEncryptionDisabled)
	return o
}

// SetIsDefaultDataAtRestEncryptionDisabledQueryParameter adds the isDefaultDataAtRestEncryptionDisabled to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetIsDefaultDataAtRestEncryptionDisabledQueryParameter(isDefaultDataAtRestEncryptionDisabled *bool) {
	o.IsDefaultDataAtRestEncryptionDisabledQueryParameter = isDefaultDataAtRestEncryptionDisabled
}

// WithMaxRecords adds the maxRecords to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithMaxRecords(maxRecords *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOnboardEnabledQueryParameter adds the onboardEnabled to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithOnboardEnabledQueryParameter(onboardEnabled *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetOnboardEnabledQueryParameter(onboardEnabled)
	return o
}

// SetOnboardEnabledQueryParameter adds the onboardEnabled to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetOnboardEnabledQueryParameter(onboardEnabled *bool) {
	o.OnboardEnabledQueryParameter = onboardEnabled
}

// WithOnboardKeyBackupQueryParameter adds the onboardKeyBackup to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithOnboardKeyBackupQueryParameter(onboardKeyBackup *string) *SecurityKeyManagerCollectionGetParams {
	o.SetOnboardKeyBackupQueryParameter(onboardKeyBackup)
	return o
}

// SetOnboardKeyBackupQueryParameter adds the onboardKeyBackup to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetOnboardKeyBackupQueryParameter(onboardKeyBackup *string) {
	o.OnboardKeyBackupQueryParameter = onboardKeyBackup
}

// WithOrderBy adds the orderBy to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithOrderBy(orderBy []string) *SecurityKeyManagerCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithReturnRecords(returnRecords *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScopeQueryParameter adds the scope to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithScopeQueryParameter(scope *string) *SecurityKeyManagerCollectionGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithStatusCodeQueryParameter adds the statusCode to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithStatusCodeQueryParameter(statusCode *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetStatusCodeQueryParameter(statusCode)
	return o
}

// SetStatusCodeQueryParameter adds the statusCode to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetStatusCodeQueryParameter(statusCode *int64) {
	o.StatusCodeQueryParameter = statusCode
}

// WithStatusMessageQueryParameter adds the statusMessage to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithStatusMessageQueryParameter(statusMessage *string) *SecurityKeyManagerCollectionGetParams {
	o.SetStatusMessageQueryParameter(statusMessage)
	return o
}

// SetStatusMessageQueryParameter adds the statusMessage to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetStatusMessageQueryParameter(statusMessage *string) {
	o.StatusMessageQueryParameter = statusMessage
}

// WithSVMNameQueryParameter adds the svmName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *SecurityKeyManagerCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *SecurityKeyManagerCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUUIDQueryParameter adds the uuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithUUIDQueryParameter(uuid *string) *SecurityKeyManagerCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WithVolumeEncryptionCodeQueryParameter adds the volumeEncryptionCode to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithVolumeEncryptionCodeQueryParameter(volumeEncryptionCode *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetVolumeEncryptionCodeQueryParameter(volumeEncryptionCode)
	return o
}

// SetVolumeEncryptionCodeQueryParameter adds the volumeEncryptionCode to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetVolumeEncryptionCodeQueryParameter(volumeEncryptionCode *int64) {
	o.VolumeEncryptionCodeQueryParameter = volumeEncryptionCode
}

// WithVolumeEncryptionMessageQueryParameter adds the volumeEncryptionMessage to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithVolumeEncryptionMessageQueryParameter(volumeEncryptionMessage *string) *SecurityKeyManagerCollectionGetParams {
	o.SetVolumeEncryptionMessageQueryParameter(volumeEncryptionMessage)
	return o
}

// SetVolumeEncryptionMessageQueryParameter adds the volumeEncryptionMessage to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetVolumeEncryptionMessageQueryParameter(volumeEncryptionMessage *string) {
	o.VolumeEncryptionMessageQueryParameter = volumeEncryptionMessage
}

// WithVolumeEncryptionSupportedQueryParameter adds the volumeEncryptionSupported to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithVolumeEncryptionSupportedQueryParameter(volumeEncryptionSupported *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetVolumeEncryptionSupportedQueryParameter(volumeEncryptionSupported)
	return o
}

// SetVolumeEncryptionSupportedQueryParameter adds the volumeEncryptionSupported to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetVolumeEncryptionSupportedQueryParameter(volumeEncryptionSupported *bool) {
	o.VolumeEncryptionSupportedQueryParameter = volumeEncryptionSupported
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeyManagerCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExternalClientCertificateUUIDQueryParameter != nil {

		// query param external.client_certificate.uuid
		var qrExternalClientCertificateUUID string

		if o.ExternalClientCertificateUUIDQueryParameter != nil {
			qrExternalClientCertificateUUID = *o.ExternalClientCertificateUUIDQueryParameter
		}
		qExternalClientCertificateUUID := qrExternalClientCertificateUUID
		if qExternalClientCertificateUUID != "" {

			if err := r.SetQueryParam("external.client_certificate.uuid", qExternalClientCertificateUUID); err != nil {
				return err
			}
		}
	}

	if o.ExternalServerCaCertificatesUUIDQueryParameter != nil {

		// query param external.server_ca_certificates.uuid
		var qrExternalServerCaCertificatesUUID string

		if o.ExternalServerCaCertificatesUUIDQueryParameter != nil {
			qrExternalServerCaCertificatesUUID = *o.ExternalServerCaCertificatesUUIDQueryParameter
		}
		qExternalServerCaCertificatesUUID := qrExternalServerCaCertificatesUUID
		if qExternalServerCaCertificatesUUID != "" {

			if err := r.SetQueryParam("external.server_ca_certificates.uuid", qExternalServerCaCertificatesUUID); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersConnectivityClusterAvailabilityQueryParameter != nil {

		// query param external.servers.connectivity.cluster_availability
		var qrExternalServersConnectivityClusterAvailability bool

		if o.ExternalServersConnectivityClusterAvailabilityQueryParameter != nil {
			qrExternalServersConnectivityClusterAvailability = *o.ExternalServersConnectivityClusterAvailabilityQueryParameter
		}
		qExternalServersConnectivityClusterAvailability := swag.FormatBool(qrExternalServersConnectivityClusterAvailability)
		if qExternalServersConnectivityClusterAvailability != "" {

			if err := r.SetQueryParam("external.servers.connectivity.cluster_availability", qExternalServersConnectivityClusterAvailability); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersConnectivityRecordsNodeNameQueryParameter != nil {

		// query param external.servers.connectivity.records.node.name
		var qrExternalServersConnectivityRecordsNodeName string

		if o.ExternalServersConnectivityRecordsNodeNameQueryParameter != nil {
			qrExternalServersConnectivityRecordsNodeName = *o.ExternalServersConnectivityRecordsNodeNameQueryParameter
		}
		qExternalServersConnectivityRecordsNodeName := qrExternalServersConnectivityRecordsNodeName
		if qExternalServersConnectivityRecordsNodeName != "" {

			if err := r.SetQueryParam("external.servers.connectivity.records.node.name", qExternalServersConnectivityRecordsNodeName); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersConnectivityRecordsNodeUUIDQueryParameter != nil {

		// query param external.servers.connectivity.records.node.uuid
		var qrExternalServersConnectivityRecordsNodeUUID string

		if o.ExternalServersConnectivityRecordsNodeUUIDQueryParameter != nil {
			qrExternalServersConnectivityRecordsNodeUUID = *o.ExternalServersConnectivityRecordsNodeUUIDQueryParameter
		}
		qExternalServersConnectivityRecordsNodeUUID := qrExternalServersConnectivityRecordsNodeUUID
		if qExternalServersConnectivityRecordsNodeUUID != "" {

			if err := r.SetQueryParam("external.servers.connectivity.records.node.uuid", qExternalServersConnectivityRecordsNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersConnectivityRecordsStateQueryParameter != nil {

		// query param external.servers.connectivity.records.state
		var qrExternalServersConnectivityRecordsState string

		if o.ExternalServersConnectivityRecordsStateQueryParameter != nil {
			qrExternalServersConnectivityRecordsState = *o.ExternalServersConnectivityRecordsStateQueryParameter
		}
		qExternalServersConnectivityRecordsState := qrExternalServersConnectivityRecordsState
		if qExternalServersConnectivityRecordsState != "" {

			if err := r.SetQueryParam("external.servers.connectivity.records.state", qExternalServersConnectivityRecordsState); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersServerQueryParameter != nil {

		// query param external.servers.server
		var qrExternalServersServer string

		if o.ExternalServersServerQueryParameter != nil {
			qrExternalServersServer = *o.ExternalServersServerQueryParameter
		}
		qExternalServersServer := qrExternalServersServer
		if qExternalServersServer != "" {

			if err := r.SetQueryParam("external.servers.server", qExternalServersServer); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersTimeoutQueryParameter != nil {

		// query param external.servers.timeout
		var qrExternalServersTimeout int64

		if o.ExternalServersTimeoutQueryParameter != nil {
			qrExternalServersTimeout = *o.ExternalServersTimeoutQueryParameter
		}
		qExternalServersTimeout := swag.FormatInt64(qrExternalServersTimeout)
		if qExternalServersTimeout != "" {

			if err := r.SetQueryParam("external.servers.timeout", qExternalServersTimeout); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersUsernameQueryParameter != nil {

		// query param external.servers.username
		var qrExternalServersUsername string

		if o.ExternalServersUsernameQueryParameter != nil {
			qrExternalServersUsername = *o.ExternalServersUsernameQueryParameter
		}
		qExternalServersUsername := qrExternalServersUsername
		if qExternalServersUsername != "" {

			if err := r.SetQueryParam("external.servers.username", qExternalServersUsername); err != nil {
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

	if o.IsDefaultDataAtRestEncryptionDisabledQueryParameter != nil {

		// query param is_default_data_at_rest_encryption_disabled
		var qrIsDefaultDataAtRestEncryptionDisabled bool

		if o.IsDefaultDataAtRestEncryptionDisabledQueryParameter != nil {
			qrIsDefaultDataAtRestEncryptionDisabled = *o.IsDefaultDataAtRestEncryptionDisabledQueryParameter
		}
		qIsDefaultDataAtRestEncryptionDisabled := swag.FormatBool(qrIsDefaultDataAtRestEncryptionDisabled)
		if qIsDefaultDataAtRestEncryptionDisabled != "" {

			if err := r.SetQueryParam("is_default_data_at_rest_encryption_disabled", qIsDefaultDataAtRestEncryptionDisabled); err != nil {
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

	if o.OnboardEnabledQueryParameter != nil {

		// query param onboard.enabled
		var qrOnboardEnabled bool

		if o.OnboardEnabledQueryParameter != nil {
			qrOnboardEnabled = *o.OnboardEnabledQueryParameter
		}
		qOnboardEnabled := swag.FormatBool(qrOnboardEnabled)
		if qOnboardEnabled != "" {

			if err := r.SetQueryParam("onboard.enabled", qOnboardEnabled); err != nil {
				return err
			}
		}
	}

	if o.OnboardKeyBackupQueryParameter != nil {

		// query param onboard.key_backup
		var qrOnboardKeyBackup string

		if o.OnboardKeyBackupQueryParameter != nil {
			qrOnboardKeyBackup = *o.OnboardKeyBackupQueryParameter
		}
		qOnboardKeyBackup := qrOnboardKeyBackup
		if qOnboardKeyBackup != "" {

			if err := r.SetQueryParam("onboard.key_backup", qOnboardKeyBackup); err != nil {
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

	if o.ScopeQueryParameter != nil {

		// query param scope
		var qrScope string

		if o.ScopeQueryParameter != nil {
			qrScope = *o.ScopeQueryParameter
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.StatusCodeQueryParameter != nil {

		// query param status.code
		var qrStatusCode int64

		if o.StatusCodeQueryParameter != nil {
			qrStatusCode = *o.StatusCodeQueryParameter
		}
		qStatusCode := swag.FormatInt64(qrStatusCode)
		if qStatusCode != "" {

			if err := r.SetQueryParam("status.code", qStatusCode); err != nil {
				return err
			}
		}
	}

	if o.StatusMessageQueryParameter != nil {

		// query param status.message
		var qrStatusMessage string

		if o.StatusMessageQueryParameter != nil {
			qrStatusMessage = *o.StatusMessageQueryParameter
		}
		qStatusMessage := qrStatusMessage
		if qStatusMessage != "" {

			if err := r.SetQueryParam("status.message", qStatusMessage); err != nil {
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

	if o.VolumeEncryptionCodeQueryParameter != nil {

		// query param volume_encryption.code
		var qrVolumeEncryptionCode int64

		if o.VolumeEncryptionCodeQueryParameter != nil {
			qrVolumeEncryptionCode = *o.VolumeEncryptionCodeQueryParameter
		}
		qVolumeEncryptionCode := swag.FormatInt64(qrVolumeEncryptionCode)
		if qVolumeEncryptionCode != "" {

			if err := r.SetQueryParam("volume_encryption.code", qVolumeEncryptionCode); err != nil {
				return err
			}
		}
	}

	if o.VolumeEncryptionMessageQueryParameter != nil {

		// query param volume_encryption.message
		var qrVolumeEncryptionMessage string

		if o.VolumeEncryptionMessageQueryParameter != nil {
			qrVolumeEncryptionMessage = *o.VolumeEncryptionMessageQueryParameter
		}
		qVolumeEncryptionMessage := qrVolumeEncryptionMessage
		if qVolumeEncryptionMessage != "" {

			if err := r.SetQueryParam("volume_encryption.message", qVolumeEncryptionMessage); err != nil {
				return err
			}
		}
	}

	if o.VolumeEncryptionSupportedQueryParameter != nil {

		// query param volume_encryption.supported
		var qrVolumeEncryptionSupported bool

		if o.VolumeEncryptionSupportedQueryParameter != nil {
			qrVolumeEncryptionSupported = *o.VolumeEncryptionSupportedQueryParameter
		}
		qVolumeEncryptionSupported := swag.FormatBool(qrVolumeEncryptionSupported)
		if qVolumeEncryptionSupported != "" {

			if err := r.SetQueryParam("volume_encryption.supported", qVolumeEncryptionSupported); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityKeyManagerCollectionGet binds the parameter fields
func (o *SecurityKeyManagerCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSecurityKeyManagerCollectionGet binds the parameter order_by
func (o *SecurityKeyManagerCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
