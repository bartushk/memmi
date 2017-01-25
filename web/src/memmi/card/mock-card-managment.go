package card

import (
	"memmi/pbuf"
)

type MockCardManagment struct {
	ReturnCardSet pbuf.CardSet
	ReturnCard    pbuf.Card
}

func (manager *MockCardManagment) GetCardSetById(id string) pbuf.CardSet {
	return manager.ReturnCardSet
}

func (manager *MockCardManagment) GetCardById(id string) pbuf.Card {
	return manager.ReturnCard
}
