package request

import (
	"net/http"
	"testing"
)

func getMockedRouter(i int) (Router, *MockLogger, []*MockHandler, *MockAuthenticator) {
	var handlers []*MockHandler
	logger := new(MockLogger)
	auth := new(MockAuthenticator)
	router := Router{
		Logger:        logger,
		Authenticator: auth,
	}
	for k := 0; k < i; k++ {
		handler := new(MockHandler)
		handler.Result.Continue = true
		handler.DoHandle = true
		handlers = append(handlers, handler)
		router.AddHandler(handler)
	}
	return router, logger, handlers, auth
}

func Test_Router_Constructor_NoHandlers(t *testing.T) {
	router := Router{}
	if len(router.handlers) != 0 {
		t.Error("Handlers member was not empty.")
	}
}

func Test_Router_AddHandler_HandlersAdded(t *testing.T) {
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

func Test_Router_AddHandler_MultipeCopies(t *testing.T) {
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

func Test_Router_HandleFunc_NotNull(t *testing.T) {
	var r, _, _, _ = getMockedRouter(0)
	var hf = r.GetHandleFunc()
	if hf == nil {
		t.Error("Handle func was nil.")
	}
}

func Test_Router_HandleFunc_NoHandlers_LoggerCalled(t *testing.T) {
	var r, l, _, _ = getMockedRouter(0)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	if len(l.CallRequests) != 1 {
		t.Error("Logger CallCount expected 1, got", len(l.CallRequests))
	}
}

func Test_Router_HandleFunc_Handlers_LoggerCalled(t *testing.T) {
	var r, l, _, _ = getMockedRouter(3)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	if len(l.CallRequests) != 2 {
		t.Error("Logger CallCount expected 2, got", len(l.CallRequests))
	}
}

func Test_Router_HandleFunc_LoggerCalled_RightRequest(t *testing.T) {
	var r, l, _, _ = getMockedRouter(3)
	var hf = r.GetHandleFunc()
	request1 := http.Request{}
	request2 := http.Request{}
	hf(nil, &request1)
	hf(nil, &request2)
	if l.CallRequests[0] != &request1 {
		t.Error("Did not pass correct request to logger at first call.")
	}
	if l.CallRequests[1] != &request2 {
		t.Error("Did not pass correct request to logger at second call.")
	}
}

func Test_Router_HandleFunc_ShouldGo_AllCalled(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)
	for i, handler := range h {
		if len(handler.HandleRequests) != 3 {
			t.Error("Logger CallCount expected 3, got", len(handler.HandleRequests),
				"on handler", i)
		}
	}
}

func Test_Router_HandleFunc_ShouldNotContinue_LoopBreaks(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4)
	h[1].Result.Continue = false
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)

	for i, handler := range h[:1] {
		if len(handler.HandleRequests) != 3 {
			t.Error("Logger CallCount expected 3, got", len(handler.HandleRequests),
				"on handler", i)
		}
	}

	for i, handler := range h[2:] {
		if len(handler.HandleRequests) != 0 {
			t.Error("Logger CallCount expected 0, got", len(handler.HandleRequests),
				"on handler", i+2)
		}
	}
}

func Test_Router_HandleFunc_ShouldNotHandle_DoesNotHandle(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4)
	h[0].DoHandle = false
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)

	if len(h[0].HandleRequests) != 0 {
		t.Error("Logger CallCount expected 0, got", len(h[0].HandleRequests),
			"on handler 0")
	}

	for i, handler := range h[1:] {
		if len(handler.HandleRequests) != 3 {
			t.Error("Logger CallCount expected 3, got", len(handler.HandleRequests),
				"on handler", i+1)
		}
	}
}

func Test_Router_HandleFunc_HandlerReceives_CorrectRequest(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	request1 := http.Request{}
	request2 := http.Request{}
	hf(nil, &request1)
	hf(nil, &request2)
	for i, handler := range h {
		if handler.HandleRequests[0] != &request1 {
			t.Error("Did not get correct value for first request on handler", i)
		}
		if handler.HandleRequests[1] != &request2 {
			t.Error("Did not get correct value for second request on handler", i)
		}
	}
}

func Test_Router_HandleFunc_HandlerReceives_CorrectWriter(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	writer1 := new(http.ResponseWriter)
	writer2 := new(http.ResponseWriter)
	hf(*writer1, nil)
	hf(*writer2, nil)
	for i, handler := range h {
		if handler.HandleWriters[0] != *writer1 {
			t.Error("Did not get correct value for first writer on handler", i)
		}
		if handler.HandleWriters[1] != *writer2 {
			t.Error("Did not get correct value for second writer on handler", i)
		}
	}
}

func Test_Router_HandleFunc_HandlerReceives_CorrectUser(t *testing.T) {
	var r, _, h, a = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	a.AuthenticatedUser.FirstName = "Kyle"
	hf(nil, nil)
	hf(nil, nil)
	for i, handler := range h {
		if handler.HandleUsers[0].FirstName != a.AuthenticatedUser.FirstName {
			t.Error("Authenticated user for first call first name was", handler.HandleUsers[0].FirstName,
				"expected", a.AuthenticatedUser.FirstName, "for handler", i)
		}
		if handler.HandleUsers[1].FirstName != a.AuthenticatedUser.FirstName {
			t.Error("Authenticated user for first call first name was", handler.HandleUsers[1].FirstName,
				"expected", a.AuthenticatedUser.FirstName, "for handler", i)
		}
	}
}

func Test_Router_HandleFunc_HandlerShould_CorrectUser(t *testing.T) {
	var r, _, h, a = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	a.AuthenticatedUser.FirstName = "Kyle"
	hf(nil, nil)
	hf(nil, nil)
	for i, handler := range h {
		if handler.ShouldUsers[0].FirstName != a.AuthenticatedUser.FirstName {
			t.Error("Authenticated user for first call first name was", handler.ShouldUsers[0].FirstName,
				"expected", a.AuthenticatedUser.FirstName, "for handler", i)
		}
		if handler.ShouldUsers[1].FirstName != a.AuthenticatedUser.FirstName {
			t.Error("Authenticated user for second call first name was", handler.ShouldUsers[1].FirstName,
				"expected", a.AuthenticatedUser.FirstName, "for handler", i)
		}
	}
}

func Test_Router_HandleFunc_HandlerShould_CorrectRequest(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	request1 := http.Request{}
	request2 := http.Request{}
	hf(nil, &request1)
	hf(nil, &request2)
	for i, handler := range h {
		if handler.ShouldRequests[0] != &request1 {
			t.Error("Did not get correct value for first request on handler", i)
		}
		if handler.ShouldRequests[1] != &request2 {
			t.Error("Did not get correct value for second request on handler", i)
		}
	}
}

func Test_Router_HandleFunc_AuthCorrectCount(t *testing.T) {
	var r, _, _, a = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)
	if len(a.CallRequests) != 3 {
		t.Error("For auth calls expected 3 got", len(a.CallRequests))
	}
}

func Test_Router_HandleFunc_AuthCorrectRequest(t *testing.T) {
	var r, _, _, a = getMockedRouter(4)
	var hf = r.GetHandleFunc()
	request1 := http.Request{}
	request2 := http.Request{}
	hf(nil, &request1)
	hf(nil, &request2)
	if a.CallRequests[0] != &request1 {
		t.Error("Request to first call of auth is incrrect.")
	}

	if a.CallRequests[1] != &request2 {
		t.Error("Request to second call of auth is incrrect.")
	}
}
