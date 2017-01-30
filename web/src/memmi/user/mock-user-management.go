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

	UpdateHistoryUsers       []pbuf.User
	UpdateHistoryCardSetIds  [][]byte
	UpdateHistoryCardUpdates []pbuf.CardUpdate
	UpdateHistoryReturn      error
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

func (management *MockUserManagement) UpdateHistory(user pbuf.User, cardSetId []byte, update pbuf.CardUpdate) error {
	management.UpdateHistoryUsers = append(management.UpdateHistoryUsers, user)
	management.UpdateHistoryCardSetIds = append(management.UpdateHistoryCardSetIds, cardSetId)
	management.UpdateHistoryCardUpdates = append(management.UpdateHistoryCardUpdates, update)
	return management.UpdateHistoryReturn
}

func (management *MockUserManagement) TotalCalls() int {
	return len(management.UpdateHistoryUsers) +
		len(management.AuthInfoUserNames) +
		len(management.AuthInfoIds) +
		len(management.GetHistoryUsers)
}
