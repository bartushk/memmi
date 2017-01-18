package handlers

import (
	"net/http"
)

type RequestHandler interface {
	ShouldHandle(r *http.Request) bool
	Handle(w http.ResponseWriter, r *http.Request) bool
}

type HttpLogger interface {
	Log(*http.Request)
}

type HttpRouter interface {
	AddHandler(RequestHandler)
	GetHandleFunc() func(http.ResponseWriter, *http.Request)
}
