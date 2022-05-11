// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125ClusterConfig hashicorp cloud vault 20201125 cluster config
//
// swagger:model hashicorp.cloud.vault_20201125.ClusterConfig
type HashicorpCloudVault20201125ClusterConfig struct {

	// audit config
	AuditConfig *HashicorpCloudVault20201125AuditConfig `json:"audit_config,omitempty"`

	// audit log export config
	AuditLogExportConfig *HashicorpCloudVault20201125ObservabilityConfig `json:"audit_log_export_config,omitempty"`

	// capacity config
	CapacityConfig *HashicorpCloudVault20201125CapacityConfig `json:"capacity_config,omitempty"`

	// maintenance config
	MaintenanceConfig HashicorpCloudVault20201125MaintenanceConfig `json:"maintenance_config,omitempty"`

	// major version upgrade config
	MajorVersionUpgradeConfig *HashicorpCloudVault20201125MajorVersionUpgradeConfig `json:"major_version_upgrade_config,omitempty"`

	// metrics config
	MetricsConfig *HashicorpCloudVault20201125ObservabilityConfig `json:"metrics_config,omitempty"`

	// network config
	NetworkConfig *HashicorpCloudVault20201125NetworkConfig `json:"network_config,omitempty"`

	// snapshot config
	SnapshotConfig *HashicorpCloudVault20201125SnapshotConfig `json:"snapshot_config,omitempty"`

	// tier
	Tier HashicorpCloudVault20201125Tier `json:"tier,omitempty"`

	// vault access
	VaultAccess *HashicorpCloudVault20201125VaultAccess `json:"vault_access,omitempty"`

	// vault config
	VaultConfig *HashicorpCloudVault20201125VaultConfig `json:"vault_config,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 cluster config
func (m *HashicorpCloudVault20201125ClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogExportConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMajorVersionUpgradeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateAuditConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditConfig) { // not required
		return nil
	}

	if m.AuditConfig != nil {
		if err := m.AuditConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audit_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateAuditLogExportConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogExportConfig) { // not required
		return nil
	}

	if m.AuditLogExportConfig != nil {
		if err := m.AuditLogExportConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audit_log_export_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateCapacityConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.CapacityConfig) { // not required
		return nil
	}

	if m.CapacityConfig != nil {
		if err := m.CapacityConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateMajorVersionUpgradeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.MajorVersionUpgradeConfig) { // not required
		return nil
	}

	if m.MajorVersionUpgradeConfig != nil {
		if err := m.MajorVersionUpgradeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("major_version_upgrade_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateMetricsConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.MetricsConfig) { // not required
		return nil
	}

	if m.MetricsConfig != nil {
		if err := m.MetricsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateNetworkConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkConfig) { // not required
		return nil
	}

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateSnapshotConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotConfig) { // not required
		return nil
	}

	if m.SnapshotConfig != nil {
		if err := m.SnapshotConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateTier(formats strfmt.Registry) error {

	if swag.IsZero(m.Tier) { // not required
		return nil
	}

	if err := m.Tier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		}
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateVaultAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.VaultAccess) { // not required
		return nil
	}

	if m.VaultAccess != nil {
		if err := m.VaultAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_access")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125ClusterConfig) validateVaultConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.VaultConfig) { // not required
		return nil
	}

	if m.VaultConfig != nil {
		if err := m.VaultConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125ClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125ClusterConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125ClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
