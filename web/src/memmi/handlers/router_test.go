package handlers

import (
	"testing"
)

func Test_Handlers_Router_GetRouter_NotNull(t *testing.T) {
	router := GetRouter()
	if router == nil {
		t.Error("Router was null.")
	}
}

func Test_Handlers_Router_GetRouter_NoHandlers(t *testing.T) {
	router := GetRouter()
	if len(router.handlers) != 0 {
		t.Error("Handlers member was not empty.")
	}
}
