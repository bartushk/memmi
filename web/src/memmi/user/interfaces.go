package user

import (
	"memmi/pbuf"
)

type UserManagement interface {
	HasHistory(userId []byte, cardSetId []byte) bool
	GetHistory(userId []byte, cardSetId []byte) pbuf.UserHistory
	GetAuthInfoByUserName(userName string) pbuf.UserAuthInfo
	GetAuthInfoById(userId []byte) pbuf.UserAuthInfo
	UpdateHistory(userId []byte, cardSetId []byte, update pbuf.CardUpdate)
}
