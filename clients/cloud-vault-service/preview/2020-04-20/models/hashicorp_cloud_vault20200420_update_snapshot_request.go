// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudVault20200420UpdateSnapshotRequest UpdateSnapshotRequest is a request to update a snapshot.
//
// swagger:model hashicorp.cloud.vault_20200420.UpdateSnapshotRequest
type HashicorpCloudVault20200420UpdateSnapshotRequest struct {

	// mask is the mask of fields to update.
	Mask *cloud.GoogleProtobufFieldMask `json:"mask,omitempty"`

	// snapshot contains the fields to update.
	//
	// Supported fields: name
	Snapshot *HashicorpCloudVault20200420Snapshot `json:"snapshot,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 update snapshot request
func (m *HashicorpCloudVault20200420UpdateSnapshotRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420UpdateSnapshotRequest) validateMask(formats strfmt.Registry) error {
	if swag.IsZero(m.Mask) { // not required
		return nil
	}

	if m.Mask != nil {
		if err := m.Mask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mask")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420UpdateSnapshotRequest) validateSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {
		if err := m.Snapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20200420 update snapshot request based on the context it is used
func (m *HashicorpCloudVault20200420UpdateSnapshotRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420UpdateSnapshotRequest) contextValidateMask(ctx context.Context, formats strfmt.Registry) error {

	if m.Mask != nil {
		if err := m.Mask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mask")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420UpdateSnapshotRequest) contextValidateSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.Snapshot != nil {
		if err := m.Snapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420UpdateSnapshotRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420UpdateSnapshotRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420UpdateSnapshotRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
