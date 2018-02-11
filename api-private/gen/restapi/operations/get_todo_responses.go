// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/todo/api-private/gen/models"
)

// GetTodoOKCode is the HTTP code returned for type GetTodoOK
const GetTodoOKCode int = 200

/*GetTodoOK ok

swagger:response getTodoOK
*/
type GetTodoOK struct {

	/*
	  In: Body
	*/
	Payload *models.TodoItem `json:"body,omitempty"`
}

// NewGetTodoOK creates GetTodoOK with default headers values
func NewGetTodoOK() *GetTodoOK {
	return &GetTodoOK{}
}

// WithPayload adds the payload to the get todo o k response
func (o *GetTodoOK) WithPayload(payload *models.TodoItem) *GetTodoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo o k response
func (o *GetTodoOK) SetPayload(payload *models.TodoItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}