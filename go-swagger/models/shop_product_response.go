// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShopProductResponse shop product response
//
// swagger:model shopProductResponse
type ShopProductResponse struct {

	// count
	Count string `json:"count,omitempty"`

	// favorite
	Favorite bool `json:"favorite,omitempty"`

	// image
	Image string `json:"image,omitempty"`

	// price
	Price float32 `json:"price,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// vote
	Vote []*ProductResponseVote `json:"vote"`
}

// Validate validates this shop product response
func (m *ShopProductResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShopProductResponse) validateVote(formats strfmt.Registry) error {

	if swag.IsZero(m.Vote) { // not required
		return nil
	}

	for i := 0; i < len(m.Vote); i++ {
		if swag.IsZero(m.Vote[i]) { // not required
			continue
		}

		if m.Vote[i] != nil {
			if err := m.Vote[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vote" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShopProductResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShopProductResponse) UnmarshalBinary(b []byte) error {
	var res ShopProductResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
