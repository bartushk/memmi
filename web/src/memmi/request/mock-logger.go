package request

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockLogger struct {
	mock.Mock
}

func (logger *MockLogger) Log(r *http.Request) {
	logger.Called(r)
}
