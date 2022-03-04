// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudPackerChannelIterationPointer A Channel Pointer is a special object that tracks channel
// history by storing both which iteration the channel points to and also when
// the channel was updated to point at said iteration.
//
// swagger:model hashicorp.cloud.packer.ChannelIterationPointer
type HashicorpCloudPackerChannelIterationPointer struct {

	// The user who pointed the channel to the iteration.
	AuthorID string `json:"author_id,omitempty"`

	// When the channel pointer was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The iteration the channel is pointing to.
	Iteration *HashicorpCloudPackerIteration `json:"iteration,omitempty"`
}

// Validate validates this hashicorp cloud packer channel iteration pointer
func (m *HashicorpCloudPackerChannelIterationPointer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIteration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerChannelIterationPointer) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPackerChannelIterationPointer) validateIteration(formats strfmt.Registry) error {

	if swag.IsZero(m.Iteration) { // not required
		return nil
	}

	if m.Iteration != nil {
		if err := m.Iteration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iteration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerChannelIterationPointer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerChannelIterationPointer) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerChannelIterationPointer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
