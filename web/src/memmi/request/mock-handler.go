package request

import (
	"memmi/pbuf"
	"net/http"
)

type MockHandler struct {
	DoHandle       bool
	ShouldContinue bool
	HandleUsers    []pbuf.User
	ShouldUsers    []pbuf.User
	HandleRequests []*http.Request
	ShouldRequests []*http.Request
	HandleWriters  []http.ResponseWriter
}

func (handler *MockHandler) ShouldHandle(r *http.Request, user pbuf.User) bool {
	handler.ShouldUsers = append(handler.ShouldUsers, user)
	handler.ShouldRequests = append(handler.ShouldRequests, r)
	return handler.DoHandle
}

func (handler *MockHandler) Handle(w http.ResponseWriter, r *http.Request, user pbuf.User) bool {
	handler.HandleUsers = append(handler.HandleUsers, user)
	handler.HandleRequests = append(handler.HandleRequests, r)
	handler.HandleWriters = append(handler.HandleWriters, w)
	return handler.ShouldContinue
}
