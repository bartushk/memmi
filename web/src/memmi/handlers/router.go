package handlers

import (
	"net/http"
)

type HttpLogger interface {
	log(*http.Request)
}

type RequestHandler interface {
	ShouldHandle(r *http.Request) bool
	Handle(w http.ResponseWriter, r *http.Request) bool
}

func GetRouter() *Router {
	defaultRouter := Router{}
	return &defaultRouter
}

type Router struct {
	Logger   HttpLogger
	handlers []RequestHandler
}

func (router *Router) AddHandler(handler RequestHandler) {
	if handler != nil {
		router.handlers = append(router.handlers, handler)
	}
}

func (router *Router) GetHandleFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if router.Logger != nil {
			router.Logger.log(r)
		}
		for _, handler := range router.handlers {
			if handler.ShouldHandle(r) {
				if handler.Handle(w, r) {
					break
				}
			}
		}
	}
}
