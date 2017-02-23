package request

import (
	"memmi/pbuf"
	"net/http"
)

type AnonAuthenticator struct {
}

func (auth *AnonAuthenticator) AuthenticateUser(r *http.Request) pbuf.User {
	return pbuf.User{
		Id:              "id",
		UserName:        "anon",
		FirstName:       "John",
		LastName:        "Smith",
		Email:           "js@memmi.net",
		IsAuthenticated: false,
		IsAnon:          false,
		JoinedDate:      "",
	}
}
