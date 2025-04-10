package auth

type LoginRequest struct {
	TelegramID int    `json:"telegram_id"`
	Password   string `json:"password"`
}

func (r LoginRequest) Validate() error {
	return nil
}

type RegisterRequest struct {
	TelegramID int    `json:"telegram_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (r RegisterRequest) Validate() error {
	return nil
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (r RefreshRequest) Validate() error {
	return nil
}
