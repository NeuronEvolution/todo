package todo_db

import "github.com/NeuronEvolution/todo/models"

func FromTodo(p *Todo) (r *models.TodoItem) {
	if p == nil {
		return nil
	}

	r = &models.TodoItem{}
	r.TodoID = p.TodoId
	r.UserID = p.UserId
	r.Category = p.TodoCategory
	r.Title = p.TodoTitle
	r.Desc = p.TodoDesc
	r.Status = p.TodoStatus
	r.Priority = p.TodoPriority

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
	r.TodoCategory = p.Category
	r.TodoTitle = p.Title
	r.TodoDesc = p.Desc
	r.TodoStatus = p.Status
	r.TodoPriority = int32(p.Priority)

	return r
}
