package user

import (
	"memmi/pbuf"
)

type UserManagement interface {
	GetHistory(user pbuf.User, cardSetId string) (pbuf.UserHistory, error)
	GetAuthInfoByUserName(userName string) (pbuf.UserAuthInfo, error)
	GetAuthInfoById(userId string) (pbuf.UserAuthInfo, error)
	GetUserByUserName(userName string) (pbuf.User, error)
	GetUserById(userId string) (pbuf.User, error)
	UpdateHistory(user pbuf.User, cardSetId string, update pbuf.CardUpdate) error
	AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) (string, error)
	DeleteUser(userId string) error
}
