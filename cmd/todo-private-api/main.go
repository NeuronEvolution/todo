package main

import (
	"github.com/NeuronEvolution/todo/api-private/gen/restapi"
	"github.com/NeuronEvolution/todo/api-private/gen/restapi/operations"
	"github.com/NeuronEvolution/todo/cmd/todo-private-api/handler"
	"github.com/NeuronFramework/restful"
	"github.com/go-openapi/loads"
	"net/http"
	"os"
)

func main() {
	os.Setenv("DEBUG", "true")
	os.Setenv("PORT", "9001")

	restful.Run(func() (http.Handler, error) {
		h, err := handler.New()
		if err != nil {
			return nil, err
		}

		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return nil, err
		}

		api := operations.NewTodoPrivateAPI(swaggerSpec)
		api.BearerAuth = h.BearerAuth
		api.GetTodoListHandler = operations.GetTodoListHandlerFunc(h.GetTodoList)
		api.GetTodoHandler = operations.GetTodoHandlerFunc(h.GetTodo)
		api.AddTodoHandler = operations.AddTodoHandlerFunc(h.AddTodo)
		api.UpdateTodoHandler = operations.UpdateTodoHandlerFunc(h.UpdateTodo)
		api.RemoveTodoHandler = operations.RemoveTodoHandlerFunc(h.RemoveTodo)
		api.GetTodoListByCategoryHandler = operations.GetTodoListByCategoryHandlerFunc(h.GetTodoListByCategory)

		return api.Serve(nil), nil
	})
}
