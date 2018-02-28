package handler

import "github.com/NeuronEvolution/todo/models"
import (
	"fmt"
	api "github.com/NeuronEvolution/todo/api-private/gen/models"
)

func fromTodoStatus(p models.TodoStatus) (r api.TodoStatus) {
	switch p {
	case models.TodoStatusOngoing:
		return api.TodoStatusOngoing
	case models.TodoStatusCompleted:
		return api.TodoStatusCompleted
	case models.TodoStatusDiscard:
		return api.TodoStatusDiscard
	default:
		panic(fmt.Errorf("fromTodoStatus unknown %v", p))
	}
}

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
	r.Status = fromTodoStatus(p.Status)
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

func toTodoStatus(p api.TodoStatus) (r models.TodoStatus) {
	switch p {
	case api.TodoStatusOngoing:
		return models.TodoStatusOngoing
	case api.TodoStatusCompleted:
		return models.TodoStatusCompleted
	case api.TodoStatusDiscard:
		return models.TodoStatusDiscard
	default:
		panic(fmt.Errorf("toTodoStatus unknown %v", p))
	}
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
	r.Status = toTodoStatus(p.Status)
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

func fromTodoVisibility(p models.TodoVisibility)(r api.TodoVisibility) {
	switch p {
	case models.TodoVisibilityPrivate:
		return api.TodoVisibilityPrivate
	case models.TodoVisibilityFriend:
		return api.TodoVisibilityFriend
	case models.TodoVisibilityPublic:
		return api.TodoVisibilityPublic
	default:
		panic("unknown TodoVisibility:" + p)
	}
}

func toTodoVisibility(p api.TodoVisibility)(r models.TodoVisibility) {
	switch p {
	case api.TodoVisibilityPrivate:
		return models.TodoVisibilityPrivate
	case api.TodoVisibilityFriend:
		return models.TodoVisibilityFriend
	case api.TodoVisibilityPublic:
		return models.TodoVisibilityPublic
	default:
		panic("unknown TodoVisibility:" + p)
	}
}

func fromUserProfile(p *models.UserProfile) (r *api.UserProfile) {
	if p == nil {
		return nil
	}

	r = &api.UserProfile{}
	r.UserID = p.UserID
	r.UserName = p.UserName
	r.TodoVisibility = fromTodoVisibility(p.TodoVisibility)

	return r
}

func toUserProfile(p *api.UserProfile) (r *models.UserProfile) {
	if p == nil {
		return nil
	}

	r = &models.UserProfile{}
	r.UserID = p.UserID
	r.UserName = p.UserName
	r.TodoVisibility = toTodoVisibility( p.TodoVisibility)

	return r
}

func fromFriendInfo(p *models.FriendInfo) (r *api.FriendInfo) {
	if p == nil {
		return nil
	}

	r = &api.FriendInfo{}
	r.UserID = p.UserID
	r.UserName = p.UserName
	r.TodoVisibility = fromTodoVisibility(p.TodoVisibility)
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
