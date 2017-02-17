package card

import (
	"github.com/golang/protobuf/proto"
	"memmi/pbuf"
	"reflect"
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

	result, err := newMan.GetCardSetById(testId)

	if !proto.Equal(&result, &blankSet) {
		t.Error("Empty CardSet was not returned when no key is present.")
	}

	if err == nil {
		t.Error("Error should have been returned.")
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

	result, err := newMan.GetCardSetById(testId)

	if !proto.Equal(&result, &testCardSet) {
		t.Error("Wrong result returned",
			"Expected:", testCardSet,
			"Got:", result)
	}

	if err != nil {
		t.Error("Error should have been returned:", err)
	}
}

func Test_InMemoryCardManagement_GetCardById_NoKey_BlankReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := []byte{1, 2, 3}
	blankCard := pbuf.Card{}

	result, err := newMan.GetCardById(testId)

	if !proto.Equal(&result, &blankCard) {
		t.Error("Empty Card was not returned when no key is present.")
	}

	if err == nil {
		t.Error("Error should have been returned.")
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

	result, err := newMan.GetCardById(testId)

	if !proto.Equal(&result, &testCard) {
		t.Error("Wrong result returned",
			"Expected:", testCard,
			"Got:", result)
	}

	if err != nil {
		t.Error("Error should have been returned:", err)
	}
}

func Test_InMemoryCardManagement_DeleteCardSet_BadKey_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()

	result := newMan.DeleteCardSet([]byte{1, 4, 6})

	if result == nil {
		t.Error("Error was not returned, it should have been.")
	}
}

func Test_InMemoryCardManagement_DeleteCardSet_GoodKey_CardSetRemoved(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := []byte{1, 2, 3}
	testCardSet := pbuf.CardSet{
		Id:      testId,
		Version: 5,
	}

	blankSet := pbuf.CardSet{}
	newMan.cardSets[newMan.getKey(testId)] = testCardSet

	result := newMan.DeleteCardSet(testId)
	mapResult := newMan.cardSets[newMan.getKey(testId)]

	if result != nil {
		t.Error("No error should have been returned by delete operation.")
	}

	if !proto.Equal(&mapResult, &blankSet) {
		t.Error("Wrong result returned",
			"Expected:", blankSet,
			"Got:", mapResult)
	}
}

func Test_InMemoryCardManagement_DeleteCard_BadKey_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()

	result := newMan.DeleteCard([]byte{1, 4, 6})

	if result == nil {
		t.Error("Error was not returned, it should have been.")
	}
}

func Test_InMemoryCardManagement_DeleteCard_GoodKey_CardRemoved(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := []byte{1, 2, 3}
	testCard := pbuf.Card{
		Id:    testId,
		Title: "Hello World!",
	}

	blankCard := pbuf.Card{}
	newMan.cards[newMan.getKey(testId)] = testCard

	result := newMan.DeleteCard(testId)
	mapResult := newMan.cards[newMan.getKey(testId)]

	if result != nil {
		t.Error("No error should have been returned by delete operation.")
	}

	if !proto.Equal(&mapResult, &blankCard) {
		t.Error("Wrong result returned",
			"Expected:", blankCard,
			"Got:", mapResult)
	}
}

func Test_InMemoryCardManagement_SaveCardSet_PassedNil_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()

	_, result := newMan.SaveCardSet(nil)

	if result == nil {
		t.Error("Error should be returned when passed nil.")
	}
}

func Test_InMemoryCardManagement_SaveCardSet_DuplicateKey_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	id := newMan.getId(newMan.cardSetCounter)
	key := newMan.getKey(id)
	newMan.cardSets[key] = pbuf.CardSet{}
	_, result := newMan.SaveCardSet(&pbuf.CardSet{})

	if result == nil {
		t.Error("Error should be returned when passed nil.")
	}
}

func Test_InMemoryCardManagement_SaveCardSet_MultipleSaves_SavedCorrectly(t *testing.T) {
	newMan := NewInMemoryManagement()
	id1 := newMan.getId(newMan.cardSetCounter)
	key1 := newMan.getKey(id1)
	id2 := newMan.getId(newMan.cardSetCounter + 1)
	key2 := newMan.getKey(id2)
	set1 := pbuf.CardSet{Version: 10}
	set2 := pbuf.CardSet{Version: 30}

	result1, err1 := newMan.SaveCardSet(&set1)
	result2, err2 := newMan.SaveCardSet(&set2)

	test1 := newMan.cardSets[key1]
	test2 := newMan.cardSets[key2]

	if err1 != nil {
		t.Error("First save returned an error, should have returned nil:", err1)
	}

	if err2 != nil {
		t.Error("Second save returned an error, should have returned nil:", err2)
	}

	if !reflect.DeepEqual(id1, result1) {
		t.Error("Wrong id1 returned",
			"Expected:", id1,
			"Got:", result1)
	}

	if !reflect.DeepEqual(id2, result2) {
		t.Error("Wrong id2 returned",
			"Expected:", id2,
			"Got:", result2)
	}

	if !proto.Equal(&test1, &set1) {
		t.Error("Wrong card set saved to dictionary for first save.",
			"Expected:", set1,
			"Got:", test1)
	}

	if !proto.Equal(&test2, &set2) {
		t.Error("Wrong card set saved to dictionary for second save.",
			"Expected:", set2,
			"Got:", test2)
	}

}

func Test_InMemoryCardManagement_SaveCard_PassedNil_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()

	_, result := newMan.SaveCard(nil)

	if result == nil {
		t.Error("Error should be returned when passed nil.")
	}
}

func Test_InMemoryCardManagement_SaveCard_DuplicateKey_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	id := newMan.getId(newMan.cardCounter)
	key := newMan.getKey(id)
	newMan.cards[key] = pbuf.Card{}
	_, result := newMan.SaveCard(&pbuf.Card{})

	if result == nil {
		t.Error("Error should be returned when passed nil.")
	}
}

func Test_InMemoryCardManagement_SaveCard_MultipleSaves_SavedCorrectly(t *testing.T) {
	newMan := NewInMemoryManagement()
	id1 := newMan.getId(newMan.cardCounter)
	key1 := newMan.getKey(id1)
	id2 := newMan.getId(newMan.cardCounter + 1)
	key2 := newMan.getKey(id2)
	card1 := pbuf.Card{Title: "Card One"}
	card2 := pbuf.Card{Title: "Card Two"}

	result1, err1 := newMan.SaveCard(&card1)
	result2, err2 := newMan.SaveCard(&card2)

	test1 := newMan.cards[key1]
	test2 := newMan.cards[key2]

	if err1 != nil {
		t.Error("First save returned an error, should have returned nil:", err1)
	}

	if err2 != nil {
		t.Error("Second save returned an error, should have returned nil:", err2)
	}

	if !reflect.DeepEqual(id1, result1) {
		t.Error("Wrong id1 returned",
			"Expected:", id1,
			"Got:", result1)
	}

	if !reflect.DeepEqual(id2, result2) {
		t.Error("Wrong id2 returned",
			"Expected:", id2,
			"Got:", result2)
	}

	if !proto.Equal(&test1, &card1) {
		t.Error("Wrong card set saved to dictionary for first save.",
			"Expected:", card1,
			"Got:", test1)
	}

	if !proto.Equal(&test2, &card2) {
		t.Error("Wrong card set saved to dictionary for second save.",
			"Expected:", card2,
			"Got:", test2)
	}

}
