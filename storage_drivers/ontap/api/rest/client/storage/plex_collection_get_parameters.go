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

// NewPlexCollectionGetParams creates a new PlexCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlexCollectionGetParams() *PlexCollectionGetParams {
	return &PlexCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlexCollectionGetParamsWithTimeout creates a new PlexCollectionGetParams object
// with the ability to set a timeout on a request.
func NewPlexCollectionGetParamsWithTimeout(timeout time.Duration) *PlexCollectionGetParams {
	return &PlexCollectionGetParams{
		timeout: timeout,
	}
}

// NewPlexCollectionGetParamsWithContext creates a new PlexCollectionGetParams object
// with the ability to set a context for a request.
func NewPlexCollectionGetParamsWithContext(ctx context.Context) *PlexCollectionGetParams {
	return &PlexCollectionGetParams{
		Context: ctx,
	}
}

// NewPlexCollectionGetParamsWithHTTPClient creates a new PlexCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlexCollectionGetParamsWithHTTPClient(client *http.Client) *PlexCollectionGetParams {
	return &PlexCollectionGetParams{
		HTTPClient: client,
	}
}

/* PlexCollectionGetParams contains all the parameters to send to the API endpoint
   for the plex collection get operation.

   Typically these are written to a http.Request.
*/
type PlexCollectionGetParams struct {

	/* AggregateName.

	   Filter by aggregate.name
	*/
	AggregateNameQueryParameter *string

	/* AggregateUUID.

	   Filter by aggregate.uuid
	*/
	AggregateUUIDQueryParameter *string

	/* AggregateUUID.

	   Aggregate UUID
	*/
	AggregateUUIDPathParameter string

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

	/* Online.

	   Filter by online
	*/
	OnlineQueryParameter *bool

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Pool.

	   Filter by pool
	*/
	PoolQueryParameter *string

	/* RaidGroupsCacheTier.

	   Filter by raid_groups.cache_tier
	*/
	RaIDGroupsCacheTierQueryParameter *bool

	/* RaidGroupsDegraded.

	   Filter by raid_groups.degraded
	*/
	RaIDGroupsDegradedQueryParameter *bool

	/* RaidGroupsDisksDiskName.

	   Filter by raid_groups.disks.disk.name
	*/
	RaIDGroupsDisksDiskNameQueryParameter *string

	/* RaidGroupsDisksPosition.

	   Filter by raid_groups.disks.position
	*/
	RaIDGroupsDisksPositionQueryParameter *string

	/* RaidGroupsDisksState.

	   Filter by raid_groups.disks.state
	*/
	RaIDGroupsDisksStateQueryParameter *string

	/* RaidGroupsDisksType.

	   Filter by raid_groups.disks.type
	*/
	RaIDGroupsDisksTypeQueryParameter *string

	/* RaidGroupsDisksUsableSize.

	   Filter by raid_groups.disks.usable_size
	*/
	RaIDGroupsDisksUsableSizeQueryParameter *int64

	/* RaidGroupsName.

	   Filter by raid_groups.name
	*/
	RaIDGroupsNameQueryParameter *string

	/* RaidGroupsRecomputingParityActive.

	   Filter by raid_groups.recomputing_parity.active
	*/
	RaIDGroupsRecomputingParityActiveQueryParameter *bool

	/* RaidGroupsRecomputingParityPercent.

	   Filter by raid_groups.recomputing_parity.percent
	*/
	RaIDGroupsRecomputingParityPercentQueryParameter *int64

	/* RaidGroupsReconstructActive.

	   Filter by raid_groups.reconstruct.active
	*/
	RaIDGroupsReconstructActiveQueryParameter *bool

	/* RaidGroupsReconstructPercent.

	   Filter by raid_groups.reconstruct.percent
	*/
	RaIDGroupsReconstructPercentQueryParameter *int64

	/* ResyncActive.

	   Filter by resync.active
	*/
	ResyncActiveQueryParameter *bool

	/* ResyncLevel.

	   Filter by resync.level
	*/
	ResyncLevelQueryParameter *string

	/* ResyncPercent.

	   Filter by resync.percent
	*/
	ResyncPercentQueryParameter *int64

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

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plex collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlexCollectionGetParams) WithDefaults() *PlexCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plex collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlexCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := PlexCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the plex collection get params
func (o *PlexCollectionGetParams) WithTimeout(timeout time.Duration) *PlexCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plex collection get params
func (o *PlexCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plex collection get params
func (o *PlexCollectionGetParams) WithContext(ctx context.Context) *PlexCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plex collection get params
func (o *PlexCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plex collection get params
func (o *PlexCollectionGetParams) WithHTTPClient(client *http.Client) *PlexCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plex collection get params
func (o *PlexCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAggregateNameQueryParameter adds the aggregateName to the plex collection get params
func (o *PlexCollectionGetParams) WithAggregateNameQueryParameter(aggregateName *string) *PlexCollectionGetParams {
	o.SetAggregateNameQueryParameter(aggregateName)
	return o
}

// SetAggregateNameQueryParameter adds the aggregateName to the plex collection get params
func (o *PlexCollectionGetParams) SetAggregateNameQueryParameter(aggregateName *string) {
	o.AggregateNameQueryParameter = aggregateName
}

// WithAggregateUUIDQueryParameter adds the aggregateUUID to the plex collection get params
func (o *PlexCollectionGetParams) WithAggregateUUIDQueryParameter(aggregateUUID *string) *PlexCollectionGetParams {
	o.SetAggregateUUIDQueryParameter(aggregateUUID)
	return o
}

// SetAggregateUUIDQueryParameter adds the aggregateUuid to the plex collection get params
func (o *PlexCollectionGetParams) SetAggregateUUIDQueryParameter(aggregateUUID *string) {
	o.AggregateUUIDQueryParameter = aggregateUUID
}

// WithAggregateUUIDPathParameter adds the aggregateUUID to the plex collection get params
func (o *PlexCollectionGetParams) WithAggregateUUIDPathParameter(aggregateUUID string) *PlexCollectionGetParams {
	o.SetAggregateUUIDPathParameter(aggregateUUID)
	return o
}

// SetAggregateUUIDPathParameter adds the aggregateUuid to the plex collection get params
func (o *PlexCollectionGetParams) SetAggregateUUIDPathParameter(aggregateUUID string) {
	o.AggregateUUIDPathParameter = aggregateUUID
}

// WithFields adds the fields to the plex collection get params
func (o *PlexCollectionGetParams) WithFields(fields []string) *PlexCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the plex collection get params
func (o *PlexCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the plex collection get params
func (o *PlexCollectionGetParams) WithMaxRecords(maxRecords *int64) *PlexCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the plex collection get params
func (o *PlexCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNameQueryParameter adds the name to the plex collection get params
func (o *PlexCollectionGetParams) WithNameQueryParameter(name *string) *PlexCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the plex collection get params
func (o *PlexCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOnlineQueryParameter adds the online to the plex collection get params
func (o *PlexCollectionGetParams) WithOnlineQueryParameter(online *bool) *PlexCollectionGetParams {
	o.SetOnlineQueryParameter(online)
	return o
}

// SetOnlineQueryParameter adds the online to the plex collection get params
func (o *PlexCollectionGetParams) SetOnlineQueryParameter(online *bool) {
	o.OnlineQueryParameter = online
}

// WithOrderBy adds the orderBy to the plex collection get params
func (o *PlexCollectionGetParams) WithOrderBy(orderBy []string) *PlexCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the plex collection get params
func (o *PlexCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPoolQueryParameter adds the pool to the plex collection get params
func (o *PlexCollectionGetParams) WithPoolQueryParameter(pool *string) *PlexCollectionGetParams {
	o.SetPoolQueryParameter(pool)
	return o
}

// SetPoolQueryParameter adds the pool to the plex collection get params
func (o *PlexCollectionGetParams) SetPoolQueryParameter(pool *string) {
	o.PoolQueryParameter = pool
}

// WithRaIDGroupsCacheTierQueryParameter adds the raidGroupsCacheTier to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsCacheTierQueryParameter(raidGroupsCacheTier *bool) *PlexCollectionGetParams {
	o.SetRaIDGroupsCacheTierQueryParameter(raidGroupsCacheTier)
	return o
}

// SetRaIDGroupsCacheTierQueryParameter adds the raidGroupsCacheTier to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsCacheTierQueryParameter(raidGroupsCacheTier *bool) {
	o.RaIDGroupsCacheTierQueryParameter = raidGroupsCacheTier
}

// WithRaIDGroupsDegradedQueryParameter adds the raidGroupsDegraded to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsDegradedQueryParameter(raidGroupsDegraded *bool) *PlexCollectionGetParams {
	o.SetRaIDGroupsDegradedQueryParameter(raidGroupsDegraded)
	return o
}

// SetRaIDGroupsDegradedQueryParameter adds the raidGroupsDegraded to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsDegradedQueryParameter(raidGroupsDegraded *bool) {
	o.RaIDGroupsDegradedQueryParameter = raidGroupsDegraded
}

// WithRaIDGroupsDisksDiskNameQueryParameter adds the raidGroupsDisksDiskName to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsDisksDiskNameQueryParameter(raidGroupsDisksDiskName *string) *PlexCollectionGetParams {
	o.SetRaIDGroupsDisksDiskNameQueryParameter(raidGroupsDisksDiskName)
	return o
}

// SetRaIDGroupsDisksDiskNameQueryParameter adds the raidGroupsDisksDiskName to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsDisksDiskNameQueryParameter(raidGroupsDisksDiskName *string) {
	o.RaIDGroupsDisksDiskNameQueryParameter = raidGroupsDisksDiskName
}

// WithRaIDGroupsDisksPositionQueryParameter adds the raidGroupsDisksPosition to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsDisksPositionQueryParameter(raidGroupsDisksPosition *string) *PlexCollectionGetParams {
	o.SetRaIDGroupsDisksPositionQueryParameter(raidGroupsDisksPosition)
	return o
}

// SetRaIDGroupsDisksPositionQueryParameter adds the raidGroupsDisksPosition to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsDisksPositionQueryParameter(raidGroupsDisksPosition *string) {
	o.RaIDGroupsDisksPositionQueryParameter = raidGroupsDisksPosition
}

// WithRaIDGroupsDisksStateQueryParameter adds the raidGroupsDisksState to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsDisksStateQueryParameter(raidGroupsDisksState *string) *PlexCollectionGetParams {
	o.SetRaIDGroupsDisksStateQueryParameter(raidGroupsDisksState)
	return o
}

// SetRaIDGroupsDisksStateQueryParameter adds the raidGroupsDisksState to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsDisksStateQueryParameter(raidGroupsDisksState *string) {
	o.RaIDGroupsDisksStateQueryParameter = raidGroupsDisksState
}

// WithRaIDGroupsDisksTypeQueryParameter adds the raidGroupsDisksType to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsDisksTypeQueryParameter(raidGroupsDisksType *string) *PlexCollectionGetParams {
	o.SetRaIDGroupsDisksTypeQueryParameter(raidGroupsDisksType)
	return o
}

// SetRaIDGroupsDisksTypeQueryParameter adds the raidGroupsDisksType to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsDisksTypeQueryParameter(raidGroupsDisksType *string) {
	o.RaIDGroupsDisksTypeQueryParameter = raidGroupsDisksType
}

