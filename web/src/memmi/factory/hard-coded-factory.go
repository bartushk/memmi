package factory

import (
	"memmi/card"
	"memmi/request"
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

	router.Logger = &request.ConsoleLogger{}
	router.Authenticator = auth

	csHandler := &request.CardSetRequestHandler{}
	cHandler := &request.CardRequestHandler{}

	csHandler.Pio = pio
	csHandler.CardMan = cMan

	cHandler.Pio = pio
	cHandler.CardSel = &card.RandomCardSelection{}
	cHandler.CardMan = cMan
	cHandler.UserMan = uMan

	router.AddHandler(cHandler)
	router.AddHandler(csHandler)

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
