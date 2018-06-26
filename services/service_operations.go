package services

import (
	"github.com/NeuronEvolution/todo/models"
	"github.com/NeuronFramework/rest"
	"go.uber.org/zap"
)

func (s *TodoService) addOperation(ctx *rest.Context, operation *models.Operation) (err error) {
	operation.UserAgent = ctx.UserAgent
	dbOperation := toOperation(operation)
	_, err = s.todoDB.Operation.Insert(ctx, nil, dbOperation)
	if err != nil {
		s.logger.Error("addOperation", zap.Error(err))
	}

	return nil
}
