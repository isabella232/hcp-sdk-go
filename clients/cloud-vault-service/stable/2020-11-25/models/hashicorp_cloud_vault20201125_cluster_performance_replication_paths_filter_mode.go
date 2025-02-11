// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode Mode represents the type of filter to be applied
//
// - CLUSTER_PERFORMANCE_REPLICATION_PATHS_FILTER_MODE_INVALID: PATHS_FILTER_MODE_INVALID is a sentinel zero value so that an uninitialized value can be
// detected.
//  - ALLOW: ALLOW allows only the specified paths for replication
//  - DENY: DENY denies the specified paths for replication
//
// swagger:model hashicorp.cloud.vault_20201125.Cluster.PerformanceReplicationPathsFilter.Mode
type HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode string

const (

	// HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeCLUSTERPERFORMANCEREPLICATIONPATHSFILTERMODEINVALID captures enum value "CLUSTER_PERFORMANCE_REPLICATION_PATHS_FILTER_MODE_INVALID"
	HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeCLUSTERPERFORMANCEREPLICATIONPATHSFILTERMODEINVALID HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode = "CLUSTER_PERFORMANCE_REPLICATION_PATHS_FILTER_MODE_INVALID"

	// HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeALLOW captures enum value "ALLOW"
	HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeALLOW HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode = "ALLOW"

	// HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeDENY captures enum value "DENY"
	HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeDENY HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode = "DENY"
)

// for schema
var hashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeEnum []interface{}

func init() {
	var res []HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode
	if err := json.Unmarshal([]byte(`["CLUSTER_PERFORMANCE_REPLICATION_PATHS_FILTER_MODE_INVALID","ALLOW","DENY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeEnum = append(hashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeEnum, v)
	}
}

func (m HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode) validateHashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeEnum(path, location string, value HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault 20201125 cluster performance replication paths filter mode
func (m HashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVault20201125ClusterPerformanceReplicationPathsFilterModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
