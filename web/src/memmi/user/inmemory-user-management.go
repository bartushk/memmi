package user

import (
	"memmi/pbuf"
)

type InMemoryUserManagement struct {
	userIds       map[string][]byte
	authInfo      map[string]pbuf.UserAuthInfo
	userHistories map[string]pbuf.UserHistory
}

func (manager *InMemoryUserManagement) GetHistory(user pbuf.User, cardSetId []byte) pbuf.UserHistory {
	return pbuf.UserHistory{}
}

func (manager *InMemoryUserManagement) GetAuthInfoByUserName(userName string) pbuf.UserAuthInfo {
	return pbuf.UserAuthInfo{}
}

func (manager *InMemoryUserManagement) UpdateHistory(user pbuf.User, cardSetId []byte, update pbuf.CardUpdate) error {
	return nil
}

func (manager *InMemoryUserManagement) GetAuthInfoById(user pbuf.User, cardSetId []byte) pbuf.UserAuthInfo {
	return pbuf.UserAuthInfo{}
}

func (management *InMemoryUserManagement) AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) error {
	return nil
}

func (management *InMemoryUserManagement) DeleteUser(userId []byte) error {
	return nil
}
