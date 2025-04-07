package api

import (
	"bytes"
	"context"
	"crypto/rsa"
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"decard/internal/infrastructure/provider"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type cardApi struct {
	client     *provider.Client
	logger     *zerolog.Logger
	publicKey  string
	privateKey *rsa.PrivateKey
}

func NewCardAPI(client *provider.Client, logger *zerolog.Logger, publicKey string, privateKey *rsa.PrivateKey) interfaces.CardService {
	return &cardApi{
		client:     client,
		logger:     logger,
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

type cardResponse struct {
	Card providerEntity.Card `json:"card"`
}

func (a *cardApi) sendRequest(url, method string, response interface{}, body []byte) error {
	endpoint := a.client.BaseURL.JoinPath(url)

	request, err := http.NewRequestWithContext(context.TODO(), method, endpoint.String(), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	if err = a.client.SendRequest(request, &response); err != nil {
		return err
	}

	return nil
}

func (a *cardApi) GetCard(card valueobject.UUID) (*providerEntity.Card, error) {
	result := new(cardResponse)

	if err := a.sendRequest(fmt.Sprintf("/cards/%s", card.String()), http.MethodGet, result, nil); err != nil {
		return nil, err
	}

	return &result.Card, nil
}

type PublicKeyPayload struct {
	PublicKey string `json:"public_key"`
}

func (a *cardApi) GetCardNumber(card valueobject.UUID) (string, error) {
	var response struct {
		Number string `json:"encrypted_card_number"`
	}

	payload, err := json.Marshal(PublicKeyPayload{
		PublicKey: a.publicKey,
	})

	if err != nil {
		return "s", err
	}

	endpoint := a.client.BaseURL.JoinPath(fmt.Sprintf("cards/%s/encrypted-card-number", card.String()))
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, endpoint.String(), bytes.NewBuffer(payload))
	if err != nil {
		return "s", err
	}

	if err = a.client.SendRequest(request, &response); err != nil {
		return "s", err
	}

	return response.Number, nil
}

func (a *cardApi) BlockCard(card valueobject.UUID) (*providerEntity.Card, error) {
	result := new(cardResponse)

	err := a.sendRequest(fmt.Sprintf("/cards/%s/block", card.String()), http.MethodPatch, result, nil)

	if err != nil {
		return nil, err
	}

	return &result.Card, nil
}

func (a *cardApi) CloseCard(card valueobject.UUID) (*providerEntity.Card, error) {
	result := new(cardResponse)

	err := a.sendRequest(fmt.Sprintf("/cards/%s/close", card.String()), http.MethodPatch, result, nil)

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
