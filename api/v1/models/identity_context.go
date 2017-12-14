// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IdentityContext Context describing a pair of source and destination identity
// swagger:model IdentityContext
type IdentityContext struct {

	// dports
	Dports IdentityContextDports `json:"dports"`

	// from
	From Labels `json:"from"`

	// to
	To Labels `json:"to"`

	// Enable verbose tracing.
	//
	Verbose bool `json:"verbose,omitempty"`
}

// Validate validates this identity context
func (m *IdentityContext) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IdentityContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityContext) UnmarshalBinary(b []byte) error {
	var res IdentityContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
