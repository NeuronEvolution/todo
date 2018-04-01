package models

import (
	"time"
)

type OperationType string

const (
	OperationAccessLog                       = OperationType("ACCESS_LOG")
	OperationUpdateUserProfile               = OperationType("UPDATE_USER_PROFILE")
	OperationUpdateUserProfileTodoVisibility = OperationType("UPDATE_USER_PROFILE_TODO_VISIBILITY")
	OperationUpdateUserProfileUserName       = OperationType("UPDATE_USER_PROFILE_USER_NAME")
	OperationAddTodo                         = OperationType("ADD_TODO")
	OperationUpdateTodo                      = OperationType("UPDATE_TODO")
	OperationRemoveTodo                      = OperationType("REMOVE_TODO")
)

type Operation struct {
	OperationId   string
	OperationTime time.Time
	OperationType OperationType
	UserAgent     string
	UserID        string
	ApiName       string
	FriendID      string
	TodoId        string
	TodoItem      *TodoItem
	UserProfile   *UserProfile
}
