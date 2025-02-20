package service

import (
	"decard/internal/domain/repository"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type BalanceService struct {
	accountRepo repository.AccountRepository
}

func NewBalanceService(repo repository.AccountRepository) *BalanceService {
	return &BalanceService{
		accountRepo: repo,
	}
}

// TODO add sync with external service
func (s BalanceService) GetByCustomer(customer uuid.UUID) (*decimal.Decimal, error) {
	account, err := s.accountRepo.GetByCustomer(customer)
	if err != nil {
		return nil, err
	}

	return (*decimal.Decimal)(&account.Balance), nil
}
