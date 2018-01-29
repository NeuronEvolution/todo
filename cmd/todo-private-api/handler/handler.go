package handler

import (
	"context"
	"github.com/NeuronEvolution/todo/api-private/gen/restapi/operations"
	"github.com/NeuronEvolution/todo/services"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type TodoHandler struct {
	logger  *zap.Logger
	service *services.TodoService
}

func New() (h *TodoHandler, err error) {
	h = &TodoHandler{}
	h.logger = log.TypedLogger(h)
	h.service, err = services.New()
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *TodoHandler) GetTodoList(p operations.GetTodoListParams) middleware.Responder {
	result, err := h.service.GetTodoList(context.Background(), "1")
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetTodoListOK().WithPayload(fromTodoItemList(result))
}

func (h *TodoHandler) GetTodo(p operations.GetTodoParams) middleware.Responder {
	todoItem, err := h.service.GetTodo(context.Background(), "1", p.TodoID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetTodoOK().WithPayload(fromTodoItem(todoItem))
}

func (h *TodoHandler) AddTodo(p operations.AddTodoParams) middleware.Responder {
	todoId, err := h.service.AddTodo(context.Background(), "1", toTodoItem(p.TodoItem))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewAddTodoOK().WithPayload(todoId)
}

func (h *TodoHandler) UpdateTodo(p operations.UpdateTodoParams) middleware.Responder {
	err := h.service.UpdateTodo(context.Background(), "1", toTodoItem(p.TodoItem))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUpdateTodoOK()
}

func (h *TodoHandler) RemoveTodo(p operations.RemoveTodoParams) middleware.Responder {
	err := h.service.RemoveTodo(context.Background(), "1", p.TodoID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewRemoveTodoOK()
}
