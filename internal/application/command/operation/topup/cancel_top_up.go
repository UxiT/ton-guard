package topup

import (
	"context"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"github.com/rs/zerolog"
)

type CancelTopUpCommand struct {
	TopUpUUID   string
	ProfileUUID valueobject.UUID
}

type CancelTopUpCommandHandler struct {
	logger     *zerolog.Logger
	repository interfaces.TopUpRepository
}

type CancelTopUpResponse struct{}

func NewCancelTopUpHandler(
	logger *zerolog.Logger,
	repository interfaces.TopUpRepository,
) *CancelTopUpCommandHandler {
	return &CancelTopUpCommandHandler{
		logger:     logger,
		repository: repository,
	}
}

func (r CancelTopUpCommandHandler) Handle(ctx context.Context, command CancelTopUpCommand) (CancelTopUpResponse, error) {
	return CancelTopUpResponse{}, nil
}
