package handlers

import (
	"memmi/pbuf"
	"net/http"
)

type Router struct {
	Logger        HttpLogger
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
		if router.Logger != nil {
			router.Logger.Log(r)
		}

		user := pbuf.User{
			Id:              0,
			UserName:        "anon",
			FirstName:       "Cool",
			LastName:        "Person",
			Email:           "none",
			IsAuthenticated: false,
		}

		if router.Authenticator != nil {
			user = router.Authenticator.AuthenticateUser(r)
		}

		for _, handler := range router.handlers {
			if handler.ShouldHandle(r, &user) {
				if !handler.Handle(w, r, &user) {
					break
				}
			}
		}
	}
}
