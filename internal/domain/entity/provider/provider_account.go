package provider_entity

import (
	"time"
)

type AccountLimits struct {
	DailyPurchase              float64 `json:"daily_purchase"`
	DailyWithdrawal            float64 `json:"daily_withdrawal"`
	DailyInternetPurchase      float64 `json:"daily_internet_purchase"`
	DailyContactlessPurchase   float64 `json:"daily_contactless_purchase"`
	WeeklyPurchase             float64 `json:"weekly_purchase"`
	WeeklyWithdrawal           float64 `json:"weekly_withdrawal"`
	WeeklyInternetPurchase     float64 `json:"weekly_internet_purchase"`
	WeeklyContactlessPurchase  float64 `json:"weekly_contactless_purchase"`
	MonthlyPurchase            float64 `json:"monthly_purchase"`
	MonthlyWithdrawal          float64 `json:"monthly_withdrawal"`
	MonthlyInternetPurchase    float64 `json:"monthly_internet_purchase"`
	MonthlyContactlessPurchase float64 `json:"monthly_contactless_purchase"`
}

type TopUpDetails struct {
	IBAN               string `json:"iban"`
	SwiftCode          string `json:"swift_code"`
	ReceiverName       string `json:"receiver_name"`
	PaymentDetails     string `json:"payment_details"`
	BankName           string `json:"bank_name"`
	BankAddress        string `json:"bank_address"`
	RegistrationNumber string `json:"registration_number"`
}

type Account struct {
	ID              string        `json:"id"`
	ProductID       string        `json:"product_id"`
	PersonID        string        `json:"person_id"`
	CompanyID       string        `json:"company_id"`
	ExternalID      string        `json:"external_id"`
	Name            string        `json:"name"`
	CurrencyCode    string        `json:"currency_code"`
	CreditLimit     float64       `json:"credit_limit"`
	UsedCredit      float64       `json:"used_credit"`
	Balance         float64       `json:"balance"`
	AvailableAmount float64       `json:"available_amount"`
	BlockedAmount   float64       `json:"blocked_amount"`
	Status          string        `json:"status"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	ClosedAt        time.Time     `json:"closed_at"`
	Limits          AccountLimits `json:"limits"`
	ClosedBy        string        `json:"closed_by"`
	CloseReason     string        `json:"close_reason"`
	IsMain          bool          `json:"is_main"`
	CardsCount      int           `json:"cards_count"`
	Viban           string        `json:"viban"`
	TopUpDetails    TopUpDetails  `json:"top_up_details"`
	ReferenceNumber string        `json:"reference_number"`
}
