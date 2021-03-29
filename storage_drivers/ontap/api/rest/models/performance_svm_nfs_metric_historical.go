// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PerformanceSvmNfsMetricHistorical Historical performance numbers, such as IOPS latency and throughput, for SVM-NFS protocol.
//
// swagger:model performance_svm_nfs_metric_historical
type PerformanceSvmNfsMetricHistorical struct {

	// links
	Links *PerformanceSvmNfsMetricHistoricalLinks `json:"_links,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// v3
	V3 *PerformanceSvmNfsMetricHistoricalV3 `json:"v3,omitempty"`
}

// Validate validates this performance svm nfs metric historical
func (m *PerformanceSvmNfsMetricHistorical) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateV3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetricHistorical) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistorical) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistorical) validateV3(formats strfmt.Registry) error {
	if swag.IsZero(m.V3) { // not required
		return nil
	}

	if m.V3 != nil {
		if err := m.V3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance svm nfs metric historical based on the context it is used
func (m *PerformanceSvmNfsMetricHistorical) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateV3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetricHistorical) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistorical) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistorical) contextValidateV3(ctx context.Context, formats strfmt.Registry) error {

	if m.V3 != nil {
		if err := m.V3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistorical) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistorical) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricHistorical
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricHistoricalLinks performance svm nfs metric historical links
//
// swagger:model PerformanceSvmNfsMetricHistoricalLinks
type PerformanceSvmNfsMetricHistoricalLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance svm nfs metric historical links
func (m *PerformanceSvmNfsMetricHistoricalLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance svm nfs metric historical links based on the context it is used
func (m *PerformanceSvmNfsMetricHistoricalLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricHistoricalLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricHistoricalV3 The NFSv3 operations
//
// swagger:model PerformanceSvmNfsMetricHistoricalV3
type PerformanceSvmNfsMetricHistoricalV3 struct {

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceSvmNfsMetricHistoricalV3Iops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceSvmNfsMetricHistoricalV3Latency `json:"latency,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: [ok error partial_no_data partial_no_uuid partial_no_response partial_other_error negative_delta backfilled_data inconsistent_delta_time inconsistent_old_data]
	Status string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceSvmNfsMetricHistoricalV3Throughput `json:"throughput,omitempty"`
}

// Validate validates this performance svm nfs metric historical v3
func (m *PerformanceSvmNfsMetricHistoricalV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var performanceSvmNfsMetricHistoricalV3TypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceSvmNfsMetricHistoricalV3TypeDurationPropEnum = append(performanceSvmNfsMetricHistoricalV3TypeDurationPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// duration
	// Duration
	// PT15S
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3DurationPT15S captures enum value "PT15S"
	PerformanceSvmNfsMetricHistoricalV3DurationPT15S string = "PT15S"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// duration
	// Duration
	// PT4M
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3DurationPT4M captures enum value "PT4M"
	PerformanceSvmNfsMetricHistoricalV3DurationPT4M string = "PT4M"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// duration
	// Duration
	// PT30M
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3DurationPT30M captures enum value "PT30M"
	PerformanceSvmNfsMetricHistoricalV3DurationPT30M string = "PT30M"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// duration
	// Duration
	// PT2H
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3DurationPT2H captures enum value "PT2H"
	PerformanceSvmNfsMetricHistoricalV3DurationPT2H string = "PT2H"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// duration
	// Duration
	// P1D
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3DurationP1D captures enum value "P1D"
	PerformanceSvmNfsMetricHistoricalV3DurationP1D string = "P1D"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// duration
	// Duration
	// PT5M
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3DurationPT5M captures enum value "PT5M"
	PerformanceSvmNfsMetricHistoricalV3DurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceSvmNfsMetricHistoricalV3) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceSvmNfsMetricHistoricalV3TypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("v3"+"."+"duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "latency")
			}
			return err
		}
	}

	return nil
}

