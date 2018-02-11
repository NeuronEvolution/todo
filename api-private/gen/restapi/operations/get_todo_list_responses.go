// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/todo/api-private/gen/models"
)

// GetTodoListOKCode is the HTTP code returned for type GetTodoListOK
const GetTodoListOKCode int = 200

/*GetTodoListOK ok

swagger:response getTodoListOK
*/
type GetTodoListOK struct {

	/*
	  In: Body
	*/
	Payload models.GetTodoListOKBody `json:"body,omitempty"`
}

// NewGetTodoListOK creates GetTodoListOK with default headers values
func NewGetTodoListOK() *GetTodoListOK {
	return &GetTodoListOK{}
}

// WithPayload adds the payload to the get todo list o k response
func (o *GetTodoListOK) WithPayload(payload models.GetTodoListOKBody) *GetTodoListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo list o k response
func (o *GetTodoListOK) SetPayload(payload models.GetTodoListOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetTodoListOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}