package card

import (
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"decard/pkg/utils/decryptor"
	"github.com/rs/zerolog"
)

type GetCard3DSQuery struct {
	CardUUID string
}

type GetCard3DSResponse struct {
	Card3DS string `json:"card_3ds"`
}

type GetCard3DSQueryHandler struct {
	cardAPI        interfaces.CardService
	decryptService *decryptor.Decryptor
	logger         *zerolog.Logger
}

func NewGetCard3DSQueryHandler(
	cardAPI interfaces.CardService,
	decryptService *decryptor.Decryptor,
	logger *zerolog.Logger,
) *GetCard3DSQueryHandler {
	return &GetCard3DSQueryHandler{
		cardAPI:        cardAPI,
		decryptService: decryptService,
		logger:         logger,
	}
}

func (h GetCard3DSQueryHandler) Handle(query GetCard3DSQuery) (GetCard3DSResponse, error) {
	cardUUID, err := valueobject.ParseUUID(query.CardUUID)

	if err != nil {
		return GetCard3DSResponse{}, err
	}

	encrypted3DS, err := h.cardAPI.GetCard3DS(cardUUID)
	if err != nil {
		h.logger.Error().Err(err).Str("cardUUID", query.CardUUID).Msg("cannot get card 3DS")

		return GetCard3DSResponse{}, err
	}

	card3DS, err := h.decryptService.Decrypt(encrypted3DS)

	return GetCard3DSResponse{
		Card3DS: card3DS,
	}, err
}
