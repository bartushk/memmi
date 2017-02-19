package card

import (
	"github.com/stretchr/testify/assert"
	"memmi/pbuf"
	"memmi/test"
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
	testId := int64(3)
	blankSet := pbuf.CardSet{}

	result, err := newMan.GetCardSetById(testId)

	test.AssertProtoEq(t, &result, &blankSet, "Should have returned empty CardSet")
	assert.NotNil(t, err)
}

func Test_InMemoryCardManagement_GetCardSetById_GoodKey_CardReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := int64(3)
	testCardSet := pbuf.CardSet{
		Id:      testId,
		Version: 5,
	}
	newMan.cardSets[testId] = testCardSet

	result, err := newMan.GetCardSetById(testId)

	test.AssertProtoEq(t, &testCardSet, &result, "Wrong result returned.")
	assert.Nil(t, err)
}

func Test_InMemoryCardManagement_GetCardById_NoKey_BlankReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := int64(3)
	blankCard := pbuf.Card{}

	result, err := newMan.GetCardById(testId)

	test.AssertProtoEq(t, &result, &blankCard, "Should have returned empty Card")
	assert.NotNil(t, err)
}

func Test_InMemoryCardManagement_GetCardById_GoodKey_CardReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := int64(3)
	testCard := pbuf.Card{
		Id:    testId,
		Title: "Hello World!",
	}
	newMan.cards[testId] = testCard

	result, err := newMan.GetCardById(testId)

	test.AssertProtoEq(t, &testCard, &result, "Wrong result returned.")
	assert.Nil(t, err)
}

func Test_InMemoryCardManagement_DeleteCardSet_BadKey_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()

	result := newMan.DeleteCardSet(int64(3))

	assert.NotNil(t, result)
}

func Test_InMemoryCardManagement_DeleteCardSet_GoodKey_CardSetRemoved(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := int64(3)
	testCardSet := pbuf.CardSet{
		Id:      testId,
		Version: 5,
	}

	blankSet := pbuf.CardSet{}
	newMan.cardSets[testId] = testCardSet

	result := newMan.DeleteCardSet(testId)
	mapResult := newMan.cardSets[testId]

	test.AssertProtoEq(t, &mapResult, &blankSet, "Wrong result returned.")
	assert.Nil(t, result)
}

func Test_InMemoryCardManagement_DeleteCard_BadKey_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()

	result := newMan.DeleteCard(int64(3))

	assert.NotNil(t, result)
}

func Test_InMemoryCardManagement_DeleteCard_GoodKey_CardRemoved(t *testing.T) {
	newMan := NewInMemoryManagement()
	testId := int64(3)
	testCard := pbuf.Card{
		Id:    testId,
		Title: "Hello World!",
	}

	blankCard := pbuf.Card{}
	newMan.cards[testId] = testCard

	result := newMan.DeleteCard(testId)
	mapResult := newMan.cards[testId]

	test.AssertProtoEq(t, &mapResult, &blankCard, "Wrong result returned.")
	assert.Nil(t, result)
}

func Test_InMemoryCardManagement_SaveCardSet_PassedNil_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()

	_, result := newMan.SaveCardSet(nil)

	assert.NotNil(t, result)
}

func Test_InMemoryCardManagement_SaveCardSet_DuplicateKey_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	id := int64(newMan.cardSetCounter)
	newMan.cardSets[id] = pbuf.CardSet{}
	_, result := newMan.SaveCardSet(&pbuf.CardSet{})

	assert.NotNil(t, result)
}

func Test_InMemoryCardManagement_SaveCardSet_MultipleSaves_SavedCorrectly(t *testing.T) {
	newMan := NewInMemoryManagement()
	id1 := int64(newMan.cardSetCounter)
	id2 := int64(newMan.cardSetCounter + 1)
	set1 := pbuf.CardSet{Version: 10}
	set2 := pbuf.CardSet{Version: 30}

	result1, err1 := newMan.SaveCardSet(&set1)
	result2, err2 := newMan.SaveCardSet(&set2)

	test1 := newMan.cardSets[id1]
	test2 := newMan.cardSets[id2]

	assert.Nil(t, err1)
	assert.Nil(t, err2)

	assert.Equal(t, id1, result1, "Wrong id1 returned.")
	assert.Equal(t, id2, result2, "Wrong id2 returned.")

	test.AssertProtoEq(t, &test1, &set1, "Wrong card set saved to dictionary for first save.")
	test.AssertProtoEq(t, &test2, &set2, "Wrong card set saved to dictionary for second save.")
}

func Test_InMemoryCardManagement_SaveCard_PassedNil_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()

	_, result := newMan.SaveCard(nil)

	assert.NotNil(t, result)
}

func Test_InMemoryCardManagement_SaveCard_DuplicateKey_ErrorReturned(t *testing.T) {
	newMan := NewInMemoryManagement()
	id := int64(newMan.cardCounter)
	newMan.cards[id] = pbuf.Card{}
	_, result := newMan.SaveCard(&pbuf.Card{})

	assert.NotNil(t, result)
}

func Test_InMemoryCardManagement_SaveCard_MultipleSaves_SavedCorrectly(t *testing.T) {
	newMan := NewInMemoryManagement()
	id1 := int64(newMan.cardCounter)
	id2 := int64(newMan.cardCounter + 1)
	card1 := pbuf.Card{Title: "Card One"}
	card2 := pbuf.Card{Title: "Card Two"}

	result1, err1 := newMan.SaveCard(&card1)
	result2, err2 := newMan.SaveCard(&card2)

	test1 := newMan.cards[id1]
	test2 := newMan.cards[id2]

	assert.Nil(t, err1)
	assert.Nil(t, err2)

	assert.Equal(t, id1, result1, "Wrong id1 returned.")
	assert.Equal(t, id2, result2, "Wrong id2 returned.")

	test.AssertProtoEq(t, &test1, &card1, "Wrong card saved to dictionary for first save.")
	test.AssertProtoEq(t, &test2, &card2, "Wrong card saved to dictionary for second save.")
}
