package topup

import (
	"context"
	"decard/internal/domain/interfaces"
	"github.com/rs/zerolog"
)

type ApproveTopUpCommand struct {
	TopUpUUID string
}

type ApproveTopUpCommandHandler struct {
	logger     *zerolog.Logger
	repository interfaces.TopUpRepository
}

type ApproveTopUpResponse struct{}

func NewApproveTopUpHandler(
	logger *zerolog.Logger,
	repository interfaces.TopUpRepository,
) *ApproveTopUpCommandHandler {
	return &ApproveTopUpCommandHandler{
		logger:     logger,
		repository: repository,
	}
}

func (r ApproveTopUpCommandHandler) Handle(ctx context.Context, command ApproveTopUpCommand) (ApproveTopUpResponse, error) {
	return ApproveTopUpResponse{}, nil
}
