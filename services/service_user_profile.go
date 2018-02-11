package services

import (
	"context"
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
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
		dbUserProfile.TodoPublicVisible = 1
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
