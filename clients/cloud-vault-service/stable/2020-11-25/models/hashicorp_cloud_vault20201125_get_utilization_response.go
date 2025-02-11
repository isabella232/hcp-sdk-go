// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125GetUtilizationResponse hashicorp cloud vault 20201125 get utilization response
//
// swagger:model hashicorp.cloud.vault_20201125.GetUtilizationResponse
type HashicorpCloudVault20201125GetUtilizationResponse struct {

	// clients
	Clients *HashicorpCloudVault20201125GetUtilizationResponseValue `json:"clients,omitempty"`

	// disk usage bytes
	DiskUsageBytes *HashicorpCloudVault20201125GetUtilizationResponseValue `json:"disk_usage_bytes,omitempty"`

	// distinct entities
	DistinctEntities *HashicorpCloudVault20201125GetUtilizationResponseValue `json:"distinct_entities,omitempty"`

	// non entity tokens
	NonEntityTokens *HashicorpCloudVault20201125GetUtilizationResponseValue `json:"non_entity_tokens,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 get utilization response
func (m *HashicorpCloudVault20201125GetUtilizationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskUsageBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistinctEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonEntityTokens(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125GetUtilizationResponse) validateClients(formats strfmt.Registry) error {

	if swag.IsZero(m.Clients) { // not required
		return nil
	}

	if m.Clients != nil {
		if err := m.Clients.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clients")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125GetUtilizationResponse) validateDiskUsageBytes(formats strfmt.Registry) error {

	if swag.IsZero(m.DiskUsageBytes) { // not required
		return nil
	}

	if m.DiskUsageBytes != nil {
		if err := m.DiskUsageBytes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_usage_bytes")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125GetUtilizationResponse) validateDistinctEntities(formats strfmt.Registry) error {

	if swag.IsZero(m.DistinctEntities) { // not required
		return nil
	}

	if m.DistinctEntities != nil {
		if err := m.DistinctEntities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distinct_entities")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125GetUtilizationResponse) validateNonEntityTokens(formats strfmt.Registry) error {

	if swag.IsZero(m.NonEntityTokens) { // not required
		return nil
	}

	if m.NonEntityTokens != nil {
		if err := m.NonEntityTokens.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("non_entity_tokens")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetUtilizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetUtilizationResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125GetUtilizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
