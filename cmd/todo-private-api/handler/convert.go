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

func fromUserProfile(p *models.UserProfile) (r *api.UserProfile) {
	if p == nil {
		return nil
	}

	r = &api.UserProfile{}
	r.UserID = p.UserID
	r.UserName = p.UserName
	r.TodoPublicVisible = p.TodoPublicVisible

	return r
}

func toUserProfile(p *api.UserProfile) (r *models.UserProfile) {
	if p == nil {
		return nil
	}

	r = &models.UserProfile{}
	r.UserID = p.UserID
	r.UserName = p.UserName
	r.TodoPublicVisible = p.TodoPublicVisible

	return r
}

func fromFriendInfo(p *models.FriendInfo) (r *api.FriendInfo) {
	if p == nil {
		return nil
	}

	r = &api.FriendInfo{}
	r.UserID = p.UserID
	r.UserName = p.UserName
	r.TodoPublicVisible = p.TodoPublicVisible
	r.TodoCount = p.TodoCount

	return r
}

func fromFriendInfoList(p []*models.FriendInfo) (r []*api.FriendInfo) {
	if p == nil {
		return nil
	}

	r = make([]*api.FriendInfo, len(p))
	for i, v := range p {
		r[i] = fromFriendInfo(v)
	}

	return r
}
