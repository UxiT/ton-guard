package provider_entity

import "time"

type CardLimits struct {
	DailyPurchase                       int64 `json:"daily_purchase"`
	DailyWithdrawal                     int64 `json:"daily_withdrawal"`
	DailyInternetPurchase               int64 `json:"daily_internet_purchase"`
	DailyContactlessPurchase            int64 `json:"daily_contactless_purchase"`
	WeeklyPurchase                      int64 `json:"weekly_purchase"`
	WeeklyWithdrawal                    int64 `json:"weekly_withdrawal"`
	WeeklyInternetPurchase              int64 `json:"weekly_internet_purchase"`
	WeeklyContactlessPurchase           int64 `json:"weekly_contactless_purchase"`
	MonthlyPurchase                     int64 `json:"monthly_purchase"`
	MonthlyWithdrawal                   int64 `json:"monthly_withdrawal"`
	MonthlyInternetPurchase             int64 `json:"monthly_internet_purchase"`
	MonthlyContactlessPurchase          int64 `json:"monthly_contactless_purchase"`
	TransactionPurchase                 int64 `json:"transaction_purchase"`
	TransactionWithdrawal               int64 `json:"transaction_withdrawal"`
	TransactionInternetPurchase         int64 `json:"transaction_internet_purchase"`
	TransactionContactlessPurchase      int64 `json:"transaction_contactless_purchase"`
	DailyOverallPurchase                int64 `json:"daily_overall_purchase"`
	WeeklyOverallPurchase               int64 `json:"weekly_overall_purchase"`
	MonthlyOverallPurchase              int64 `json:"monthly_overall_purchase"`
	AllTimePurchase                     int64 `json:"all_time_purchase"`
	AllTimeWithdrawal                   int64 `json:"all_time_withdrawal"`
	AllTimeInternetPurchase             int64 `json:"all_time_internet_purchase"`
	AllTimeContactlessPurchase          int64 `json:"all_time_contactless_purchase"`
	AllTimeContactlessPurchaseAvailable int64 `json:"all_time_contactless_purchase_available"`
	AllTimeContactlessPurchaseUsed      int64 `json:"all_time_contactless_purchase_used"`
	AllTimeInternetPurchaseAvailable    int64 `json:"all_time_internet_purchase_available"`
	AllTimeInternetPurchaseUsed         int64 `json:"all_time_internet_purchase_used"`
	AllTimePurchaseAvailable            int64 `json:"all_time_purchase_available"`
	AllTimePurchaseUsed                 int64 `json:"all_time_purchase_used"`
	AllTimeWithdrawalAvailable          int64 `json:"all_time_withdrawal_available"`
	AllTimeWithdrawalUsed               int64 `json:"all_time_withdrawal_used"`
	DailyContactlessPurchaseAvailable   int64 `json:"daily_contactless_purchase_available"`
	DailyContactlessPurchaseUsed        int64 `json:"daily_contactless_purchase_used"`
	DailyInternetPurchaseAvailable      int64 `json:"daily_internet_purchase_available"`
	DailyInternetPurchaseUsed           int64 `json:"daily_internet_purchase_used"`
	DailyOverallPurchaseAvailable       int64 `json:"daily_overall_purchase_available"`
	DailyOverallPurchaseUsed            int64 `json:"daily_overall_purchase_used"`
	DailyPurchaseAvailable              int64 `json:"daily_purchase_available"`
	DailyPurchaseUsed                   int64 `json:"daily_purchase_used"`
	DailyWithdrawalAvailable            int64 `json:"daily_withdrawal_available"`
	DailyWithdrawalUsed                 int64 `json:"daily_withdrawal_used"`
	MonthlyContactlessPurchaseAvailable int64 `json:"monthly_contactless_purchase_available"`
	MonthlyContactlessPurchaseUsed      int64 `json:"monthly_contactless_purchase_used"`
	MonthlyInternetPurchaseAvailable    int64 `json:"monthly_internet_purchase_available"`
	MonthlyInternetPurchaseUsed         int64 `json:"monthly_internet_purchase_used"`
	MonthlyOverallPurchaseAvailable     int64 `json:"monthly_overall_purchase_available"`
	MonthlyOverallPurchaseUsed          int64 `json:"monthly_overall_purchase_used"`
	MonthlyPurchaseAvailable            int64 `json:"monthly_purchase_available"`
	MonthlyPurchaseUsed                 int64 `json:"monthly_purchase_used"`
	MonthlyWithdrawalAvailable          int64 `json:"monthly_withdrawal_available"`
	MonthlyWithdrawalUsed               int64 `json:"monthly_withdrawal_used"`
	WeeklyContactlessPurchaseAvailable  int64 `json:"weekly_contactless_purchase_available"`
	WeeklyContactlessPurchaseUsed       int64 `json:"weekly_contactless_purchase_used"`
	WeeklyInternetPurchaseAvailable     int64 `json:"weekly_internet_purchase_available"`
	WeeklyInternetPurchaseUsed          int64 `json:"weekly_internet_purchase_used"`
	WeeklyOverallPurchaseAvailable      int64 `json:"weekly_overall_purchase_available"`
	WeeklyOverallPurchaseUsed           int64 `json:"weekly_overall_purchase_used"`
	WeeklyPurchaseAvailable             int64 `json:"weekly_purchase_available"`
	WeeklyPurchaseUsed                  int64 `json:"weekly_purchase_used"`
	WeeklyWithdrawalAvailable           int64 `json:"weekly_withdrawal_available"`
	WeeklyWithdrawalUsed                int64 `json:"weekly_withdrawal_used"`
}

