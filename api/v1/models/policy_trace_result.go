// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PolicyTraceResult Response to a policy resolution process
// swagger:model PolicyTraceResult
type PolicyTraceResult struct {

	// log
	Log string `json:"log,omitempty"`

	// verdict
	Verdict string `json:"verdict,omitempty"`
}

// Validate validates this policy trace result
func (m *PolicyTraceResult) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTraceResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTraceResult) UnmarshalBinary(b []byte) error {
	var res PolicyTraceResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
