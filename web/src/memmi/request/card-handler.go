package request

import (
	"memmi/pbuf"
	"net/http"
	"strings"
)

const CARD_API_URL = "/api/card"
const CARD_NEXT_URL = CARD_API_URL + "/get-next"
const CARD_REPORT_URL = CARD_API_URL + "/report"
const CARD_REPORT_NEXT_URL = CARD_API_URL + "/report-get-next"

type CardRequestHandler struct {
}

func (handler *CardRequestHandler) ShouldHandle(r *http.Request, user pbuf.User) bool {
	return strings.HasPrefix(r.URL.EscapedPath(), CARD_API_URL)
}

func (handler *CardRequestHandler) Handle(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	switch r.URL.EscapedPath() {
	case CARD_NEXT_URL:
		return handleNext(w, r, user)
	case CARD_REPORT_URL:
		return handleReport(w, r, user)
	case CARD_REPORT_NEXT_URL:
		return handleReportNext(w, r, user)
	default:
		return false
	}
}

func handleNext(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	return true
}

func handleReport(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	return true
}

func handleReportNext(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	return true
}
