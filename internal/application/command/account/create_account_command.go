package account

import "context"

type CreateAccountCommand struct {
}

type CreateAccountCommandResponse struct{}

type CreateAccountCommandHandler struct {
}

func (h CreateAccountCommand) Handle(ctx context.Context, cmd CreateAccountCommand) (CreateAccountCommandResponse, error) {
	return CreateAccountCommandResponse{}, nil
}
