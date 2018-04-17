package services

import (
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/rand"
	"github.com/NeuronFramework/rest"
)

func (s *TodoService) GetUserProfile(ctx *rest.Context, userID string) (userProfile *models.UserProfile, err error) {
	dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(userID).QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbUserProfile == nil {
		dbUserProfile = &todo_db.UserProfile{}
		dbUserProfile.UserId = userID
		dbUserProfile.UserName = "无名氏" + rand.NextNumberFixedLength(8)
		dbUserProfile.TodoVisibility = string(models.TodoVisibilityPublic)
		_, err = s.todoDB.UserProfile.Insert(ctx, nil, dbUserProfile)
		if err != nil {
			return nil, err
		}
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationAccessLog,
		UserID:        userID,
		ApiName:       "GetUserProfile",
	})

	return todo_db.FromUserProfile(dbUserProfile), nil
}

func (s *TodoService) UpdateUserProfileTodoVisibility(ctx *rest.Context, userID string, visibility models.TodoVisibility) (err error) {
	dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(userID).QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbUserProfile == nil {
		return errors.NotFound("UserProfile不存在")
	}

	err = s.todoDB.UserProfile.GetUpdate().
		TodoVisibility(string(visibility)).
		Update(ctx, nil, dbUserProfile.Id)
	if err != nil {
		return err
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationUpdateUserProfileTodoVisibility,
		UserID:        userID,
		UserProfile: &models.UserProfile{
			TodoVisibility: visibility,
		},
	})

	return nil
}

func (s *TodoService) UpdateUserProfileUserName(ctx *rest.Context, userID string, userName string) (err error) {
	dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(userID).QueryOne(ctx, nil)
	if err != nil {
		return err
	}

	if dbUserProfile == nil {
		return errors.NotFound("UserProfile不存在")
	}

	err = s.todoDB.UserProfile.GetUpdate().UserName(userName).Update(ctx, nil, dbUserProfile.Id)
	if err != nil {
		return err
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationUpdateUserProfileUserName,
		UserID:        userID,
		UserProfile: &models.UserProfile{
			UserName: userName,
		},
	})

	return nil
}
