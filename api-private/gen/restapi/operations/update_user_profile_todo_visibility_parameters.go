// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/NeuronEvolution/todo/api-private/gen/models"
)

// NewUpdateUserProfileTodoVisibilityParams creates a new UpdateUserProfileTodoVisibilityParams object
// no default values defined in spec.
func NewUpdateUserProfileTodoVisibilityParams() UpdateUserProfileTodoVisibilityParams {

	return UpdateUserProfileTodoVisibilityParams{}
}

// UpdateUserProfileTodoVisibilityParams contains all the bound params for the update user profile todo visibility operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateUserProfileTodoVisibility
type UpdateUserProfileTodoVisibilityParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Visibility models.TodoVisibility
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateUserProfileTodoVisibilityParams() beforehand.
func (o *UpdateUserProfileTodoVisibilityParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.TodoVisibility
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("visibility", "body"))
			} else {
				res = append(res, errors.NewParseError("visibility", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Visibility = body
			}
		}
	} else {
		res = append(res, errors.Required("visibility", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
