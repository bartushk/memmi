package request

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/mock"
	"memmi/card"
	"memmi/pbuf"
	"memmi/user"
	"testing"
)

func getMockedHandler() (RequestHandler, *MockProtoIO, *card.MockCardSelection, *user.MockUserManagement, *card.MockCardManagement) {
	handler := &CardRequestHandler{}
	pio := &MockProtoIO{}
	pio.On("WriteProtoResponse", mock.Anything, mock.Anything).Return(nil)
	cardSel := &card.MockCardSelection{}
	cardMan := &card.MockCardManagement{}
	userMan := &user.MockUserManagement{}
	handler.Pio = pio
	handler.CardSel = cardSel
	handler.UserMan = userMan
	handler.CardMan = cardMan
	return handler, pio, cardSel, userMan, cardMan
}

func Test_CardHandler_ExactUrl_ShouldHanlde(t *testing.T) {
	var req = RequestFromURL(CARD_API_URL)
	handler := CardRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle request with URL:", CARD_API_URL)
	}
}

func Test_CardHandler_ExactUrl_IfPreviouslyWritten_ShouldNotHandle(t *testing.T) {
	var req = RequestFromURL(CARD_API_URL)
	handler := CardRequestHandler{}
	if handler.ShouldHandle(req, pbuf.User{}, true) {
		t.Error("Handler should not handle request with responseWritten true.")
	}
}

func Test_CardHandler_UrlPlusQuery_ShouldHanlde(t *testing.T) {
	test_url := CARD_API_URL + "?asdf"
	var req = RequestFromURL(test_url)
	handler := CardRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle request with URL:", test_url)
	}
}

func Test_CardHandler_UrlSubUrl_ShouldHanlde(t *testing.T) {
	test_url := CARD_API_URL + "/asdf/fffa"
	var req = RequestFromURL(test_url)
	handler := CardRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle request with URL:", test_url)
	}
}

func Test_CardHandler_AnyDifferentPrefix_ShouldNotHanlde(t *testing.T) {
	test_url := "/tes" + CARD_API_URL
	var req = RequestFromURL(test_url)
	handler := CardRequestHandler{}
	if handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should not handle request with URL:", test_url)
	}
}

func Test_CardHandler_HandleNext_ProtoReadError_WriteError(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_NEXT_URL)

	pio.On("ReadNextCardRequest", req).Return(pbuf.NextCardRequest{}, errors.New(""))

	handler.Handle(nil, req, pbuf.User{})

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, BODY_READ_ERROR)
	}))

	um.AssertNotCalled(t, "UpdateHistory", mock.Anything)
	um.AssertNotCalled(t, "GetHistory", mock.Anything)
	cs.AssertNotCalled(t, "SelectCard", mock.Anything)
	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardHandler_HandleNext_ProtoReadGood_HandledCorrectly(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_NEXT_URL)
	testUser := pbuf.User{UserName: "bartushk"}
	nextCardRequest := pbuf.NextCardRequest{CardSetId: "setId", PreviousCardId: "prevCard"}
	nextCard := pbuf.Card{Title: "TestCard"}
	testHistory := pbuf.UserHistory{PlayIndex: 123}
	csReturn := "setId"

	cm.On("GetCardById", mock.Anything).Return(nextCard, nil)
	cs.On("SelectCard", mock.Anything, mock.Anything).Return(csReturn)
	um.On("GetHistory", mock.Anything, mock.Anything).Return(testHistory, nil)
	pio.On("ReadNextCardRequest", req).Return(nextCardRequest, nil)

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, &nextCard)
	}))

	um.AssertCalled(t, "GetHistory", mock.MatchedBy(func(u pbuf.User) bool {
		return proto.Equal(&u, &testUser)
	}), nextCardRequest.CardSetId)

	cs.AssertCalled(t, "SelectCard", mock.MatchedBy(func(h *pbuf.UserHistory) bool {
		return proto.Equal(h, &testHistory)
	}), nextCardRequest.PreviousCardId)

	cm.AssertCalled(t, "GetCardById", csReturn)
}

