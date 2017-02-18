package user

import (
	"memmi/pbuf"
)

type UserManagement interface {
	GetHistory(user pbuf.User, cardSetId []byte) (pbuf.UserHistory, error)
	GetAuthInfoByUserName(userName string) (pbuf.UserAuthInfo, error)
	GetAuthInfoById(userId []byte) (pbuf.UserAuthInfo, error)
	GetUserByUserName(userName string) (pbuf.User, error)
	GetUserById(userId []byte) (pbuf.User, error)
	UpdateHistory(user pbuf.User, cardSetId []byte, update pbuf.CardUpdate) error
	AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) error
	DeleteUser(userId []byte) error
}
