package request

import (
	"memmi/pbuf"
	"net/http"
)

type MockAuthenticator struct {
	AuthenticatedUser pbuf.User
	CallRequests      []*http.Request
}

func (auth *MockAuthenticator) AuthenticateUser(r *http.Request) pbuf.User {
	auth.CallRequests = append(auth.CallRequests, r)
	return auth.AuthenticatedUser
}
