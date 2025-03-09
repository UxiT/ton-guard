package handlers

import (
	"decard/internal/domain/service"
	"decard/internal/presentation/http/common"
	"encoding/json"
	"net/http"
)

type PaymentHandler struct {
	paymentService service.PaymentService
}

func NewPaymentHandler(paymentService service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

func (h *PaymentHandler) Transfer(w http.ResponseWriter, r *http.Request) error {
	var request struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
		From        string  `json:"from_account_id"`
		To          string  `json:"to_account_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	err := h.paymentService.Transfer(request.Amount, request.Description, request.From, request.To)

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, nil)
}
