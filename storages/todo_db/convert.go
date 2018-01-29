package todo_db

import "github.com/NeuronEvolution/todo/models"

func FromTodo(p *Todo) (r *models.TodoItem) {
	if p == nil {
		return nil
	}

	r = &models.TodoItem{}
	r.TodoID = p.TodoId
	r.UserID = p.UserId
	r.Title = p.TodoTitle
	r.Desc = p.TodoDesc
	r.Priority = p.TodoPriority
	r.Status = p.TodoStatus

	return r
}

func FromTodoList(p []*Todo) (r []*models.TodoItem) {
	if p == nil {
		return nil
	}

	r = make([]*models.TodoItem, len(p))
	for i, v := range p {
		r[i] = FromTodo(v)
	}

	return r
}

func ToTodo(p *models.TodoItem) (r *Todo) {
	if p == nil {
		return nil
	}

	r = &Todo{}
	r.TodoId = p.TodoID
	r.UserId = p.UserID
	r.TodoTitle = p.Title
	r.TodoDesc = p.Desc
	r.TodoPriority = int32(p.Priority)
	r.TodoStatus = int32(p.Status)

	return r
}
