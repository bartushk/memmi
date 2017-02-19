package request

import (
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
	"net/http"
)

type MockHandler struct {
	mock.Mock
}

func (handler *MockHandler) ShouldHandle(r *http.Request, user pbuf.User, responseWritten bool) bool {
	args := handler.Called(r, user, responseWritten)
	return args.Bool(0)
}

func (handler *MockHandler) Handle(w http.ResponseWriter, r *http.Request, user pbuf.User) HandleResult {
	args := handler.Called(w, r, user)
	return args.Get(0).(HandleResult)
}
