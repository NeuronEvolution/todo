// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TodoItemGroup todo item group
// swagger:model TodoItemGroup
type TodoItemGroup struct {

	// category
	// Required: true
	Category *string `json:"category"`

	// todo item list
	// Required: true
	TodoItemList []*TodoItem `json:"todoItemList"`
}

// Validate validates this todo item group
func (m *TodoItemGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTodoItemList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TodoItemGroup) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *TodoItemGroup) validateTodoItemList(formats strfmt.Registry) error {

	if err := validate.Required("todoItemList", "body", m.TodoItemList); err != nil {
		return err
	}

	for i := 0; i < len(m.TodoItemList); i++ {
		if swag.IsZero(m.TodoItemList[i]) { // not required
			continue
		}

		if m.TodoItemList[i] != nil {
			if err := m.TodoItemList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("todoItemList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TodoItemGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TodoItemGroup) UnmarshalBinary(b []byte) error {
	var res TodoItemGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
