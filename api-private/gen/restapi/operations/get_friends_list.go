// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

// GetFriendsListHandlerFunc turns a function with the right signature into a get friends list handler
type GetFriendsListHandlerFunc func(GetFriendsListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFriendsListHandlerFunc) Handle(params GetFriendsListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetFriendsListHandler interface for that can handle valid get friends list params
type GetFriendsListHandler interface {
	Handle(GetFriendsListParams, interface{}) middleware.Responder
}

// NewGetFriendsList creates a new http.Handler for the get friends list operation
func NewGetFriendsList(ctx *middleware.Context, handler GetFriendsListHandler) *GetFriendsList {
	return &GetFriendsList{Context: ctx, Handler: handler}
}

/*GetFriendsList swagger:route GET /friends getFriendsList

GetFriendsList get friends list API

*/
type GetFriendsList struct {
	Context *middleware.Context
	Handler GetFriendsListHandler
}

func (o *GetFriendsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	zap.L().Named("api").Info("GetFriendsList")

	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFriendsListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		zap.L().Named("api").Info("GetFriendsList", zap.Error(err))
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
		zap.L().Named("api").Info("GetFriendsList", zap.Error(err))
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	zap.L().Named("api").Info("GetFriendsList", zap.Any("request", &Params))

	res := o.Handler.Handle(Params, principal) // actually handle the request

	zap.L().Named("api").Info("GetFriendsList", zap.Any("response", res))

	o.Context.Respond(rw, r, route.Produces, route, res)

}
