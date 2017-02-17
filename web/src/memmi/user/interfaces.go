package user

import (
	"memmi/pbuf"
)

//TODO: Make these return errors properly.
type UserManagement interface {
	GetHistory(user pbuf.User, cardSetId []byte) pbuf.UserHistory
	GetAuthInfoByUserName(userName string) pbuf.UserAuthInfo
	GetAuthInfoById(userId []byte) pbuf.UserAuthInfo
	GetUserByUserName(userName string) pbuf.User
	GetUserById(userId []byte) pbuf.User
	UpdateHistory(user pbuf.User, cardSetId []byte, update pbuf.CardUpdate) error
	AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) error
	DeleteUser(userId []byte) error
}
