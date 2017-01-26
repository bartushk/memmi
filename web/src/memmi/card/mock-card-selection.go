package card

import (
	"memmi/pbuf"
)

type MockCardSelection struct {
	UserHistories   []*pbuf.UserHistory
	PreviousCardIds [][]byte
}

func (selection *MockCardSelection) SelectCard(history *pbuf.UserHistory, previousCard []byte) pbuf.Card {
	selection.UserHistories = append(selection.UserHistories, history)
	selection.PreviousCardIds = append(selection.PreviousCardIds, previousCard)
	return pbuf.Card{}
}
