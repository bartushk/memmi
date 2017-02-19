package request

import (
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
	"net/http"
)

type MockAuthenticator struct {
	mock.Mock
}

func (auth *MockAuthenticator) AuthenticateUser(r *http.Request) pbuf.User {
	args := auth.Called(r)
	return args.Get(0).(pbuf.User)
}
