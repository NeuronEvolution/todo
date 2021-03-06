package handler

import (
	api "github.com/NeuronEvolution/todo/api-private/gen/models"
	"github.com/NeuronEvolution/todo/api-private/gen/restapi/operations"
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/services"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/rest"
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

func (h *TodoHandler) BearerAuth(token string) (userId interface{}, err error) {
	claims := jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("0123456789"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims.Subject == "" {
		return nil, errors.Unknown("验证失败： claims.Subject nil")
	}

	return claims.Subject, nil
}

func (h *TodoHandler) GetTodoList(p operations.GetTodoListParams, userId interface{}) middleware.Responder {
	uid := userId.(string)
	if p.FriendID != nil {
		uid = *p.FriendID
	}

	result, err := h.service.GetTodoList(rest.NewContext(p.HTTPRequest), uid)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetTodoListOK().WithPayload(fromTodoItemList(result))
}

func (h *TodoHandler) GetTodo(p operations.GetTodoParams, userId interface{}) middleware.Responder {
	todoItem, err := h.service.GetTodo(rest.NewContext(p.HTTPRequest), userId.(string), p.TodoID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetTodoOK().WithPayload(fromTodoItem(todoItem))
}

func (h *TodoHandler) AddTodo(p operations.AddTodoParams, userId interface{}) middleware.Responder {
	todoId, err := h.service.AddTodo(rest.NewContext(p.HTTPRequest), userId.(string), toTodoItem(p.TodoItem))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewAddTodoOK().WithPayload(todoId)
}

func (h *TodoHandler) UpdateTodo(p operations.UpdateTodoParams, userId interface{}) middleware.Responder {
	todoItem := toTodoItem(p.TodoItem)

	err := h.service.UpdateTodo(rest.NewContext(p.HTTPRequest), userId.(string), p.TodoID, todoItem)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUpdateTodoOK()
}

func (h *TodoHandler) RemoveTodo(p operations.RemoveTodoParams, userId interface{}) middleware.Responder {
	err := h.service.RemoveTodo(rest.NewContext(p.HTTPRequest), userId.(string), p.TodoID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewRemoveTodoOK()
}

func (h *TodoHandler) GetTodoListByCategory(p operations.GetTodoListByCategoryParams, userId interface{}) middleware.Responder {
	friendId := ""
	if p.FriendID != nil {
		friendId = *p.FriendID
	}

	result, err := h.service.GetTodoListByCategory(rest.NewContext(p.HTTPRequest), userId.(string), friendId)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetTodoListByCategoryOK().WithPayload(fromTodoItemGroupList(result))
}

func (h *TodoHandler) GetUserProfile(p operations.GetUserProfileParams, userId interface{}) middleware.Responder {
	userProfile, err := h.service.GetUserProfile(rest.NewContext(p.HTTPRequest), userId.(string))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetUserProfileOK().WithPayload(fromUserProfile(userProfile))
}

func (h *TodoHandler) UpdateUserProfileTodoVisibility(p operations.UpdateUserProfileTodoVisibilityParams, userID interface{}) middleware.Responder {
	err := h.service.UpdateUserProfileTodoVisibility(rest.NewContext(p.HTTPRequest), userID.(string), toTodoVisibility(p.Visibility))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUpdateUserProfileTodoVisibilityOK()
}

func (h *TodoHandler) UpdateUserProfileUserName(p operations.UpdateUserProfileUserNameParams, userID interface{}) middleware.Responder {
	err := h.service.UpdateUserProfileUserName(rest.NewContext(p.HTTPRequest), userID.(string), p.UserName)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUpdateUserProfileUserNameOK()
}

func (h *TodoHandler) GetFriendsList(p operations.GetFriendsListParams, userId interface{}) middleware.Responder {
	query := &models.FriendsQuery{}
	if p.PageSize != nil {
		query.PageSize = *p.PageSize
	}
	if p.PageToken != nil {
		query.PageToken = *p.PageToken
	}

	result, nextPageToken, err := h.service.GetFriendsList(rest.NewContext(p.HTTPRequest), userId.(string), query)
	if err != nil {
		return errors.Wrap(err)
	}

	response := &api.FriendInfoList{}
	response.NextPageToken = &nextPageToken
	response.Items = fromFriendInfoList(result)

	return operations.NewGetFriendsListOK().WithPayload(response)
}

func (h *TodoHandler) GetFriend(p operations.GetFriendParams, userId interface{}) middleware.Responder {
	friendInfo, err := h.service.GetFriend(rest.NewContext(p.HTTPRequest), userId.(string), p.FriendID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetFriendOK().WithPayload(fromFriendInfo(friendInfo))
}

func (h *TodoHandler) GetCategoryNameList(p operations.GetCategoryNameListParams, userId interface{}) middleware.Responder {
	result, err := h.service.GetCategoryNameList(rest.NewContext(p.HTTPRequest), userId.(string))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewGetCategoryNameListOK().WithPayload(fromCategoryInfoList(result))
}
