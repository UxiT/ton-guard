package api

import (
	"context"
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/infrastructure/provider"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

type transactionApi struct {
	client *provider.Client
}

func NewTransactionApi(client *provider.Client) interfaces.TransactionService {
	return &transactionApi{
		client: client,
	}
}

func (a *transactionApi) GetCardTransactions(card uuid.UUID) (*[]providerEntity.Transaction, error) {
	var result struct {
		Transactions []providerEntity.Transaction `json:"transactions"`
	}

	query := url.Values{
		"from_record":   {fmt.Sprintf("%d", 0)},
		"records_count": {fmt.Sprintf("%d", 10)},
	}

	endpoint := a.client.BaseURL.JoinPath(fmt.Sprintf("cards/%s/transactions", card))
	endpoint.RawQuery = query.Encode()
	ctx := context.TODO()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	err = a.client.SendRequest(request, &result)
	if err != nil {
		return nil, err
	}

	return &result.Transactions, nil
}
