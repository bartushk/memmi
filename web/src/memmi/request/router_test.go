package request

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
	"net/http"
	"testing"
)

func getMockedRouter(i, skip int) (Router, *MockLogger, []*MockHandler, *MockAuthenticator) {
	var handlers []*MockHandler
	logger := &MockLogger{}
	logger.On("Log", mock.Anything).Return()
	auth := &MockAuthenticator{}
	auth.On("AuthenticateUser", mock.Anything).Return(pbuf.User{})
	router := Router{
		Logger:        logger,
		Authenticator: auth,
	}
	for k := 0; k < i; k++ {
		handler := &MockHandler{}
		if k != skip {
			hRes := HandleResult{Continue: true, ResponseWritten: false}
			handler.On("Handle", mock.Anything, mock.Anything, mock.Anything).Return(hRes)
			handler.On("ShouldHandle", mock.Anything, mock.Anything, mock.Anything).Return(true)
		}
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
	h1 := &MockHandler{}
	h2 := &MockHandler{}
	h3 := &MockHandler{}
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

func Test_Router_AddHandler_MultipleCopies(t *testing.T) {
	router := Router{}
	h1 := &MockHandler{}
	h2 := &MockHandler{}
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
	var r, _, _, _ = getMockedRouter(0, -1)
	var hf = r.GetHandleFunc()
	if hf == nil {
		t.Error("Handle func was nil.")
	}
}

func Test_Router_HandleFunc_NoHandlers_LoggerCalled(t *testing.T) {
	var r, l, _, _ = getMockedRouter(0, -1)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	l.AssertCalled(t, "Log", mock.Anything)
}

func Test_Router_HandleFunc_Handlers_LoggerCalled(t *testing.T) {
	var r, l, _, _ = getMockedRouter(3, -1)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	l.AssertNumberOfCalls(t, "Log", 2)
}

func Test_Router_HandleFunc_LoggerCalled_RightRequest(t *testing.T) {
	var r, l, _, _ = getMockedRouter(3, -1)
	var hf = r.GetHandleFunc()
	request1 := http.Request{}
	request2 := http.Request{}
	hf(nil, &request1)
	hf(nil, &request2)

	l.AssertNumberOfCalls(t, "Log", 2)
	l.AssertCalled(t, "Log", &request1)
	l.AssertCalled(t, "Log", &request2)
}

func Test_Router_HandleFunc_ShouldGo_AllCalled(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, -1)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)
	for _, handler := range h {
		handler.AssertNumberOfCalls(t, "Handle", 3)
	}
}

func Test_Router_HandleFunc_ShouldNotContinue_LoopBreaks(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, 1)
	hRes := HandleResult{Continue: false, ResponseWritten: false}
	h[1].On("Handle", mock.Anything, mock.Anything, mock.Anything).Return(hRes)
	h[1].On("ShouldHandle", mock.Anything, mock.Anything, mock.Anything).Return(true)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)

	for _, handler := range h[:1] {
		handler.AssertNumberOfCalls(t, "Handle", 3)
	}

	for _, handler := range h[2:] {
		handler.AssertNotCalled(t, "Handle")
	}
}

func Test_Router_HandleFunc_ResponseWritten_PassedCorrectly(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, 1)
	hRes := HandleResult{Continue: true, ResponseWritten: true}
	h[1].On("Handle", mock.Anything, mock.Anything, mock.Anything).Return(hRes)
	h[1].On("ShouldHandle", mock.Anything, mock.Anything, mock.Anything).Return(true)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)

	for _, handler := range h[:1] {
		handler.AssertCalled(t, "ShouldHandle", mock.Anything, mock.Anything, false)
	}

	for _, handler := range h[2:] {
		handler.AssertCalled(t, "ShouldHandle", mock.Anything, mock.Anything, true)
	}
}

func Test_Router_HandleFunc_ShouldNotHandle_DoesNotHandle(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, 0)
	hRes := HandleResult{Continue: true, ResponseWritten: false}
	h[0].On("Handle", mock.Anything, mock.Anything, mock.Anything).Return(hRes)
	h[0].On("ShouldHandle", mock.Anything, mock.Anything, mock.Anything).Return(false)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)

	h[0].AssertNumberOfCalls(t, "Handle", 0)

	for _, handler := range h[1:] {
		handler.AssertNumberOfCalls(t, "Handle", 3)
	}
}