// WithRaIDGroupsDisksUsableSizeQueryParameter adds the raidGroupsDisksUsableSize to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsDisksUsableSizeQueryParameter(raidGroupsDisksUsableSize *int64) *PlexCollectionGetParams {
	o.SetRaIDGroupsDisksUsableSizeQueryParameter(raidGroupsDisksUsableSize)
	return o
}

// SetRaIDGroupsDisksUsableSizeQueryParameter adds the raidGroupsDisksUsableSize to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsDisksUsableSizeQueryParameter(raidGroupsDisksUsableSize *int64) {
	o.RaIDGroupsDisksUsableSizeQueryParameter = raidGroupsDisksUsableSize
}

// WithRaIDGroupsNameQueryParameter adds the raidGroupsName to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsNameQueryParameter(raidGroupsName *string) *PlexCollectionGetParams {
	o.SetRaIDGroupsNameQueryParameter(raidGroupsName)
	return o
}

// SetRaIDGroupsNameQueryParameter adds the raidGroupsName to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsNameQueryParameter(raidGroupsName *string) {
	o.RaIDGroupsNameQueryParameter = raidGroupsName
}

// WithRaIDGroupsRecomputingParityActiveQueryParameter adds the raidGroupsRecomputingParityActive to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsRecomputingParityActiveQueryParameter(raidGroupsRecomputingParityActive *bool) *PlexCollectionGetParams {
	o.SetRaIDGroupsRecomputingParityActiveQueryParameter(raidGroupsRecomputingParityActive)
	return o
}

