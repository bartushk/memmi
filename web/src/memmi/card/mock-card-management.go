package card

import (
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
)

type MockCardManagement struct {
	mock.Mock
}

func (manager *MockCardManagement) GetCardSetById(id int64) (pbuf.CardSet, error) {
	args := manager.Called(id)
	return args.Get(0).(pbuf.CardSet), args.Error(1)
}

func (manager *MockCardManagement) GetCardById(id int64) (pbuf.Card, error) {
	args := manager.Called(id)
	return args.Get(0).(pbuf.Card), args.Error(1)
}

func (manager *MockCardManagement) SaveCardSet(set *pbuf.CardSet) (int64, error) {
	args := manager.Called(set)
	return args.Get(0).(int64), args.Error(1)
}

func (manager *MockCardManagement) SaveCard(card *pbuf.Card) (int64, error) {
	args := manager.Called(card)
	return args.Get(0).(int64), args.Error(1)
}

func (manager *MockCardManagement) DeleteCardSet(id int64) error {
	args := manager.Called(id)
	return args.Error(0)
}

func (manager *MockCardManagement) DeleteCard(id int64) error {
	args := manager.Called(id)
	return args.Error(0)
}
