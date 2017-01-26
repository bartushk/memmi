package request

import (
	"github.com/golang/protobuf/proto"
	"memmi/pbuf"
	"net/http"
)

type MockProtoIO struct {
	MessageWrites   []*proto.Message
	ResponseWriters []http.ResponseWriter
	WriteReturn     error

	NextCardRequests []*http.Request
	NextCardReturn   pbuf.NextCardRequest
	NextCardError    error

	ReportRequests []*http.Request
	ReportReturn   pbuf.CardScoreReport
	ReportError    error

	ReportNextRequests []*http.Request
	ReportNextReturn   pbuf.ReportAndNext
	ReportNextError    error
}

func (io *MockProtoIO) WriteProtoResponse(w http.ResponseWriter, message *proto.Message) error {
	io.ResponseWriters = append(io.ResponseWriters, w)
	io.MessageWrites = append(io.MessageWrites, message)
	return io.WriteReturn
}

func (io *MockProtoIO) ReadNextCardRequest(r *http.Request) (pbuf.NextCardRequest, error) {
	io.NextCardRequests = append(io.NextCardRequests, r)
	return io.NextCardReturn, io.NextCardError
}

func (io *MockProtoIO) ReadCardScoreReport(r *http.Request) (pbuf.CardScoreReport, error) {
	io.ReportRequests = append(io.ReportRequests, r)
	return io.ReportReturn, io.ReportError
}

func (io *MockProtoIO) ReadReportAndNext(r *http.Request) (pbuf.ReportAndNext, error) {
	io.ReportNextRequests = append(io.ReportNextRequests, r)
	return io.ReportNextReturn, io.ReportNextError
}
