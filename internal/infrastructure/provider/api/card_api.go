package api

import (
	"context"
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"decard/internal/infrastructure/provider"
	"fmt"
	"net/http"
	"time"
)

type cardApi struct {
	client *provider.Client
}

func NewCardAPI(client *provider.Client) interfaces.CardService {
	return &cardApi{
		client: client,
	}
}

type cardResponse struct {
	Card providerEntity.Card `json:"card"`
}

func (a *cardApi) sendRequest(url, method string, response interface{}) error {
	endpoint := a.client.BaseURL.JoinPath(url)

	request, err := http.NewRequestWithContext(context.TODO(), method, endpoint.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	err = a.client.SendRequest(request, &response)

	if err != nil {
		return err
	}

	return nil
}

func (a *cardApi) GetCard(card valueobject.UUID) (*providerEntity.Card, error) {
	result := new(cardResponse)

	err := a.sendRequest(fmt.Sprintf("/cards/%s", card.String()), http.MethodGet, result)

	if err != nil {
		return nil, err
	}

	return &result.Card, nil
}

func (a *cardApi) BlockCard(card valueobject.UUID) (*providerEntity.Card, error) {
	result := new(cardResponse)

	err := a.sendRequest(fmt.Sprintf("/cards/%s/block", card.String()), http.MethodPatch, result)

	if err != nil {
		return nil, err
	}

	return &result.Card, nil
}

func (a *cardApi) CloseCard(card valueobject.UUID) (*providerEntity.Card, error) {
	result := new(cardResponse)

	err := a.sendRequest(fmt.Sprintf("/cards/%s/close", card.String()), http.MethodPatch, result)

	if err != nil {
		return nil, err
	}

	return &result.Card, nil
}

type CreateCardRequest struct {
	ThreeDSecureSettings       providerEntity.ThreeDSecureSettings `json:"3d_secure_settings"`
	AccountID                  string                              `json:"account_id"`
	DeliveryAddress            providerEntity.DeliveryAddress      `json:"delivery_address"`
	EmbossingCompanyName       string                              `json:"embossing_company_name"`
	EmbossingName              string                              `json:"embossing_name"`
	ExpirationDate             time.Time                           `json:"expiration_date"`
	IsDisposable               bool                                `json:"is_disposable"`
	Limits                     providerEntity.CardLimits           `json:"limits"`
	Name                       string                              `json:"name"`
	PersonalizationProductCode string                              `json:"personalization_product_code"`
	Security                   providerEntity.Security             `json:"security"`
	Type                       string                              `json:"type"`
}

func (a *cardApi) CreateCard() (*providerEntity.Card, error) {
	var result cardResponse

	endpoint := a.client.BaseURL.JoinPath(fmt.Sprintf("/cards"))

	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	err = a.client.SendRequest(request, &result)

	if err != nil {
		return nil, err
	}

	return &result.Card, nil
}
