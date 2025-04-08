package card

import (
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"decard/pkg/utils/decryptor"
	"github.com/rs/zerolog"
)

type GetCardPINQuery struct {
	CardUUID string
}

type GetCardPINResponse struct {
	CardPIN string `json:"card_pin"`
}

type GetCardPINQueryHandler struct {
	cardAPI        interfaces.CardService
	decryptService *decryptor.Decryptor
	logger         *zerolog.Logger
}

func NewGetCardPINQueryHandler(
	cardAPI interfaces.CardService,
	decryptService *decryptor.Decryptor,
	logger *zerolog.Logger,
) *GetCardPINQueryHandler {
	return &GetCardPINQueryHandler{
		cardAPI:        cardAPI,
		decryptService: decryptService,
		logger:         logger,
	}
}

func (h GetCardPINQueryHandler) Handle(query GetCardPINQuery) (GetCardPINResponse, error) {
	cardUUID, err := valueobject.ParseUUID(query.CardUUID)

	if err != nil {
		return GetCardPINResponse{}, err
	}

	encryptedPIN, err := h.cardAPI.GetCardPIN(cardUUID)
	if err != nil {
		h.logger.Error().Err(err).Str("cardUUID", query.CardUUID).Msg("cannot get card PIN")

		return GetCardPINResponse{}, err
	}

	cardPIN, err := h.decryptService.Decrypt(encryptedPIN)

	return GetCardPINResponse{
		CardPIN: cardPIN,
	}, err
}
