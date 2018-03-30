package services

import (
	"encoding/json"
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
)

func toOperation(p *models.Operation) (r *todo_db.Operation) {
	if p == nil {
		return nil
	}

	r = &todo_db.Operation{}
	r.OperationType = string(p.OperationType)
	r.UserAgent = p.UserAgent
	r.UserId = p.UserID
	r.ApiName = p.ApiName
	r.FriendId = p.FriendID
	r.TodoId = p.TodoId
	todoItemJson, _ := json.Marshal(p.TodoItem)
	r.TodoItem = string(todoItemJson)
	userProfileJson, _ := json.Marshal(p.UserProfile)
	r.UserProfile = string(userProfileJson)

	return r
}
