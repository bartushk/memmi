package card

import (
	"github.com/stretchr/testify/assert"
	"memmi/pbuf"
	"memmi/test"
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

func Test_RandomCardSelection_WhenPassedNill_ReturnsZero(t *testing.T) {
	randSel := RandomCardSelection{}
	expected := int64(0)

	result := randSel.SelectCard(nil, 0)

	assert.Equal(t, expected, result, "Wrong card id returned.")
}

func Test_RandomCardSelection_WhenHistoryIsEmpty_ReturnsZero(t *testing.T) {
	randSel := RandomCardSelection{}
	input := &pbuf.UserHistory{}
	expected := int64(0)

	result := randSel.SelectCard(input, 0)

	assert.Equal(t, expected, result, "Wrong card id returned.")
}

func Test_RandomCardSelection_ReturnsCorrectItem(t *testing.T) {
	testHistory := test.GetTestHistory()
	mockSource := &MockSource{returnVal: 1}
	randSel := RandomCardSelection{src: mockSource}
	expected := testHistory.History[1].CardId

	result := randSel.SelectCard(&testHistory, 0)

	assert.Equal(t, expected, result, "Wrong card id returned.")
}

func Tetst_RandomCardSelectoin_HandlesOverflow(t *testing.T) {
	testHistory := test.GetTestHistory()
	mockSource := &MockSource{returnVal: 4}
	randSel := RandomCardSelection{src: mockSource}
	expected := testHistory.History[1].CardId

	result := randSel.SelectCard(&testHistory, 0)

	assert.Equal(t, expected, result, "Wrong card id returned.")
}
