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
		t.Error("There should be one write to proto io, got:", len(pio.MessageWrites))
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
		t.Error("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != nextCard.String() {
		t.Error("Next card should have been written to proto io",
			"Expected:", nextCard.String(),
			"Got:", pio.MessageWrites[0].String())
	}

	if len(cs.UserHistories) != 1 {
		t.Error("Card selection should not be run. Times run:", len(cs.UserHistories))
	}

	if cs.UserHistories[0].PlayIndex != testHistory.PlayIndex {
		t.Error("Wrong history passed to card selection",
			"Expected:", testHistory,
			"Got:", cs.UserHistories[0])
		t.Error("Card selection should not be run. Times run:", len(cs.UserHistories))
	}

	if um.TotalCalls() != 1 {
		t.Error("User managment should be called once. Times called:", um.TotalCalls)
	}

	if !CompareByteSlices(um.GetHistoryCardSetIds[0], nextCardRequest.CardSetId) {
		t.Error("Wrong card set Id passed to GetHistory")
	}

}
