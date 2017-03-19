package request

import (
	"memmi/card"
	"memmi/pbuf"
	"net/http"
	"strings"
)

const CARD_SET_API_URL = "/api/card"
const GET_CARDSET_URL = CARD_SET_API_URL + "/get-card-set"
const GET_CARD_URL = CARD_SET_API_URL + "/get-card"

var CARD_NOT_FOUND_ERROR = &pbuf.RequestError{Reason: "Could not get the requested card."}
var CARD_SET_NOT_FOUND_ERROR = &pbuf.RequestError{Reason: "Could not get the requested card set."}

type CardSetRequestHandler struct {
	Pio     ProtoIO
	CardMan card.CardManagement
}

func (handler *CardSetRequestHandler) ShouldHandle(r *http.Request, user pbuf.User, responseWritten bool) bool {
	return strings.HasPrefix(r.URL.EscapedPath(), CARD_SET_API_URL) && !responseWritten
}

func (handler *CardSetRequestHandler) Handle(w http.ResponseWriter, r *http.Request, user pbuf.User) HandleResult {
	result := HandleResult{Continue: true, ResponseWritten: false}
	switch r.URL.EscapedPath() {
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
		handler.Pio.WriteCodedProtoResponse(w, BODY_READ_ERROR, http.StatusBadRequest)
		return true
	}
	cardSet, getErr := handler.CardMan.GetCardSetById(cardSetRequest.Id)
	if getErr != nil {
		handler.Pio.WriteCodedProtoResponse(w, CARD_SET_NOT_FOUND_ERROR, http.StatusBadRequest)
		return true
	}
	handler.Pio.WriteProtoResponse(w, &cardSet)
	return true
}

func (handler *CardSetRequestHandler) handleCard(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	cardRequest, err := handler.Pio.ReadCardRequest(r)
	if err != nil {
		handler.Pio.WriteCodedProtoResponse(w, BODY_READ_ERROR, http.StatusBadRequest)
		return true
	}
	card, getErr := handler.CardMan.GetCardById(cardRequest.Id)
	if getErr != nil {
		handler.Pio.WriteCodedProtoResponse(w, CARD_NOT_FOUND_ERROR, http.StatusBadRequest)
		return true
	}
	handler.Pio.WriteProtoResponse(w, &card)
	return true
}
