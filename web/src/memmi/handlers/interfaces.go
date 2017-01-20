package handlers

import (
	"memmi/pbuf"
	"net/http"
)

type HttpAuthenticator interface {
	AuthenticateUser(r *http.Request) pbuf.User
}

type RequestHandler interface {
	ShouldHandle(r *http.Request, u pbuf.User) bool
	Handle(w http.ResponseWriter, r *http.Request, u pbuf.User) bool
}

type HttpLogger interface {
	Log(*http.Request)
}

type HttpRouter interface {
	AddHandler(RequestHandler)
	GetHandleFunc() func(http.ResponseWriter, *http.Request)
}
