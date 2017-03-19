package request

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"memmi/pbuf"
	"net/http"
)

type ProtoIoImpl struct {
}

func (io *ProtoIoImpl) WriteCodedProtoResponse(w http.ResponseWriter, message proto.Message, statusCode int) error {
	w.WriteHeader(statusCode)
	return io.WriteProtoResponse(w, message)
}

func (io *ProtoIoImpl) WriteProtoResponse(w http.ResponseWriter, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	_, wErr := w.Write(data)
	return wErr
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
