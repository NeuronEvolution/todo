// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFriendParams creates a new GetFriendParams object
// no default values defined in spec.
func NewGetFriendParams() GetFriendParams {

	return GetFriendParams{}
}

// GetFriendParams contains all the bound params for the get friend operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFriend
type GetFriendParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	FriendID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFriendParams() beforehand.
func (o *GetFriendParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rFriendID, rhkFriendID, _ := route.Params.GetOK("friendID")
	if err := o.bindFriendID(rFriendID, rhkFriendID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFriendID binds and validates parameter FriendID from path.
func (o *GetFriendParams) bindFriendID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.FriendID = raw

	return nil
}
