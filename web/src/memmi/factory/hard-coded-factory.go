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
	reportHandler := &request.ReportRequestHandler{}

	csHandler.Pio = pio
	csHandler.CardMan = cMan

	reportHandler.Pio = pio
	reportHandler.UserMan = uMan

	router.AddHandler(reportHandler)
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
