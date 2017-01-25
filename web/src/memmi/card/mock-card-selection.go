package card

import (
	"memmi/pbuf"
)

type MockCardSelection struct {
	PlayerHistories []*pbuf.PlayerHistory
	PreviousCards   []*pbuf.Card
}

func (selection *MockCardSelection) SelectCard(history *pbuf.PlayerHistory, previousCard *pbuf.Card) pbuf.Card {
	selection.PlayerHistories = append(selection.PlayerHistories, history)
	selection.PreviousCards = append(selection.PreviousCards, previousCard)
	return pbuf.Card{}
}
