package card

import (
	"memmi/pbuf"
)

type MockCardManagement struct {
	GetCardIds   []int64
	GetCardError error
	ReturnCard   pbuf.Card

	GetCardSetIds   []int64
	GetCardSetError error
	ReturnCardSet   pbuf.CardSet

	SavedCardSets    []pbuf.CardSet
	SaveCardSetError error
	SaveCardSetId    int64

	SavedCards    []pbuf.Card
	SaveCardError error
	SaveCardId    int64

	DeleteCardSetIds   []int64
	DeleteCardSetError error

	DeleteCardIds   []int64
	DeleteCardError error
}

func (manager *MockCardManagement) GetCardSetById(id int64) (pbuf.CardSet, error) {
	manager.GetCardSetIds = append(manager.GetCardSetIds, id)
	return manager.ReturnCardSet, manager.GetCardSetError
}

func (manager *MockCardManagement) GetCardById(id int64) (pbuf.Card, error) {
	manager.GetCardIds = append(manager.GetCardIds, id)
	return manager.ReturnCard, manager.GetCardError
}

func (manager *MockCardManagement) TotalCalls() int {
	return len(manager.GetCardIds) + len(manager.GetCardSetIds)
}

func (manager *MockCardManagement) SaveCardSet(set *pbuf.CardSet) (int64, error) {
	manager.SavedCardSets = append(manager.SavedCardSets, *set)
	return manager.SaveCardSetId, manager.SaveCardSetError
}

func (manager *MockCardManagement) SaveCard(card *pbuf.Card) (int64, error) {
	manager.SavedCards = append(manager.SavedCards, *card)
	return manager.SaveCardId, manager.SaveCardError
}

func (manager *MockCardManagement) DeleteCardSet(id int64) error {
	manager.DeleteCardSetIds = append(manager.DeleteCardSetIds, id)
	return manager.DeleteCardSetError
}

func (manager *MockCardManagement) DeleteCard(id int64) error {
	manager.DeleteCardIds = append(manager.DeleteCardIds, id)
	return manager.DeleteCardError
}
