// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/NeuronFramework/restful"
	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewTodoPrivateAPI creates a new TodoPrivate instance
func NewTodoPrivateAPI(spec *loads.Document) *TodoPrivateAPI {
	return &TodoPrivateAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          restful.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		AddTodoHandler: AddTodoHandlerFunc(func(params AddTodoParams, principal interface{}) middleware.Responder {
			panic("operation AddTodo has not yet been implemented")
		}),
		GetFriendHandler: GetFriendHandlerFunc(func(params GetFriendParams, principal interface{}) middleware.Responder {
			panic("operation GetFriend has not yet been implemented")
		}),
		GetFriendsListHandler: GetFriendsListHandlerFunc(func(params GetFriendsListParams, principal interface{}) middleware.Responder {
			panic("operation GetFriendsList has not yet been implemented")
		}),
		GetTodoHandler: GetTodoHandlerFunc(func(params GetTodoParams, principal interface{}) middleware.Responder {
			panic("operation GetTodo has not yet been implemented")
		}),
		GetTodoListHandler: GetTodoListHandlerFunc(func(params GetTodoListParams, principal interface{}) middleware.Responder {
			panic("operation GetTodoList has not yet been implemented")
		}),
		GetTodoListByCategoryHandler: GetTodoListByCategoryHandlerFunc(func(params GetTodoListByCategoryParams, principal interface{}) middleware.Responder {
			panic("operation GetTodoListByCategory has not yet been implemented")
		}),
		GetUserProfileHandler: GetUserProfileHandlerFunc(func(params GetUserProfileParams, principal interface{}) middleware.Responder {
			panic("operation GetUserProfile has not yet been implemented")
		}),
		RemoveTodoHandler: RemoveTodoHandlerFunc(func(params RemoveTodoParams, principal interface{}) middleware.Responder {
			panic("operation RemoveTodo has not yet been implemented")
		}),
		UpdateTodoHandler: UpdateTodoHandlerFunc(func(params UpdateTodoParams, principal interface{}) middleware.Responder {
			panic("operation UpdateTodo has not yet been implemented")
		}),
		UpdateUserProfileHandler: UpdateUserProfileHandlerFunc(func(params UpdateUserProfileParams, principal interface{}) middleware.Responder {
			panic("operation UpdateUserProfile has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*TodoPrivateAPI the todo private API */
type TodoPrivateAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// AddTodoHandler sets the operation handler for the add todo operation
	AddTodoHandler AddTodoHandler
	// GetFriendHandler sets the operation handler for the get friend operation
	GetFriendHandler GetFriendHandler
	// GetFriendsListHandler sets the operation handler for the get friends list operation
	GetFriendsListHandler GetFriendsListHandler
	// GetTodoHandler sets the operation handler for the get todo operation
	GetTodoHandler GetTodoHandler
	// GetTodoListHandler sets the operation handler for the get todo list operation
	GetTodoListHandler GetTodoListHandler
	// GetTodoListByCategoryHandler sets the operation handler for the get todo list by category operation
	GetTodoListByCategoryHandler GetTodoListByCategoryHandler
	// GetUserProfileHandler sets the operation handler for the get user profile operation
	GetUserProfileHandler GetUserProfileHandler
	// RemoveTodoHandler sets the operation handler for the remove todo operation
	RemoveTodoHandler RemoveTodoHandler
	// UpdateTodoHandler sets the operation handler for the update todo operation
	UpdateTodoHandler UpdateTodoHandler
	// UpdateUserProfileHandler sets the operation handler for the update user profile operation
	UpdateUserProfileHandler UpdateUserProfileHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TodoPrivateAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TodoPrivateAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TodoPrivateAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TodoPrivateAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TodoPrivateAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TodoPrivateAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TodoPrivateAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TodoPrivateAPI
func (o *TodoPrivateAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.AddTodoHandler == nil {
		unregistered = append(unregistered, "AddTodoHandler")
	}

	if o.GetFriendHandler == nil {
		unregistered = append(unregistered, "GetFriendHandler")
	}

	if o.GetFriendsListHandler == nil {
		unregistered = append(unregistered, "GetFriendsListHandler")
	}

	if o.GetTodoHandler == nil {
		unregistered = append(unregistered, "GetTodoHandler")
	}

	if o.GetTodoListHandler == nil {
		unregistered = append(unregistered, "GetTodoListHandler")
	}

	if o.GetTodoListByCategoryHandler == nil {
		unregistered = append(unregistered, "GetTodoListByCategoryHandler")
	}

	if o.GetUserProfileHandler == nil {
		unregistered = append(unregistered, "GetUserProfileHandler")
	}

	if o.RemoveTodoHandler == nil {
		unregistered = append(unregistered, "RemoveTodoHandler")
	}

	if o.UpdateTodoHandler == nil {
		unregistered = append(unregistered, "UpdateTodoHandler")
	}

	if o.UpdateUserProfileHandler == nil {
		unregistered = append(unregistered, "UpdateUserProfileHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TodoPrivateAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TodoPrivateAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "Bearer":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *TodoPrivateAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *TodoPrivateAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TodoPrivateAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TodoPrivateAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the todo private API
func (o *TodoPrivateAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TodoPrivateAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"][""] = NewAddTodo(o.context, o.AddTodoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/friends/{friendID}"] = NewGetFriend(o.context, o.GetFriendHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/friends"] = NewGetFriendsList(o.context, o.GetFriendsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{todoId}"] = NewGetTodo(o.context, o.GetTodoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = NewGetTodoList(o.context, o.GetTodoListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/listByCategory"] = NewGetTodoListByCategory(o.context, o.GetTodoListByCategoryHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/userProfile"] = NewGetUserProfile(o.context, o.GetUserProfileHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/{todoId}"] = NewRemoveTodo(o.context, o.RemoveTodoHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/{todoId}"] = NewUpdateTodo(o.context, o.UpdateTodoHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/userProfile"] = NewUpdateUserProfile(o.context, o.UpdateUserProfileHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TodoPrivateAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TodoPrivateAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
