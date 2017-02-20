package factory

import (
	"memmi/card"
	"memmi/request"
	"memmi/user"
)

type ConfigFactory struct {
}

func (fact *ConfigFactory) GetRouter() request.Router {
	return request.Router{}
}

func (fact *ConfigFactory) GetCardSelection() card.CardManagement {
	return &card.InMemoryCardManagement{}
}

func (fact *ConfigFactory) GetUserManagement() user.UserManagement {
	return &user.InMemoryUserManagement{}
}
