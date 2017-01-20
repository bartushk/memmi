package handlers

import (
	"memmi/pbuf"
	"net/http"
)

type MockAuthenticator struct {
	AuthenticatedUser pbuf.User
	CallCount         int
}

func (auth *MockAuthenticator) AuthenticateUser(r *http.Request) pbuf.User {
	auth.CallCount += 1
	return auth.AuthenticatedUser
}
