package request

import (
	"memmi/card"
	"memmi/pbuf"
	"net/http"
	"strings"
)

const CARD_SET_API_URL = "/api/cardset"
const GET_CARDSET_URL = CARD_SET_API_URL + "/get-card-set"
const GET_CARD_URL = CARD_SET_API_URL + "/get-card"

type CardSetRequestHandler struct {
	Pio     ProtoIO
	CardMan card.CardManagement
}

func (handler *CardSetRequestHandler) ShouldHandle(r *http.Request, user pbuf.User, responseWritten bool) bool {
	return strings.HasPrefix(r.URL.EscapedPath(), CARD_API_URL) && !responseWritten
}

func (handler *CardSetRequestHandler) Handle(w http.ResponseWriter, r *http.Request, user pbuf.User) HandleResult {
	result := HandleResult{Continue: true, ResponseWritten: false}
	switch r.URL.EscapedPath() {
	case CARD_SET_API_URL:
	case GET_CARDSET_URL:
		result.ResponseWritten = handler.handleCardSet(w, r, user)
	case GET_CARD_URL:
		result.ResponseWritten = handler.handleCard(w, r, user)
	}
	return result
}

func (handler *CardSetRequestHandler) handleCardSet(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	cardSetRequest, err := handler.Pio.ReadCardSetRequest(r)
	if err != nil {
		handler.Pio.WriteProtoResponse(w, BODY_READ_ERROR)
		return true
	}
	var cardSet = handler.CardMan.GetCardSetById(cardSetRequest.Id)
	handler.Pio.WriteProtoResponse(w, &cardSet)
	return true
}

func (handler *CardSetRequestHandler) handleCard(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	cardRequest, err := handler.Pio.ReadCardRequest(r)
	if err != nil {
		handler.Pio.WriteProtoResponse(w, BODY_READ_ERROR)
		return true
	}
	var card = handler.CardMan.GetCardById(cardRequest.Id)
	handler.Pio.WriteProtoResponse(w, &card)
	return true
}
