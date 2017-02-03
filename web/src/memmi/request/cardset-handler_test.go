package request

import (
	"memmi/card"
	"memmi/pbuf"
	"testing"
)

func getCardSetMocked() (RequestHandler, *MockProtoIO, *card.MockCardManagement) {
	handler := &CardSetRequestHandler{}
	pio := &MockProtoIO{}
	cardMan := &card.MockCardManagement{}
	handler.Pio = pio
	handler.CardMan = cardMan
	return handler, pio, cardMan
}

func Test_CardSetHandler_ExactUrl_ShouldHandle(t *testing.T) {
	var req = RequestFromURL(CARD_SET_API_URL)
	handler := CardSetRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle with URL:", CARD_API_URL)
	}
}

func Test_CardSetHandler_URLPlusQuery_ShouldHandle(t *testing.T) {
	test_url := CARD_SET_API_URL + "?asdf"
	var req = RequestFromURL(test_url)
	handler := CardSetRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle with URL:", test_url)
	}
}

func Test_CardSetHandler_URLPlusSubUrl_ShouldHandle(t *testing.T) {
	test_url := CARD_SET_API_URL + "/asdf/fffa"
	var req = RequestFromURL(test_url)
	handler := CardSetRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle with URL:", test_url)
	}
}

func Test_CardSetHandler_AnyDifferentPrefix_ShouldNotHandle(t *testing.T) {
	test_url := "/test" + CARD_SET_API_URL
	var req = RequestFromURL(test_url)
	handler := CardSetRequestHandler{}
	if handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should not handle with URL:", test_url)
	}
}
