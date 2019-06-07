// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BPFMapHistoryEntry BPF map operations history entry
// swagger:model BPFMapHistoryEntry
type BPFMapHistoryEntry struct {

	// Action which was performed
	// Enum: [ok insert delete]
	Action string `json:"action,omitempty"`

	// Error while performing action
	Error string `json:"error,omitempty"`

	// Key of map entry
	Key string `json:"key,omitempty"`

	// Timestamp of operation
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// Value of map entry
	Value string `json:"value,omitempty"`
}

// Validate validates this b p f map history entry
func (m *BPFMapHistoryEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
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

var bPFMapHistoryEntryTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","insert","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bPFMapHistoryEntryTypeActionPropEnum = append(bPFMapHistoryEntryTypeActionPropEnum, v)
	}
}

const (

	// BPFMapHistoryEntryActionOk captures enum value "ok"
	BPFMapHistoryEntryActionOk string = "ok"

	// BPFMapHistoryEntryActionInsert captures enum value "insert"
	BPFMapHistoryEntryActionInsert string = "insert"

	// BPFMapHistoryEntryActionDelete captures enum value "delete"
	BPFMapHistoryEntryActionDelete string = "delete"
)

// prop value enum
func (m *BPFMapHistoryEntry) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bPFMapHistoryEntryTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BPFMapHistoryEntry) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *BPFMapHistoryEntry) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BPFMapHistoryEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BPFMapHistoryEntry) UnmarshalBinary(b []byte) error {
	var res BPFMapHistoryEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
