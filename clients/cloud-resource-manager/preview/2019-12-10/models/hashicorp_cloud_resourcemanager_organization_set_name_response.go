// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudResourcemanagerOrganizationSetNameResponse see OrganizationService.SetName
//
// swagger:model hashicorp.cloud.resourcemanager.OrganizationSetNameResponse
type HashicorpCloudResourcemanagerOrganizationSetNameResponse struct {

	// Organization is the specified organization with an updated name.
	Organization *HashicorpCloudResourcemanagerOrganization `json:"organization,omitempty"`
}

// Validate validates this hashicorp cloud resourcemanager organization set name response
func (m *HashicorpCloudResourcemanagerOrganizationSetNameResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudResourcemanagerOrganizationSetNameResponse) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationSetNameResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationSetNameResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudResourcemanagerOrganizationSetNameResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
