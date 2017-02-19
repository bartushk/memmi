package card

import (
	"memmi/pbuf"
)

func GenerateEmptyHistory(cardSet *pbuf.CardSet) pbuf.UserHistory {
	if cardSet == nil {
		return pbuf.UserHistory{}
	}
	history := pbuf.UserHistory{
		UserId:    3,
		CardSetId: cardSet.Id,
		PlayIndex: 0,
		History:   []*pbuf.CardHistory{},
	}
	for i, cardId := range cardSet.CardIds {
		cardHistory := &pbuf.CardHistory{
			CardId:       cardId,
			CurrentScore: 0,
			CardIndex:    int32(i),
			Scores:       []int32{},
			Indicies:     []int32{},
		}
		history.History = append(history.History, cardHistory)
	}
	return history
}
