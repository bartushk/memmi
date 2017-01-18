package handlers

import (
	"net/http"
)

type MockHandler struct {
	DoHandle       bool
	CallCount      int
	ShouldContinue bool
}

func (handler *MockHandler) ShouldHandle(r *http.Request) bool {
	return handler.DoHandle
}

func (handler *MockHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	handler.CallCount += 1
	return handler.ShouldContinue
}
