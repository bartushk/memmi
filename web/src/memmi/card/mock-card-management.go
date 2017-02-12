package card

import (
	"memmi/pbuf"
)

type MockCardManagement struct {
	GetCardIds [][]byte
	ReturnCard pbuf.Card

	GetCardSetIds [][]byte
	ReturnCardSet pbuf.CardSet

	SavedCardSets    []pbuf.CardSet
	SaveCardSetError error
	SaveCardSetId    []byte

	SavedCards    []pbuf.Card
	SaveCardError error
	SaveCardId    []byte

	DeleteCardSetIds   [][]byte
	DeleteCardSetError error

	DeleteCardIds   [][]byte
	DeleteCardError error
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

func (manager *MockCardManagement) SaveCardSet(set *pbuf.CardSet) ([]byte, error) {
	manager.SavedCardSets = append(manager.SavedCardSets, *set)
	return manager.SaveCardSetId, manager.SaveCardSetError
}

func (manager *MockCardManagement) SaveCard(card *pbuf.Card) ([]byte, error) {
	manager.SavedCards = append(manager.SavedCards, *card)
	return manager.SaveCardId, manager.SaveCardError
}

func (manager *MockCardManagement) DeleteCardSet(id []byte) error {
	manager.DeleteCardSetIds = append(manager.DeleteCardSetIds, id)
	return manager.DeleteCardSetError
}

func (manager *MockCardManagement) DeleteCard(id []byte) error {
	manager.DeleteCardIds = append(manager.DeleteCardIds, id)
	return manager.DeleteCardError
}
