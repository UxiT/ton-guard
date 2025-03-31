package account

import (
	"context"
	"decard/internal/domain/entity"
	"decard/internal/domain/service"
	domaintype "decard/internal/domain/type"
	"decard/internal/domain/valueobject"
)

type CreateAccountCommand struct {
	Profile     valueobject.UUID
	AccountName string
}

type CreateAccountCommandResponse struct {
	UUID     valueobject.UUID         `json:"uuid"`
	Currency domaintype.Currency      `json:"currency"`
	Status   domaintype.AccountStatus `json:"status"`
	Balance  entity.Balance           `json:"balance"`
}

type CreateAccountCommandHandler struct {
	accountService service.AccountService
}

func (h CreateAccountCommandHandler) Handle(ctx context.Context, cmd CreateAccountCommand) (CreateAccountCommandResponse, error) {
	account, err := h.accountService.CreateProviderAccount(cmd.Profile, cmd.AccountName)

	if err != nil {
		return CreateAccountCommandResponse{}, err
	}

	return CreateAccountCommandResponse{
		UUID:     account.UUID,
		Currency: account.Currency,
		Status:   account.Status,
		Balance:  account.Balance,
	}, nil
}
