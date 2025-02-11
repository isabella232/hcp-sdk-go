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

// HashicorpCloudVault20201125GetUtilizationResponseValueUnit hashicorp cloud vault 20201125 get utilization response value unit
//
// swagger:model hashicorp.cloud.vault_20201125.GetUtilizationResponse.Value.Unit
type HashicorpCloudVault20201125GetUtilizationResponseValueUnit string

const (

	// HashicorpCloudVault20201125GetUtilizationResponseValueUnitGETUTILIZATIONRESPONSEVALUEUNITINVALID captures enum value "GET_UTILIZATION_RESPONSE_VALUE_UNIT_INVALID"
	HashicorpCloudVault20201125GetUtilizationResponseValueUnitGETUTILIZATIONRESPONSEVALUEUNITINVALID HashicorpCloudVault20201125GetUtilizationResponseValueUnit = "GET_UTILIZATION_RESPONSE_VALUE_UNIT_INVALID"

	// HashicorpCloudVault20201125GetUtilizationResponseValueUnitBYTES captures enum value "BYTES"
	HashicorpCloudVault20201125GetUtilizationResponseValueUnitBYTES HashicorpCloudVault20201125GetUtilizationResponseValueUnit = "BYTES"

	// HashicorpCloudVault20201125GetUtilizationResponseValueUnitCOUNT captures enum value "COUNT"
	HashicorpCloudVault20201125GetUtilizationResponseValueUnitCOUNT HashicorpCloudVault20201125GetUtilizationResponseValueUnit = "COUNT"
)

// for schema
var hashicorpCloudVault20201125GetUtilizationResponseValueUnitEnum []interface{}

func init() {
	var res []HashicorpCloudVault20201125GetUtilizationResponseValueUnit
	if err := json.Unmarshal([]byte(`["GET_UTILIZATION_RESPONSE_VALUE_UNIT_INVALID","BYTES","COUNT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVault20201125GetUtilizationResponseValueUnitEnum = append(hashicorpCloudVault20201125GetUtilizationResponseValueUnitEnum, v)
	}
}

func (m HashicorpCloudVault20201125GetUtilizationResponseValueUnit) validateHashicorpCloudVault20201125GetUtilizationResponseValueUnitEnum(path, location string, value HashicorpCloudVault20201125GetUtilizationResponseValueUnit) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVault20201125GetUtilizationResponseValueUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault 20201125 get utilization response value unit
func (m HashicorpCloudVault20201125GetUtilizationResponseValueUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVault20201125GetUtilizationResponseValueUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
