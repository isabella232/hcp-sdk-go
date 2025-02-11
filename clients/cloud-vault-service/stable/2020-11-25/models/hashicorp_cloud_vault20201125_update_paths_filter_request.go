// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudVault20201125UpdatePathsFilterRequest hashicorp cloud vault 20201125 update paths filter request
//
// swagger:model hashicorp.cloud.vault_20201125.UpdatePathsFilterRequest
type HashicorpCloudVault20201125UpdatePathsFilterRequest struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// mode
	Mode HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode `json:"mode,omitempty"`

	// paths
	Paths []string `json:"paths"`
}

// Validate validates this hashicorp cloud vault 20201125 update paths filter request
func (m *HashicorpCloudVault20201125UpdatePathsFilterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125UpdatePathsFilterRequest) validateLocation(formats strfmt.Registry) error {

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

func (m *HashicorpCloudVault20201125UpdatePathsFilterRequest) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125UpdatePathsFilterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125UpdatePathsFilterRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125UpdatePathsFilterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
