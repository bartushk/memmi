package card

import (
	"memmi/pbuf"
)

func getTestHistory() pbuf.UserHistory {
	ret := pbuf.UserHistory{
		UserId:    3,
		CardSetId: 3,
		PlayIndex: 4,
		History:   []*pbuf.CardHistory{},
	}

	history1 := &pbuf.CardHistory{
		CardId:       0,
		CurrentScore: 2,
		CardIndex:    0,
		Scores:       []int32{1, 1},
		Indicies:     []int32{1, 3},
	}

	history2 := &pbuf.CardHistory{
		CardId:       1,
		CurrentScore: -1,
		CardIndex:    1,
		Scores:       []int32{-1},
		Indicies:     []int32{2},
	}

	history3 := &pbuf.CardHistory{
		CardId:       3,
		CurrentScore: 0,
		CardIndex:    0,
		Scores:       []int32{0},
		Indicies:     []int32{4},
	}

	ret.History = append(ret.History, history1)
	ret.History = append(ret.History, history2)
	ret.History = append(ret.History, history3)
	return ret
}
