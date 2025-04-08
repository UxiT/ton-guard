package provider

import (
	"crypto/rsa"
	"decard/config"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type authTransport struct {
	http.RoundTripper
	PrivateKey *rsa.PrivateKey
	ApiKey     string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := generateAuthJWT(t.ApiKey, t.PrivateKey)
	if err != nil {
		log.Fatalf("couldn't create the JWT: %s", err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	// Call the next RoundTripper (the actual HTTP request)
	return t.RoundTripper.RoundTrip(req)
}

func generateAuthJWT(apiKey string, privateKey *rsa.PrivateKey) (string, error) {
	ts := time.Now().Unix()

	return jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"api_key": apiKey,
		"ts":      ts,
	}).SignedString(privateKey)
}

type Client struct {
	HttpClient *http.Client
	BaseURL    *url.URL

	ApiKey     string
	PrivateKey any

	logger *zerolog.Logger
}

func NewClient(cfg config.Config, logger *zerolog.Logger) *Client {
	httpClient := &http.Client{
		Transport: &authTransport{
			RoundTripper: http.DefaultTransport,
			ApiKey:       cfg.ProviderApiKey,
			PrivateKey:   cfg.PrivateKey,
		},
	}

	return &Client{
		HttpClient: httpClient,
		BaseURL:    cfg.ProviderBaseApiURL,
		ApiKey:     cfg.ProviderApiKey,
		PrivateKey: cfg.PrivateKey,
		logger:     logger,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (c *Client) SendRequest(
	request *http.Request,
	result interface{},
) error {
	resp, err := c.HttpClient.Do(request)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		c.logger.Error().Str("response", string(body)).Msg("provider error")

		var errorResponse ErrorResponse

		err = json.Unmarshal(body, &errorResponse)

		if err != nil {
			return fmt.Errorf("error unmarshalling response body: %w", err)
		}

		return fmt.Errorf("%d: %s", resp.StatusCode, errorResponse.Message)
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return nil
}
