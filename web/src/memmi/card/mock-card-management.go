package card

import (
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
)

type MockCardManagement struct {
	mock.Mock
}

func (manager *MockCardManagement) GetCardSetById(id string) (pbuf.CardSet, error) {
	args := manager.Called(id)
	return args.Get(0).(pbuf.CardSet), args.Error(1)
}

func (manager *MockCardManagement) GetCardById(id string) (pbuf.Card, error) {
	args := manager.Called(id)
	return args.Get(0).(pbuf.Card), args.Error(1)
}

func (manager *MockCardManagement) SaveCardSet(set *pbuf.CardSet) (string, error) {
	args := manager.Called(set)
	return args.String(0), args.Error(1)
}

func (manager *MockCardManagement) SaveCard(card *pbuf.Card) (string, error) {
	args := manager.Called(card)
	return args.String(0), args.Error(1)
}

func (manager *MockCardManagement) DeleteCardSet(id string) error {
	args := manager.Called(id)
	return args.Error(0)
}

func (manager *MockCardManagement) DeleteCard(id string) error {
	args := manager.Called(id)
	return args.Error(0)
}
