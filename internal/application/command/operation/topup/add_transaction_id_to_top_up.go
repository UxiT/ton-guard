package topup

import (
	"context"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"fmt"
	"github.com/rs/zerolog"
)

type AddTransactionToTopUpCommand struct {
	TopUpUUID     string
	Profile       valueobject.UUID
	TransactionID string
}

type AddTransactionToTopUpCommandHandler struct {
	logger             *zerolog.Logger
	topUpRepository    interfaces.TopUpRepository
	customerRepository interfaces.CustomerRepository
}

type AddTransactionToTopUpResponse struct{}

func NewAddTransactionToTopUpHandler(
	logger *zerolog.Logger,
	topUpRepository interfaces.TopUpRepository,
	customerRepository interfaces.CustomerRepository,
) *AddTransactionToTopUpCommandHandler {
	return &AddTransactionToTopUpCommandHandler{
		logger:             logger,
		topUpRepository:    topUpRepository,
		customerRepository: customerRepository,
	}
}

func (h AddTransactionToTopUpCommandHandler) Handle(ctx context.Context, cmd AddTransactionToTopUpCommand) (AddTransactionToTopUpResponse, error) {
	customer, err := h.customerRepository.FindByProfileUUID(cmd.Profile)
	if err != nil {
		return AddTransactionToTopUpResponse{}, err
	}

	topUpUUID, err := valueobject.ParseUUID(cmd.TopUpUUID)
	if err != nil {
		return AddTransactionToTopUpResponse{}, err
	}

	topUp, err := h.topUpRepository.GetCustomerCurrentTopUp(topUpUUID)
	if err != nil {
		return AddTransactionToTopUpResponse{}, err
	}

	if topUp.Customer.String() != customer.UUID.String() {
		return AddTransactionToTopUpResponse{}, fmt.Errorf("customer has no access")
	}

	err = h.topUpRepository.SetStatus(topUpUUID, entity.Validating)

	return AddTransactionToTopUpResponse{}, nil
}
