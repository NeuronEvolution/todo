package services

import (
	"context"
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"github.com/NeuronFramework/errors"
	"go.uber.org/zap"
	"strconv"
)

func (s *TodoService) GetFriendsList(ctx context.Context, userID string, query *models.FriendsQuery) (result []*models.FriendInfo, nextPageToken string, err error) {
	limitStart := int64(0)
	if query.PageToken != "" {
		limitStart, err = strconv.ParseInt(query.PageToken, 10, 64)
		if err != nil {
			return nil, "", errors.InvalidParam("invalid PageToken")
		}
	}
	limitCount := query.PageSize
	if limitCount == 0 {
		limitCount = models.DefaultPageSize
	}

	rows, err := s.todoDB.Todo.GetQuery().
		GroupBy(todo_db.TODO_FIELD_USER_ID).
		OrderByGroupCount(false).
		Limit(int64(limitStart), limitCount).
		QueryGroupBy(ctx, nil)
	if err != nil {
		return nil, "", err
	}

	result = make([]*models.FriendInfo, 0)
	for rows.Next() {
		e := &models.FriendInfo{}
		err = rows.Scan(&e.UserID, &e.TodoCount)
		if err != nil {
			return nil, "", err
		}

		dbUserProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(e.UserID).QueryOne(ctx, nil)
		if err != nil {
			s.logger.Warn("user profile not exist", zap.String("UserID", e.UserID))
			continue
		}

		e.UserName = dbUserProfile.UserName
		if dbUserProfile.TodoPublicVisible == 0 {
			e.TodoPublicVisible = false
		} else {
			e.TodoPublicVisible = true
		}

		result = append(result, e)
	}
	if rows.Err() != nil {
		return nil, "", rows.Err()
	}

	return result, strconv.FormatInt(limitStart+limitCount, 10), nil
}

func (s *TodoService) GetFriend(ctx context.Context, userID string, friendID string) (friend *models.FriendInfo, err error) {
	dbFriendProfile, err := s.todoDB.UserProfile.GetQuery().UserId_Equal(friendID).QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbFriendProfile == nil {
		return nil, errors.NotFound("用户不存在")
	}

	todoCount, err := s.todoDB.Todo.GetQuery().UserId_Equal(friendID).QueryCount(ctx, nil)
	if err != nil {
		return nil, err
	}

	friend = &models.FriendInfo{}
	friend.UserID = dbFriendProfile.UserId
	friend.UserName = dbFriendProfile.UserName
	if dbFriendProfile.TodoPublicVisible == 0 {
		friend.TodoPublicVisible = false
	} else {
		friend.TodoPublicVisible = true
	}
	friend.TodoCount = todoCount

	return friend, nil
}
