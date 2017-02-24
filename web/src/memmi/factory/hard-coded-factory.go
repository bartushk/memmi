package factory

import (
	"memmi/card"
	"memmi/request"
	"memmi/test"
	"memmi/user"
)

type HardCodedFactory struct {
}

func (fact *HardCodedFactory) GetRouter() request.Router {
	router := request.Router{}
	cMan := card.NewInMemoryManagement()
	uMan := user.NewInMemoryManagement()
	uMan.CardMan = cMan
	pio := &request.ProtoIoImpl{}
	auth := &request.AnonAuthenticator{}

	router.Authenticator = auth

	csHandler := &request.CardSetRequestHandler{}
	cHandler := &request.CardRequestHandler{}

	csHandler.Pio = pio
	csHandler.CardMan = cMan

	cHandler.Pio = pio
	cHandler.CardSel = card.NewRandomCardSelection()
	cHandler.CardMan = cMan
	cHandler.UserMan = uMan

	router.AddHandler(cHandler)
	router.AddHandler(csHandler)

	// Seed some data.
	cardSet := test.GetFakeCardSet()
	cardSet.CardIds = []string{}
	cards := test.GetFakeCards()
	for _, card := range cards {
		id, _ := cMan.SaveCard(&card)
		cardSet.CardIds = append(cardSet.CardIds, id)
	}
	cMan.SaveCardSet(&cardSet)

	return router
}

func (fact *HardCodedFactory) GetCardManagement() card.CardManagement {
	cMan := card.NewInMemoryManagement()
	return cMan
}

func (fact *HardCodedFactory) GetUserManagement() user.UserManagement {
	uMan := user.NewInMemoryManagement()
	return uMan
}
