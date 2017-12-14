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

// K8sStatus Status of Kubernetes integration
// swagger:model K8sStatus
type K8sStatus struct {

	// k8s api versions
	K8sAPIVersions []string `json:"k8s-api-versions"`

	// Human readable status/error/warning message
	Msg string `json:"msg,omitempty"`

	// State the component is in
	State string `json:"state,omitempty"`
}

// Validate validates this k8s status
func (m *K8sStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateK8sAPIVersions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sStatus) validateK8sAPIVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.K8sAPIVersions) { // not required
		return nil
	}

	return nil
}

var k8sStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ok","Warning","Failure","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		k8sStatusTypeStatePropEnum = append(k8sStatusTypeStatePropEnum, v)
	}
}

const (
	// K8sStatusStateOk captures enum value "Ok"
	K8sStatusStateOk string = "Ok"
	// K8sStatusStateWarning captures enum value "Warning"
	K8sStatusStateWarning string = "Warning"
	// K8sStatusStateFailure captures enum value "Failure"
	K8sStatusStateFailure string = "Failure"
	// K8sStatusStateDisabled captures enum value "Disabled"
	K8sStatusStateDisabled string = "Disabled"
)

// prop value enum
func (m *K8sStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, k8sStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *K8sStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sStatus) UnmarshalBinary(b []byte) error {
	var res K8sStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
