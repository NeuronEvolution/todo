// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

// UpdateTodoHandlerFunc turns a function with the right signature into a update todo handler
type UpdateTodoHandlerFunc func(UpdateTodoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateTodoHandlerFunc) Handle(params UpdateTodoParams) middleware.Responder {
	return fn(params)
}

// UpdateTodoHandler interface for that can handle valid update todo params
type UpdateTodoHandler interface {
	Handle(UpdateTodoParams) middleware.Responder
}

// NewUpdateTodo creates a new http.Handler for the update todo operation
func NewUpdateTodo(ctx *middleware.Context, handler UpdateTodoHandler) *UpdateTodo {
	return &UpdateTodo{Context: ctx, Handler: handler}
}

/*UpdateTodo swagger:route POST /{todoId} updateTodo

UpdateTodo update todo API

*/
type UpdateTodo struct {
	Context *middleware.Context
	Handler UpdateTodoHandler
}

func (o *UpdateTodo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateTodoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	zap.L().Named("api").Info("UpdateTodo", zap.Any("request", &Params))

	res := o.Handler.Handle(Params) // actually handle the request

	zap.L().Named("api").Info("UpdateTodo", zap.Any("response", res))

	o.Context.Respond(rw, r, route.Produces, route, res)

}
