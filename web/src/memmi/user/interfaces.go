package user

import (
	"memmi/pbuf"
)

type UserManagement interface {
	GetHistory(user pbuf.User, cardSetId []byte) pbuf.UserHistory
	GetAuthInfoByUserName(userName string) pbuf.UserAuthInfo
	GetAuthInfoById(userId []byte) pbuf.UserAuthInfo
	UpdateHistory(user pbuf.User, cardSetId []byte, update pbuf.CardUpdate) error
	AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) error
	DeleteUser(userId []byte) error
}
