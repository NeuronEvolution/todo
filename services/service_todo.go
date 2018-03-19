package services

import (
	"context"
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/rand"
)

func (s *TodoService) GetTodoList(ctx context.Context, userId string) (result []*models.TodoItem, err error) {
	dbTodoList, err := s.todoDB.Todo.GetQuery().UserId_Equal(userId).QueryList(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbTodoList == nil {
		return nil, nil
	}

	return todo_db.FromTodoList(dbTodoList), nil
}

func (s *TodoService) GetTodo(ctx context.Context, userId string, todoId string) (todoItem *models.TodoItem, err error) {
	dbTodo, err := s.todoDB.Todo.GetQuery().
		TodoId_Equal(todoId).And().UserId_Equal(userId).
		QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbTodo == nil {
		return nil, errors.NotFound("todo not exists")
	}

	return todo_db.FromTodo(dbTodo), nil
}

func (s *TodoService) AddTodo(ctx context.Context, userId string, todoItem *models.TodoItem) (todoId string, err error) {
	dbTodo := todo_db.ToTodo(todoItem)
	dbTodo.UserId = userId
	dbTodo.TodoId = rand.NextHex(16)
	_, err = s.todoDB.Todo.Insert(ctx, nil, dbTodo)
	if err != nil {
		return "", err
	}

	return dbTodo.TodoId, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, userId string, todoID string, todoItem *models.TodoItem) error {
	if todoItem.Category == "" {
		return errors.InvalidParam("分类不能为空")
	}

	dbTodo, err := s.todoDB.Todo.GetQuery().
		UserId_Equal(userId).And().TodoId_Equal(todoID).
		QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbTodo == nil {
		return errors.NotFound("todo not exists")
	}

	dbTodo.TodoCategory = todoItem.Category
	dbTodo.TodoTitle = todoItem.Title
	dbTodo.TodoDesc = todoItem.Desc
	dbTodo.TodoStatus = string(todoItem.Status)
	dbTodo.TodoPriority = todoItem.Priority

	err = s.todoDB.Todo.Update(ctx, nil, dbTodo)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) RemoveTodo(ctx context.Context, userId string, todoId string) error {
	dbTodo, err := s.todoDB.Todo.GetQuery().
		UserId_Equal(userId).And().TodoId_Equal(todoId).
		QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbTodo == nil {
		return errors.NotFound("todo not exists")
	}

	err = s.todoDB.Todo.Delete(ctx, nil, dbTodo.Id)
	if err != nil {
		return err
	}

	return nil
}
