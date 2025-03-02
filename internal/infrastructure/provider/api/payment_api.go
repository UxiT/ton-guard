package api

import (
	"bytes"
	"context"
	"decard/internal/domain/interfaces"
	"decard/internal/infrastructure/provider"
	"encoding/json"
	"fmt"
	"net/http"
)

type paymentApi struct {
	client *provider.Client
}

func NewPaymentApi(client *provider.Client) interfaces.PaymentService {
	return &paymentApi{
		client: client,
	}
}

type CreatePaymentRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	From        string  `json:"from_account_id"`
	To          string  `json:"to_account_id"`
}

func (a *paymentApi) CreateAccountTransfer(amount float64, description, from, to string) error {
	payload, err := json.Marshal(CreatePaymentRequest{
		Amount:      amount,
		Description: description,
		From:        from,
		To:          to,
	})

	if err != nil {
		return err
	}

	endpoint := a.client.BaseURL.JoinPath("payments/account-transfer")
	ctx := context.TODO()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), bytes.NewBuffer(payload))

	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	return a.client.SendRequest(request, nil)
}
