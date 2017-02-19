package card

import (
	"memmi/pbuf"
)

type CardSelection interface {
	SelectCard(history *pbuf.UserHistory, previousCardId int64) int64
}

type CardManagement interface {
	GetCardSetById(id int64) (pbuf.CardSet, error)
	GetCardById(id int64) (pbuf.Card, error)
	SaveCardSet(*pbuf.CardSet) (int64, error)
	SaveCard(*pbuf.Card) (int64, error)
	DeleteCardSet(id int64) error
	DeleteCard(id int64) error
}
