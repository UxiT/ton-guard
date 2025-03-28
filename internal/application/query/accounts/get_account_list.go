package accounts

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"github.com/rs/zerolog"
)

type GetAccountListQueryHandler struct {
	logger         *zerolog.Logger
	accountService interfaces.AccountService
}

func NewGetAccountListQueryHandler(
	logger *zerolog.Logger,
	accountService interfaces.AccountService,
) GetAccountListQueryHandler {
	return GetAccountListQueryHandler{
		logger:         logger,
		accountService: accountService,
	}
}

func (h GetAccountListQueryHandler) Handle() ([]providerEntity.Account, error) {
	const op = "application.query.GetAccountList"

	logger := h.logger.With().Str("operation", op).Logger()

	account, err := h.accountService.GetAccountsList()

	if err != nil {
		logger.Error().Err(err).Msg("error getting account list")

		return nil, err
	}

	return account, nil
}
