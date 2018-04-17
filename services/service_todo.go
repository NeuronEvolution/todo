package services

import (
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/rand"
	"github.com/NeuronFramework/rest"
)

func (s *TodoService) GetTodoList(ctx *rest.Context, userId string) (result []*models.TodoItem, err error) {
	dbTodoList, err := s.todoDB.Todo.GetQuery().UserId_Equal(userId).QueryList(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbTodoList == nil {
		return nil, nil
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationAccessLog,
		UserID:        userId,
		ApiName:       "GetTodoList",
	})

	return todo_db.FromTodoList(dbTodoList), nil
}

func (s *TodoService) GetTodo(ctx *rest.Context, userId string, todoId string) (todoItem *models.TodoItem, err error) {
	dbTodo, err := s.todoDB.Todo.GetQuery().
		TodoId_Equal(todoId).And().UserId_Equal(userId).
		QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbTodo == nil {
		return nil, errors.NotFound("计划不存在å")
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationAccessLog,
		UserID:        userId,
		ApiName:       "GetTodo",
	})

	return todo_db.FromTodo(dbTodo), nil
}

func (s *TodoService) AddTodo(ctx *rest.Context, userId string, todoItem *models.TodoItem) (todoId string, err error) {
	if todoItem == nil {
		return "", errors.InvalidParam("todoItem不能为空")
	}

	err = todoItem.ValidateParams()
	if err != nil {
		return "", err
	}

	dbTodo := todo_db.ToTodo(todoItem)
	dbTodo.UserId = userId
	dbTodo.TodoId = rand.NextHex(16)
	_, err = s.todoDB.Todo.Insert(ctx, nil, dbTodo)
	if err != nil {
		return "", err
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationAddTodo,
		UserID:        userId,
		TodoItem:      todoItem,
	})

	return dbTodo.TodoId, nil
}

func (s *TodoService) UpdateTodo(ctx *rest.Context, userId string, todoID string, todoItem *models.TodoItem) (err error) {
	if todoItem == nil {
		return errors.InvalidParam("todoItem不能为空")
	}

	err = todoItem.ValidateParams()
	if err != nil {
		return err
	}

	dbTodo, err := s.todoDB.Todo.GetQuery().
		UserId_Equal(userId).And().TodoId_Equal(todoID).
		QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbTodo == nil {
		return errors.NotFound("计划不存在")
	}

	err = s.todoDB.Todo.GetUpdate().
		TodoCategory(todoItem.Category).
		TodoTitle(todoItem.Title).
		TodoDesc(todoItem.Desc).
		TodoStatus(string(todoItem.Status)).
		TodoPriority(todoItem.Priority).
		Update(ctx, nil, dbTodo.Id)
	if err != nil {
		return err
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationUpdateTodo,
		UserID:        userId,
		TodoId:        todoID,
		TodoItem:      todoItem,
	})

	return nil
}

func (s *TodoService) RemoveTodo(ctx *rest.Context, userId string, todoId string) error {
	dbTodo, err := s.todoDB.Todo.GetQuery().
		UserId_Equal(userId).And().TodoId_Equal(todoId).
		QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbTodo == nil {
		return errors.NotFound("计划不存在")
	}

	err = s.todoDB.Todo.Delete(ctx, nil, dbTodo.Id)
	if err != nil {
		return err
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationRemoveTodo,
		UserID:        userId,
		TodoId:        todoId,
	})

	return nil
}