// SetRaIDGroupsRecomputingParityActiveQueryParameter adds the raidGroupsRecomputingParityActive to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsRecomputingParityActiveQueryParameter(raidGroupsRecomputingParityActive *bool) {
	o.RaIDGroupsRecomputingParityActiveQueryParameter = raidGroupsRecomputingParityActive
}

// WithRaIDGroupsRecomputingParityPercentQueryParameter adds the raidGroupsRecomputingParityPercent to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsRecomputingParityPercentQueryParameter(raidGroupsRecomputingParityPercent *int64) *PlexCollectionGetParams {
	o.SetRaIDGroupsRecomputingParityPercentQueryParameter(raidGroupsRecomputingParityPercent)
	return o
}

// SetRaIDGroupsRecomputingParityPercentQueryParameter adds the raidGroupsRecomputingParityPercent to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsRecomputingParityPercentQueryParameter(raidGroupsRecomputingParityPercent *int64) {
	o.RaIDGroupsRecomputingParityPercentQueryParameter = raidGroupsRecomputingParityPercent
}

// WithRaIDGroupsReconstructActiveQueryParameter adds the raidGroupsReconstructActive to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsReconstructActiveQueryParameter(raidGroupsReconstructActive *bool) *PlexCollectionGetParams {
	o.SetRaIDGroupsReconstructActiveQueryParameter(raidGroupsReconstructActive)
	return o
}

// SetRaIDGroupsReconstructActiveQueryParameter adds the raidGroupsReconstructActive to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsReconstructActiveQueryParameter(raidGroupsReconstructActive *bool) {
	o.RaIDGroupsReconstructActiveQueryParameter = raidGroupsReconstructActive
}

// WithRaIDGroupsReconstructPercentQueryParameter adds the raidGroupsReconstructPercent to the plex collection get params
func (o *PlexCollectionGetParams) WithRaIDGroupsReconstructPercentQueryParameter(raidGroupsReconstructPercent *int64) *PlexCollectionGetParams {
	o.SetRaIDGroupsReconstructPercentQueryParameter(raidGroupsReconstructPercent)
	return o
}

// SetRaIDGroupsReconstructPercentQueryParameter adds the raidGroupsReconstructPercent to the plex collection get params
func (o *PlexCollectionGetParams) SetRaIDGroupsReconstructPercentQueryParameter(raidGroupsReconstructPercent *int64) {
	o.RaIDGroupsReconstructPercentQueryParameter = raidGroupsReconstructPercent
}

// WithResyncActiveQueryParameter adds the resyncActive to the plex collection get params
func (o *PlexCollectionGetParams) WithResyncActiveQueryParameter(resyncActive *bool) *PlexCollectionGetParams {
	o.SetResyncActiveQueryParameter(resyncActive)
	return o
}

// SetResyncActiveQueryParameter adds the resyncActive to the plex collection get params
func (o *PlexCollectionGetParams) SetResyncActiveQueryParameter(resyncActive *bool) {
	o.ResyncActiveQueryParameter = resyncActive
}

