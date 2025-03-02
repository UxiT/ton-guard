package service

import (
	"decard/internal/domain/interfaces"
)

type PaymentService struct {
	providerPaymentService interfaces.PaymentService
}

func NewPaymentService(providerPaymentService interfaces.PaymentService) *PaymentService {
	return &PaymentService{
		providerPaymentService: providerPaymentService,
	}
}

func (s *PaymentService) Transfer(amount float64, description, from, to string) error {
	return s.providerPaymentService.CreateAccountTransfer(amount, description, from, to)
}
