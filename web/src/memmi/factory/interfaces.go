package factory

import (
	"memmi/card"
	"memmi/request"
	"memmi/user"
)

type Factory interface {
	GetProtoIO() request.ProtoIO
	GetRouter() request.Router
	GetCardManagment() card.CardManagement
	GetCardSelection() card.CardSelection
	GetUserManagement() user.UserManagement
}
