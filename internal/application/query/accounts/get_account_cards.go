package accounts

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
)

type GetAccountCards struct {
	AccountUUID string
}

type GetAccountCardsHandler struct {
	accountService interfaces.AccountService
}

func NewGetAccountCardsHandler(accountService interfaces.AccountService) GetAccountCardsHandler {
	return GetAccountCardsHandler{
		accountService: accountService,
	}
}

func (h *GetAccountCardsHandler) Handle(q GetAccountCards) ([]providerEntity.Card, error) {
	accountUUID, err := valueobject.ParseUUID(q.AccountUUID)
	if err != nil {
		return nil, err
	}

	return h.accountService.GetAccountCards(accountUUID)
}
