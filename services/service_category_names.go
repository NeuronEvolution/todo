package services

import (
	"context"
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
)

func (s *TodoService) GetCategoryNameList(ctx context.Context, userId string) (result []*models.CategoryInfo, err error) {
	rows, err := s.todoDB.Todo.GetQuery().UserId_Equal(userId).
		GroupBy(todo_db.TODO_FIELD_TODO_CATEGORY).
		OrderByGroupCount(false).QueryGroupBy(ctx, nil)
	if err != nil {
		return nil, err
	}

	result = make([]*models.CategoryInfo, 0)
	for rows.Next() {
		var categoryInfo = &models.CategoryInfo{}
		err = rows.Scan(&(categoryInfo.Category), &(categoryInfo.TodoCount))
		if err != nil {
			return nil, err
		}

		result = append(result, categoryInfo)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}
