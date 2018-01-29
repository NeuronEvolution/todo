// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/NeuronEvolution/todo/api-private/gen/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../gen --name  --spec ../swagger.json --template-dir /Users/god/work/neuron/src/github.com/NeuronFramework/restful/go_template/

func configureFlags(api *operations.TodoPrivateAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TodoPrivateAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AddTodoHandler = operations.AddTodoHandlerFunc(func(params operations.AddTodoParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddTodo has not yet been implemented")
	})
	api.GetTodoHandler = operations.GetTodoHandlerFunc(func(params operations.GetTodoParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetTodo has not yet been implemented")
	})
	api.GetTodoListHandler = operations.GetTodoListHandlerFunc(func(params operations.GetTodoListParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetTodoList has not yet been implemented")
	})
	api.RemoveTodoHandler = operations.RemoveTodoHandlerFunc(func(params operations.RemoveTodoParams) middleware.Responder {
		return middleware.NotImplemented("operation .RemoveTodo has not yet been implemented")
	})
	api.UpdateTodoHandler = operations.UpdateTodoHandlerFunc(func(params operations.UpdateTodoParams) middleware.Responder {
		return middleware.NotImplemented("operation .UpdateTodo has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
