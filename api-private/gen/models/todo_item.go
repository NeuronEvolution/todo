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

// TodoItem todo item
// swagger:model TodoItem
type TodoItem struct {

	// category
	// Required: true
	Category *string `json:"category"`

	// desc
	Desc string `json:"desc,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// status
	// Required: true
	Status TodoStatus `json:"status"`

	// title
	// Required: true
	Title *string `json:"title"`

	// todo Id
	// Required: true
	TodoID *string `json:"todoId"`
}

// Validate validates this todo item
func (m *TodoItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTodoID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TodoItem) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *TodoItem) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *TodoItem) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *TodoItem) validateTodoID(formats strfmt.Registry) error {

	if err := validate.Required("todoId", "body", m.TodoID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TodoItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TodoItem) UnmarshalBinary(b []byte) error {
	var res TodoItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
