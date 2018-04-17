package services

import (
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/rest"
	"go.uber.org/zap"
	"strconv"
)

func (s *TodoService) GetFriendsList(ctx *rest.Context, userID string, query *models.FriendsQuery) (result []*models.FriendInfo, nextPageToken string, err error) {
	limitStart := int64(0)
	if query.PageToken != "" {
		limitStart, err = strconv.ParseInt(query.PageToken, 10, 64)
		if err != nil {
			return nil, "", errors.InvalidParam("PageToken无效")
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
		e.TodoVisibility = models.TodoVisibility(dbUserProfile.TodoVisibility)

		result = append(result, e)
	}
	if rows.Err() != nil {
		return nil, "", rows.Err()
	}

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationAccessLog,
		UserID:        userID,
		ApiName:       "GetFriendsList",
	})

	return result, strconv.FormatInt(limitStart+limitCount, 10), nil
}

func (s *TodoService) GetFriend(ctx *rest.Context, userID string, friendID string) (friend *models.FriendInfo, err error) {
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
	friend.TodoVisibility = models.TodoVisibility(dbFriendProfile.TodoVisibility)
	friend.TodoCount = todoCount

	s.addOperation(ctx, &models.Operation{
		OperationType: models.OperationAccessLog,
		UserID:        userID,
		FriendID:      friendID,
		ApiName:       "GetFriend",
	})

	return friend, nil
}
