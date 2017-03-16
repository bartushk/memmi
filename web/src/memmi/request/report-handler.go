package request

import (
	"memmi/pbuf"
	"memmi/user"
	"net/http"
	"strings"
)

const REPORT_API_URL = "/api/report"
const SCORE_REPORT_URL = REPORT_API_URL + "/card_score"

type ReportRequestHandler struct {
	Pio     ProtoIO
	UserMan user.UserManagement
}

func (handler *ReportRequestHandler) ShouldHandle(r *http.Request, user pbuf.User, responseWritten bool) bool {
	return strings.HasPrefix(r.URL.EscapedPath(), REPORT_API_URL) && !responseWritten
}

func (handler *ReportRequestHandler) Handle(w http.ResponseWriter, r *http.Request, user pbuf.User) HandleResult {
	result := HandleResult{Continue: true, ResponseWritten: false}
	switch r.URL.EscapedPath() {
	case SCORE_REPORT_URL:
		result.ResponseWritten = handler.handleReport(w, r, user)
	}
	return result
}

func (handler *ReportRequestHandler) handleReport(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	cardScoreReport, csErr := handler.Pio.ReadCardScoreReport(r)
	if csErr != nil {
		handler.Pio.WriteProtoResponse(w, BODY_READ_ERROR)
		return true
	}
	updateErr := handler.UserMan.UpdateHistory(user, cardScoreReport.CardSetId, *cardScoreReport.Update)
	if updateErr != nil {
		handler.Pio.WriteProtoResponse(w, USER_HISTORY_UPDATE_ERROR)
		return true
	}
	handler.Pio.WriteProtoResponse(w, &pbuf.UpdateResponse{Status: 1})
	return true
}
