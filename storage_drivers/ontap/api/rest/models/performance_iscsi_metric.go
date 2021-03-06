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

// PerformanceIscsiMetric Performance numbers, such as IOPS latency and throughput.
//
// swagger:model performance_iscsi_metric
type PerformanceIscsiMetric struct {

	// links
	Links *PerformanceIscsiMetricLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceIscsiMetricIops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceIscsiMetricLatency `json:"latency,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: [ok error partial_no_data partial_no_uuid partial_no_response partial_other_error negative_delta backfilled_data inconsistent_delta_time inconsistent_old_data]
	Status string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceIscsiMetricThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance iscsi metric
func (m *PerformanceIscsiMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceIscsiMetric) validateLinks(formats strfmt.Registry) error {
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

var performanceIscsiMetricTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceIscsiMetricTypeDurationPropEnum = append(performanceIscsiMetricTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// duration
	// Duration
	// PT15S
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricDurationPT15S captures enum value "PT15S"
	PerformanceIscsiMetricDurationPT15S string = "PT15S"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// duration
	// Duration
	// PT4M
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricDurationPT4M captures enum value "PT4M"
	PerformanceIscsiMetricDurationPT4M string = "PT4M"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// duration
	// Duration
	// PT30M
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricDurationPT30M captures enum value "PT30M"
	PerformanceIscsiMetricDurationPT30M string = "PT30M"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// duration
	// Duration
	// PT2H
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricDurationPT2H captures enum value "PT2H"
	PerformanceIscsiMetricDurationPT2H string = "PT2H"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// duration
	// Duration
	// P1D
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricDurationP1D captures enum value "P1D"
	PerformanceIscsiMetricDurationP1D string = "P1D"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// duration
	// Duration
	// PT5M
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricDurationPT5M captures enum value "PT5M"
	PerformanceIscsiMetricDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceIscsiMetric) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceIscsiMetricTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceIscsiMetric) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceIscsiMetric) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceIscsiMetric) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

var performanceIscsiMetricTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceIscsiMetricTypeStatusPropEnum = append(performanceIscsiMetricTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// ok
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusOk captures enum value "ok"
	PerformanceIscsiMetricStatusOk string = "ok"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// error
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusError captures enum value "error"
	PerformanceIscsiMetricStatusError string = "error"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// partial_no_data
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusPartialNoData captures enum value "partial_no_data"
	PerformanceIscsiMetricStatusPartialNoData string = "partial_no_data"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// partial_no_uuid
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceIscsiMetricStatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// partial_no_response
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceIscsiMetricStatusPartialNoResponse string = "partial_no_response"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// partial_other_error
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceIscsiMetricStatusPartialOtherError string = "partial_other_error"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// negative_delta
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusNegativeDelta captures enum value "negative_delta"
	PerformanceIscsiMetricStatusNegativeDelta string = "negative_delta"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// backfilled_data
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusBackfilledData captures enum value "backfilled_data"
	PerformanceIscsiMetricStatusBackfilledData string = "backfilled_data"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// inconsistent_delta_time
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceIscsiMetricStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN RIPPY DEBUGGING
	// performance_iscsi_metric
	// PerformanceIscsiMetric
	// status
	// Status
	// inconsistent_old_data
	// END RIPPY DEBUGGING
	// PerformanceIscsiMetricStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceIscsiMetricStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *PerformanceIscsiMetric) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceIscsiMetricTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceIscsiMetric) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceIscsiMetric) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceIscsiMetric) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance iscsi metric based on the context it is used
func (m *PerformanceIscsiMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceIscsiMetric) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceIscsiMetric) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", string(m.Duration)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceIscsiMetric) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {
		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceIscsiMetric) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {
		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceIscsiMetric) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceIscsiMetric) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceIscsiMetric) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceIscsiMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceIscsiMetric) UnmarshalBinary(b []byte) error {
	var res PerformanceIscsiMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceIscsiMetricIops The rate of I/O operations observed at the storage object.
//
// swagger:model PerformanceIscsiMetricIops
type PerformanceIscsiMetricIops struct {

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

// Validate validates this performance iscsi metric iops
func (m *PerformanceIscsiMetricIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance iscsi metric iops based on the context it is used
func (m *PerformanceIscsiMetricIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceIscsiMetricIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceIscsiMetricIops) UnmarshalBinary(b []byte) error {
	var res PerformanceIscsiMetricIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceIscsiMetricLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model PerformanceIscsiMetricLatency
type PerformanceIscsiMetricLatency struct {

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

// Validate validates this performance iscsi metric latency
func (m *PerformanceIscsiMetricLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance iscsi metric latency based on the context it is used
func (m *PerformanceIscsiMetricLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceIscsiMetricLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceIscsiMetricLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceIscsiMetricLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceIscsiMetricLinks performance iscsi metric links
//
// swagger:model PerformanceIscsiMetricLinks
type PerformanceIscsiMetricLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance iscsi metric links
func (m *PerformanceIscsiMetricLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceIscsiMetricLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance iscsi metric links based on the context it is used
func (m *PerformanceIscsiMetricLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceIscsiMetricLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceIscsiMetricLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceIscsiMetricLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceIscsiMetricLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceIscsiMetricThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model PerformanceIscsiMetricThroughput
type PerformanceIscsiMetricThroughput struct {

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

// Validate validates this performance iscsi metric throughput
func (m *PerformanceIscsiMetricThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance iscsi metric throughput based on the context it is used
func (m *PerformanceIscsiMetricThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceIscsiMetricThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceIscsiMetricThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceIscsiMetricThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
