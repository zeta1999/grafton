package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"

	manifold "github.com/manifoldco/go-manifold"
)

// ResourceRequest The information sent along to a Provider to provision a resource.
//
// swagger:model ResourceRequest
type ResourceRequest struct {

	// A map of feature labels to selected values for customizable features
	Features map[string]interface{} `json:"features,omitempty"`

	// id
	// Required: true
	ID manifold.ID `json:"id"`

	// plan
	// Required: true
	Plan manifold.Label `json:"plan"`

	// product
	// Required: true
	Product manifold.Label `json:"product"`

	// region
	// Required: true
	Region RegionSlug `json:"region"`
}

// Validate validates this resource request
func (m *ResourceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceRequest) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *ResourceRequest) validatePlan(formats strfmt.Registry) error {

	if err := m.Plan.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("plan")
		}
		return err
	}

	return nil
}

func (m *ResourceRequest) validateProduct(formats strfmt.Registry) error {

	if err := m.Product.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("product")
		}
		return err
	}

	return nil
}

func (m *ResourceRequest) validateRegion(formats strfmt.Registry) error {

	if err := m.Region.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("region")
		}
		return err
	}

	return nil
}
