package operation

type TopUpRequest struct {
	Amount  string `json:"amount"`
	Network string `json:"network"`
}

func (r TopUpRequest) Validate() error {
	return nil
}

type AddTransactionRequest struct {
	TopUp         string `json:"-" uri:"uuid"`
	TransactionID string `json:"transaction_id"`
}

func (r AddTransactionRequest) Validate() error {
	return nil
}

type TopUpUUIDRequest struct {
	UUID string `json:"-" uri:"uuid"`
}

func (r TopUpUUIDRequest) Validate() error {
	return nil
}
