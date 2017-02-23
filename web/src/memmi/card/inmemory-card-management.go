package card

import (
	"errors"
	"memmi/pbuf"
	"strconv"
)

func NewInMemoryManagement() *InMemoryCardManagement {
	newVal := &InMemoryCardManagement{
		cardSets: make(map[string]pbuf.CardSet),
		cards:    make(map[string]pbuf.Card),
	}
	return newVal
}

type InMemoryCardManagement struct {
	cardSets       map[string]pbuf.CardSet
	cards          map[string]pbuf.Card
	cardCounter    int
	cardSetCounter int
}

func (manager *InMemoryCardManagement) GetCardSetById(id string) (pbuf.CardSet, error) {
	cardSet, ok := manager.cardSets[id]
	if !ok {
		return pbuf.CardSet{}, errors.New("Card set not found.")
	}
	return cardSet, nil
}

func (manager *InMemoryCardManagement) GetCardById(id string) (pbuf.Card, error) {
	card, ok := manager.cards[id]
	if !ok {
		return pbuf.Card{}, errors.New("Card not found.")
	}
	return card, nil
}

func (manager *InMemoryCardManagement) DeleteCardSet(id string) error {
	_, ok := manager.cardSets[id]
	if !ok {
		return errors.New("CardSet with that ID does not exist and could not be deleted.")
	}
	delete(manager.cardSets, id)
	return nil
}

func (manager *InMemoryCardManagement) DeleteCard(id string) error {
	_, ok := manager.cards[id]
	if !ok {
		return errors.New("Card with that ID does not exist and could not be deleted.")
	}
	delete(manager.cards, id)
	return nil
}

func (manager *InMemoryCardManagement) SaveCardSet(set *pbuf.CardSet) (string, error) {
	if set == nil {
		return "", errors.New("Cannot save nil.")
	}
	id := strconv.Itoa(manager.cardSetCounter)
	_, ok := manager.cardSets[id]
	if ok {
		return "", errors.New("There was a key collision and your card set could not be saved.")
	}
	set.Id = id
	manager.cardSets[id] = *set
	manager.cardSetCounter += 1
	return id, nil
}

func (manager *InMemoryCardManagement) SaveCard(card *pbuf.Card) (string, error) {
	if card == nil {
		return "", errors.New("Cannot save nil.")
	}
	id := strconv.Itoa(manager.cardCounter)
	_, ok := manager.cards[id]
	if ok {
		return "", errors.New("There was a key collision and your card could not be saved.")
	}
	card.Id = id
	manager.cards[id] = *card
	manager.cardCounter += 1
	return id, nil
}
