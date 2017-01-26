package card

import (
	"memmi/pbuf"
)

type CardSelection interface {
	SelectCard(history *pbuf.UserHistory, previousCardId []byte) pbuf.Card
}

type CardManagement interface {
	GetCardSetById(id []byte) pbuf.CardSet
	GetCardById(id []byte) pbuf.Card
}
