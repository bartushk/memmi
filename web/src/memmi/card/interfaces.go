package card

import (
	"memmi/pbuf"
)

type CardSelection interface {
	SelectCard(history *pbuf.UserHistory, previousCardId []byte) []byte
}

type CardManagement interface {
	GetCardSetById(id []byte) pbuf.CardSet
	GetCardById(id []byte) pbuf.Card
	SaveCardSet(*pbuf.CardSet) ([]byte, error)
	SaveCard(*pbuf.Card) ([]byte, error)
}
