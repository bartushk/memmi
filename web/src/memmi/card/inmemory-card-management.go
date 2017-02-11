package card

import (
	"fmt"
	"memmi/pbuf"
)

func NewInMemoryManagement() *InMemoryCardManagement {
	newVal := &InMemoryCardManagement{
		cardSets: make(map[string]pbuf.CardSet),
		cards:    make(map[string]pbuf.Card),
	}
	return newVal
}

type InMemoryCardManagement struct {
	cardSets map[string]pbuf.CardSet
	cards    map[string]pbuf.Card
}

func (manager *InMemoryCardManagement) GetCardSetById(id []byte) pbuf.CardSet {
	key := fmt.Sprintf("%x", id)
	return manager.cardSets[key]
}

func (manager *InMemoryCardManagement) GetCardById(id []byte) pbuf.Card {
	key := fmt.Sprintf("%x", id)
	return manager.cards[key]
}

func (manager *InMemoryCardManagement) SaveCardSet(set *pbuf.CardSet) ([]byte, error) {
	return nil, nil
}

func (manager *MockCardManagement) SaveCard(card *pbuf.Card) ([]byte, error) {
	return nil, nil
}
