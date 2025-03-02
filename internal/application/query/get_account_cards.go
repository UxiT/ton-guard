package query

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/service"
	"github.com/google/uuid"
)

type GetAccountCards struct {
	AccountUUID string
}

type GetAccountCardsHandler struct {
	accountService *service.AccountService
}

func NewGetAccountCardsHandler(accountService *service.AccountService) *GetAccountCardsHandler {
	return &GetAccountCardsHandler{
		accountService: accountService,
	}
}

func (h *GetAccountCardsHandler) Handle(q GetAccountCards) ([]providerEntity.Card, error) {
	accountUUID, err := uuid.Parse(q.AccountUUID)
	if err != nil {
		return nil, err
	}

	return h.accountService.GetAccountCards(accountUUID)
}
