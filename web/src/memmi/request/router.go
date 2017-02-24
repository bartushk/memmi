package request

import (
	"net/http"
)

type Router struct {
	Authenticator HttpAuthenticator
	handlers      []RequestHandler
}

func (router *Router) AddHandler(handler RequestHandler) {
	if handler != nil {
		router.handlers = append(router.handlers, handler)
	}
}

func (router *Router) GetHandleFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := router.Authenticator.AuthenticateUser(r)
		var responseWritten = false
		for _, handler := range router.handlers {
			if handler.ShouldHandle(r, user, responseWritten) {
				handleResult := handler.Handle(w, r, user)
				if !handleResult.Continue {
					break
				}
				responseWritten = handleResult.ResponseWritten || responseWritten
			}
		}
	}
}
