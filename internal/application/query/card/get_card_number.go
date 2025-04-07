package card

import (
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"decard/pkg/utils/decryptor"
)

type GetCardNumberQuery struct {
	CardUUID string
}

type GetCardNumberResponse struct {
	Number string `json:"number"`
}

type GetCardNumberQueryHandler struct {
	cardAPI        interfaces.CardService
	decryptService *decryptor.Decryptor
}

func NewGetCardNumberQueryHandler(cardAPI interfaces.CardService, decryptService *decryptor.Decryptor) *GetCardNumberQueryHandler {
	return &GetCardNumberQueryHandler{
		cardAPI:        cardAPI,
		decryptService: decryptService,
	}
}

func (h GetCardNumberQueryHandler) Handle(query GetCardNumberQuery) (GetCardNumberResponse, error) {
	cardUUID, err := valueobject.ParseUUID(query.CardUUID)

	if err != nil {
		return GetCardNumberResponse{}, err
	}

	encryptedNumber, err := h.cardAPI.GetCardNumber(cardUUID)
	if err != nil {
		return GetCardNumberResponse{}, err
	}

	number, err := h.decryptService.Decrypt(encryptedNumber)

	return GetCardNumberResponse{
		Number: number,
	}, err
}
