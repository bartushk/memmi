package card

import (
	"memmi/pbuf"
)

type CardManagement interface {
	GetCardSetById(id string) (pbuf.CardSet, error)
	GetCardById(id string) (pbuf.Card, error)
	SaveCardSet(*pbuf.CardSet) (string, error)
	SaveCard(*pbuf.Card) (string, error)
	DeleteCardSet(id string) error
	DeleteCard(id string) error
}
