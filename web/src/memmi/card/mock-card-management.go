package card

import (
	"memmi/pbuf"
)

type MockCardManagement struct {
	GetCardSetIds [][]byte
	GetCardIds    [][]byte

	ReturnCardSet pbuf.CardSet
	ReturnCard    pbuf.Card
}

func (manager *MockCardManagement) GetCardSetById(id []byte) pbuf.CardSet {
	manager.GetCardSetIds = append(manager.GetCardSetIds, id)
	return manager.ReturnCardSet
}

func (manager *MockCardManagement) GetCardById(id []byte) pbuf.Card {
	manager.GetCardIds = append(manager.GetCardIds, id)
	return manager.ReturnCard
}

func (manager *MockCardManagement) TotalCalls() int {
	return len(manager.GetCardIds) + len(manager.GetCardSetIds)
}
