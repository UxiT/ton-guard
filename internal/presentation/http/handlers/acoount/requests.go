package acoount

type GetAccountCardsRequest struct {
	Account string `json:"account"`
}

func (r GetAccountCardsRequest) Validate() error {
	return nil
}
