package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func CompareByteSlices(first []byte, second []byte) bool {
	if len(first) != len(second) {
		return false
	}
	for i, _ := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

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
