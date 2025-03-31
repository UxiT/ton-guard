package domaintype

type AccountStatus string

const (
	AccountActive  AccountStatus = "ACTIVE"
	AccountBlocked AccountStatus = "BLOCKED"
)

type Currency string

const (
	CurrencyEUR Currency = "EUR"
)
