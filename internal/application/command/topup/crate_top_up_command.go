package topup

import (
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

type CreateCommand struct {
	Customer string
	Amount   string
	Network  string
}

type CreateCommandHandler struct {
	logger *zerolog.Logger
	repo   interfaces.TopUpRepository
}

func (h CreateCommandHandler) Handle(cmd CreateCommand) error {
	customerUUID, err := valueobject.ParseUUID(cmd.Customer)
	if err != nil {
		return err
	}

	amount, err := decimal.NewFromString(cmd.Amount)
	if err != nil {
		return err
	}

	topUp := entity.NewTopUp(customerUUID, amount, cmd.Network)

	return h.repo.Create(topUp)
}
