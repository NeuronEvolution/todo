package services

import (
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/restful"
	"sort"
)

func (s *TodoService) GetTodoListByCategory(ctx *restful.Context, userId string, friendId string) (result []*models.TodoItemGroup, err error) {
	targetUserID := userId
	if friendId != "" && friendId != userId {
		targetUserID = friendId
		dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(targetUserID).QueryOne(ctx, nil)
		if err != nil {
			return nil, err
		}

		if dbUserProfile.TodoVisibility != string(models.TodoVisibilityPublic) {
			return nil, errors.BadRequest("", "该用户的计划不公开")
		}
	}

	dbTodoList, err := s.todoDB.Todo.GetQuery().UserId_Equal(targetUserID).QueryList(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbTodoList == nil {
		return nil, nil
	}

	resultMap := make(map[string]*models.TodoItemGroup)
	for _, v := range dbTodoList {
		m, _ := resultMap[v.TodoCategory]
		if m == nil {
			m = &models.TodoItemGroup{
				Category:     v.TodoCategory,
				TodoItemList: make([]*models.TodoItem, 0),
			}
			resultMap[v.TodoCategory] = m
		}

		m.TodoItemList = append(m.TodoItemList, todo_db.FromTodo(v))
	}

	result = make([]*models.TodoItemGroup, 0)
	for _, v := range resultMap {
		result = append(result, v)
	}

	sort.Sort(models.TodoItemGroupArray(result))

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationAccessLog,
		UserID:        userId,
		ApiName:       "GetTodoListByCategory",
	})

	return result, nil
}
