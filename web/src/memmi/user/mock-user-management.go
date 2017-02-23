package user

import (
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
)

type MockUserManagement struct {
	mock.Mock
}

func (management *MockUserManagement) GetHistory(user pbuf.User, cardSetId string) (pbuf.UserHistory, error) {
	args := management.Called(user, cardSetId)
	return args.Get(0).(pbuf.UserHistory), args.Error(1)
}

func (management *MockUserManagement) GetAuthInfoByUserName(userName string) (pbuf.UserAuthInfo, error) {
	args := management.Called(userName)
	return args.Get(0).(pbuf.UserAuthInfo), args.Error(1)
}

func (management *MockUserManagement) GetAuthInfoById(userId string) (pbuf.UserAuthInfo, error) {
	args := management.Called(userId)
	return args.Get(0).(pbuf.UserAuthInfo), args.Error(1)
}

func (management *MockUserManagement) GetUserByUserName(userName string) (pbuf.User, error) {
	args := management.Called(userName)
	return args.Get(0).(pbuf.User), args.Error(1)
}

func (management *MockUserManagement) GetUserById(userId string) (pbuf.User, error) {
	args := management.Called(userId)
	return args.Get(0).(pbuf.User), args.Error(1)
}

func (management *MockUserManagement) UpdateHistory(user pbuf.User, cardSetId string, update pbuf.CardUpdate) error {
	args := management.Called(user, cardSetId, update)
	return args.Error(0)
}

func (management *MockUserManagement) AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) (string, error) {
	args := management.Called(user, authInfo)
	return args.String(0), args.Error(1)
}

func (management *MockUserManagement) DeleteUser(userId string) error {
	args := management.Called(userId)
	return args.Error(0)
}
