package user

import (
	"memmi/pbuf"
)

type MockUserManagement struct {
	GetHistoryReturn     pbuf.UserHistory
	GetHistoryUsers      []pbuf.User
	GetHistoryCardSetIds [][]byte

	AuthInfoReturn    pbuf.UserAuthInfo
	AuthInfoUserNames []string
	AuthInfoIds       [][]byte

	GetUserReturn    pbuf.User
	GetUserUserNames []string
	GetUserIds       [][]byte

	UpdateHistoryUsers       []pbuf.User
	UpdateHistoryCardSetIds  [][]byte
	UpdateHistoryCardUpdates []pbuf.CardUpdate
	UpdateHistoryReturn      error

	AddUserUsers  []pbuf.User
	AddUserAuths  []pbuf.UserAuthInfo
	AddUserReturn error

	DeleteUserIds    [][]byte
	DeleteUserReturn error
}

func (management *MockUserManagement) GetHistory(user pbuf.User, cardSetId []byte) pbuf.UserHistory {
	management.GetHistoryUsers = append(management.GetHistoryUsers, user)
	management.GetHistoryCardSetIds = append(management.GetHistoryCardSetIds, cardSetId)
	return management.GetHistoryReturn
}

func (management *MockUserManagement) GetAuthInfoByUserName(userName string) pbuf.UserAuthInfo {
	management.AuthInfoUserNames = append(management.AuthInfoUserNames, userName)
	return management.AuthInfoReturn
}

func (management *MockUserManagement) GetAuthInfoById(userId []byte) pbuf.UserAuthInfo {
	management.AuthInfoIds = append(management.AuthInfoIds, userId)
	return management.AuthInfoReturn
}

func (management *MockUserManagement) GetUserByUserName(userName string) pbuf.User {
	management.GetUserUserNames = append(management.GetUserUserNames, userName)
	return management.GetUserReturn
}

func (management *MockUserManagement) GetUserById(userId []byte) pbuf.User {
	management.GetUserIds = append(management.GetUserIds, userId)
	return management.GetUserReturn
}

func (management *MockUserManagement) UpdateHistory(user pbuf.User, cardSetId []byte, update pbuf.CardUpdate) error {
	management.UpdateHistoryUsers = append(management.UpdateHistoryUsers, user)
	management.UpdateHistoryCardSetIds = append(management.UpdateHistoryCardSetIds, cardSetId)
	management.UpdateHistoryCardUpdates = append(management.UpdateHistoryCardUpdates, update)
	return management.UpdateHistoryReturn
}

func (management *MockUserManagement) AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) error {
	management.AddUserUsers = append(management.AddUserUsers, user)
	management.AddUserAuths = append(management.AddUserAuths, authInfo)
	return management.AddUserReturn
}

func (management *MockUserManagement) DeleteUser(userId []byte) error {
	management.DeleteUserIds = append(management.DeleteUserIds, userId)
	return management.DeleteUserReturn
}

func (management *MockUserManagement) TotalCalls() int {
	return len(management.UpdateHistoryUsers) +
		len(management.AuthInfoUserNames) +
		len(management.AuthInfoIds) +
		len(management.AddUserUsers) +
		len(management.GetUserIds) +
		len(management.GetUserUserNames) +
		len(management.DeleteUserIds) +
		len(management.GetHistoryUsers)
}
