package handlers

import (
	"decard/internal/domain/service"
	presentation "decard/internal/presentation/http"
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

func (h *PaymentHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
		From        string  `json:"from_account_id"`
		To          string  `json:"to_account_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presentation.WriteError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.paymentService.Transfer(request.Amount, request.Description, request.From, request.To)

	if err != nil {
		presentation.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	presentation.WriteJSONResponse(w, nil)
}
