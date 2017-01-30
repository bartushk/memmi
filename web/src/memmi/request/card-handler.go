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
	CardSel card.CardSelection
	UserMan user.UserManagement
}

func (handler *CardRequestHandler) ShouldHandle(r *http.Request, user pbuf.User, responseWritten bool) bool {
	return strings.HasPrefix(r.URL.EscapedPath(), CARD_API_URL) && !responseWritten
}

func (handler *CardRequestHandler) Handle(w http.ResponseWriter, r *http.Request, user pbuf.User) HandleResult {
	result := HandleResult{Continue: true, ResponseWritten: false}
	switch r.URL.EscapedPath() {
	case CARD_NEXT_URL:
		result.ResponseWritten = handler.handleNext(w, r, user)
	case CARD_REPORT_URL:
		result.ResponseWritten = handler.handleReport(w, r, user)
	case CARD_REPORT_NEXT_URL:
		handler.handleReportNext(w, r, user)
	}
	return result
}

func (handler *CardRequestHandler) handleNext(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	nextRequest, err := handler.Pio.ReadNextCardRequest(r)
	if err != nil {
		handler.Pio.WriteProtoResponse(w, BODY_READ_ERROR)
		return true
	}
	history := handler.UserMan.GetHistory(user, nextRequest.CardSetId)
	nextCard := handler.CardSel.SelectCard(&history, nextRequest.PreviousCardId)
	handler.Pio.WriteProtoResponse(w, &nextCard)
	return true
}

func (handler *CardRequestHandler) handleReport(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	cardScoreReport, csErr := handler.Pio.ReadCardScoreReport(r)
	if csErr != nil {
		handler.Pio.WriteProtoResponse(w, BODY_READ_ERROR)
		return true
	}
	updateErr := handler.UserMan.UpdateHistory(user, cardScoreReport.CardSetId, *cardScoreReport.Update)
	if updateErr != nil {
		handler.Pio.WriteProtoResponse(w, USER_HISTORY_UPDATE_ERROR)
	}
	return true
}

func (handler *CardRequestHandler) handleReportNext(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	return false
}
