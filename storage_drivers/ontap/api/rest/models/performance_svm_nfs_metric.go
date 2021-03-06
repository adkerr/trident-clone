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

// PerformanceSvmNfsMetric Historical performance numbers, such as IOPS latency and throughput, for SVM-NFS protocol.
//
// swagger:model performance_svm_nfs_metric
type PerformanceSvmNfsMetric struct {

	// v3
	V3 *PerformanceSvmNfsMetricV3 `json:"v3,omitempty"`
}

// Validate validates this performance svm nfs metric
func (m *PerformanceSvmNfsMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateV3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetric) validateV3(formats strfmt.Registry) error {
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

// ContextValidate validate this performance svm nfs metric based on the context it is used
func (m *PerformanceSvmNfsMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateV3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetric) contextValidateV3(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceSvmNfsMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetric) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricV3 The NFSv3 operations
//
// swagger:model PerformanceSvmNfsMetricV3
type PerformanceSvmNfsMetricV3 struct {

	// links
	Links *PerformanceSvmNfsMetricV3Links `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceSvmNfsMetricV3Iops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceSvmNfsMetricV3Latency `json:"latency,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: [ok error partial_no_data partial_no_uuid partial_no_response partial_other_error negative_delta backfilled_data inconsistent_delta_time inconsistent_old_data]
	Status string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceSvmNfsMetricV3Throughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance svm nfs metric v3
func (m *PerformanceSvmNfsMetricV3) Validate(formats strfmt.Registry) error {
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

func (m *PerformanceSvmNfsMetricV3) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var performanceSvmNfsMetricV3TypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceSvmNfsMetricV3TypeDurationPropEnum = append(performanceSvmNfsMetricV3TypeDurationPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// duration
	// Duration
	// PT15S
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3DurationPT15S captures enum value "PT15S"
	PerformanceSvmNfsMetricV3DurationPT15S string = "PT15S"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// duration
	// Duration
	// PT4M
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3DurationPT4M captures enum value "PT4M"
	PerformanceSvmNfsMetricV3DurationPT4M string = "PT4M"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// duration
	// Duration
	// PT30M
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3DurationPT30M captures enum value "PT30M"
	PerformanceSvmNfsMetricV3DurationPT30M string = "PT30M"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// duration
	// Duration
	// PT2H
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3DurationPT2H captures enum value "PT2H"
	PerformanceSvmNfsMetricV3DurationPT2H string = "PT2H"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// duration
	// Duration
	// P1D
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3DurationP1D captures enum value "P1D"
	PerformanceSvmNfsMetricV3DurationP1D string = "P1D"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// duration
	// Duration
	// PT5M
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3DurationPT5M captures enum value "PT5M"
	PerformanceSvmNfsMetricV3DurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceSvmNfsMetricV3) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceSvmNfsMetricV3TypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceSvmNfsMetricV3) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("v3"+"."+"duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricV3) validateIops(formats strfmt.Registry) error {
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

func (m *PerformanceSvmNfsMetricV3) validateLatency(formats strfmt.Registry) error {
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

var performanceSvmNfsMetricV3TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceSvmNfsMetricV3TypeStatusPropEnum = append(performanceSvmNfsMetricV3TypeStatusPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// ok
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusOk captures enum value "ok"
	PerformanceSvmNfsMetricV3StatusOk string = "ok"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// error
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusError captures enum value "error"
	PerformanceSvmNfsMetricV3StatusError string = "error"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// partial_no_data
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusPartialNoData captures enum value "partial_no_data"
	PerformanceSvmNfsMetricV3StatusPartialNoData string = "partial_no_data"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// partial_no_uuid
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceSvmNfsMetricV3StatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// partial_no_response
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceSvmNfsMetricV3StatusPartialNoResponse string = "partial_no_response"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// partial_other_error
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusPartialOtherError captures enum value "partial_other_error"
	PerformanceSvmNfsMetricV3StatusPartialOtherError string = "partial_other_error"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// negative_delta
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusNegativeDelta captures enum value "negative_delta"
	PerformanceSvmNfsMetricV3StatusNegativeDelta string = "negative_delta"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// backfilled_data
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusBackfilledData captures enum value "backfilled_data"
	PerformanceSvmNfsMetricV3StatusBackfilledData string = "backfilled_data"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// inconsistent_delta_time
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceSvmNfsMetricV3StatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3
	// PerformanceSvmNfsMetricV3
	// status
	// Status
	// inconsistent_old_data
	// END RIPPY DEBUGGING
	// PerformanceSvmNfsMetricV3StatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceSvmNfsMetricV3StatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *PerformanceSvmNfsMetricV3) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceSvmNfsMetricV3TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceSvmNfsMetricV3) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("v3"+"."+"status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricV3) validateThroughput(formats strfmt.Registry) error {
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

func (m *PerformanceSvmNfsMetricV3) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("v3"+"."+"timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance svm nfs metric v3 based on the context it is used
func (m *PerformanceSvmNfsMetricV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PerformanceSvmNfsMetricV3) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmNfsMetricV3) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "v3"+"."+"duration", "body", string(m.Duration)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricV3) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceSvmNfsMetricV3) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceSvmNfsMetricV3) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "v3"+"."+"status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmNfsMetricV3) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceSvmNfsMetricV3) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "v3"+"."+"timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricV3Iops The rate of I/O operations observed at the storage object.
//
// swagger:model PerformanceSvmNfsMetricV3Iops
type PerformanceSvmNfsMetricV3Iops struct {

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

// Validate validates this performance svm nfs metric v3 iops
func (m *PerformanceSvmNfsMetricV3Iops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm nfs metric v3 iops based on the context it is used
func (m *PerformanceSvmNfsMetricV3Iops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3Iops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3Iops) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricV3Iops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricV3Latency The round trip latency in microseconds observed at the storage object.
//
// swagger:model PerformanceSvmNfsMetricV3Latency
type PerformanceSvmNfsMetricV3Latency struct {

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

// Validate validates this performance svm nfs metric v3 latency
func (m *PerformanceSvmNfsMetricV3Latency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm nfs metric v3 latency based on the context it is used
func (m *PerformanceSvmNfsMetricV3Latency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3Latency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3Latency) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricV3Latency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricV3Links performance svm nfs metric v3 links
//
// swagger:model PerformanceSvmNfsMetricV3Links
type PerformanceSvmNfsMetricV3Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance svm nfs metric v3 links
func (m *PerformanceSvmNfsMetricV3Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetricV3Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance svm nfs metric v3 links based on the context it is used
func (m *PerformanceSvmNfsMetricV3Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmNfsMetricV3Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3Links) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricV3Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmNfsMetricV3Throughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model PerformanceSvmNfsMetricV3Throughput
type PerformanceSvmNfsMetricV3Throughput struct {

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

// Validate validates this performance svm nfs metric v3 throughput
func (m *PerformanceSvmNfsMetricV3Throughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm nfs metric v3 throughput based on the context it is used
func (m *PerformanceSvmNfsMetricV3Throughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3Throughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmNfsMetricV3Throughput) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmNfsMetricV3Throughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
