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

// HashicorpCloudVault20201125InputCluster hashicorp cloud vault 20201125 input cluster
//
// swagger:model hashicorp.cloud.vault_20201125.InputCluster
type HashicorpCloudVault20201125InputCluster struct {

	// config for the cluster
	Config *HashicorpCloudVault20201125InputClusterConfig `json:"config,omitempty"`

	// id is ID of the Vault cluster.
	ID string `json:"id,omitempty"`

	// location is the location of the cluster.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// performance_replication_paths_filter is the information about what paths should be
	// filtered for a performance replication secondary.
	PerformanceReplicationPathsFilter *HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilter `json:"performance_replication_paths_filter,omitempty"`

	// performance_replication_primary_cluster holds the link information of the
	// primary cluster under performance replication.
	PerformanceReplicationPrimaryCluster *cloud.HashicorpCloudLocationLink `json:"performance_replication_primary_cluster,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 input cluster
func (m *HashicorpCloudVault20201125InputCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceReplicationPathsFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceReplicationPrimaryCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125InputCluster) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputCluster) validateLocation(formats strfmt.Registry) error {

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

func (m *HashicorpCloudVault20201125InputCluster) validatePerformanceReplicationPathsFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.PerformanceReplicationPathsFilter) { // not required
		return nil
	}

	if m.PerformanceReplicationPathsFilter != nil {
		if err := m.PerformanceReplicationPathsFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performance_replication_paths_filter")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputCluster) validatePerformanceReplicationPrimaryCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.PerformanceReplicationPrimaryCluster) { // not required
		return nil
	}

	if m.PerformanceReplicationPrimaryCluster != nil {
		if err := m.PerformanceReplicationPrimaryCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performance_replication_primary_cluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125InputCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125InputCluster) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125InputCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
