package handlers

import (
	"memmi/pbuf"
	"net/http"
)

type RequestHandler interface {
	ShouldHandle(r *http.Request, u *pbuf.User) bool
	Handle(w http.ResponseWriter, r *http.Request, u *pbuf.User) bool
}

type HttpLogger interface {
	Log(*http.Request)
}

type HttpRouter interface {
	AddHandler(RequestHandler)
	GetHandleFunc() func(http.ResponseWriter, *http.Request)
}
