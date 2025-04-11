package operation

import (
	"context"
	"decard/internal/application/command/operation/topup"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
	"github.com/rs/zerolog"
	"net/http"
)

type FinancialOperationHandler struct {
	logger                       *zerolog.Logger
	createTopUpHandler           topup.CreateCommandHandler
	declineTopUpHandler          topup.DeclineTopUpCommandHandler
	approveTopUpHandler          topup.ApproveTopUpCommandHandler
	closeTopUpHandler            topup.CloseTopUpCommandHandler
	cancelTopUpHandler           topup.CancelTopUpCommandHandler
	addTransactionToTopUpHandler topup.AddTransactionToTopUpCommandHandler
}

func NewFinancialOperationHandler(
	logger *zerolog.Logger,
	createTopUpHandler topup.CreateCommandHandler,
	declineTopUpHandler topup.DeclineTopUpCommandHandler,
	approveTopUpHandler topup.ApproveTopUpCommandHandler,
	closeTopUpHandler topup.CloseTopUpCommandHandler,
	cancelTopUpHandler topup.CancelTopUpCommandHandler,
	addTransactionToTopUpHandler topup.AddTransactionToTopUpCommandHandler,
) *FinancialOperationHandler {
	return &FinancialOperationHandler{
		logger:                       logger,
		createTopUpHandler:           createTopUpHandler,
		declineTopUpHandler:          declineTopUpHandler,
		approveTopUpHandler:          approveTopUpHandler,
		closeTopUpHandler:            closeTopUpHandler,
		cancelTopUpHandler:           cancelTopUpHandler,
		addTransactionToTopUpHandler: addTransactionToTopUpHandler,
	}
}

func (h FinancialOperationHandler) CreateTopUp(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	req := r.(TopUpRequest)

	response, err := h.createTopUpHandler.Handle(topup.CreateCommand{
		Profile: profile,
		Amount:  req.Amount,
		Network: req.Network,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusCreated, response)
}

func (h FinancialOperationHandler) Cancel(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	req := r.(TopUpUUIDRequest)

	response, err := h.cancelTopUpHandler.Handle(context.Background(), topup.CancelTopUpCommand{
		TopUpUUID:   req.UUID,
		ProfileUUID: profile,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusCreated, response)
}

func (h FinancialOperationHandler) AddTransactionID(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	req := r.(AddTransactionRequest)

	response, err := h.addTransactionToTopUpHandler.Handle(context.Background(), topup.AddTransactionToTopUpCommand{
		TopUpUUID:     req.TopUp,
		Customer:      profile,
		TransactionID: req.TransactionID,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusCreated, response)
}

func (h FinancialOperationHandler) Decline(w http.ResponseWriter, r any) error {
	req := r.(TopUpUUIDRequest)

	response, err := h.declineTopUpHandler.Handle(context.Background(), topup.DeclineTopUpCommand{
		TopUpUUID: req.UUID,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusCreated, response)
}

func (h FinancialOperationHandler) Approve(w http.ResponseWriter, r any) error {
	req := r.(TopUpUUIDRequest)

	response, err := h.approveTopUpHandler.Handle(context.Background(), topup.ApproveTopUpCommand{
		TopUpUUID: req.UUID,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusCreated, response)
}

func (h FinancialOperationHandler) Close(w http.ResponseWriter, r any) error {
	req := r.(TopUpUUIDRequest)

	response, err := h.closeTopUpHandler.Handle(context.Background(), topup.CloseTopUpCommand{
		TopUpUUID: req.UUID,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusCreated, response)
}
