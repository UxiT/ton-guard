package service

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"github.com/google/uuid"
)

type AccountService struct {
	accountService interfaces.AccountService
}

func NewAccountService(accService interfaces.AccountService) *AccountService {
	return &AccountService{
		accountService: accService,
	}
}

func (s *AccountService) GetAccountList() ([]providerEntity.Account, error) {
	return s.accountService.GetAccountsList()
}

func (s *AccountService) GetAccountCards(account uuid.UUID) ([]providerEntity.Card, error) {
	return s.accountService.GetAccountCards(account)
}
