package topup

import (
	"context"
	"decard/internal/domain/interfaces"
	"github.com/rs/zerolog"
)

type CloseTopUpCommand struct {
	TopUpUUID string
}

type CloseTopUpCommandHandler struct {
	logger     *zerolog.Logger
	repository interfaces.TopUpRepository
}

type CloseTopUpResponse struct{}

func NewCloseTopUpHandler(
	logger *zerolog.Logger,
	repository interfaces.TopUpRepository,
) *CloseTopUpCommandHandler {
	return &CloseTopUpCommandHandler{
		logger:     logger,
		repository: repository,
	}
}

func (r CloseTopUpCommandHandler) Handle(ctx context.Context, command CloseTopUpCommand) (CloseTopUpResponse, error) {
	return CloseTopUpResponse{}, nil
}
