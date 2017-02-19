package user

import (
	"memmi/pbuf"
)

type UserManagement interface {
	GetHistory(user pbuf.User, cardSetId int64) (pbuf.UserHistory, error)
	GetAuthInfoByUserName(userName string) (pbuf.UserAuthInfo, error)
	GetAuthInfoById(userId int64) (pbuf.UserAuthInfo, error)
	GetUserByUserName(userName string) (pbuf.User, error)
	GetUserById(userId int64) (pbuf.User, error)
	UpdateHistory(user pbuf.User, cardSetId int64, update pbuf.CardUpdate) error
	AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) error
	DeleteUser(userId int64) error
}