func Test_Router_HandleFunc_HandlerReceives_CorrectRequest(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, -1)
	var hf = r.GetHandleFunc()
	request1 := http.Request{}
	request2 := http.Request{}
	hf(nil, &request1)
	hf(nil, &request2)
	for _, handler := range h {
		handler.AssertNumberOfCalls(t, "Handle", 2)
		handler.AssertCalled(t, "Handle", mock.Anything, &request1, mock.Anything)
		handler.AssertCalled(t, "Handle", mock.Anything, &request2, mock.Anything)
	}
}

func Test_Router_HandleFunc_HandlerReceives_CorrectWriter(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, -1)
	var hf = r.GetHandleFunc()
	writer1 := new(http.ResponseWriter)
	writer2 := new(http.ResponseWriter)
	hf(*writer1, nil)
	hf(*writer2, nil)
	for _, handler := range h {
		handler.AssertNumberOfCalls(t, "Handle", 2)
		handler.AssertCalled(t, "Handle", *writer1, mock.Anything, mock.Anything)
		handler.AssertCalled(t, "Handle", *writer2, mock.Anything, mock.Anything)
	}
}

func Test_Router_HandleFunc_HandlerReceives_CorrectUser(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, -1)
	var hf = r.GetHandleFunc()
	testUser := pbuf.User{FirstName: "Kyle"}
	testAuth := &MockAuthenticator{}
	testAuth.On("AuthenticateUser", mock.Anything).Return(testUser)
	r.Authenticator = testAuth
	hf(nil, nil)
	hf(nil, nil)
	for _, handler := range h {
		handler.AssertNumberOfCalls(t, "Handle", 2)
		handler.AssertCalled(t, "Handle", mock.Anything, mock.Anything,
			mock.MatchedBy(func(u pbuf.User) bool { return proto.Equal(&u, &testUser) }))

	}
}

func Test_Router_HandleFunc_HandlerShould_CorrectUser(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, -1)
	var hf = r.GetHandleFunc()
	testUser := pbuf.User{FirstName: "Kyle"}
	testAuth := &MockAuthenticator{}
	testAuth.On("AuthenticateUser", mock.Anything).Return(testUser)
	r.Authenticator = testAuth
	hf(nil, nil)
	hf(nil, nil)
	for _, handler := range h {
		handler.AssertNumberOfCalls(t, "ShouldHandle", 2)
		handler.AssertCalled(t, "ShouldHandle", mock.Anything, mock.MatchedBy(func(u pbuf.User) bool {
			return proto.Equal(&u, &testUser)
		}), mock.Anything)
	}
}

func Test_Router_HandleFunc_HandlerShould_CorrectRequest(t *testing.T) {
	var r, _, h, _ = getMockedRouter(4, -1)
	var hf = r.GetHandleFunc()
	request1 := &http.Request{}
	request2 := &http.Request{}
	hf(nil, request1)
	hf(nil, request2)
	for _, handler := range h {
		handler.AssertNumberOfCalls(t, "ShouldHandle", 2)
		handler.AssertCalled(t, "ShouldHandle", request2, mock.Anything, mock.Anything)
		handler.AssertCalled(t, "ShouldHandle", request1, mock.Anything, mock.Anything)
	}
}

func Test_Router_HandleFunc_AuthCorrectCount(t *testing.T) {
	var r, _, _, a = getMockedRouter(4, -1)
	var hf = r.GetHandleFunc()
	hf(nil, nil)
	hf(nil, nil)
	hf(nil, nil)
	a.AssertNumberOfCalls(t, "AuthenticateUser", 3)
}

func Test_Router_HandleFunc_AuthCorrectRequest(t *testing.T) {
	var r, _, _, a = getMockedRouter(4, -1)
	var hf = r.GetHandleFunc()
	request1 := &http.Request{}
	request2 := &http.Request{}
	hf(nil, request1)
	hf(nil, request2)
	a.AssertCalled(t, "AuthenticateUser", request1)
	a.AssertCalled(t, "AuthenticateUser", request2)
}
