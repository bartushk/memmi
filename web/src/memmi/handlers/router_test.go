package handlers

import (
	"testing"
)

func getMockedRouter(i int) (Router, *MockLogger, []*MockHandler) {
	var handlers []*MockHandler
	logger := new(MockLogger)
	router := Router{
		Logger: logger,
	}
	for k := 0; k < i; k++ {
		handler := new(MockHandler)
		handler.ShouldContinue = true
		handler.DoHandle = true
		handlers = append(handlers, handler)
		router.AddHandler(handler)
	}
	return router, logger, handlers
}

func Test_Handlers_Router_Constructor_NoHandlers(t *testing.T) {
	router := Router{}
	if len(router.handlers) != 0 {
		t.Error("Handlers member was not empty.")
	}
}

func Test_Handlers_Router_AddHandler_HandlersAdded(t *testing.T) {
	router := Router{}
	h1 := new(MockHandler)
	h2 := new(MockHandler)
	h3 := new(MockHandler)
	router.AddHandler(h1)
	router.AddHandler(h2)
	router.AddHandler(h3)
	if len(router.handlers) != 3 {
		t.Error("For handlers expected 3 got", len(router.handlers))
	}

	if router.handlers[0] != h1 {
		t.Error("h1 not found at handlers[0]")
	}

	if router.handlers[1] != h2 {
		t.Error("h2 not found at handlers[1]")
	}

	if router.handlers[2] != h3 {
		t.Error("h3 not found at handlers[2]")
	}
}

func Test_Handlers_Router_AddHandler_MultipeCopies(t *testing.T) {
	router := Router{}
	h1 := new(MockHandler)
	h2 := new(MockHandler)
	router.AddHandler(h1)
	router.AddHandler(h2)
	router.AddHandler(h1)
	if len(router.handlers) != 3 {
		t.Error("For handlers expected 3 got", len(router.handlers))
	}

	if router.handlers[0] != h1 {
		t.Error("h1 not found at handlers[0]")
	}

	if router.handlers[1] != h2 {
		t.Error("h2 not found at handlers[1]")
	}

	if router.handlers[2] != h1 {
		t.Error("h1 not found at handlers[2]")
	}
}

func Test_Handlers_Router_HandleFunc_NotNull(t *testing.T) {
	var r, _, _ = getMockedRouter(0)
	var hf = r.GetHandleFunc()
	if hf == nil {
		t.Error("Handle func was nil.")
	}
}

func Test_Handlers_Router_HandleFunc_NoHandlers_LoggerCalled(t *testing.T) {
	var r, l, _ = getMockedRouter(0)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	if l.CallCount != 1 {
		t.Error("Logger CallCount expected 1, got", l.CallCount)
	}
}

func Test_Handlers_Router_HandleFunc_Handlers_LoggerCalled(t *testing.T) {
	var r, l, _ = getMockedRouter(3)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	if l.CallCount != 2 {
		t.Error("Logger CallCount expected 2, got", l.CallCount)
	}
}

func Test_Handlers_Router_HandleFunc_ShouldGo_AllCalled(t *testing.T) {
	var r, _, h = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)
	for i, handler := range h {
		if handler.CallCount != 3 {
			t.Error("Logger CallCount expected 3, got", handler.CallCount,
				"on handler", i)
		}
	}
}

func Test_Handlers_Router_HandleFunc_ShouldNotContinue_LoopBreaks(t *testing.T) {
	var r, _, h = getMockedRouter(4)
	h[1].ShouldContinue = false
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)

	for i, handler := range h[:1] {
		if handler.CallCount != 3 {
			t.Error("Logger CallCount expected 3, got", handler.CallCount,
				"on handler", i)
		}
	}

	for i, handler := range h[2:] {
		if handler.CallCount != 0 {
			t.Error("Logger CallCount expected 0, got", handler.CallCount,
				"on handler", i+2)
		}
	}
}

func Test_Handlers_Router_HandleFunc_ShouldNotHandle_DoesNotHandle(t *testing.T) {
	var r, _, h = getMockedRouter(4)
	h[0].DoHandle = false
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)

	if h[0].CallCount != 0 {
		t.Error("Logger CallCount expected 0, got", h[0].CallCount,
			"on handler 0")
	}

	for i, handler := range h[1:] {
		if handler.CallCount != 3 {
			t.Error("Logger CallCount expected 3, got", handler.CallCount,
				"on handler", i+1)
		}
	}
}
