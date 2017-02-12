package card

import (
	"github.com/golang/protobuf/proto"
	"memmi/pbuf"
	"testing"
)

func Test_InMemoryCardManagement_NewInMemoryManagement(t *testing.T) {
	newMan := NewInMemoryManagement()
	if newMan.cards == nil {
		t.Error("cards should be initialized")
	}

	if newMan.cardSets == nil {
		t.Error("cardSets should be initialized")
	}
}

func Test_InMemoryCardManagement_GetCardSetById_NoKey_BlankReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := []byte{1, 2, 3}
	blankSet := pbuf.CardSet{}

	result := newMan.GetCardSetById(testId)

	if !proto.Equal(&result, &blankSet) {
		t.Error("Empty CardSet was not returned when no key is present.")
	}
}

func Test_InMemoryCardManagement_GetCardSetById_GoodKey_CardReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := []byte{1, 2, 3}
	testCardSet := pbuf.CardSet{
		Id:      testId,
		Version: 5,
	}
	newMan.cardSets[newMan.getKey(testId)] = testCardSet

	result := newMan.GetCardSetById(testId)

	if !proto.Equal(&result, &testCardSet) {
		t.Error("Wrong result returned",
			"Expected:", testCardSet,
			"Got:", result)
	}
}

func Test_InMemoryCardManagement_GetCardById_NoKey_BlankReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := []byte{1, 2, 3}
	blankCard := pbuf.Card{}

	result := newMan.GetCardById(testId)

	if !proto.Equal(&result, &blankCard) {
		t.Error("Empty Card was not returned when no key is present.")
	}
}

func Test_InMemoryCardManagement_GetCardById_GoodKey_CardReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := []byte{1, 2, 3}
	testCard := pbuf.Card{
		Id:    testId,
		Title: "Hello World!",
	}
	newMan.cards[newMan.getKey(testId)] = testCard

	result := newMan.GetCardById(testId)

	if !proto.Equal(&result, &testCard) {
		t.Error("Wrong result returned",
			"Expected:", testCard,
			"Got:", result)
	}
}
