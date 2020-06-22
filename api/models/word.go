// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Word word
//
// swagger:model Word
type Word struct {

	// id
	// Read Only: true
	ID uint64 `json:"id,omitempty"`

	// origin
	// Required: true
	Origin *string `json:"origin"`

	// samples
	// Required: true
	Samples []string `json:"samples"`

	// transcription
	// Required: true
	Transcription *string `json:"transcription"`

	// translations
	// Required: true
	Translations []*Translation `json:"translations"`

	// word
	// Required: true
	Word *string `json:"word"`
}

// Validate validates this word
func (m *Word) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamples(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranscription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranslations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWord(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Word) validateOrigin(formats strfmt.Registry) error {

	if err := validate.Required("origin", "body", m.Origin); err != nil {
		return err
	}

	return nil
}

func (m *Word) validateSamples(formats strfmt.Registry) error {

	if err := validate.Required("samples", "body", m.Samples); err != nil {
		return err
	}

	return nil
}

func (m *Word) validateTranscription(formats strfmt.Registry) error {

	if err := validate.Required("transcription", "body", m.Transcription); err != nil {
		return err
	}

	return nil
}

func (m *Word) validateTranslations(formats strfmt.Registry) error {

	if err := validate.Required("translations", "body", m.Translations); err != nil {
		return err
	}

	for i := 0; i < len(m.Translations); i++ {
		if swag.IsZero(m.Translations[i]) { // not required
			continue
		}

		if m.Translations[i] != nil {
			if err := m.Translations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("translations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Word) validateWord(formats strfmt.Registry) error {

	if err := validate.Required("word", "body", m.Word); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Word) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Word) UnmarshalBinary(b []byte) error {
	var res Word
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}