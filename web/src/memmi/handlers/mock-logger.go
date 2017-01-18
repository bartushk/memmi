package handlers

import (
	"net/http"
)

type MockLogger struct {
	CallCount int
}

func (logger *MockLogger) Log(r *http.Request) {
	logger.CallCount += 1
}
