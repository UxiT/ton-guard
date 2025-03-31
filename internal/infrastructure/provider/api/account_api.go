package api

import (
	"bytes"
	"context"
	provider_entity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"decard/internal/infrastructure/provider"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
	"net/url"
)

type accountAPI struct {
	logger *zerolog.Logger
	client *provider.Client
}

func NewAccountApi(logger *zerolog.Logger, client *provider.Client) interfaces.AccountService {
	return &accountAPI{
		logger: logger,
		client: client,
	}
}

func (a *accountAPI) GetAccount(account valueobject.UUID) (*provider_entity.Account, error) {
	var result struct {
		Account provider_entity.Account `json:"account"`
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

func (a *accountAPI) GetAccountCards(account valueobject.UUID) ([]provider_entity.Card, error) {
	var result struct {
		Cards []provider_entity.Card `json:"cards"`
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

func (a *accountAPI) GetAccountsList() ([]provider_entity.Account, error) {
	var result struct {
		Accounts []provider_entity.Account `json:"accounts"`
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

func (a *accountAPI) CreateAccount(name string) (*provider_entity.Account, error) {
	var result struct {
		Account provider_entity.Account `json:"account"`
	}

	payload, err := json.Marshal(struct {
		Name string `json:"name"`
	}{
		Name: name,
	})

	endpoint := a.client.BaseURL.JoinPath("accounts")
	ctx := context.TODO()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), bytes.NewBuffer(payload))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err = a.client.SendRequest(request, &result); err != nil {
		a.logger.Error().Err(err).Msg("API call failed")

		return nil, fmt.Errorf("API call failed")
	}

	return &result.Account, nil
}
