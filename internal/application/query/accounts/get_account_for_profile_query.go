package accounts

import (
	"decard/internal/domain"
	providerentity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"log/slog"
)

type GetAccountForProfileQuery struct {
	ProfileUUID valueobject.UUID
}

type GetAccountForProfileQueryHandler struct {
	logger             *slog.Logger
	accountService     interfaces.AccountService
	customerRepository interfaces.CustomerRepository
	accountRepository  interfaces.AccountRepository
}

func NewGetAccountForProfileQueryHandler(
	logger *slog.Logger,
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

	logger := h.logger.With(slog.String("operation", op))

	customer, err := h.customerRepository.FindByProfileUUID(q.ProfileUUID)
	if err != nil {
		logger.Error("error getting customer", slog.String("error", err.Error()))
		return nil, domain.ErrCustomerNotFound
	}

	systemAccount, err := h.accountRepository.GetByCustomer(customer.UUID)
	if err != nil {
		logger.Error("error getting account", slog.String("error", err.Error()))
		return nil, domain.ErrAccountNotFound
	}

	return h.accountService.GetAccount(systemAccount.ExternalUUID)
}
