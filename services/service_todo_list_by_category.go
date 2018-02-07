package services

import (
	"context"
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"sort"
)

func (s *TodoService) GetTodoListByCategory(ctx context.Context, userId string) (result []*models.TodoItemGroup, err error) {
	dbTodoList, err := s.todoDB.Todo.GetQuery().UserId_Equal(userId).QueryList(ctx, nil)
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

	return result, nil
}
