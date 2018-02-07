// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

// GetTodoListHandlerFunc turns a function with the right signature into a get todo list handler
type GetTodoListHandlerFunc func(GetTodoListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTodoListHandlerFunc) Handle(params GetTodoListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTodoListHandler interface for that can handle valid get todo list params
type GetTodoListHandler interface {
	Handle(GetTodoListParams, interface{}) middleware.Responder
}

// NewGetTodoList creates a new http.Handler for the get todo list operation
func NewGetTodoList(ctx *middleware.Context, handler GetTodoListHandler) *GetTodoList {
	return &GetTodoList{Context: ctx, Handler: handler}
}

/*GetTodoList swagger:route GET / getTodoList

GetTodoList get todo list API

*/
type GetTodoList struct {
	Context *middleware.Context
	Handler GetTodoListHandler
}

func (o *GetTodoList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTodoListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	zap.L().Named("api").Info("GetTodoList", zap.Any("request", &Params))

	res := o.Handler.Handle(Params, principal) // actually handle the request

	zap.L().Named("api").Info("GetTodoList", zap.Any("response", res))

	o.Context.Respond(rw, r, route.Produces, route, res)

}
