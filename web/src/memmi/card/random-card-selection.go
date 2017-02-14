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

func (selection *RandomCardSelection) SelectCard(history *pbuf.UserHistory, previousCard []byte) []byte {
	if history == nil {
		return []byte{}
	}
	if len(history.History) == 0 {
		return []byte{}
	}
	i := selection.src.Int63() % int64(len(history.History))
	return history.History[i].CardId
}
