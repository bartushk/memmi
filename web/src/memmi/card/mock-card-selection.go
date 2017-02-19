package card

import (
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
)

type MockCardSelection struct {
	mock.Mock
}

func (selection *MockCardSelection) SelectCard(history *pbuf.UserHistory, previousCard int64) int64 {
	args := selection.Called(history, previousCard)
	return args.Get(0).(int64)
}
