package api

import (
	"context"
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"decard/internal/infrastructure/provider"
	"fmt"
	"net/http"
	"net/url"
)

type accountAPI struct {
	client *provider.Client
}

func NewAccountApi(client *provider.Client) interfaces.AccountService {
	return &accountAPI{
		client: client,
	}
}

func (a *accountAPI) GetAccount(account valueobject.UUID) (*providerEntity.Account, error) {
	var result struct {
		Account providerEntity.Account `json:"account"`
	}

	endpoint := a.client.BaseURL.JoinPath(fmt.Sprintf("/accounts/%s", account.String()))

	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	err = a.client.SendRequest(request, &result)

	if err != nil {
		return nil, err
	}

	return &result.Account, nil
}

func (a *accountAPI) GetAccountCards(account valueobject.UUID) ([]providerEntity.Card, error) {
	var result struct {
		Cards []providerEntity.Card `json:"cards"`
	}

	query := url.Values{
		"from_record":   {fmt.Sprintf("%d", 0)},
		"records_count": {fmt.Sprintf("%d", 10)},
	}

	endpoint := a.client.BaseURL.JoinPath(fmt.Sprintf("accounts/%s/cards", account.String()))
	endpoint.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	err = a.client.SendRequest(request, &result)

	if err != nil {
		return nil, err
	}

	return result.Cards, nil
}

func (a *accountAPI) GetAccountsList() ([]providerEntity.Account, error) {
	var result struct {
		Accounts []providerEntity.Account `json:"accounts"`
	}

	query := url.Values{
		"from_record":   {fmt.Sprintf("%d", 0)},
		"records_count": {fmt.Sprintf("%d", 10)},
	}

	endpoint := a.client.BaseURL.JoinPath("/accounts")
	endpoint.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	err = a.client.SendRequest(request, &result)

	if err != nil {
		return nil, err
	}

	return result.Accounts, nil
}
