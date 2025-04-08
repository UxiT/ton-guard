package card

import (
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"decard/pkg/utils/decryptor"
	"github.com/rs/zerolog"
)

type GetCardCVVQuery struct {
	CardUUID string
}

type GetCardCVVResponse struct {
	CVV string `json:"cvv"`
}

type GetCardCVVQueryHandler struct {
	cardAPI        interfaces.CardService
	decryptService *decryptor.Decryptor
	logger         *zerolog.Logger
}

func NewGetCardCVVQueryHandler(
	cardAPI interfaces.CardService,
	decryptService *decryptor.Decryptor,
	logger *zerolog.Logger,
) *GetCardCVVQueryHandler {
	return &GetCardCVVQueryHandler{
		cardAPI:        cardAPI,
		decryptService: decryptService,
		logger:         logger,
	}
}

func (h GetCardCVVQueryHandler) Handle(query GetCardCVVQuery) (GetCardCVVResponse, error) {
	cardUUID, err := valueobject.ParseUUID(query.CardUUID)

	if err != nil {
		return GetCardCVVResponse{}, err
	}

	encryptedCVV, err := h.cardAPI.GetCardCVV(cardUUID)
	if err != nil {
		h.logger.Error().Err(err).Str("cardUUID", query.CardUUID).Msg("cannot get card cvv")

		return GetCardCVVResponse{}, err
	}

	cvv, err := h.decryptService.Decrypt(encryptedCVV)

	return GetCardCVVResponse{
		CVV: cvv,
	}, err
}
