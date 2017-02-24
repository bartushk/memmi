package request

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"memmi/pbuf"
	"net/http"
)

type ProtoIoImpl struct {
}

func (io *ProtoIoImpl) WriteProtoResponse(w http.ResponseWriter, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	_, wErr := w.Write(data)
	return wErr
}

func (io *ProtoIoImpl) ReadNextCardRequest(r *http.Request) (pbuf.NextCardRequest, error) {
	var retNextCardRequest pbuf.NextCardRequest
	data, readErr := ioutil.ReadAll(r.Body)
	if len(data) == 0 {
		fmt.Println(readErr)
		return retNextCardRequest, errors.New("Cannot read from empty request body.")
	}
	if readErr != nil {
		return retNextCardRequest, readErr
	}
	marshalError := proto.Unmarshal(data, &retNextCardRequest)
	return retNextCardRequest, marshalError
}

func (io *ProtoIoImpl) ReadCardScoreReport(r *http.Request) (pbuf.CardScoreReport, error) {
	var retReport pbuf.CardScoreReport
	data, readErr := ioutil.ReadAll(r.Body)
	if len(data) == 0 {
		return retReport, errors.New("Cannot read from empty request body.")
	}
	if readErr != nil {
		return retReport, readErr
	}
	marshalError := proto.Unmarshal(data, &retReport)
	return retReport, marshalError
}

func (io *ProtoIoImpl) ReadReportAndNext(r *http.Request) (pbuf.ReportAndNext, error) {
	var retReportNext pbuf.ReportAndNext
	data, readErr := ioutil.ReadAll(r.Body)
	if len(data) == 0 {
		return retReportNext, errors.New("Cannot read from empty request body.")
	}
	if readErr != nil {
		return retReportNext, readErr
	}
	marshalError := proto.Unmarshal(data, &retReportNext)
	return retReportNext, marshalError
}

func (io *ProtoIoImpl) ReadCardSetRequest(r *http.Request) (pbuf.CardSetRequest, error) {
	var retCardSetRequest pbuf.CardSetRequest
	data, readErr := ioutil.ReadAll(r.Body)
	if len(data) == 0 {
		return retCardSetRequest, errors.New("Cannot read from empty request body.")
	}
	if readErr != nil {
		return retCardSetRequest, readErr
	}
	marshalError := proto.Unmarshal(data, &retCardSetRequest)
	return retCardSetRequest, marshalError
}

func (io *ProtoIoImpl) ReadCardRequest(r *http.Request) (pbuf.CardRequest, error) {
	var retCardRequest pbuf.CardRequest
	data, readErr := ioutil.ReadAll(r.Body)
	if len(data) == 0 {
		return retCardRequest, errors.New("Cannot read from empty request body.")
	}
	if readErr != nil {
		return retCardRequest, readErr
	}
	marshalError := proto.Unmarshal(data, &retCardRequest)
	return retCardRequest, marshalError
}
