package transfer

type MakeAccountTransferCommand struct {
	Amount      float64
	Description string
	AccountFrom string
	AccountTo   string
}

type MakeAccountTransferCommandHandler struct {
}

func NewMakeAccountTransferCommandHandler() *MakeAccountTransferCommandHandler {
	return &MakeAccountTransferCommandHandler{}
}

func (h MakeAccountTransferCommandHandler) Handle(cmd MakeAccountTransferCommand) error {
	return nil
}
