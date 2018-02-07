package handler

import (
	"context"
	"github.com/NeuronEvolution/todo/api-private/gen/restapi/operations"
	"github.com/NeuronEvolution/todo/services"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/dgrijalva/jwt-go"
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

func (h *TodoHandler) BearerAuth(token string) (accountId interface{}, err error) {
	claims := jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("0123456789"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims.Subject == "" {
		return nil, errors.Unknown("claims.Subject nil")
	}

	return claims.Subject, nil
}

func (h *TodoHandler) GetTodoList(p operations.GetTodoListParams, principal interface{}) middleware.Responder {
	result, err := h.service.GetTodoList(context.Background(), principal.(string))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetTodoListOK().WithPayload(fromTodoItemList(result))
}

func (h *TodoHandler) GetTodo(p operations.GetTodoParams, principal interface{}) middleware.Responder {
	todoItem, err := h.service.GetTodo(context.Background(), principal.(string), p.TodoID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetTodoOK().WithPayload(fromTodoItem(todoItem))
}

func (h *TodoHandler) AddTodo(p operations.AddTodoParams, principal interface{}) middleware.Responder {
	todoId, err := h.service.AddTodo(context.Background(), principal.(string), toTodoItem(p.TodoItem))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewAddTodoOK().WithPayload(todoId)
}

func (h *TodoHandler) UpdateTodo(p operations.UpdateTodoParams, principal interface{}) middleware.Responder {
	err := h.service.UpdateTodo(context.Background(), principal.(string), toTodoItem(p.TodoItem))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUpdateTodoOK()
}

func (h *TodoHandler) RemoveTodo(p operations.RemoveTodoParams, principal interface{}) middleware.Responder {
	err := h.service.RemoveTodo(context.Background(), principal.(string), p.TodoID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewRemoveTodoOK()
}

func (h *TodoHandler) GetTodoListByCategory(p operations.GetTodoListByCategoryParams, principal interface{}) middleware.Responder {
	result, err := h.service.GetTodoListByCategory(context.Background(), principal.(string))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetTodoListByCategoryOK().WithPayload(fromTodoItemGroupList(result))
}
