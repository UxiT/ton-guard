package transaction

type GetCardTransactionRequest struct {
	Card string `json:"-" uri:"card"`
}

func (r GetCardTransactionRequest) Validate() error {
	return nil
}
