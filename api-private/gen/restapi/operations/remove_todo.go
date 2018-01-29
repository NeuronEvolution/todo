// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

// RemoveTodoHandlerFunc turns a function with the right signature into a remove todo handler
type RemoveTodoHandlerFunc func(RemoveTodoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveTodoHandlerFunc) Handle(params RemoveTodoParams) middleware.Responder {
	return fn(params)
}

// RemoveTodoHandler interface for that can handle valid remove todo params
type RemoveTodoHandler interface {
	Handle(RemoveTodoParams) middleware.Responder
}

// NewRemoveTodo creates a new http.Handler for the remove todo operation
func NewRemoveTodo(ctx *middleware.Context, handler RemoveTodoHandler) *RemoveTodo {
	return &RemoveTodo{Context: ctx, Handler: handler}
}

/*RemoveTodo swagger:route DELETE /{todoId} removeTodo

RemoveTodo remove todo API

*/
type RemoveTodo struct {
	Context *middleware.Context
	Handler RemoveTodoHandler
}

func (o *RemoveTodo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveTodoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	zap.L().Named("api").Info("RemoveTodo", zap.Any("request", &Params))

	res := o.Handler.Handle(Params) // actually handle the request

	zap.L().Named("api").Info("RemoveTodo", zap.Any("response", res))

	o.Context.Respond(rw, r, route.Produces, route, res)

}
