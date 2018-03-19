package services

import (
	"context"
	"github.com/NeuronEvolution/todo/storages/todo_db"
)

func (s *TodoService)GetCategoryNameList(ctx context.Context,userId string)(result [] string,err error) {
	rows, err := s.todoDB.Todo.GetQuery().UserId_Equal(userId).
		GroupBy(todo_db.TODO_FIELD_TODO_CATEGORY).
		OrderByGroupCount(false).QueryGroupBy(ctx, nil)
	if err != nil {
		return nil, err
	}

	result = make([]string, 0)
	for rows.Next() {
		var categoryName string
		var count int64
		err = rows.Scan(&categoryName,&count)
		if err != nil {
			return nil, err
		}

		result = append(result, categoryName)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result,nil
}
