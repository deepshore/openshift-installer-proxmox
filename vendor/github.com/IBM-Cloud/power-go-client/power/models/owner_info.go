// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OwnerInfo owner info
// swagger:model OwnerInfo
type OwnerInfo struct {

	// Country code of user
	// Required: true
	CountryCode *string `json:"countryCode"`

	// Currency code of user
	// Required: true
	CurrencyCode *string `json:"currencyCode"`

	// Email address of user
	// Required: true
	Email *string `json:"email"`

	// IAM id of user
	// Required: true
	IamID *string `json:"iamID"`

	// Indicates if user is an IBMer
	// Required: true
	IsIBMer *bool `json:"isIBMer"`

	// Name of user
	// Required: true
	Name *string `json:"name"`

	// Array of Soft Layer IDs
	// Required: true
	SoftlayerIds []string `json:"softlayerIDs"`

	// User id of user
	// Required: true
	UserID *string `json:"userID"`
}

// Validate validates this owner info
func (m *OwnerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrencyCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIamID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsIBMer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftlayerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OwnerInfo) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	return nil
}

func (m *OwnerInfo) validateCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("currencyCode", "body", m.CurrencyCode); err != nil {
		return err
	}

	return nil
}

func (m *OwnerInfo) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *OwnerInfo) validateIamID(formats strfmt.Registry) error {

	if err := validate.Required("iamID", "body", m.IamID); err != nil {
		return err
	}

	return nil
}

func (m *OwnerInfo) validateIsIBMer(formats strfmt.Registry) error {

	if err := validate.Required("isIBMer", "body", m.IsIBMer); err != nil {
		return err
	}

	return nil
}

func (m *OwnerInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OwnerInfo) validateSoftlayerIds(formats strfmt.Registry) error {

	if err := validate.Required("softlayerIDs", "body", m.SoftlayerIds); err != nil {
		return err
	}

	return nil
}

func (m *OwnerInfo) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userID", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OwnerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OwnerInfo) UnmarshalBinary(b []byte) error {
	var res OwnerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}