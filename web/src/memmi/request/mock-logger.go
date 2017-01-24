package request

import (
	"net/http"
)

type MockLogger struct {
	CallRequests []*http.Request
}

func (logger *MockLogger) Log(r *http.Request) {
	logger.CallRequests = append(logger.CallRequests, r)
}