func Test_CardHandler_Report_ProtoReadError_WriteError(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_URL)
	pio.On("ReadCardScoreReport", req).Return(pbuf.CardScoreReport{}, errors.New(""))

	handler.Handle(nil, req, pbuf.User{})

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, BODY_READ_ERROR)
	}))

	um.AssertNotCalled(t, "UpdateHistory", mock.Anything)
	um.AssertNotCalled(t, "GetHistory", mock.Anything)
	cs.AssertNotCalled(t, "SelectCard", mock.Anything)
	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardHandler_Report_UpdateError_WriteError(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_URL)
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

	cs.AssertNotCalled(t, "SelectCard", mock.Anything)
	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardHandler_Report_Success_HandledCorrectly(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_URL)
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

	cs.AssertNotCalled(t, "SelectCard", mock.Anything)
	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardHandler_ReportNext_ProtoIO_ErrorWritten(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_NEXT_URL)
	pio.On("ReadReportAndNext", req).Return(pbuf.ReportAndNext{}, errors.New(""))

	handler.Handle(nil, req, pbuf.User{})

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, BODY_READ_ERROR)
	}))

	um.AssertNotCalled(t, "UpdateHistory", mock.Anything)
	um.AssertNotCalled(t, "GetHistory", mock.Anything)
	cs.AssertNotCalled(t, "SelectCard", mock.Anything)
	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardHandler_ReportNext_HandledCorrectly(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_NEXT_URL)
	testUpdate := pbuf.CardUpdate{CardId: "cardId"}
	reportAndNext := pbuf.ReportAndNext{
		NextRequest: &pbuf.NextCardRequest{CardSetId: "setId", PreviousCardId: "cardId"},
		Report:      &pbuf.CardScoreReport{CardSetId: "setId", Update: &testUpdate},
	}
	nextCard := pbuf.Card{Title: "TestCard"}
	testHistory := pbuf.UserHistory{PlayIndex: 123}
	testUser := pbuf.User{Id: "suerId"}
	csReturn := "setId"

	um.On("UpdateHistory", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	um.On("GetHistory", mock.Anything, mock.Anything).Return(testHistory, nil)
	cm.On("GetCardById", mock.Anything).Return(nextCard, nil)
	cs.On("SelectCard", mock.Anything, mock.Anything).Return(csReturn)
	pio.On("ReadReportAndNext", req).Return(reportAndNext, nil)

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, &nextCard)
	}))

	um.AssertCalled(t, "UpdateHistory", mock.MatchedBy(func(u pbuf.User) bool {
		return proto.Equal(&u, &testUser)
	}), reportAndNext.NextRequest.CardSetId, mock.MatchedBy(func(up pbuf.CardUpdate) bool {
		return proto.Equal(&up, &testUpdate)
	}))

	um.AssertCalled(t, "GetHistory", mock.MatchedBy(func(u pbuf.User) bool {
		return proto.Equal(&u, &testUser)
	}), reportAndNext.NextRequest.CardSetId)

	cs.AssertCalled(t, "SelectCard", mock.MatchedBy(func(h *pbuf.UserHistory) bool {
		return proto.Equal(h, &testHistory)
	}), reportAndNext.NextRequest.PreviousCardId)

	cm.AssertCalled(t, "GetCardById", csReturn)
}

func Test_CardHandler_ReportNext_WithUpdateError_ErrorSilent(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_NEXT_URL)
	testUpdate := pbuf.CardUpdate{CardId: "cardId"}
	reportAndNext := pbuf.ReportAndNext{
		NextRequest: &pbuf.NextCardRequest{CardSetId: "setId", PreviousCardId: "cardId"},
		Report:      &pbuf.CardScoreReport{CardSetId: "setId", Update: &testUpdate},
	}
	nextCard := pbuf.Card{Title: "TestCard"}
	testUser := pbuf.User{Id: "userId"}
	testHistory := pbuf.UserHistory{PlayIndex: 123}
	csReturn := "setId"

	um.On("GetHistory", mock.Anything, mock.Anything).Return(testHistory, nil)
	um.On("UpdateHistory", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	cm.On("GetCardById", mock.Anything).Return(nextCard, nil)
	cs.On("SelectCard", mock.Anything, mock.Anything).Return(csReturn)
	pio.On("ReadReportAndNext", req).Return(reportAndNext, nil)

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, &nextCard)
	}))

	um.AssertCalled(t, "UpdateHistory", mock.MatchedBy(func(u pbuf.User) bool {
		return proto.Equal(&u, &testUser)
	}), reportAndNext.NextRequest.CardSetId, mock.MatchedBy(func(up pbuf.CardUpdate) bool {
		return proto.Equal(&up, &testUpdate)
	}))

	um.AssertCalled(t, "GetHistory", mock.MatchedBy(func(u pbuf.User) bool {
		return proto.Equal(&u, &testUser)
	}), reportAndNext.NextRequest.CardSetId)

	cs.AssertCalled(t, "SelectCard", mock.MatchedBy(func(h *pbuf.UserHistory) bool {
		return proto.Equal(h, &testHistory)
	}), reportAndNext.NextRequest.PreviousCardId)

	cm.AssertCalled(t, "GetCardById", csReturn)
}
