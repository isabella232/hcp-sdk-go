// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudNetwork20200907PeeringTarget hashicorp cloud network 20200907 peering target
//
// swagger:model hashicorp.cloud.network_20200907.Peering.Target
type HashicorpCloudNetwork20200907PeeringTarget struct {

	// aws target
	AwsTarget *HashicorpCloudNetwork20200907AWSPeeringTarget `json:"aws_target,omitempty"`

	// azure target
	AzureTarget *HashicorpCloudNetwork20200907AzurePeeringTarget `json:"azure_target,omitempty"`

	// hvn target
	HvnTarget *HashicorpCloudNetwork20200907NetworkTarget `json:"hvn_target,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 peering target
func (m *HashicorpCloudNetwork20200907PeeringTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHvnTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907PeeringTarget) validateAwsTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsTarget) { // not required
		return nil
	}

	if m.AwsTarget != nil {
		if err := m.AwsTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_target")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907PeeringTarget) validateAzureTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.AzureTarget) { // not required
		return nil
	}

	if m.AzureTarget != nil {
		if err := m.AzureTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_target")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907PeeringTarget) validateHvnTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.HvnTarget) { // not required
		return nil
	}

	if m.HvnTarget != nil {
		if err := m.HvnTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hvn_target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907PeeringTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907PeeringTarget) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907PeeringTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
