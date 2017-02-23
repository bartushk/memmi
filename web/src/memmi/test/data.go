package test

import (
	"memmi/pbuf"
)

func GetFakeUser() pbuf.User {
	return pbuf.User{
		Id:              "kbart",
		UserName:        "bartushk",
		FirstName:       "Kyle",
		LastName:        "Bartush",
		Email:           "bartushk@gmail.com",
		IsAuthenticated: false,
		IsAnon:          false,
	}
}

func GetFakeAuthInfo() pbuf.UserAuthInfo {
	return pbuf.UserAuthInfo{
		UserId:   "kbart",
		UserName: "bartushk",
		PassHash: "48175AC",
	}
}

func GetFakeCardSet() pbuf.CardSet {
	return pbuf.CardSet{
		Id:          "cardSet",
		Version:     int32(6),
		CreatedDate: "",
		AuthorId:    "kbart",
		Title:       "This set is cool.",
		CardIds:     []string{"c1", "c2", "c3"},
		Tags:        []string{"Cool", "Funny", "Test"},
	}
}

func GetFakeCards() []pbuf.Card {
	card1 := pbuf.Card{
		Id:          "c1",
		Title:       "First Card",
		Front:       &pbuf.CardInfo{Type: "html", Value: "<h1> First Card Front </h1>"},
		Back:        &pbuf.CardInfo{Type: "html", Value: "<h1> First Card Back </h1>"},
		CreatedDate: "",
		Tags:        []string{"First", "Card"},
	}

	card2 := pbuf.Card{
		Id:          "c2",
		Title:       "Second Card",
		Front:       &pbuf.CardInfo{Type: "html", Value: "<h1> Second Card Front </h1>"},
		Back:        &pbuf.CardInfo{Type: "html", Value: "<h1> Second Card Back </h1>"},
		CreatedDate: "",
		Tags:        []string{"Second", "Card"},
	}

	card3 := pbuf.Card{
		Id:          "c3",
		Title:       "Third Card",
		Front:       &pbuf.CardInfo{Type: "html", Value: "<h1> Third Card Front </h1>"},
		Back:        &pbuf.CardInfo{Type: "html", Value: "<h1> Third Card Back </h1>"},
		CreatedDate: "",
		Tags:        []string{"Third", "Card"},
	}
	return []pbuf.Card{card1, card2, card3}
}

func GetFakeHistory() pbuf.UserHistory {
	ret := pbuf.UserHistory{
		UserId:     "kbart",
		CardSetId:  "cardSet",
		SetVersion: int32(6),
		PlayIndex:  4,
		History:    []*pbuf.CardHistory{},
	}

	history1 := &pbuf.CardHistory{
		CardId:       "c1",
		CurrentScore: 2,
		CardIndex:    0,
		Scores:       []int32{1, 1},
		Indicies:     []int32{1, 3},
	}

	history2 := &pbuf.CardHistory{
		CardId:       "c2",
		CurrentScore: -1,
		CardIndex:    1,
		Scores:       []int32{-1},
		Indicies:     []int32{2},
	}

	history3 := &pbuf.CardHistory{
		CardId:       "c3",
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