type ThreeDSecureSettings struct {
	Type             string `json:"type"`
	Mobile           string `json:"mobile"`
	LanguageCode     string `json:"language_code"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	OutOfBandEnabled bool   `json:"out_of_band_enabled"`
	OutOfBandID      string `json:"out_of_band_id"`
}

type DeliveryAddress struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	CompanyName    string `json:"company_name"`
	Address1       string `json:"address1"`
	Address2       string `json:"address2"`
	PostalCode     string `json:"postal_code"`
	City           string `json:"city"`
	CountryCode    string `json:"country_code"`
	DispatchMethod string `json:"dispatch_method"`
	Phone          string `json:"phone"`
	TrackingNumber string `json:"tracking_number"`
}

type CardNotificationSettings struct {
	ReceiptsReminderEnabled   bool `json:"receipts_reminder_enabled"`
	InstantSpendUpdateEnabled bool `json:"instant_spend_update_enabled"`
}

type Card struct {
	ID                      string               `json:"id"`
	PredecessorCardID       string               `json:"predecessor_card_id"`
	AccountID               string               `json:"account_id"`
	PersonID                string               `json:"person_id"`
	ExternalID              string               `json:"external_id"`
	Type                    string               `json:"type"`
	Name                    string               `json:"name"`
	MaskedCardNumber        string               `json:"masked_card_number"`
	ReferenceNumber         string               `json:"reference_number"`
	ExpiryDate              time.Time            `json:"expiry_date"`
	BlockType               string               `json:"block_type"`
	BlockedAt               time.Time            `json:"blocked_at"`
	BlockedBy               string               `json:"blocked_by"`
	Status                  string               `json:"status"`
	EmbossingName           string               `json:"embossing_name"`
	EmbossingFirstName      string               `json:"embossing_first_name"`
	EmbossingLastName       string               `json:"embossing_last_name"`
	EmbossingCompanyName    string               `json:"embossing_company_name"`
	Limits                  CardLimits           `json:"limits"`
	ThreeDSecureSettings    ThreeDSecureSettings `json:"3d_secure_settings"`
	DeliveryAddress         DeliveryAddress      `json:"delivery_address"`
	IsEnrolledFor3DSecure   bool                 `json:"is_enrolled_for_3d_secure"`
	IsCard3DSecureActivated bool                 `json:"is_card_3d_secure_activated"`
	RenewAutomatically      bool                 `json:"renew_automatically"`
	IsDisposable            bool                 `json:"is_disposable"`
	Security                struct {
		ContactlessEnabled      bool `json:"contactless_enabled"`
		WithdrawalEnabled       bool `json:"withdrawal_enabled"`
		InternetPurchaseEnabled bool `json:"internet_purchase_enabled"`
		OverallLimitsEnabled    bool `json:"overall_limits_enabled"`
		AllTimeLimitsEnabled    bool `json:"all_time_limits_enabled"`
	} `json:"security"`
	PersonalizationProductCode string                   `json:"personalization_product_code"`
	CarrierType                string                   `json:"carrier_type"`
	CardMetadataProfileID      string                   `json:"card_metadata_profile_id"`
	ActivatedAt                time.Time                `json:"activated_at"`
	CreatedAt                  time.Time                `json:"created_at"`
	UpdatedAt                  time.Time                `json:"updated_at"`
	ClosedAt                   time.Time                `json:"closed_at"`
	ClosedBy                   string                   `json:"closed_by"`
	CloseReason                string                   `json:"close_reason"`
	CompanyID                  string                   `json:"company_id"`
	DispatchedAt               time.Time                `json:"dispatched_at"`
	CardNotificationSettings   CardNotificationSettings `json:"card_notification_settings"`
	DisposableType             string                   `json:"disposable_type"`
}
