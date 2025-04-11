package topup

import (
	"database/sql"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

type CreateCommand struct {
	Profile valueobject.UUID
	Amount  string
	Network string
}

type CreateResponse struct {
	TopUp valueobject.UUID `json:"uuid"`
}

type CreateCommandHandler struct {
	logger       *zerolog.Logger
	topUpRepo    interfaces.TopUpRepository
	customerRepo interfaces.CustomerRepository
}

func (h CreateCommandHandler) Handle(cmd CreateCommand) (*CreateResponse, error) {
	amount, err := decimal.NewFromString(cmd.Amount)
	if err != nil {
		return nil, err
	}

	customer, err := h.customerRepo.FindByProfileUUID(cmd.Profile)
	if err != nil {
		return nil, err
	}

	currentTopUp, err := h.topUpRepo.GetCustomerCurrentTopUp(customer.UUID)

	if (err != nil && !errors.Is(err, sql.ErrNoRows)) || currentTopUp != nil {
		h.logger.Err(err).Msg("customer already has topUp request")

		return nil, err
	}

	topUp := entity.NewTopUp(customer.UUID, amount, cmd.Network)

	if err = h.topUpRepo.Create(topUp); err != nil {
		return nil, err
	}

	return &CreateResponse{
		TopUp: topUp.UUID,
	}, err
}
