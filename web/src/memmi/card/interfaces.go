package card

import (
	"memmi/pbuf"
)

type CardSelection interface {
	SelectCard(history *pbuf.PlayerHistory, previousCard *pbuf.Card) pbuf.Card
}

type CardManagment interface {
	GetCardSetById(id string) pbuf.CardSet
	GetCardById(id string) pbuf.Card
}
