package request

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/mock"
	"memmi/card"
	"memmi/pbuf"
	"net/http"
	"testing"
)

func getCardSetMocked() (RequestHandler, *MockProtoIO, *card.MockCardManagement) {
	handler := &CardSetRequestHandler{}
	pio := &MockProtoIO{}
	pio.On("WriteProtoResponse", mock.Anything, mock.Anything).Return(nil)
	pio.On("WriteCodedProtoResponse", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	cardMan := &card.MockCardManagement{}
	handler.Pio = pio
	handler.CardMan = cardMan
	return handler, pio, cardMan
}

func Test_CardSetHandler_ExactUrl_ShouldHandle(t *testing.T) {
	var req = RequestFromURL(CARD_SET_API_URL)
	handler := CardSetRequestHandler{}
	if !handler.ShouldHandle(req, pbuf.User{}, false) {
		t.Error("Handler should handle with URL:", CARD_SET_API_URL)
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

	pio.On("ReadCardSetRequest", req).Return(pbuf.CardSetRequest{}, errors.New(""))

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteCodedProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, BODY_READ_ERROR)
	}), http.StatusBadRequest)

	cm.AssertNotCalled(t, "GetCardSetById", mock.Anything)
}

func Test_CardSetHandler_GetCardSet_GetByIdError_WriteError(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARDSET_URL)
	testCardSetRequest := pbuf.CardSetRequest{Id: "asdf"}
	testCardSet := pbuf.CardSet{Title: "TestCard"}
	testUser := pbuf.User{}

	pio.On("ReadCardSetRequest", req).Return(testCardSetRequest, nil)
	cm.On("GetCardSetById", mock.Anything).Return(testCardSet, errors.New(""))
	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteCodedProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, CARD_SET_NOT_FOUND_ERROR)
	}), http.StatusBadRequest)

	cm.AssertCalled(t, "GetCardSetById", testCardSetRequest.Id)
}

func Test_CardSetHandler_GetCardSet_NoError_HandledCorrectly(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARDSET_URL)
	testUser := pbuf.User{}
	testCardSetRequest := pbuf.CardSetRequest{Id: "asdf"}
	testCardSet := pbuf.CardSet{Title: "TestCard"}

	cm.On("GetCardSetById", mock.Anything).Return(testCardSet, nil)
	pio.On("ReadCardSetRequest", req).Return(testCardSetRequest, nil)
	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, &testCardSet)
	}))

	cm.AssertCalled(t, "GetCardSetById", testCardSetRequest.Id)
}

func Test_CardSetHandler_GetCard_ProtoIOReadError_WriteError(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARD_URL)
	testUser := pbuf.User{}
	pio.On("ReadCardRequest", req).Return(pbuf.CardRequest{}, errors.New(""))

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteCodedProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, BODY_READ_ERROR)
	}), http.StatusBadRequest)

	cm.AssertNotCalled(t, "GetCardById", mock.Anything)
}

func Test_CardSetHandler_GetCard_GetByIdError_WritesError(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARD_URL)
	testUser := pbuf.User{}
	testCardRequest := pbuf.CardRequest{Id: "asdf"}
	testCard := pbuf.Card{Title: "TestCard"}

	cm.On("GetCardById", mock.Anything).Return(testCard, errors.New(""))
	pio.On("ReadCardRequest", req).Return(testCardRequest, nil)

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteCodedProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, CARD_NOT_FOUND_ERROR)
	}), http.StatusBadRequest)

	cm.AssertCalled(t, "GetCardById", testCardRequest.Id)
}

func Test_CardSetHandler_GetCard_NoError_HandledCorrectly(t *testing.T) {
	var handler, pio, cm = getCardSetMocked()
	var req = RequestFromURL(GET_CARD_URL)
	testUser := pbuf.User{}
	testCardRequest := pbuf.CardRequest{Id: "asdf"}
	testCard := pbuf.Card{Title: "TestCard"}

	cm.On("GetCardById", mock.Anything).Return(testCard, nil)
	pio.On("ReadCardRequest", req).Return(testCardRequest, nil)

	handler.Handle(nil, req, testUser)

	pio.AssertCalled(t, "WriteProtoResponse", mock.Anything, mock.MatchedBy(func(m proto.Message) bool {
		return proto.Equal(m, &testCard)
	}))

	cm.AssertCalled(t, "GetCardById", testCardRequest.Id)
}
