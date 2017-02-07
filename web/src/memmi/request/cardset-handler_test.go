package request

import (
	"errors"
	"memmi/card"
	"memmi/pbuf"
	"testing"
)

func getCardSetMocked() (RequestHandler, *MockProtoIO, *card.MockCardManagement) {
	handler := &CardSetRequestHandler{}
	pio := &MockProtoIO{}
	cardMan := &card.MockCardManagement{}
	handler.Pio = pio
	handler.CardMan = cardMan
	return handler, pio, cardMan
}

func Test_CardSetHandler_ExactUrl_ShouldHandle(t *testing.T) {
	var req = RequestFromURL(CARD_SET_API_URL)
	handler := CardSetRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle with URL:", CARD_API_URL)
	}
}

func Test_CardSetHandler_URLPlusQuery_ShouldHandle(t *testing.T) {
	test_url := CARD_SET_API_URL + "?asdf"
	var req = RequestFromURL(test_url)
	handler := CardSetRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle with URL:", test_url)
	}
}

func Test_CardSetHandler_URLPlusSubUrl_ShouldHandle(t *testing.T) {
	test_url := CARD_SET_API_URL + "/asdf/fffa"
	var req = RequestFromURL(test_url)
	handler := CardSetRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle with URL:", test_url)
	}
}

func Test_CardSetHandler_AnyDifferentPrefix_ShouldNotHandle(t *testing.T) {
	test_url := "/test" + CARD_SET_API_URL
	var req = RequestFromURL(test_url)
	handler := CardSetRequestHandler{}
	if handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should not handle with URL:", test_url)
	}
}

func Test_CardSetHandler_GetCardSet_ProtoIOReadError_WriteError(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARDSET_URL)
	testUser := pbuf.User{}
	pio.CardSetError = errors.New("")

	handler.Handle(nil, req, testUser)

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0] != BODY_READ_ERROR {
		t.Error("Wrong error type written to proto io.",
			"Expected:", BODY_READ_ERROR,
			"Got:", pio.MessageWrites[0])
	}

	if cm.TotalCalls() != 0 {
		t.Error("Expected total calls to card management to be zero. Got: ", cm.TotalCalls())
	}

}

func Test_CardSetHandler_GetCardSet_RequestPassed(t *testing.T) {
	var handler, pio, _ = getCardSetMocked()
	var req = RequestFromURL(GET_CARDSET_URL)
	testUser := pbuf.User{}
	handler.Handle(nil, req, testUser)

	if len(pio.CardSetRequests) != 1 {
		t.Fatal("There should have been one request passed to proto io. Received: ", len(pio.CardSetRequests))
	}

	if pio.CardSetRequests[0] != req {
		t.Error("Wrong request passed to proto io.",
			"Expected:", req,
			"Got:", pio.CardSetRequests[0])
	}
}

func Test_CardSetHandler_GetCardSet_NoError_HandledCorrectly(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARDSET_URL)
	testUser := pbuf.User{}
	testCardSetRequest := pbuf.CardSetRequest{Id: []byte{3, 7, 9}}
	testCardSet := pbuf.CardSet{SetName: "TestCard"}

	cm.ReturnCardSet = testCardSet
	pio.CardSetReturn = testCardSetRequest
	handler.Handle(nil, req, testUser)

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != testCardSet.String() {
		t.Error("Wrong message written to proto io.",
			"Expected:", testCardSet.String(),
			"Got:", pio.MessageWrites[0].String())
	}

	if cm.TotalCalls() != 1 {
		t.Fatal("Expected total calls to card management to be one. Got: ", cm.TotalCalls())
	}

	if !CompareByteSlices(testCardSetRequest.Id, cm.GetCardSetIds[0]) {
		t.Error("Wrong cardSetId passed to card management.",
			"Expected:", testCardSetRequest.Id,
			"Got:", cm.GetCardSetIds[0])
	}
}

func Test_CardSetHandler_GetCard_ProtoIOReadError_WriteError(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARD_URL)
	testUser := pbuf.User{}
	pio.CardError = errors.New("")

	handler.Handle(nil, req, testUser)

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0] != BODY_READ_ERROR {
		t.Error("Wrong error type written to proto io.",
			"Expected:", BODY_READ_ERROR,
			"Got:", pio.MessageWrites[0])
	}

	if cm.TotalCalls() != 0 {
		t.Error("Expected total calls to card management to be zero. Got: ", cm.TotalCalls())
	}
}

func Test_CardSetHandler_GetCard_RequestPassed(t *testing.T) {
	var handler, pio, _ = getCardSetMocked()
	var req = RequestFromURL(GET_CARD_URL)
	testUser := pbuf.User{}
	handler.Handle(nil, req, testUser)

	if len(pio.CardRequests) != 1 {
		t.Fatal("There should have been one request passed to proto io. Received: ", len(pio.CardRequests))
	}

	if pio.CardRequests[0] != req {
		t.Error("Wrong request passed to proto io.",
			"Expected:", req,
			"Got:", pio.CardRequests[0])
	}
}

func Test_CardSetHandler_GetCard_NoError_HandledCorrectly(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARD_URL)
	testUser := pbuf.User{}
	testCardRequest := pbuf.CardRequest{Id: []byte{3, 7, 9}}
	testCard := pbuf.Card{Title: "TestCard"}

	cm.ReturnCard = testCard
	pio.CardReturn = testCardRequest
	handler.Handle(nil, req, testUser)

	if len(pio.MessageWrites) != 1 {
		t.Fatal("There should be one write to proto io, got:", len(pio.MessageWrites))
	}

	if pio.MessageWrites[0].String() != testCard.String() {
		t.Error("Wrong message written to proto io.",
			"Expected:", testCard.String(),
			"Got:", pio.MessageWrites[0].String())
	}

	if cm.TotalCalls() != 1 {
		t.Fatal("Expected total calls to card management to be one. Got: ", cm.TotalCalls())
	}

	if !CompareByteSlices(testCardRequest.Id, cm.GetCardIds[0]) {
		t.Error("Wrong cardId passed to card management.",
			"Expected:", testCardRequest.Id,
			"Got:", cm.GetCardIds[0])
	}
}