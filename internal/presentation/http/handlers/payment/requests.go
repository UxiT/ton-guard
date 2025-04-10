package payment

type TransferRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	From        string  `json:"from_account_id"`
	To          string  `json:"to_account_id"`
}

func (r TransferRequest) Validate() error {
	return nil
}
