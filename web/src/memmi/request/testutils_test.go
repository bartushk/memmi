package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func RequestFromURL(url_string string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, url_string, nil)
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte{}))
	return req
}

type MockResponseWriter struct {
	ReturnHeader      http.Header
	WriteHeaderInputs []int

	WriteReturn int
	WriteError  error
	WriteInputs [][]byte
}

func (writer *MockResponseWriter) Header() http.Header {
	return writer.ReturnHeader
}

func (writer *MockResponseWriter) Write(toWrite []byte) (int, error) {
	writer.WriteInputs = append(writer.WriteInputs, toWrite)
	return writer.WriteReturn, writer.WriteError
}

func (writer *MockResponseWriter) WriteHeader(toWrite int) {
	writer.WriteHeaderInputs = append(writer.WriteHeaderInputs, toWrite)
}
