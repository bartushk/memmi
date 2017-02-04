package request

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"memmi/pbuf"
	"net/http"
)

type ProtoIoImpl struct {
}

func (io *ProtoIoImpl) WriteProtoResponse(w http.ResponseWriter, message proto.Message) error {
	// data, err := proto.Marshal(test)
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	_, wErr := w.Write(data)
	return wErr
}

func (io *ProtoIoImpl) ReadNextCardRequest(r *http.Request) (pbuf.NextCardRequest, error) {
	var retCardRequest pbuf.NextCardRequest
	_, err := ioutil.ReadAll(r.Body)
	return retCardRequest, err
}

func (io *ProtoIoImpl) ReadCardScoreReport(r *http.Request) (pbuf.CardScoreReport, error) {
	var retReport pbuf.CardScoreReport
	return retReport, nil
}

func (io *ProtoIoImpl) ReadReportAndNext(r *http.Request) (pbuf.ReportAndNext, error) {
	var retReportNext pbuf.ReportAndNext
	return retReportNext, nil
}

func (io *ProtoIoImpl) ReadCardSetRequest(r *http.Request) (pbuf.CardSetRequest, error) {
	var retCardSetRequest pbuf.CardSetRequest
	return retCardSetRequest, nil
}

func (io *ProtoIoImpl) ReadCardRequest(r *http.Request) (pbuf.CardRequest, error) {
	var retCardRequest pbuf.CardRequest
	return retCardRequest, nil
}
