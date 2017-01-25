package card

import (
	"memmi/pbuf"
)

type MockCardSelection struct {
	UserHistories []*pbuf.UserHistory
	PreviousCards []*pbuf.Card
}

func (selection *MockCardSelection) SelectCard(history *pbuf.UserHistory, previousCard *pbuf.Card) pbuf.Card {
	selection.UserHistories = append(selection.UserHistories, history)
	selection.PreviousCards = append(selection.PreviousCards, previousCard)
	return pbuf.Card{}
}
