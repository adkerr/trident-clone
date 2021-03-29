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

// LdapService ldap service
//
// swagger:model ldap_service
type LdapService struct {

	// links
	Links *LdapServiceLinks `json:"_links,omitempty"`

	// This parameter specifies the name of the Active Directory domain
	// used to discover LDAP servers for use by this client.
	// This is mutually exclusive with `servers` during POST and PATCH.
	//
	AdDomain string `json:"ad_domain,omitempty"`

	// Specifies the default base DN for all searches.
	BaseDn string `json:"base_dn,omitempty"`

	// Specifies the default search scope for LDAP queries:
	// * base - search the named entry only
	// * onelevel - search all entries immediately below the DN
	// * subtree - search the named DN entry and the entire subtree below the DN
	//
	// Enum: [base onelevel subtree]
	BaseScope *string `json:"base_scope,omitempty"`

	// Specifies the user that binds to the LDAP servers.
	BindDn string `json:"bind_dn,omitempty"`

	// Specifies the bind password for the LDAP servers.
	BindPassword string `json:"bind_password,omitempty"`

	// The minimum bind authentication level. Possible values are:
	// * anonymous - anonymous bind
	// * simple - simple bind
	// * sasl - Simple Authentication and Security Layer (SASL) bind
	//
	// Enum: [anonymous simple sasl]
	MinBindLevel *string `json:"min_bind_level,omitempty"`

	// The port used to connect to the LDAP Servers.
	// Example: 389
	// Maximum: 65535
	// Minimum: 1
	Port int64 `json:"port,omitempty"`

	// preferred ad servers
	PreferredAdServers []string `json:"preferred_ad_servers,omitempty"`

	// The name of the schema template used by the SVM.
	// * AD-IDMU - Active Directory Identity Management for UNIX
	// * AD-SFU - Active Directory Services for UNIX
	// * MS-AD-BIS - Active Directory Identity Management for UNIX
	// * RFC-2307 - Schema based on RFC 2307
	// * Custom schema
	//
	Schema *string `json:"schema,omitempty"`

	// servers
	Servers []string `json:"servers,omitempty"`

	// Specifies the level of security to be used for LDAP communications:
	// * none - no signing or sealing
	// * sign - sign LDAP traffic
	// * seal - seal and sign LDAP traffic
	//
	// Enum: [none sign seal]
	SessionSecurity *string `json:"session_security,omitempty"`

	// svm
	Svm *LdapServiceSvm `json:"svm,omitempty"`

	// Specifies whether or not to use Start TLS over LDAP connections.
	//
	UseStartTLS *bool `json:"use_start_tls,omitempty"`
}

// Validate validates this ldap service
func (m *LdapService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinBindLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionSecurity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapService) validateLinks(formats strfmt.Registry) error {
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

var ldapServiceTypeBaseScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["base","onelevel","subtree"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServiceTypeBaseScopePropEnum = append(ldapServiceTypeBaseScopePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// base_scope
	// BaseScope
	// base
	// END RIPPY DEBUGGING
	// LdapServiceBaseScopeBase captures enum value "base"
	LdapServiceBaseScopeBase string = "base"

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// base_scope
	// BaseScope
	// onelevel
	// END RIPPY DEBUGGING
	// LdapServiceBaseScopeOnelevel captures enum value "onelevel"
	LdapServiceBaseScopeOnelevel string = "onelevel"

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// base_scope
	// BaseScope
	// subtree
	// END RIPPY DEBUGGING
	// LdapServiceBaseScopeSubtree captures enum value "subtree"
	LdapServiceBaseScopeSubtree string = "subtree"
)

// prop value enum
func (m *LdapService) validateBaseScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServiceTypeBaseScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapService) validateBaseScope(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateBaseScopeEnum("base_scope", "body", *m.BaseScope); err != nil {
		return err
	}

	return nil
}

var ldapServiceTypeMinBindLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["anonymous","simple","sasl"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServiceTypeMinBindLevelPropEnum = append(ldapServiceTypeMinBindLevelPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// min_bind_level
	// MinBindLevel
	// anonymous
	// END RIPPY DEBUGGING
	// LdapServiceMinBindLevelAnonymous captures enum value "anonymous"
	LdapServiceMinBindLevelAnonymous string = "anonymous"

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// min_bind_level
	// MinBindLevel
	// simple
	// END RIPPY DEBUGGING
	// LdapServiceMinBindLevelSimple captures enum value "simple"
	LdapServiceMinBindLevelSimple string = "simple"

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// min_bind_level
	// MinBindLevel
	// sasl
	// END RIPPY DEBUGGING
	// LdapServiceMinBindLevelSasl captures enum value "sasl"
	LdapServiceMinBindLevelSasl string = "sasl"
)

// prop value enum
func (m *LdapService) validateMinBindLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServiceTypeMinBindLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapService) validateMinBindLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.MinBindLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateMinBindLevelEnum("min_bind_level", "body", *m.MinBindLevel); err != nil {
		return err
	}

	return nil
}

func (m *LdapService) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", m.Port, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

var ldapServiceTypeSessionSecurityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","sign","seal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServiceTypeSessionSecurityPropEnum = append(ldapServiceTypeSessionSecurityPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// session_security
	// SessionSecurity
	// none
	// END RIPPY DEBUGGING
	// LdapServiceSessionSecurityNone captures enum value "none"
	LdapServiceSessionSecurityNone string = "none"

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// session_security
	// SessionSecurity
	// sign
	// END RIPPY DEBUGGING
	// LdapServiceSessionSecuritySign captures enum value "sign"
	LdapServiceSessionSecuritySign string = "sign"

	// BEGIN RIPPY DEBUGGING
	// ldap_service
	// LdapService
	// session_security
	// SessionSecurity
	// seal
	// END RIPPY DEBUGGING
	// LdapServiceSessionSecuritySeal captures enum value "seal"
	LdapServiceSessionSecuritySeal string = "seal"
)

// prop value enum
func (m *LdapService) validateSessionSecurityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServiceTypeSessionSecurityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapService) validateSessionSecurity(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionSecurity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSessionSecurityEnum("session_security", "body", *m.SessionSecurity); err != nil {
		return err
	}

	return nil
}

func (m *LdapService) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap service based on the context it is used
func (m *LdapService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapService) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LdapService) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapService) UnmarshalBinary(b []byte) error {
	var res LdapService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServiceLinks ldap service links
//
// swagger:model LdapServiceLinks
type LdapServiceLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ldap service links
func (m *LdapServiceLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServiceLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ldap service links based on the context it is used
func (m *LdapServiceLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServiceLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LdapServiceLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServiceLinks) UnmarshalBinary(b []byte) error {
	var res LdapServiceLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServiceSvm ldap service svm
//
// swagger:model LdapServiceSvm
type LdapServiceSvm struct {

	// links
	Links *LdapServiceSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this ldap service svm
func (m *LdapServiceSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServiceSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap service svm based on the context it is used
func (m *LdapServiceSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServiceSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapServiceSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServiceSvm) UnmarshalBinary(b []byte) error {
	var res LdapServiceSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServiceSvmLinks ldap service svm links
//
// swagger:model LdapServiceSvmLinks
type LdapServiceSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ldap service svm links
func (m *LdapServiceSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServiceSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap service svm links based on the context it is used
func (m *LdapServiceSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServiceSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapServiceSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServiceSvmLinks) UnmarshalBinary(b []byte) error {
	var res LdapServiceSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY