package service

import (
	"decard/internal/domain"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"github.com/rs/zerolog"
)

type AccountService struct {
	providerAccountAPI interfaces.AccountService
	accountRepository  interfaces.AccountRepository
	customerRepository interfaces.CustomerRepository
	logger             *zerolog.Logger
}

func NewAccountService(providerAccountAPI interfaces.AccountService) *AccountService {
	return &AccountService{
		providerAccountAPI: providerAccountAPI,
	}
}

func (s AccountService) CreateProviderAccount(profile valueobject.UUID, name string) (*entity.Account, error) {
	customer, err := s.customerRepository.FindByProfileUUID(profile)

	if err != nil {
		s.logger.Error().Err(err).Msg("error finding customer")

		return nil, domain.ErrInvalidProfileType
	}

	account, err := s.accountRepository.GetByCustomer(customer.UUID)

	if account != nil && err == nil {
		return nil, domain.ErrAccountExists
	}

	providerAccount, err := s.providerAccountAPI.CreateAccount(name)
	if err != nil {
		s.logger.Error().Err(err).Msg("error creating account in provider")

		return nil, err
	}

	externalUUID, err := valueobject.ParseUUID(providerAccount.ID)
	if err != nil {
		s.logger.Error().Err(err).Msg("error parsing uuid")

		return nil, domain.ErrProviderAccount
	}

	internalAccount, err := entity.CreateAccount(externalUUID, providerAccount.Balance, providerAccount.CurrencyCode)
	if err != nil {
		s.logger.Error().Err(err).Msg("error creating account via factory")

		return nil, domain.ErrSystemAccount
	}

	if err = s.accountRepository.Create(*internalAccount); err != nil {
		s.logger.Error().Err(err).Msg("error creating account via repository")

		return nil, domain.ErrSystemAccount
	}

	return internalAccount, nil
}
