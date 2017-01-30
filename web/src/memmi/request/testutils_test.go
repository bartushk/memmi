package request

import (
	"net/http"
	"net/url"
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
	test_url := &url.URL{}
	test_request := &http.Request{}
	test_url.Path = url_string
	test_request.URL = test_url
	return test_request
}
