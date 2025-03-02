package service

import (
	"decard/internal/domain/interfaces"
)

type CardService struct {
	cardRepository interfaces.CardRepository
}

func NewCardService(repo interfaces.CardRepository) *CardService {
	return &CardService{
		cardRepository: repo,
	}
}

//func (s CardService) GetCardsByCustomer(customer uuid.UUID) ([]entity.Card, error) {
//	return s.cardRepository.GetByCustomer(customer)
//}
