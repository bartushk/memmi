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

func Test_RandomCardSelection_WhenPassedNill_ReturnsEmpty(t *testing.T) {
	randSel := RandomCardSelection{}
	expected := ""

	result := randSel.SelectCard(nil, "")

	assert.Equal(t, expected, result, "Wrong card id returned.")
}

func Test_RandomCardSelection_WhenHistoryIsEmpty_ReturnsEmpty(t *testing.T) {
	randSel := RandomCardSelection{}
	input := &pbuf.UserHistory{}
	expected := ""

	result := randSel.SelectCard(input, "")

	assert.Equal(t, expected, result, "Wrong card id returned.")
}

func Test_RandomCardSelection_ReturnsCorrectItem(t *testing.T) {
	testHistory := test.GetFakeHistory()
	mockSource := &MockSource{returnVal: 1}
	randSel := RandomCardSelection{src: mockSource}
	expected := testHistory.History[1].CardId

	result := randSel.SelectCard(&testHistory, "")

	assert.Equal(t, expected, result, "Wrong card id returned.")
}

func Tetst_RandomCardSelectoin_HandlesOverflow(t *testing.T) {
	testHistory := test.GetFakeHistory()
	mockSource := &MockSource{returnVal: 4}
	randSel := RandomCardSelection{src: mockSource}
	expected := testHistory.History[1].CardId

	result := randSel.SelectCard(&testHistory, "")

	assert.Equal(t, expected, result, "Wrong card id returned.")
}
