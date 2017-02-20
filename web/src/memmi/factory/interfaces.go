package factory

import (
	"memmi/card"
	"memmi/request"
	"memmi/user"
)

type Factory interface {
	GetRouter() request.Router
	GetCardManagement() card.CardManagement
	GetUserManagement() user.UserManagement
}
