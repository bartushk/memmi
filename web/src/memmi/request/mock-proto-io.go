package request

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
	"net/http"
)

type MockProtoIO struct {
	mock.Mock
}

func (io *MockProtoIO) WriteProtoResponse(w http.ResponseWriter, message proto.Message) error {
	args := io.Called(w, message)
	return args.Error(0)
}

func (io *MockProtoIO) ReadNextCardRequest(r *http.Request) (pbuf.NextCardRequest, error) {
	args := io.Called(r)
	return args.Get(0).(pbuf.NextCardRequest), args.Error(1)
}

func (io *MockProtoIO) ReadCardScoreReport(r *http.Request) (pbuf.CardScoreReport, error) {
	args := io.Called(r)
	return args.Get(0).(pbuf.CardScoreReport), args.Error(1)
}

func (io *MockProtoIO) ReadReportAndNext(r *http.Request) (pbuf.ReportAndNext, error) {
	args := io.Called(r)
	return args.Get(0).(pbuf.ReportAndNext), args.Error(1)
}

func (io *MockProtoIO) ReadCardSetRequest(r *http.Request) (pbuf.CardSetRequest, error) {
	args := io.Called(r)
	return args.Get(0).(pbuf.CardSetRequest), args.Error(1)
}

func (io *MockProtoIO) ReadCardRequest(r *http.Request) (pbuf.CardRequest, error) {
	args := io.Called(r)
	return args.Get(0).(pbuf.CardRequest), args.Error(1)
}
