package services

import (
	"github.com/NeuronEvolution/todo/storages/todo_db"
	"github.com/NeuronFramework/log"
	"go.uber.org/zap"
)

type TodoService struct {
	logger *zap.Logger
	todoDB *todo_db.DB
}

func New() (s *TodoService, err error) {
	s = &TodoService{}
	s.logger = log.TypedLogger(s)
	s.todoDB, err = todo_db.NewDB()
	if err != nil {
		return nil, err
	}

	return s, nil
}
