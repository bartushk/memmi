package card

import (
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
)

type MockCardSelection struct {
	mock.Mock
}

func (selection *MockCardSelection) SelectCard(history *pbuf.UserHistory, previousCard string) string {
	args := selection.Called(history, previousCard)
	return args.String(0)
}