// WithResyncLevelQueryParameter adds the resyncLevel to the plex collection get params
func (o *PlexCollectionGetParams) WithResyncLevelQueryParameter(resyncLevel *string) *PlexCollectionGetParams {
	o.SetResyncLevelQueryParameter(resyncLevel)
	return o
}

// SetResyncLevelQueryParameter adds the resyncLevel to the plex collection get params
func (o *PlexCollectionGetParams) SetResyncLevelQueryParameter(resyncLevel *string) {
	o.ResyncLevelQueryParameter = resyncLevel
}

// WithResyncPercentQueryParameter adds the resyncPercent to the plex collection get params
func (o *PlexCollectionGetParams) WithResyncPercentQueryParameter(resyncPercent *int64) *PlexCollectionGetParams {
	o.SetResyncPercentQueryParameter(resyncPercent)
	return o
}

// SetResyncPercentQueryParameter adds the resyncPercent to the plex collection get params
func (o *PlexCollectionGetParams) SetResyncPercentQueryParameter(resyncPercent *int64) {
	o.ResyncPercentQueryParameter = resyncPercent
}

// WithReturnRecords adds the returnRecords to the plex collection get params
func (o *PlexCollectionGetParams) WithReturnRecords(returnRecords *bool) *PlexCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the plex collection get params
func (o *PlexCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the plex collection get params
func (o *PlexCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *PlexCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the plex collection get params
func (o *PlexCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithStateQueryParameter adds the state to the plex collection get params
func (o *PlexCollectionGetParams) WithStateQueryParameter(state *string) *PlexCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the plex collection get params
func (o *PlexCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WriteToRequest writes these params to a swagger request
func (o *PlexCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AggregateNameQueryParameter != nil {

		// query param aggregate.name
		var qrAggregateName string

		if o.AggregateNameQueryParameter != nil {
			qrAggregateName = *o.AggregateNameQueryParameter
		}
		qAggregateName := qrAggregateName
		if qAggregateName != "" {

			if err := r.SetQueryParam("aggregate.name", qAggregateName); err != nil {
				return err
			}
		}
	}

	if o.AggregateUUIDQueryParameter != nil {

		// query param aggregate.uuid
		var qrAggregateUUID string

		if o.AggregateUUIDQueryParameter != nil {
			qrAggregateUUID = *o.AggregateUUIDQueryParameter
		}
		qAggregateUUID := qrAggregateUUID
		if qAggregateUUID != "" {

			if err := r.SetQueryParam("aggregate.uuid", qAggregateUUID); err != nil {
				return err
			}
		}
	}

	// path param aggregate.uuid
	if err := r.SetPathParam("aggregate.uuid", o.AggregateUUIDPathParameter); err != nil {
		return err
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

	if o.OnlineQueryParameter != nil {

		// query param online
		var qrOnline bool

		if o.OnlineQueryParameter != nil {
			qrOnline = *o.OnlineQueryParameter
		}
		qOnline := swag.FormatBool(qrOnline)
		if qOnline != "" {

			if err := r.SetQueryParam("online", qOnline); err != nil {
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

	if o.PoolQueryParameter != nil {

		// query param pool
		var qrPool string

		if o.PoolQueryParameter != nil {
			qrPool = *o.PoolQueryParameter
		}
		qPool := qrPool
		if qPool != "" {

			if err := r.SetQueryParam("pool", qPool); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsCacheTierQueryParameter != nil {

		// query param raid_groups.cache_tier
		var qrRaidGroupsCacheTier bool

		if o.RaIDGroupsCacheTierQueryParameter != nil {
			qrRaidGroupsCacheTier = *o.RaIDGroupsCacheTierQueryParameter
		}
		qRaidGroupsCacheTier := swag.FormatBool(qrRaidGroupsCacheTier)
		if qRaidGroupsCacheTier != "" {

			if err := r.SetQueryParam("raid_groups.cache_tier", qRaidGroupsCacheTier); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsDegradedQueryParameter != nil {

		// query param raid_groups.degraded
		var qrRaidGroupsDegraded bool

		if o.RaIDGroupsDegradedQueryParameter != nil {
			qrRaidGroupsDegraded = *o.RaIDGroupsDegradedQueryParameter
		}
		qRaidGroupsDegraded := swag.FormatBool(qrRaidGroupsDegraded)
		if qRaidGroupsDegraded != "" {

			if err := r.SetQueryParam("raid_groups.degraded", qRaidGroupsDegraded); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsDisksDiskNameQueryParameter != nil {

		// query param raid_groups.disks.disk.name
		var qrRaidGroupsDisksDiskName string

		if o.RaIDGroupsDisksDiskNameQueryParameter != nil {
			qrRaidGroupsDisksDiskName = *o.RaIDGroupsDisksDiskNameQueryParameter
		}
		qRaidGroupsDisksDiskName := qrRaidGroupsDisksDiskName
		if qRaidGroupsDisksDiskName != "" {

			if err := r.SetQueryParam("raid_groups.disks.disk.name", qRaidGroupsDisksDiskName); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsDisksPositionQueryParameter != nil {

		// query param raid_groups.disks.position
		var qrRaidGroupsDisksPosition string

		if o.RaIDGroupsDisksPositionQueryParameter != nil {
			qrRaidGroupsDisksPosition = *o.RaIDGroupsDisksPositionQueryParameter
		}
		qRaidGroupsDisksPosition := qrRaidGroupsDisksPosition
		if qRaidGroupsDisksPosition != "" {

			if err := r.SetQueryParam("raid_groups.disks.position", qRaidGroupsDisksPosition); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsDisksStateQueryParameter != nil {

		// query param raid_groups.disks.state
		var qrRaidGroupsDisksState string

		if o.RaIDGroupsDisksStateQueryParameter != nil {
			qrRaidGroupsDisksState = *o.RaIDGroupsDisksStateQueryParameter
		}
		qRaidGroupsDisksState := qrRaidGroupsDisksState
		if qRaidGroupsDisksState != "" {

			if err := r.SetQueryParam("raid_groups.disks.state", qRaidGroupsDisksState); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsDisksTypeQueryParameter != nil {

		// query param raid_groups.disks.type
		var qrRaidGroupsDisksType string

		if o.RaIDGroupsDisksTypeQueryParameter != nil {
			qrRaidGroupsDisksType = *o.RaIDGroupsDisksTypeQueryParameter
		}
		qRaidGroupsDisksType := qrRaidGroupsDisksType
		if qRaidGroupsDisksType != "" {

			if err := r.SetQueryParam("raid_groups.disks.type", qRaidGroupsDisksType); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsDisksUsableSizeQueryParameter != nil {

		// query param raid_groups.disks.usable_size
		var qrRaidGroupsDisksUsableSize int64

		if o.RaIDGroupsDisksUsableSizeQueryParameter != nil {
			qrRaidGroupsDisksUsableSize = *o.RaIDGroupsDisksUsableSizeQueryParameter
		}
		qRaidGroupsDisksUsableSize := swag.FormatInt64(qrRaidGroupsDisksUsableSize)
		if qRaidGroupsDisksUsableSize != "" {

			if err := r.SetQueryParam("raid_groups.disks.usable_size", qRaidGroupsDisksUsableSize); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsNameQueryParameter != nil {

		// query param raid_groups.name
		var qrRaidGroupsName string

		if o.RaIDGroupsNameQueryParameter != nil {
			qrRaidGroupsName = *o.RaIDGroupsNameQueryParameter
		}
		qRaidGroupsName := qrRaidGroupsName
		if qRaidGroupsName != "" {

			if err := r.SetQueryParam("raid_groups.name", qRaidGroupsName); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsRecomputingParityActiveQueryParameter != nil {

		// query param raid_groups.recomputing_parity.active
		var qrRaidGroupsRecomputingParityActive bool

		if o.RaIDGroupsRecomputingParityActiveQueryParameter != nil {
			qrRaidGroupsRecomputingParityActive = *o.RaIDGroupsRecomputingParityActiveQueryParameter
		}
		qRaidGroupsRecomputingParityActive := swag.FormatBool(qrRaidGroupsRecomputingParityActive)
		if qRaidGroupsRecomputingParityActive != "" {

			if err := r.SetQueryParam("raid_groups.recomputing_parity.active", qRaidGroupsRecomputingParityActive); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsRecomputingParityPercentQueryParameter != nil {

		// query param raid_groups.recomputing_parity.percent
		var qrRaidGroupsRecomputingParityPercent int64

		if o.RaIDGroupsRecomputingParityPercentQueryParameter != nil {
			qrRaidGroupsRecomputingParityPercent = *o.RaIDGroupsRecomputingParityPercentQueryParameter
		}
		qRaidGroupsRecomputingParityPercent := swag.FormatInt64(qrRaidGroupsRecomputingParityPercent)
		if qRaidGroupsRecomputingParityPercent != "" {

			if err := r.SetQueryParam("raid_groups.recomputing_parity.percent", qRaidGroupsRecomputingParityPercent); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsReconstructActiveQueryParameter != nil {

		// query param raid_groups.reconstruct.active
		var qrRaidGroupsReconstructActive bool

		if o.RaIDGroupsReconstructActiveQueryParameter != nil {
			qrRaidGroupsReconstructActive = *o.RaIDGroupsReconstructActiveQueryParameter
		}
		qRaidGroupsReconstructActive := swag.FormatBool(qrRaidGroupsReconstructActive)
		if qRaidGroupsReconstructActive != "" {

			if err := r.SetQueryParam("raid_groups.reconstruct.active", qRaidGroupsReconstructActive); err != nil {
				return err
			}
		}
	}

	if o.RaIDGroupsReconstructPercentQueryParameter != nil {

		// query param raid_groups.reconstruct.percent
		var qrRaidGroupsReconstructPercent int64

		if o.RaIDGroupsReconstructPercentQueryParameter != nil {
			qrRaidGroupsReconstructPercent = *o.RaIDGroupsReconstructPercentQueryParameter
		}
		qRaidGroupsReconstructPercent := swag.FormatInt64(qrRaidGroupsReconstructPercent)
		if qRaidGroupsReconstructPercent != "" {

			if err := r.SetQueryParam("raid_groups.reconstruct.percent", qRaidGroupsReconstructPercent); err != nil {
				return err
			}
		}
	}

	if o.ResyncActiveQueryParameter != nil {

		// query param resync.active
		var qrResyncActive bool

		if o.ResyncActiveQueryParameter != nil {
			qrResyncActive = *o.ResyncActiveQueryParameter
		}
		qResyncActive := swag.FormatBool(qrResyncActive)
		if qResyncActive != "" {

			if err := r.SetQueryParam("resync.active", qResyncActive); err != nil {
				return err
			}
		}
	}

	if o.ResyncLevelQueryParameter != nil {

		// query param resync.level
		var qrResyncLevel string

		if o.ResyncLevelQueryParameter != nil {
			qrResyncLevel = *o.ResyncLevelQueryParameter
		}
		qResyncLevel := qrResyncLevel
		if qResyncLevel != "" {

			if err := r.SetQueryParam("resync.level", qResyncLevel); err != nil {
				return err
			}
		}
	}

	if o.ResyncPercentQueryParameter != nil {

		// query param resync.percent
		var qrResyncPercent int64

		if o.ResyncPercentQueryParameter != nil {
			qrResyncPercent = *o.ResyncPercentQueryParameter
		}
		qResyncPercent := swag.FormatInt64(qrResyncPercent)
		if qResyncPercent != "" {

			if err := r.SetQueryParam("resync.percent", qResyncPercent); err != nil {
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

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPlexCollectionGet binds the parameter fields
func (o *PlexCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamPlexCollectionGet binds the parameter order_by
func (o *PlexCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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