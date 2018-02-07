package handler

import "github.com/NeuronEvolution/todo/models"
import api "github.com/NeuronEvolution/todo/api-private/gen/models"

func fromTodoItem(p *models.TodoItem) (r *api.TodoItem) {
	if p == nil {
		return nil
	}

	r = &api.TodoItem{}
	r.TodoID = p.TodoID
	r.UserID = p.UserID
	r.Category = p.Category
	r.Title = p.Title
	r.Desc = p.Desc
	r.Status = p.Status
	r.Priority = p.Priority

	return r
}

func fromTodoItemList(p []*models.TodoItem) (r []*api.TodoItem) {
	if p == nil {
		return nil
	}

	r = make([]*api.TodoItem, len(p))
	for i, v := range p {
		r[i] = fromTodoItem(v)
	}

	return r
}

func toTodoItem(p *api.TodoItem) (r *models.TodoItem) {
	if p == nil {
		return nil
	}

	r = &models.TodoItem{}
	r.TodoID = p.TodoID
	r.UserID = p.UserID
	r.Category = p.Category
	r.Title = p.Title
	r.Desc = p.Desc
	r.Status = p.Status
	r.Priority = p.Priority

	return r
}

func fromTodoItemGroup(p *models.TodoItemGroup) (r *api.TodoItemGroup) {
	if p == nil {
		return nil
	}

	r = &api.TodoItemGroup{}
	r.Category = p.Category
	r.TodoItemList = fromTodoItemList(p.TodoItemList)

	return r
}

func fromTodoItemGroupList(p []*models.TodoItemGroup) (r []*api.TodoItemGroup) {
	if p == nil {
		return nil
	}

	r = make([]*api.TodoItemGroup, len(p))
	for i, v := range p {
		r[i] = fromTodoItemGroup(v)
	}

	return r
}
