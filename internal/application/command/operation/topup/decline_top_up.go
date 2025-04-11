package topup

import (
	"context"
	"decard/internal/domain/interfaces"
	"github.com/rs/zerolog"
)

type DeclineTopUpCommand struct {
	TopUpUUID string
}

type DeclineTopUpCommandHandler struct {
	logger     *zerolog.Logger
	repository interfaces.TopUpRepository
}

type DeclineTopUpResponse struct{}

func NewDeclineTopUpHandler(
	logger *zerolog.Logger,
	repository interfaces.TopUpRepository,
) *DeclineTopUpCommandHandler {
	return &DeclineTopUpCommandHandler{
		logger:     logger,
		repository: repository,
	}
}

func (r DeclineTopUpCommandHandler) Handle(ctx context.Context, command DeclineTopUpCommand) (DeclineTopUpResponse, error) {
	return DeclineTopUpResponse{}, nil
}
