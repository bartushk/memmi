package card

import (
	"memmi/pbuf"
)

type MockCardManagement struct {
	ReturnCardSet pbuf.CardSet
	ReturnCard    pbuf.Card
}

func (manager *MockCardManagement) GetCardSetById(id []byte) pbuf.CardSet {
	return manager.ReturnCardSet
}

func (manager *MockCardManagement) GetCardById(id []byte) pbuf.Card {
	return manager.ReturnCard
}
