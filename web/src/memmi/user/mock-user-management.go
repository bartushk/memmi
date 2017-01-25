package user

import (
	"memmi/pbuf"
)

type MockUserManagement struct {
	HasHistoryReturn     bool
	HasHistoryUserIds    [][]byte
	HasHistoryCardSetIds [][]byte

	GetHistoryReturn     pbuf.UserHistory
	GetHistoryUserIds    [][]byte
	GetHistoryCardSetIds [][]byte

	AuthInfoReturn    pbuf.UserAuthInfo
	AuthInfoUserNames []string
	AuthInfoIds       [][]byte

	UpdateHistoryUserIds     [][]byte
	UpdateHistoryCardSetIds  [][]byte
	UpdateHistoryCardUpdates []pbuf.CardUpdate
}

func (management *MockUserManagement) HasHistory(userId []byte, cardSetId []byte) bool {
	management.HasHistoryUserIds = append(management.HasHistoryUserIds, userId)
	management.HasHistoryCardSetIds = append(management.HasHistoryCardSetIds, cardSetId)
	return management.HasHistoryReturn
}

func (management *MockUserManagement) GetHistory(userId []byte, cardSetId []byte) pbuf.UserHistory {
	management.GetHistoryUserIds = append(management.GetHistoryUserIds, userId)
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

func (management *MockUserManagement) UpdateHistory(userId []byte, cardSetId []byte, update pbuf.CardUpdate) {
	management.UpdateHistoryUserIds = append(management.UpdateHistoryUserIds, userId)
	management.UpdateHistoryCardSetIds = append(management.UpdateHistoryCardSetIds, cardSetId)
	management.UpdateHistoryCardUpdates = append(management.UpdateHistoryCardUpdates, update)
}
