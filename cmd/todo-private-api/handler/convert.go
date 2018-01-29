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
	r.Title = p.Title
	r.Desc = p.Desc
	r.Priority = p.Priority
	r.Status = p.Status

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
	r.Title = p.Title
	r.Desc = p.Desc
	r.Priority = p.Priority
	r.Status = p.Status

	return r
}
