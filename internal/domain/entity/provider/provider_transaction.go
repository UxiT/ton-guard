package provider_entity

import "time"

type Transaction struct {
	AccountAmount           float64   `json:"account_amount"`
	AccountCurrencyCode     string    `json:"account_currency_code"`
	AcquirerBusinessID      string    `json:"acquirer_business_id"`
	AcquirerReferenceNumber string    `json:"acquirer_reference_number"`
	AuthorizationCode       string    `json:"authorization_code"`
	AuthorizationID         string    `json:"authorization_id"`
	CardID                  string    `json:"card_id"`
	CentralProcessingDate   time.Time `json:"central_processing_date"`
	CreatedAt               time.Time `json:"created_at"`
	ExchangeRate            float64   `json:"exchange_rate"`
	Group                   string    `json:"group"`
	HasPaymentDocumentFiles bool      `json:"has_payment_document_files"`
	HasPaymentNotes         bool      `json:"has_payment_notes"`
	ID                      string    `json:"id"`
	IsFailed                bool      `json:"is_failed"`
	MarkedForDisputeAt      time.Time `json:"marked_for_dispute_at"`
	MarkedForDisputeBy      string    `json:"marked_for_dispute_by"`
	MerchantCategoryCode    string    `json:"merchant_category_code"`
	MerchantCity            string    `json:"merchant_city"`
	MerchantCountryCode     string    `json:"merchant_country_code"`
	MerchantID              string    `json:"merchant_id"`
	MerchantName            string    `json:"merchant_name"`
	MerchantPostalCode      string    `json:"merchant_postal_code"`
	ProcessedAt             time.Time `json:"processed_at"`
	PurchaseDate            time.Time `json:"purchase_date"`
	TerminalID              string    `json:"terminal_id"`
	TotalAmount             float64   `json:"total_amount"`
	TransactionAmount       float64   `json:"transaction_amount"`
	TransactionCode         string    `json:"transaction_code"`
	TransactionCurrencyCode string    `json:"transaction_currency_code"`
	TransactionIdentifier   string    `json:"transaction_identifier"`
}
