package request

import (
	"errors"
	"memmi/card"
	"memmi/pbuf"
	"memmi/user"
	"testing"
)

func getMockedHandler() (RequestHandler, *MockProtoIO, *card.MockCardSelection, *user.MockUserManagement) {
	handler := &CardRequestHandler{}
	pio := &MockProtoIO{}
	cardSel := &card.MockCardSelection{}
	userMan := &user.MockUserManagement{}
	handler.Pio = pio
	handler.CardSel = cardSel
	handler.UserMan = userMan
	return handler, pio, cardSel, userMan
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
	var handler, pio, cs, um = getMockedHandler()
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
	if len(cs.UserHistories) != 0 {
		t.Error("Card selection should not be run. Times run:", len(cs.UserHistories))
	}
	if um.TotalCalls() != 0 {
		t.Error("User managment should not be called. Times called:", um.TotalCalls)
	}
}

func Test_CardHandler_HandleNext_ProtoReadGood_HandledCorrectly(t *testing.T) {
	var handler, pio, cs, um = getMockedHandler()
	var req = RequestFromURL(CARD_NEXT_URL)
	nextCardRequest := pbuf.NextCardRequest{CardSetId: []byte{0, 1, 2}}
	nextCard := pbuf.Card{Title: "TestCard"}
	testHistory := pbuf.UserHistory{PlayIndex: 123}
	cs.NextCard = nextCard
	pio.NextCardReturn = nextCardRequest
	um.GetHistoryReturn = testHistory
	handler.Handle(nil, req, pbuf.User{})

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != nextCard.String() {
		t.Error("Next card should have been written to proto io",
			"Expected:", nextCard.String(),
			"Got:", pio.MessageWrites[0].String())
	}

	if len(cs.UserHistories) != 1 {
		t.Fatal("Card selection should be run once. Times run:", len(cs.UserHistories))
	}

	if cs.UserHistories[0].PlayIndex != testHistory.PlayIndex {
		t.Error("Wrong history passed to card selection",
			"Expected:", testHistory,
			"Got:", cs.UserHistories[0])
		t.Error("Card selection should not be run. Times run:", len(cs.UserHistories))
	}

	if um.TotalCalls() != 1 {
		t.Fatal("User managment should be called once. Times called:", um.TotalCalls)
	}

	if !CompareByteSlices(um.GetHistoryCardSetIds[0], nextCardRequest.CardSetId) {
		t.Error("Wrong card set Id passed to GetHistory",
			"Expected:", um.GetHistoryCardSetIds[0],
			"Got:", nextCardRequest.CardSetId)
	}

}

func Test_CardHandler_Report_ProtoReadError_WriteError(t *testing.T) {
	var handler, pio, cs, um = getMockedHandler()
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
	if len(cs.UserHistories) != 0 {
		t.Error("Card selection should not be run. Times run:", len(cs.UserHistories))
	}
	if um.TotalCalls() != 0 {
		t.Error("User managment should not be called. Times called:", um.TotalCalls)
	}
}

func Test_CardHandler_Report_UpdateError_WriteError(t *testing.T) {
	var handler, pio, cs, um = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_URL)
	testUser := pbuf.User{Id: []byte{2, 3, 7}}
	testCardReport := pbuf.CardScoreReport{CardSetId: []byte{1, 2, 10}}
	testCardReport.Update = &pbuf.CardUpdate{}
	um.UpdateHistoryReturn = errors.New("")
	pio.ReportReturn = testCardReport
	handler.Handle(nil, req, testUser)
	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}
	if pio.MessageWrites[0] != USER_HISTORY_UPDATE_ERROR {
		t.Error("Wrong error type written to proto io.",
			"Expected:", USER_HISTORY_UPDATE_ERROR,
			"Got:", pio.MessageWrites[0])
	}
	if len(cs.UserHistories) != 0 {
		t.Error("Card selection should not be run. Times run:", len(cs.UserHistories))
	}

	if um.TotalCalls() != 1 {
		t.Fatal("User managment should be called once. Times called:", um.TotalCalls)
	}

	if !CompareByteSlices(um.UpdateHistoryUsers[0].Id, testUser.Id) {
		t.Error("Wrong user passed to update history.",
			"Expected:", um.UpdateHistoryUsers[0].Id,
			"Got:", testUser.Id)
	}

	if !CompareByteSlices(um.UpdateHistoryCardSetIds[0], testCardReport.CardSetId) {
		t.Error("Wrong user passed to update history.",
			"Expected:", um.UpdateHistoryCardSetIds[0],
			"Got:", testCardReport.CardSetId)
	}

}

func Test_CardHandler_Report_Success_HandledCorrectly(t *testing.T) {
	var handler, pio, cs, um = getMockedHandler()
	var req = RequestFromURL(CARD_REPORT_URL)
	expectedResponseWrite := pbuf.UpdateResponse{Status: 1}
	testUser := pbuf.User{Id: []byte{2, 3, 7}}
	testCardReport := pbuf.CardScoreReport{CardSetId: []byte{1, 2, 10}}
	testCardReport.Update = &pbuf.CardUpdate{}
	pio.ReportReturn = testCardReport
	handler.Handle(nil, req, testUser)
	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != expectedResponseWrite.String() {
		t.Error("Wrong error type written to proto io.",
			"Expected:", expectedResponseWrite,
			"Got:", pio.MessageWrites[0])
	}

	if len(cs.UserHistories) != 0 {
		t.Error("Card selection should not be run. Times run:", len(cs.UserHistories))
	}

	if um.TotalCalls() != 1 {
		t.Fatal("User managment should be called once. Times called:", um.TotalCalls)
	}

	if !CompareByteSlices(um.UpdateHistoryUsers[0].Id, testUser.Id) {
		t.Error("Wrong user passed to update history.",
			"Expected:", um.UpdateHistoryUsers[0].Id,
			"Got:", testUser.Id)
	}

	if !CompareByteSlices(um.UpdateHistoryCardSetIds[0], testCardReport.CardSetId) {
		t.Error("Wrong user passed to update history.",
			"Expected:", um.UpdateHistoryCardSetIds[0],
			"Got:", testCardReport.CardSetId)
	}
}

func Test_CardHandler_ReportNext_ProtoIO_ErrorWritten(t *testing.T) {
	var handler, pio, cs, um = getMockedHandler()
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
	if len(cs.UserHistories) != 0 {
		t.Error("Card selection should not be run. Times run:", len(cs.UserHistories))
	}
	if um.TotalCalls() != 0 {
		t.Error("User managment should not be called. Times called:", um.TotalCalls)
	}
}
