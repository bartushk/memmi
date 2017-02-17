package card

import (
	"errors"
	"fmt"
	"hash/fnv"
	"memmi/pbuf"
)

func (manager *InMemoryCardManagement) getKey(id []byte) string {
	return fmt.Sprintf("%x", id)
}

func (manager *InMemoryCardManagement) getId(input uint32) []byte {
	hash := fnv.New64()
	hash.Write([]byte(fmt.Sprintf("%x", input)))
	return hash.Sum(nil)
}

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
	cardCounter    uint32
	cardSetCounter uint32
}

func (manager *InMemoryCardManagement) GetCardSetById(id []byte) (pbuf.CardSet, error) {
	key := manager.getKey(id)
	cardSet, ok := manager.cardSets[key]
	if !ok {
		return pbuf.CardSet{}, errors.New("Card set not found.")
	}
	return cardSet, nil
}

func (manager *InMemoryCardManagement) GetCardById(id []byte) (pbuf.Card, error) {
	key := manager.getKey(id)
	card, ok := manager.cards[key]
	if !ok {
		return pbuf.Card{}, errors.New("Card not found.")
	}
	return card, nil
}

func (manager *InMemoryCardManagement) DeleteCardSet(id []byte) error {
	key := fmt.Sprintf("%x", id)
	_, ok := manager.cardSets[key]
	if !ok {
		return errors.New("CardSet with that ID does not exist and could not be deleted.")
	}
	delete(manager.cardSets, key)
	return nil
}

func (manager *InMemoryCardManagement) DeleteCard(id []byte) error {
	key := fmt.Sprintf("%x", id)
	_, ok := manager.cards[key]
	if !ok {
		return errors.New("Card with that ID does not exist and could not be deleted.")
	}
	delete(manager.cards, key)
	return nil
}

func (manager *InMemoryCardManagement) SaveCardSet(set *pbuf.CardSet) ([]byte, error) {
	if set == nil {
		return nil, errors.New("Cannot save nil.")
	}
	id := manager.getId(manager.cardSetCounter)
	key := manager.getKey(id)
	_, ok := manager.cardSets[key]
	if ok {
		return nil, errors.New("There was a key collision and your card set could not be saved.")
	}
	set.Id = id
	manager.cardSets[key] = *set
	manager.cardSetCounter += 1
	return id, nil
}

func (manager *InMemoryCardManagement) SaveCard(card *pbuf.Card) ([]byte, error) {
	if card == nil {
		return nil, errors.New("Cannot save nil.")
	}
	id := manager.getId(manager.cardCounter)
	key := manager.getKey(id)
	_, ok := manager.cards[key]
	if ok {
		return nil, errors.New("There was a key collision and your card could not be saved.")
	}
	card.Id = id
	manager.cards[key] = *card
	manager.cardCounter += 1
	return id, nil
}