var performanceSvmNfsMetricHistoricalV3TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceSvmNfsMetricHistoricalV3TypeStatusPropEnum = append(performanceSvmNfsMetricHistoricalV3TypeStatusPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// ok
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusOk captures enum value "ok"
	PerformanceSvmNfsMetricHistoricalV3StatusOk string = "ok"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// error
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusError captures enum value "error"
	PerformanceSvmNfsMetricHistoricalV3StatusError string = "error"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// partial_no_data
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusPartialNoData captures enum value "partial_no_data"
	PerformanceSvmNfsMetricHistoricalV3StatusPartialNoData string = "partial_no_data"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// partial_no_uuid
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceSvmNfsMetricHistoricalV3StatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// partial_no_response
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceSvmNfsMetricHistoricalV3StatusPartialNoResponse string = "partial_no_response"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// partial_other_error
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusPartialOtherError captures enum value "partial_other_error"
	PerformanceSvmNfsMetricHistoricalV3StatusPartialOtherError string = "partial_other_error"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// negative_delta
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusNegativeDelta captures enum value "negative_delta"
	PerformanceSvmNfsMetricHistoricalV3StatusNegativeDelta string = "negative_delta"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// backfilled_data
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusBackfilledData captures enum value "backfilled_data"
	PerformanceSvmNfsMetricHistoricalV3StatusBackfilledData string = "backfilled_data"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// inconsistent_delta_time
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceSvmNfsMetricHistoricalV3StatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3
	// PerformanceSvmNfsMetricHistoricalV3
	// status
	// Status
	// inconsistent_old_data
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricHistoricalV3StatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceSvmNfsMetricHistoricalV3StatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *PerformanceSvmNfsMetricHistoricalV3) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceSvmNfsMetricHistoricalV3TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("v3"+"."+"status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "throughput")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance svm nfs metric historical v3 based on the context it is used
func (m *PerformanceSvmNfsMetricHistoricalV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "v3"+"."+"duration", "body", string(m.Duration)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {
		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {
		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "v3"+"."+"status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricHistoricalV3) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "throughput")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalV3) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricHistoricalV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricHistoricalV3Iops The rate of I/O operations observed at the storage object.
//
// swagger:model PerformanceSvmNfsMetricHistoricalV3Iops
type PerformanceSvmNfsMetricHistoricalV3Iops struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write int64 `json:"write,omitempty"`
}

// Validate validates this performance svm nfs metric historical v3 iops
func (m *PerformanceSvmNfsMetricHistoricalV3Iops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm nfs metric historical v3 iops based on the context it is used
func (m *PerformanceSvmNfsMetricHistoricalV3Iops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalV3Iops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalV3Iops) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricHistoricalV3Iops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricHistoricalV3Latency The round trip latency in microseconds observed at the storage object.
//
// swagger:model PerformanceSvmNfsMetricHistoricalV3Latency
type PerformanceSvmNfsMetricHistoricalV3Latency struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write int64 `json:"write,omitempty"`
}

// Validate validates this performance svm nfs metric historical v3 latency
func (m *PerformanceSvmNfsMetricHistoricalV3Latency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm nfs metric historical v3 latency based on the context it is used
func (m *PerformanceSvmNfsMetricHistoricalV3Latency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalV3Latency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalV3Latency) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricHistoricalV3Latency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricHistoricalV3Throughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model PerformanceSvmNfsMetricHistoricalV3Throughput
type PerformanceSvmNfsMetricHistoricalV3Throughput struct {

	// Performance metric for read I/O operations.
	// Example: 200
	Read int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write int64 `json:"write,omitempty"`
}

// Validate validates this performance svm nfs metric historical v3 throughput
func (m *PerformanceSvmNfsMetricHistoricalV3Throughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm nfs metric historical v3 throughput based on the context it is used
func (m *PerformanceSvmNfsMetricHistoricalV3Throughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalV3Throughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricHistoricalV3Throughput) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricHistoricalV3Throughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY