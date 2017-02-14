package card

import (
	"reflect"
	"testing"
)

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
