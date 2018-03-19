package services

import (
	"context"
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"github.com/NeuronFramework/errors"
)

func (s *TodoService) GetUserProfile(ctx context.Context, userID string) (userProfile *models.UserProfile, err error) {
	dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(userID).QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbUserProfile == nil {
		dbUserProfile = &todo_db.UserProfile{}
		dbUserProfile.UserId = userID
		dbUserProfile.UserName = "无名氏"
		dbUserProfile.TodoVisibility = string(models.TodoVisibilityPublic)
		_, err = s.todoDB.UserProfile.Insert(ctx, nil, dbUserProfile)
		if err != nil {
			return nil, err
		}
	}

	return todo_db.FromUserProfile(dbUserProfile), nil
}

func (s *TodoService) UpdateUserProfile(ctx context.Context, userID string, userProfile *models.UserProfile) (err error) {
	dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(userID).QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbUserProfile == nil {
		dbUserProfile = todo_db.ToUserProfile(userProfile)
		dbUserProfile.UserId = userID
		_, err = s.todoDB.UserProfile.Insert(ctx, nil, dbUserProfile)
		if err != nil {
			return err
		}
	} else {
		dbUserProfile = todo_db.ToUserProfile(userProfile)
		dbUserProfile.UserId = userID
		err = s.todoDB.UserProfile.Update(ctx, nil, dbUserProfile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TodoService) UpdateUserProfileTodoVisibility(ctx context.Context, userID string, visibility models.TodoVisibility) (err error) {
	dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(userID).QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbUserProfile == nil {
		return errors.NotFound("UserProfile不存在")
	}

	dbUserProfile.TodoVisibility = string(visibility)
	err = s.todoDB.UserProfile.Update(ctx, nil, dbUserProfile)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) UpdateUserProfileUserName(ctx context.Context, userID string, userName string) (err error) {
	dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(userID).QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbUserProfile == nil {
		return errors.NotFound("UserProfile不存在")
	}

	dbUserProfile.UserName = userName
	err = s.todoDB.UserProfile.Update(ctx, nil, dbUserProfile)
	if err != nil {
		return err
	}

	return nil
}
