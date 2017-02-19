package card

import (
	"memmi/pbuf"
)

type MockCardSelection struct {
	UserHistories   []*pbuf.UserHistory
	PreviousCardIds []int64
	NextCard        int64
}

func (selection *MockCardSelection) SelectCard(history *pbuf.UserHistory, previousCard int64) int64 {
	selection.UserHistories = append(selection.UserHistories, history)
	selection.PreviousCardIds = append(selection.PreviousCardIds, previousCard)
	return selection.NextCard
}
