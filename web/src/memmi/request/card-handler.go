package request

import (
	"memmi/card"
	"memmi/pbuf"
	"memmi/user"
	"net/http"
	"strings"
)

const CARD_API_URL = "/api/card"
const CARD_NEXT_URL = CARD_API_URL + "/get-next"
const CARD_REPORT_URL = CARD_API_URL + "/report"
const CARD_REPORT_NEXT_URL = CARD_API_URL + "/report-get-next"

type CardRequestHandler struct {
	Pio     ProtoIO
	CardMan card.CardManagement
	CardSel card.CardSelection
	UserMan user.UserManagement
}

func (handler *CardRequestHandler) ShouldHandle(r *http.Request, user pbuf.User) bool {
	return strings.HasPrefix(r.URL.EscapedPath(), CARD_API_URL)
}

func (handler *CardRequestHandler) Handle(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	switch r.URL.EscapedPath() {
	case CARD_NEXT_URL:
		handler.handleNext(w, r, user)
	case CARD_REPORT_URL:
		if handler.handleReport(w, r, user) {
			var toWrite []byte
			w.Write(toWrite)
		}
	case CARD_REPORT_NEXT_URL:
		handler.handleReportNext(w, r, user)
	}
	return false
}

func (handler *CardRequestHandler) handleNext(w http.ResponseWriter, r *http.Request, user pbuf.User) {
	nextRequest, err := handler.Pio.ReadNextCardRequest(r)
	if err != nil {
		pErr := &pbuf.RequestError{Reason: "Could not read request body."}
		handler.Pio.WriteProtoResponse(w, pErr)
		return
	}
	history := handler.UserMan.GetHistory(user, nextRequest.CardSetId)
	nextCard := handler.CardSel.SelectCard(&history, nextRequest.PreviousCardId)
	handler.Pio.WriteProtoResponse(w, &nextCard)
}

func (handler *CardRequestHandler) handleReport(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	cardScoreReport, csErr := handler.Pio.ReadCardScoreReport(r)
	if csErr != nil {
		pErr := &pbuf.RequestError{Reason: "Could not read request body."}
		handler.Pio.WriteProtoResponse(w, pErr)
		return false
	}
	updateErr := handler.UserMan.UpdateHistory(user, cardScoreReport.CardSetId, *cardScoreReport.Update)
	if updateErr != nil {
		pErr := &pbuf.RequestError{Reason: "Could not update user history."}
		handler.Pio.WriteProtoResponse(w, pErr)
		return false
	}
	return true
}

func (handler *CardRequestHandler) handleReportNext(w http.ResponseWriter, r *http.Request, user pbuf.User) {
	handler.handleReport(w, r, user)
	handler.handleNext(w, r, user)
}
