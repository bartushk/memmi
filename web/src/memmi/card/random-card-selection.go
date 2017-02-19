package card

import (
	"math/rand"
	"memmi/pbuf"
)

func NewRandomCardSelection() RandomCardSelection {
	return RandomCardSelection{src: rand.NewSource(0)}
}

type RandomCardSelection struct {
	src rand.Source
}

func (selection *RandomCardSelection) SelectCard(history *pbuf.UserHistory, previousCard int64) int64 {
	if history == nil {
		return 0
	}
	if len(history.History) == 0 {
		return 0
	}
	i := selection.src.Int63() % int64(len(history.History))
	return history.History[i].CardId
}
