package handlers

import (
	"memmi/pbuf"
	"net/http"
)

type MockHandler struct {
	DoHandle       bool
	CallCount      int
	ShouldContinue bool
}

func (handler *MockHandler) ShouldHandle(r *http.Request, user *pbuf.User) bool {
	return handler.DoHandle
}

func (handler *MockHandler) Handle(w http.ResponseWriter, r *http.Request, user *pbuf.User) bool {
	handler.CallCount += 1
	return handler.ShouldContinue
}
