package handlers

import (
	"decard/internal/domain/service"
	"net/http"
)

type BalanceHandler struct {
	balanceService service.BalanceService
}

func NewBalanceService(balanceService *service.BalanceService) *BalanceHandler {
	return &BalanceHandler{
		balanceService: *balanceService,
	}
}

func (h *BalanceHandler) GetByCustomer(w http.ResponseWriter, r *http.Request) {

}
