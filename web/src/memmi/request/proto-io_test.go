package request

import (
	"bytes"
	"errors"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"memmi/pbuf"
	"testing"
)

// Write proto response tests

func Test_ProtoIO_WriteProtoResponse_BadMessage_ReturnsError(t *testing.T) {
	pio := ProtoIoImpl{}
	writer := &MockResponseWriter{}
	err := pio.WriteProtoResponse(writer, nil)
	if err == nil {
		t.Error("Should have received error.")
	}
}

func Test_ProtoIO_WriteProtoResponse_BadWrite_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	writer := &MockResponseWriter{}
	testMessage := &pbuf.Card{}
	testError := errors.New("Test Error")

	writer.WriteError = testError
	err := pio.WriteProtoResponse(writer, testMessage)
	if err != testError {
		t.Error("Wrong error returned. Expected:", testError,
			"Received:", err)
	}
}

func Test_ProtoIO_WriteProtoResponse_GoodMessage_WrittenCorrectly(t *testing.T) {
	pio := ProtoIoImpl{}
	writer := &MockResponseWriter{}
	testMessage := &pbuf.Card{Title: "Card Title"}
	testData, _ := proto.Marshal(testMessage)

	err := pio.WriteProtoResponse(writer, testMessage)
	if err != nil {
		t.Error("Should not have returned an error.")
	}

	if len(writer.WriteInputs) != 1 {
		t.Fatal("Expeceted one write to responsewriter got:", len(writer.WriteInputs))
	}

	if !CompareByteSlices(writer.WriteInputs[0], testData) {
		t.Error("Did not write the correct data. Expected:", testData,
			"Received:", writer.WriteInputs[0])
	}
}

// Next Card Tests

func Test_ProtoIO_ReadNextCardRequest_EmptyRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	_, err := pio.ReadNextCardRequest(req)
	if err == nil {
		t.Error("Expected error from empty request body")
	}
}

func Test_ProtoIO_ReadNextCardRequest_BadRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte{2, 3, 5}))
	_, err := pio.ReadNextCardRequest(req)
	if err == nil {
		t.Error("Expected error from bad request type.")
	}
}

func Test_ProtoIO_ReadNextCardRequest_GoodRequestBody_ReadCorrectly(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	goodMessage := &pbuf.NextCardRequest{CardSetId: []byte{5, 10, 3}}
	goodData, _ := proto.Marshal(goodMessage)
	req.Body = ioutil.NopCloser(bytes.NewReader(goodData))
	result, err := pio.ReadNextCardRequest(req)
	if err != nil {
		t.Error("Expected no error.")
	}

	if !CompareByteSlices(goodMessage.CardSetId, result.CardSetId) {
		t.Error("Did not read request body correctly. Expected:", goodMessage,
			"Received:", result)
	}
}

// Get Card Tests

func Test_ProtoIO_ReadCardRequest_EmptyRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	_, err := pio.ReadCardRequest(req)
	if err == nil {
		t.Error("Expected error from empty request body")
	}
}

func Test_ProtoIO_ReadCardRequest_BadRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte{2, 3, 5}))
	_, err := pio.ReadCardRequest(req)
	if err == nil {
		t.Error("Expected error from bad request type")
	}
}

func Test_ProtoIO_ReadCardRequest_GoodRequestBody_ReadCorrectly(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	goodMessage := &pbuf.CardRequest{Id: []byte{2, 3, 6}}
	goodData, _ := proto.Marshal(goodMessage)
	req.Body = ioutil.NopCloser(bytes.NewReader(goodData))
	result, err := pio.ReadCardRequest(req)
	if err != nil {
		t.Error("Expected no error.")
	}

	if !CompareByteSlices(goodMessage.Id, result.Id) {
		t.Error("Did not read request body correctly. Expected:", goodMessage,
			"Received:", result)
	}
}

// Get Card Set Tests

func Test_ProtoIO_ReadCardSetRequest_EmptyRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	_, err := pio.ReadCardSetRequest(req)
	if err == nil {
		t.Error("Expected error from empty request body")
	}
}

func Test_ProtoIO_ReadCardSetRequest_BadRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte{2, 3, 5}))
	_, err := pio.ReadCardSetRequest(req)
	if err == nil {
		t.Error("Expected error from bad request type")
	}
}

func Test_ProtoIO_ReadCardSetRequest_GoodRequestBody_ReadCorrectly(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	goodMessage := &pbuf.CardSetRequest{Id: []byte{2, 3, 6}}
	goodData, _ := proto.Marshal(goodMessage)
	req.Body = ioutil.NopCloser(bytes.NewReader(goodData))
	result, err := pio.ReadCardSetRequest(req)
	if err != nil {
		t.Error("Expected no error.")
	}

	if !CompareByteSlices(goodMessage.Id, result.Id) {
		t.Error("Did not read request body correctly. Expected:", goodMessage,
			"Received:", result)
	}
}

// Report and Next Tests

func Test_ProtoIO_ReadReportAndNext_EmptyRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	_, err := pio.ReadReportAndNext(req)
	if err == nil {
		t.Error("Expected error from empty request body")
	}
}

func Test_ProtoIO_ReadReportAndNext_BadRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte{2, 3, 5}))
	_, err := pio.ReadReportAndNext(req)
	if err == nil {
		t.Error("Expected error from bad request type")
	}
}

func Test_ProtoIO_ReadReportAndNext_GoodRequestBody_ReadCorrectly(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	goodMessage := &pbuf.ReportAndNext{}
	goodNextRequest := &pbuf.NextCardRequest{CardSetId: []byte{8, 78, 3}}
	goodMessage.NextRequest = goodNextRequest
	goodData, _ := proto.Marshal(goodMessage)
	req.Body = ioutil.NopCloser(bytes.NewReader(goodData))
	result, err := pio.ReadReportAndNext(req)
	if err != nil {
		t.Error("Expected no error.")
	}

	if !CompareByteSlices(goodMessage.NextRequest.CardSetId, result.NextRequest.CardSetId) {
		t.Error("Did not read request body correctly. Expected:", goodMessage,
			"Received:", result)
	}
}

// Report Card Score Tests

func Test_ProtoIO_ReadCardScoreReport_EmptyRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	_, err := pio.ReadCardScoreReport(req)
	if err == nil {
		t.Error("Expected error from empty request body")
	}
}

func Test_ProtoIO_ReadCardScoreReport_BadRequestBody_GetError(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte{2, 3, 5}))
	_, err := pio.ReadCardScoreReport(req)
	if err == nil {
		t.Error("Expected error from bad request type")
	}
}

func Test_ProtoIO_ReadCardScoreReport_GoodRequestBody_ReadCorrectly(t *testing.T) {
	pio := ProtoIoImpl{}
	req := RequestFromURL("sadf")
	goodMessage := &pbuf.CardScoreReport{CardSetId: []byte{66, 23, 12}}
	goodData, _ := proto.Marshal(goodMessage)
	req.Body = ioutil.NopCloser(bytes.NewReader(goodData))
	result, err := pio.ReadCardScoreReport(req)
	if err != nil {
		t.Error("Expected no error.")
	}

	if !CompareByteSlices(goodMessage.CardSetId, result.CardSetId) {
		t.Error("Did not read request body correctly. Expected:", goodMessage,
			"Received:", result)
	}
}