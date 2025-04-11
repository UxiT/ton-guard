package card

type GetCardInfoRequest struct {
	Card string `json:"-" uri:"card"`
}

func (r GetCardInfoRequest) Validate() error {
	return nil
}
