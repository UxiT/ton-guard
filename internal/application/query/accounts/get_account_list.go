package accounts

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"log/slog"
)

type GetAccountListQueryHandler struct {
	logger         *slog.Logger
	accountService interfaces.AccountService
}

func NewGetAccountListQueryHandler(
	logger *slog.Logger,
	accountService interfaces.AccountService,
) GetAccountListQueryHandler {
	return GetAccountListQueryHandler{
		logger:         logger,
		accountService: accountService,
	}
}

func (h GetAccountListQueryHandler) Handle() ([]providerEntity.Account, error) {
	const op = "application.query.GetAccountList"

	logger := h.logger.With(slog.String("op", op))

	account, err := h.accountService.GetAccountsList()

	if err != nil {
		logger.Error("error getting account list", "error", err.Error())

		return nil, err
	}

	return account, nil
}
