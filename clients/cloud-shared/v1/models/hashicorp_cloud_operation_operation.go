// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudOperationOperation Operation represents a single operation.
//
// swagger:model hashicorp.cloud.operation.Operation
type HashicorpCloudOperationOperation struct {

	// CreatedAt is the timestamp of when the operation was first created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// error is the error that occurred in the operation.
	Error *GoogleRPCStatus `json:"error,omitempty"`

	// id is the unique ID for this operation used in other RPC calls.
	// This ID is only guaranteed to be unique within the region that
	// the operation is running in.
	ID string `json:"id,omitempty"`

	// Link is the resource link the operation is associated with.
	Link *HashicorpCloudLocationLink `json:"link,omitempty"`

	// Location is location of the resource that this operation belongs to.
	Location *HashicorpCloudLocationLocation `json:"location,omitempty"`

	// response is the result of the operation. See the documentation for the API
	// call creating the operation to understand what the value of this is.
	Response *GoogleProtobufAny `json:"response,omitempty"`

	// state is the current state of the operation. This is a simple tri-state:
	// PENDING means the operation is created but not yet started, RUNNING means
	// the operation is currently running (though it may be very long-running),
	// and DONE means the operation is complete whether successfully or not.
	State *HashicorpCloudOperationOperationState `json:"state,omitempty"`

	// UpdatedAt is the timestamp of when the operation was last updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud operation operation
func (m *HashicorpCloudOperationOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *HashicorpCloudOperationOperation) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) validateLink(formats strfmt.Registry) error {
	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud operation operation based on the context it is used
func (m *HashicorpCloudOperationOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudOperationOperation) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	if m.Link != nil {
		if err := m.Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.Response != nil {
		if err := m.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudOperationOperation) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudOperationOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudOperationOperation) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudOperationOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
