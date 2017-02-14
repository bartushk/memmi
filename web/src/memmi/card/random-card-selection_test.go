package card

import (
	"memmi/pbuf"
	"reflect"
	"testing"
)

type MockSource struct {
	returnVal int64
	seeds     []int64
}

func (src *MockSource) Int63() int64 {
	return src.returnVal
}

func (src *MockSource) Seed(seed int64) {
	src.seeds = append(src.seeds, seed)
}

func Test_NewRandomCardSlection(t *testing.T) {
	randSel := NewRandomCardSelection()

	if randSel.src == nil {
		t.Error("Source is nil.")
	}
}

func Test_RandomCardSelection_WhenPassedNill_ReturnsEmptyId(t *testing.T) {
	randSel := RandomCardSelection{}
	expected := []byte{}

	result := randSel.SelectCard(nil, nil)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Wrong card id returned",
			"Expected:", expected,
			"Got:", result)
	}
}

func Test_RandomCardSelection_WhenHistoryIsEmpty_ReturnEMptyId(t *testing.T) {
	randSel := RandomCardSelection{}
	input := &pbuf.UserHistory{}
	expected := []byte{}

	result := randSel.SelectCard(input, nil)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Wrong card id returned",
			"Expected:", expected,
			"Got:", result)
	}
}

func Test_RandomCardSelection_ReturnsCorrectItem(t *testing.T) {
	testHistory := getTestHistory()
	mockSource := &MockSource{returnVal: 1}
	randSel := RandomCardSelection{src: mockSource}
	expected := testHistory.History[1].CardId

	result := randSel.SelectCard(&testHistory, nil)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Wrong card id returned",
			"Expected:", expected,
			"Got:", result)
	}
}

func Tetst_RandomCardSelectoin_HandlesOverflow(t *testing.T) {
	testHistory := getTestHistory()
	mockSource := &MockSource{returnVal: 4}
	randSel := RandomCardSelection{src: mockSource}
	expected := testHistory.History[1].CardId

	result := randSel.SelectCard(&testHistory, nil)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Wrong card id returned",
			"Expected:", expected,
			"Got:", result)
	}
}
