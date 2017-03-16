package request

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
	"memmi/user"
	"testing"
)

func getMockedHandler() (RequestHandler, *MockProtoIO, *user.MockUserManagement) {
	handler := &ReportRequestHandler{}
	pio := &MockProtoIO{}
	pio.On("WriteProtoResponse", mock.Anything, mock.Anything).Return(nil)
	userMan := &user.MockUserManagement{}
	handler.Pio = pio
	handler.UserMan = userMan
	return handler, pio, userMan
}

func Test_CardHandler_ExactUrl_ShouldHanlde(t *testing.T) {
	var req = RequestFromURL(REPORT_API_URL)
	handler := ReportRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle request with URL:", REPORT_API_URL)
	}
}

func Test_CardHandler_ExactUrl_IfPreviouslyWritten_ShouldNotHandle(t *testing.T) {
	var req = RequestFromURL(REPORT_API_URL)
	handler := ReportRequestHandler{}
	if handler.ShouldHandle(req, pbuf.User{}, true) {
		t.Error("Handler should not handle request with responseWritten true.")
	}
}

func Test_CardHandler_UrlPlusQuery_ShouldHanlde(t *testing.T) {
	test_url := REPORT_API_URL + "?asdf"
	var req = RequestFromURL(test_url)
	handler := ReportRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle request with URL:", test_url)
	}
}

func Test_CardHandler_UrlSubUrl_ShouldHanlde(t *testing.T) {
	test_url := REPORT_API_URL + "/asdf/fffa"
	var req = RequestFromURL(test_url)
	handler := ReportRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle request with URL:", test_url)
	}
}

func Test_CardHandler_AnyDifferentPrefix_ShouldNotHanlde(t *testing.T) {
	test_url := "/tes" + REPORT_API_URL
	var req = RequestFromURL(test_url)
	handler := ReportRequestHandler{}
	if handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should not handle request with URL:", test_url)
	}
}

func Test_ReportHandler_Report_ProtoReadError_WriteError(t *testing.T) {
	var handler, pio, um = getMockedHandler()
	var req = RequestFromURL(SCORE_REPORT_URL)
	pio.On("ReadCardScoreReport", req).Return(pbuf.CardScoreReport{}, errors.New(""))

	handler.Handle(nil, req, pbuf.User{})

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, BODY_READ_ERROR)
	}))

	um.AssertNotCalled(t, "UpdateHistory", mock.Anything)
	um.AssertNotCalled(t, "GetHistory", mock.Anything)
}

func Test_ReportHandler_Report_UpdateError_WriteError(t *testing.T) {
	var handler, pio, um = getMockedHandler()
	var req = RequestFromURL(SCORE_REPORT_URL)
	testUser := pbuf.User{Id: "userId"}
	testUpdate := pbuf.CardUpdate{CardId: "cardId"}
	testCardReport := pbuf.CardScoreReport{CardSetId: "setId"}
	testCardReport.Update = &testUpdate

	um.On("UpdateHistory", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	pio.On("ReadCardScoreReport", req).Return(testCardReport, nil)

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, USER_HISTORY_UPDATE_ERROR)
	}))

	um.AssertCalled(t, "UpdateHistory", mock.MatchedBy(func(u pbuf.User) bool {
		return proto.Equal(&u, &testUser)
	}), testCardReport.CardSetId, mock.MatchedBy(func(up pbuf.CardUpdate) bool {
		return proto.Equal(&up, &testUpdate)
	}))
}

func Test_ReportHandler_Report_Success_HandledCorrectly(t *testing.T) {
	var handler, pio, um = getMockedHandler()
	var req = RequestFromURL(SCORE_REPORT_URL)
	expectedResponseWrite := pbuf.UpdateResponse{Status: 1}
	testUser := pbuf.User{Id: "userId"}
	testUpdate := pbuf.CardUpdate{CardId: "cardId"}
	testCardReport := pbuf.CardScoreReport{CardSetId: "setId"}
	testCardReport.Update = &testUpdate

	um.On("UpdateHistory", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	pio.On("ReadCardScoreReport", req).Return(testCardReport, nil)

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, &expectedResponseWrite)
	}))

	um.AssertCalled(t, "UpdateHistory", mock.MatchedBy(func(u pbuf.User) bool {
		return proto.Equal(&u, &testUser)
	}), testCardReport.CardSetId, mock.MatchedBy(func(up pbuf.CardUpdate) bool {
		return proto.Equal(&up, &testUpdate)
	}))
}
