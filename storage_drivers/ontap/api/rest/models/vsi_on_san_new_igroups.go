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

// VsiOnSanNewIgroups The list of initiator groups to create.
//
// swagger:model vsi_on_san_new_igroups
type VsiOnSanNewIgroups struct {

	// initiators
	Initiators []string `json:"initiators,omitempty"`

	// The name of the new initiator group.
	// Required: true
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name"`

	// The protocol of the new initiator group.
	// Enum: [fcp iscsi mixed]
	Protocol *string `json:"protocol,omitempty"`
}

// Validate validates this vsi on san new igroups
func (m *VsiOnSanNewIgroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnSanNewIgroups) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

var vsiOnSanNewIgroupsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fcp","iscsi","mixed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vsiOnSanNewIgroupsTypeProtocolPropEnum = append(vsiOnSanNewIgroupsTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// vsi_on_san_new_igroups
	// VsiOnSanNewIgroups
	// protocol
	// Protocol
	// fcp
	// END RIPPY DEBUGGING
	// VsiOnSanNewIgroupsProtocolFcp captures enum value "fcp"
	VsiOnSanNewIgroupsProtocolFcp string = "fcp"

	// BEGIN RIPPY DEBUGGING
	// vsi_on_san_new_igroups
	// VsiOnSanNewIgroups
	// protocol
	// Protocol
	// iscsi
	// END RIPPY DEBUGGING
	// VsiOnSanNewIgroupsProtocolIscsi captures enum value "iscsi"
	VsiOnSanNewIgroupsProtocolIscsi string = "iscsi"

	// BEGIN RIPPY DEBUGGING
	// vsi_on_san_new_igroups
	// VsiOnSanNewIgroups
	// protocol
	// Protocol
	// mixed
	// END RIPPY DEBUGGING
	// VsiOnSanNewIgroupsProtocolMixed captures enum value "mixed"
	VsiOnSanNewIgroupsProtocolMixed string = "mixed"
)

// prop value enum
func (m *VsiOnSanNewIgroups) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vsiOnSanNewIgroupsTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VsiOnSanNewIgroups) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vsi on san new igroups based on context it is used
func (m *VsiOnSanNewIgroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnSanNewIgroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnSanNewIgroups) UnmarshalBinary(b []byte) error {
	var res VsiOnSanNewIgroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
