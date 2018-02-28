// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/NeuronEvolution/todo/api-private/gen/models"
)

// GetUserProfileOKCode is the HTTP code returned for type GetUserProfileOK
const GetUserProfileOKCode int = 200

/*GetUserProfileOK ok

swagger:response getUserProfileOK
*/
type GetUserProfileOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserProfile `json:"body,omitempty"`
}

// NewGetUserProfileOK creates GetUserProfileOK with default headers values
func NewGetUserProfileOK() *GetUserProfileOK {

	return &GetUserProfileOK{}
}

// WithPayload adds the payload to the get user profile o k response
func (o *GetUserProfileOK) WithPayload(payload *models.UserProfile) *GetUserProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user profile o k response
func (o *GetUserProfileOK) SetPayload(payload *models.UserProfile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
