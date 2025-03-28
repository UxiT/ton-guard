package accounts

import (
	"decard/internal/domain"
	providerentity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"github.com/rs/zerolog"
)

type GetAccountForProfileQuery struct {
	ProfileUUID valueobject.UUID
}

type GetAccountForProfileQueryHandler struct {
	logger             *zerolog.Logger
	accountService     interfaces.AccountService
	customerRepository interfaces.CustomerRepository
	accountRepository  interfaces.AccountRepository
}

func NewGetAccountForProfileQueryHandler(
	logger *zerolog.Logger,
	accountService interfaces.AccountService,
	customerRepository interfaces.CustomerRepository,
	accountRepository interfaces.AccountRepository,
) GetAccountForProfileQueryHandler {
	return GetAccountForProfileQueryHandler{
		logger:             logger,
		accountService:     accountService,
		customerRepository: customerRepository,
		accountRepository:  accountRepository,
	}
}

func (h GetAccountForProfileQueryHandler) Handle(q GetAccountForProfileQuery) (*providerentity.Account, error) {
	const op = "application.query.getAccountForProfile"

	logger := h.logger.With().Str("operation", op).Logger()

	customer, err := h.customerRepository.FindByProfileUUID(q.ProfileUUID)
	if err != nil {
		logger.Error().Err(err).Msg("error getting customer")

		return nil, domain.ErrCustomerNotFound
	}

	systemAccount, err := h.accountRepository.GetByCustomer(customer.UUID)
	if err != nil {
		logger.Error().Err(err).Msg("error getting account")

		return nil, domain.ErrAccountNotFound
	}

	return h.accountService.GetAccount(systemAccount.ExternalUUID)
}
