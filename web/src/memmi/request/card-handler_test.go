package request

import (
	"memmi/pbuf"
	"net/http"
	"net/url"
	"testing"
)

func requestFromURL(url_string string) *http.Request {
	test_url := &url.URL{}
	test_request := &http.Request{}
	test_url.Path = url_string
	test_request.URL = test_url
	return test_request
}

func Test_CardHandler_ExactUrl_ShouldHanlde(t *testing.T) {
	var req = requestFromURL(CARD_API_URL)
	handler := CardRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}) {
		t.Error("Handler should handle request with URL:", CARD_API_URL)
	}
}

func Test_CardHandler_UrlPlusQuery_ShouldHanlde(t *testing.T) {
	test_url := CARD_API_URL + "?asdf"
	var req = requestFromURL(test_url)
	handler := CardRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}) {
		t.Error("Handler should handle request with URL:", test_url)
	}
}

func Test_CardHandler_UrlSubUrl_ShouldHanlde(t *testing.T) {
	test_url := CARD_API_URL + "/asdf/fffa"
	var req = requestFromURL(test_url)
	handler := CardRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}) {
		t.Error("Handler should handle request with URL:", test_url)
	}
}

func Test_CardHandler_AnyDifferentPrefix_ShouldNotHanlde(t *testing.T) {
	test_url := "/tes" + CARD_API_URL
	var req = requestFromURL(test_url)
	handler := CardRequestHandler{}
	if handler.ShouldHandle(req, pbuf.User{}) {
		t.Error("Handler should not handle request with URL:", test_url)
	}
}
