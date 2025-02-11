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

// HashicorpCloudVault20201125SnapshotType Type represents the type of snapshots.
//
// swagger:model hashicorp.cloud.vault_20201125.Snapshot.Type
type HashicorpCloudVault20201125SnapshotType string

const (

	// HashicorpCloudVault20201125SnapshotTypeSNAPSHOTTYPEINVALID captures enum value "SNAPSHOT_TYPE_INVALID"
	HashicorpCloudVault20201125SnapshotTypeSNAPSHOTTYPEINVALID HashicorpCloudVault20201125SnapshotType = "SNAPSHOT_TYPE_INVALID"

	// HashicorpCloudVault20201125SnapshotTypeAUTOMATIC captures enum value "AUTOMATIC"
	HashicorpCloudVault20201125SnapshotTypeAUTOMATIC HashicorpCloudVault20201125SnapshotType = "AUTOMATIC"

	// HashicorpCloudVault20201125SnapshotTypeSCHEDULED captures enum value "SCHEDULED"
	HashicorpCloudVault20201125SnapshotTypeSCHEDULED HashicorpCloudVault20201125SnapshotType = "SCHEDULED"

	// HashicorpCloudVault20201125SnapshotTypeMANUAL captures enum value "MANUAL"
	HashicorpCloudVault20201125SnapshotTypeMANUAL HashicorpCloudVault20201125SnapshotType = "MANUAL"

	// HashicorpCloudVault20201125SnapshotTypeBEFOREUPGRADE captures enum value "BEFORE_UPGRADE"
	HashicorpCloudVault20201125SnapshotTypeBEFOREUPGRADE HashicorpCloudVault20201125SnapshotType = "BEFORE_UPGRADE"
)

// for schema
var hashicorpCloudVault20201125SnapshotTypeEnum []interface{}

func init() {
	var res []HashicorpCloudVault20201125SnapshotType
	if err := json.Unmarshal([]byte(`["SNAPSHOT_TYPE_INVALID","AUTOMATIC","SCHEDULED","MANUAL","BEFORE_UPGRADE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVault20201125SnapshotTypeEnum = append(hashicorpCloudVault20201125SnapshotTypeEnum, v)
	}
}

func (m HashicorpCloudVault20201125SnapshotType) validateHashicorpCloudVault20201125SnapshotTypeEnum(path, location string, value HashicorpCloudVault20201125SnapshotType) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVault20201125SnapshotTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault 20201125 snapshot type
func (m HashicorpCloudVault20201125SnapshotType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVault20201125SnapshotTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
