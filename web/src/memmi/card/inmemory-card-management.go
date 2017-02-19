package card

import (
	"errors"
	"memmi/pbuf"
)

func NewInMemoryManagement() *InMemoryCardManagement {
	newVal := &InMemoryCardManagement{
		cardSets: make(map[int64]pbuf.CardSet),
		cards:    make(map[int64]pbuf.Card),
	}
	return newVal
}

type InMemoryCardManagement struct {
	cardSets       map[int64]pbuf.CardSet
	cards          map[int64]pbuf.Card
	cardCounter    int64
	cardSetCounter int64
}

func (manager *InMemoryCardManagement) GetCardSetById(id int64) (pbuf.CardSet, error) {
	cardSet, ok := manager.cardSets[id]
	if !ok {
		return pbuf.CardSet{}, errors.New("Card set not found.")
	}
	return cardSet, nil
}

func (manager *InMemoryCardManagement) GetCardById(id int64) (pbuf.Card, error) {
	card, ok := manager.cards[id]
	if !ok {
		return pbuf.Card{}, errors.New("Card not found.")
	}
	return card, nil
}

func (manager *InMemoryCardManagement) DeleteCardSet(id int64) error {
	_, ok := manager.cardSets[id]
	if !ok {
		return errors.New("CardSet with that ID does not exist and could not be deleted.")
	}
	delete(manager.cardSets, id)
	return nil
}

func (manager *InMemoryCardManagement) DeleteCard(id int64) error {
	_, ok := manager.cards[id]
	if !ok {
		return errors.New("Card with that ID does not exist and could not be deleted.")
	}
	delete(manager.cards, id)
	return nil
}

func (manager *InMemoryCardManagement) SaveCardSet(set *pbuf.CardSet) (int64, error) {
	if set == nil {
		return 0, errors.New("Cannot save nil.")
	}
	id := manager.cardSetCounter
	_, ok := manager.cardSets[id]
	if ok {
		return 0, errors.New("There was a key collision and your card set could not be saved.")
	}
	set.Id = id
	manager.cardSets[id] = *set
	manager.cardSetCounter += 1
	return id, nil
}

func (manager *InMemoryCardManagement) SaveCard(card *pbuf.Card) (int64, error) {
	if card == nil {
		return 0, errors.New("Cannot save nil.")
	}
	id := manager.cardCounter
	_, ok := manager.cards[id]
	if ok {
		return 0, errors.New("There was a key collision and your card could not be saved.")
	}
	card.Id = id
	manager.cards[id] = *card
	manager.cardCounter += 1
	return id, nil
}
