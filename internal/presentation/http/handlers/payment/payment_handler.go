package payment

import (
	"decard/internal/domain/service"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
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

func (h *PaymentHandler) Transfer(w http.ResponseWriter, r any, profileUUID valueobject.UUID) error {
	req := r.(*TransferRequest)

	err := h.paymentService.Transfer(req.Amount, req.Description, req.From, req.To)

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, nil)
}
