package service

import (
	"decard/internal/domain/entity"
	"decard/internal/domain/repository"

	"github.com/google/uuid"
)

type CardService struct {
	cardRepository repository.CardRepository
}

func NewCardService(repo repository.CardRepository) *CardService {
	return &CardService{
		cardRepository: repo,
	}
}

func (s CardService) GetByCustomerCards(customer uuid.UUID) ([]entity.Card, error) {
	return s.cardRepository.GetByCustomer(customer)
}
