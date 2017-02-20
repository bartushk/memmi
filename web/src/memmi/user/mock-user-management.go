package user

import (
	"github.com/stretchr/testify/mock"
	"memmi/pbuf"
)

type MockUserManagement struct {
	mock.Mock
}

func (management *MockUserManagement) GetHistory(user pbuf.User, cardSetId int64) (pbuf.UserHistory, error) {
	args := management.Called(user, cardSetId)
	return args.Get(0).(pbuf.UserHistory), args.Error(1)
}

func (management *MockUserManagement) GetAuthInfoByUserName(userName string) (pbuf.UserAuthInfo, error) {
	args := management.Called(userName)
	return args.Get(0).(pbuf.UserAuthInfo), args.Error(1)
}

func (management *MockUserManagement) GetAuthInfoById(userId int64) (pbuf.UserAuthInfo, error) {
	args := management.Called(userId)
	return args.Get(0).(pbuf.UserAuthInfo), args.Error(1)
}

func (management *MockUserManagement) GetUserByUserName(userName string) (pbuf.User, error) {
	args := management.Called(userName)
	return args.Get(0).(pbuf.User), args.Error(1)
}

func (management *MockUserManagement) GetUserById(userId int64) (pbuf.User, error) {
	args := management.Called(userId)
	return args.Get(0).(pbuf.User), args.Error(1)
}

func (management *MockUserManagement) UpdateHistory(user pbuf.User, cardSetId int64, update pbuf.CardUpdate) error {
	args := management.Called(user, cardSetId, update)
	return args.Error(0)
}

func (management *MockUserManagement) AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) (int64, error) {
	args := management.Called(user, authInfo)
	return args.Get(0).(int64), args.Error(1)
}

func (management *MockUserManagement) DeleteUser(userId int64) error {
	args := management.Called(userId)
	return args.Error(0)
}
