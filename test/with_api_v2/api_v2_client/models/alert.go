// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Alert alert
// swagger:model alert
type Alert struct {

	// annotations
	Annotations LabelSet `json:"annotations,omitempty"`

	// ends at
	// Format: date-time
	EndsAt strfmt.DateTime `json:"endsAt,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// generator URL
	// Format: uri
	GeneratorURL strfmt.URI `json:"generatorURL,omitempty"`

	// labels
	Labels LabelSet `json:"labels,omitempty"`

	// receivers
	Receivers []*Receiver `json:"receivers"`

	// starts at
	// Format: date-time
	StartsAt strfmt.DateTime `json:"startsAt,omitempty"`

	// status
	Status *AlertStatus `json:"status,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneratorURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) validateAnnotations(formats strfmt.Registry) error {

	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	if err := m.Annotations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *Alert) validateEndsAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EndsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("endsAt", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateGeneratorURL(formats strfmt.Registry) error {

	if swag.IsZero(m.GeneratorURL) { // not required
		return nil
	}

	if err := validate.FormatOf("generatorURL", "body", "uri", m.GeneratorURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		}
		return err
	}

	return nil
}

func (m *Alert) validateReceivers(formats strfmt.Registry) error {

	if swag.IsZero(m.Receivers) { // not required
		return nil
	}

	for i := 0; i < len(m.Receivers); i++ {
		if swag.IsZero(m.Receivers[i]) { // not required
			continue
		}

		if m.Receivers[i] != nil {
			if err := m.Receivers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Alert) validateStartsAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("startsAt", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Alert) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AlertStatus alert status
// swagger:model AlertStatus
type AlertStatus struct {

	// inhibited by
	InhibitedBy []string `json:"inhibitedBy"`

	// silenced by
	SilencedBy []string `json:"silencedBy"`

	// state
	// Enum: [unprocessed active suppressed]
	State string `json:"state,omitempty"`
}

// Validate validates this alert status
func (m *AlertStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var alertStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unprocessed","active","suppressed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertStatusTypeStatePropEnum = append(alertStatusTypeStatePropEnum, v)
	}
}

const (

	// AlertStatusStateUnprocessed captures enum value "unprocessed"
	AlertStatusStateUnprocessed string = "unprocessed"

	// AlertStatusStateActive captures enum value "active"
	AlertStatusStateActive string = "active"

	// AlertStatusStateSuppressed captures enum value "suppressed"
	AlertStatusStateSuppressed string = "suppressed"
)

// prop value enum
func (m *AlertStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, alertStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AlertStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("status"+"."+"state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertStatus) UnmarshalBinary(b []byte) error {
	var res AlertStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}