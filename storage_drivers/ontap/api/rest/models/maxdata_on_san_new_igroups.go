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

// MaxdataOnSanNewIgroups The list of initiator groups to create.
//
// swagger:model maxdata_on_san_new_igroups
type MaxdataOnSanNewIgroups struct {

	// initiators
	Initiators []string `json:"initiators,omitempty"`

	// The name of the new initiator group.
	// Required: true
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name"`

	// The name of the host OS accessing the application. The default value is the host OS that is running the application.
	// Enum: [aix hpux hyper_v linux netware openvms solaris vmware windows xen]
	OsType string `json:"os_type,omitempty"`

	// The protocol of the new initiator group.
	// Enum: [fcp iscsi]
	Protocol *string `json:"protocol,omitempty"`
}

// Validate validates this maxdata on san new igroups
func (m *MaxdataOnSanNewIgroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
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

func (m *MaxdataOnSanNewIgroups) validateName(formats strfmt.Registry) error {

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

var maxdataOnSanNewIgroupsTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aix","hpux","hyper_v","linux","netware","openvms","solaris","vmware","windows","xen"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		maxdataOnSanNewIgroupsTypeOsTypePropEnum = append(maxdataOnSanNewIgroupsTypeOsTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// aix
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeAix captures enum value "aix"
	MaxdataOnSanNewIgroupsOsTypeAix string = "aix"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// hpux
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeHpux captures enum value "hpux"
	MaxdataOnSanNewIgroupsOsTypeHpux string = "hpux"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// hyper_v
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeHyperv captures enum value "hyper_v"
	MaxdataOnSanNewIgroupsOsTypeHyperv string = "hyper_v"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// linux
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeLinux captures enum value "linux"
	MaxdataOnSanNewIgroupsOsTypeLinux string = "linux"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// netware
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeNetware captures enum value "netware"
	MaxdataOnSanNewIgroupsOsTypeNetware string = "netware"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// openvms
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeOpenvms captures enum value "openvms"
	MaxdataOnSanNewIgroupsOsTypeOpenvms string = "openvms"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// solaris
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeSolaris captures enum value "solaris"
	MaxdataOnSanNewIgroupsOsTypeSolaris string = "solaris"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// vmware
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeVmware captures enum value "vmware"
	MaxdataOnSanNewIgroupsOsTypeVmware string = "vmware"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// windows
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeWindows captures enum value "windows"
	MaxdataOnSanNewIgroupsOsTypeWindows string = "windows"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// os_type
	// OsType
	// xen
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsOsTypeXen captures enum value "xen"
	MaxdataOnSanNewIgroupsOsTypeXen string = "xen"
)

// prop value enum
func (m *MaxdataOnSanNewIgroups) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, maxdataOnSanNewIgroupsTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MaxdataOnSanNewIgroups) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

var maxdataOnSanNewIgroupsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fcp","iscsi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		maxdataOnSanNewIgroupsTypeProtocolPropEnum = append(maxdataOnSanNewIgroupsTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// protocol
	// Protocol
	// fcp
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsProtocolFcp captures enum value "fcp"
	MaxdataOnSanNewIgroupsProtocolFcp string = "fcp"

	// BEGIN RIPPY DEBUGGING
	// maxdata_on_san_new_igroups
	// MaxdataOnSanNewIgroups
	// protocol
	// Protocol
	// iscsi
	// END RIPPY DEBUGGING
	// MaxdataOnSanNewIgroupsProtocolIscsi captures enum value "iscsi"
	MaxdataOnSanNewIgroupsProtocolIscsi string = "iscsi"
)

// prop value enum
func (m *MaxdataOnSanNewIgroups) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, maxdataOnSanNewIgroupsTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MaxdataOnSanNewIgroups) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this maxdata on san new igroups based on context it is used
func (m *MaxdataOnSanNewIgroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MaxdataOnSanNewIgroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaxdataOnSanNewIgroups) UnmarshalBinary(b []byte) error {
	var res MaxdataOnSanNewIgroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
