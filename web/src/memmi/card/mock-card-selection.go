package card

import (
	"memmi/pbuf"
)

type MockCardSelection struct {
	UserHistories   []*pbuf.UserHistory
	PreviousCardIds [][]byte
	NextCard        []byte
}

func (selection *MockCardSelection) SelectCard(history *pbuf.UserHistory, previousCard []byte) []byte {
	selection.UserHistories = append(selection.UserHistories, history)
	selection.PreviousCardIds = append(selection.PreviousCardIds, previousCard)
	return selection.NextCard
}
