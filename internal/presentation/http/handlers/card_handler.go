package handlers

import (
	"decard/internal/domain/service"
)

type CardHandler struct {
	accountService service.CardService
}

func NewCardHandler() *CardHandler {
	return &CardHandler{}
}

//func (h *CardHandler) GetCardList(w http.ResponseWriter, r *http.Request) {
//	userUUID := r.Context().Value(middleware.UserUUIDKey).(uuid.UUID)
//
//	cards, err := h.cardService.GetCardsByCustomer(userUUID)
//}
