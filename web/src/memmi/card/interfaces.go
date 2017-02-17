package card

import (
	"memmi/pbuf"
)

type CardSelection interface {
	SelectCard(history *pbuf.UserHistory, previousCardId []byte) []byte
}

type CardManagement interface {
	GetCardSetById(id []byte) (pbuf.CardSet, error)
	GetCardById(id []byte) (pbuf.Card, error)
	SaveCardSet(*pbuf.CardSet) ([]byte, error)
	SaveCard(*pbuf.Card) ([]byte, error)
	DeleteCardSet(id []byte) error
	DeleteCard(id []byte) error
}
