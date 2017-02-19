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
	pio.NextCardError = errors.New("")
	handler.Handle(nil, req, pbuf.User{})
	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}
	if pio.MessageWrites[0] != BODY_READ_ERROR {
		t.Error("Wrong error type written to proto io.",
			"Expected:", BODY_READ_ERROR,
			"Got:", pio.MessageWrites[0])
	}

	um.AssertNotCalled(t, "UpdateHistory", mock.Anything)
	um.AssertNotCalled(t, "GetHistory", mock.Anything)
	cs.AssertNotCalled(t, "SelectCard", mock.Anything)
	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardHandler_HandleNext_ProtoReadGood_HandledCorrectly(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_NEXT_URL)
	testUser := pbuf.User{UserName: "bartushk"}
	nextCardRequest := pbuf.NextCardRequest{CardSetId: int64(3), PreviousCardId: int64(7)}
	nextCard := pbuf.Card{Title: "TestCard"}
	testHistory := pbuf.UserHistory{PlayIndex: 123}
	csReturn := int64(3)

	cm.On("GetCardById", mock.Anything).Return(nextCard, nil)
	cs.On("SelectCard", mock.Anything, mock.Anything).Return(csReturn)
	um.On("GetHistory", mock.Anything, mock.Anything).Return(testHistory, nil)
	pio.NextCardReturn = nextCardRequest

	handler.Handle(nil, req, testUser)

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != nextCard.String() {
		t.Error("Next card should have been written to proto io",
			"Expected:", nextCard.String(),
			"Got:", pio.MessageWrites[0].String())
	}

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
	pio.ReportError = errors.New("")
	handler.Handle(nil, req, pbuf.User{})
	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}
	if pio.MessageWrites[0] != BODY_READ_ERROR {
		t.Error("Wrong error type written to proto io.",
			"Expected:", BODY_READ_ERROR,
			"Got:", pio.MessageWrites[0])
	}

	um.AssertNotCalled(t, "UpdateHistory", mock.Anything)
	um.AssertNotCalled(t, "GetHistory", mock.Anything)
	cs.AssertNotCalled(t, "SelectCard", mock.Anything)
	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardHandler_Report_UpdateError_WriteError(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_URL)
	testUser := pbuf.User{Id: int64(3)}
	testUpdate := pbuf.CardUpdate{CardId: int64(3)}
	testCardReport := pbuf.CardScoreReport{CardSetId: int64(3)}
	testCardReport.Update = &testUpdate
	pio.ReportReturn = testCardReport

	um.On("UpdateHistory", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))

	handler.Handle(nil, req, testUser)

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}
	if pio.MessageWrites[0] != USER_HISTORY_UPDATE_ERROR {
		t.Error("Wrong error type written to proto io.",
			"Expected:", USER_HISTORY_UPDATE_ERROR,
			"Got:", pio.MessageWrites[0])
	}

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
	testUser := pbuf.User{Id: int64(3)}
	testUpdate := pbuf.CardUpdate{CardId: int64(3)}
	testCardReport := pbuf.CardScoreReport{CardSetId: int64(3)}
	testCardReport.Update = &testUpdate
	pio.ReportReturn = testCardReport

	um.On("UpdateHistory", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	handler.Handle(nil, req, testUser)
	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != expectedResponseWrite.String() {
		t.Error("Wrong error type written to proto io.",
			"Expected:", expectedResponseWrite,
			"Got:", pio.MessageWrites[0])
	}

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
	pio.ReportNextError = errors.New("")
	handler.Handle(nil, req, pbuf.User{})
	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}
	if pio.MessageWrites[0] != BODY_READ_ERROR {
		t.Error("Wrong error type written to proto io.",
			"Expected:", BODY_READ_ERROR,
			"Got:", pio.MessageWrites[0])
	}
	um.AssertNotCalled(t, "UpdateHistory", mock.Anything)
	um.AssertNotCalled(t, "GetHistory", mock.Anything)
	cs.AssertNotCalled(t, "SelectCard", mock.Anything)
	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardHandler_ReportNext_HandledCorrectly(t *testing.T) {
	var handler, pio, cs, um, cm = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_NEXT_URL)
	testUpdate := pbuf.CardUpdate{CardId: int64(3)}
	reportAndNext := pbuf.ReportAndNext{
		NextRequest: &pbuf.NextCardRequest{CardSetId: int64(10), PreviousCardId: int64(18)},
		Report:      &pbuf.CardScoreReport{CardSetId: int64(10), Update: &testUpdate},
	}
	nextCard := pbuf.Card{Title: "TestCard"}
	testHistory := pbuf.UserHistory{PlayIndex: 123}
	testUser := pbuf.User{Id: int64(3)}
	csReturn := int64(3)

	um.On("UpdateHistory", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	um.On("GetHistory", mock.Anything, mock.Anything).Return(testHistory, nil)
	cm.On("GetCardById", mock.Anything).Return(nextCard, nil)
	cs.On("SelectCard", mock.Anything, mock.Anything).Return(csReturn)
	pio.ReportNextReturn = reportAndNext

	handler.Handle(nil, req, testUser)

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != nextCard.String() {
		t.Error("Next card should have been written to proto io",
			"Expected:", nextCard.String(),
			"Got:", pio.MessageWrites[0].String())
	}

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
	testUpdate := pbuf.CardUpdate{CardId: int64(3)}
	reportAndNext := pbuf.ReportAndNext{
		NextRequest: &pbuf.NextCardRequest{CardSetId: int64(10), PreviousCardId: int64(18)},
		Report:      &pbuf.CardScoreReport{CardSetId: int64(10), Update: &testUpdate},
	}
	nextCard := pbuf.Card{Title: "TestCard"}
	testUser := pbuf.User{Id: int64(3)}
	testHistory := pbuf.UserHistory{PlayIndex: 123}
	csReturn := int64(3)

	um.On("GetHistory", mock.Anything, mock.Anything).Return(testHistory, nil)
	um.On("UpdateHistory", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	cm.On("GetCardById", mock.Anything).Return(nextCard, nil)
	cs.On("SelectCard", mock.Anything, mock.Anything).Return(csReturn)
	pio.ReportNextReturn = reportAndNext

	handler.Handle(nil, req, testUser)

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != nextCard.String() {
		t.Error("Next card should have been written to proto io",
			"Expected:", nextCard.String(),
			"Got:", pio.MessageWrites[0].String())
	}

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
