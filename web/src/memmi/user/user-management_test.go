package user

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"memmi/pbuf"
	"memmi/test"
	"testing"
)

type NewMan func() UserManagement

type UserManagementTestSuite struct {
	suite.Suite
	New  NewMan
	Desc string
}

func NewInMemory() UserManagement {
	uMan := NewInMemoryManagement()
	toAdd := test.GetFakeCardSet()
	uMan.CardMan.SaveCardSet(&pbuf.CardSet{})
	uMan.CardMan.SaveCardSet(&pbuf.CardSet{})
	uMan.CardMan.SaveCardSet(&pbuf.CardSet{})
	uMan.CardMan.SaveCardSet(&toAdd)
	return uMan
}

func Test_UserManagement_Suite(t *testing.T) {
	inMemSuite := &UserManagementTestSuite{
		Desc: "InMemoryManagement", New: NewMan(NewInMemory)}
	suite.Run(t, inMemSuite)
}

func (suite *UserManagementTestSuite) Test_UserManagement_GetHistory_BadId_ReturnError() {
	t := suite.T()
	uMan := suite.New()
	testUser := test.GetFakeUser()
	testUser.Id = -12
	blank := pbuf.UserHistory{}

	result, err := uMan.GetHistory(testUser, 12)

	test.AssertProtoEq(t, &blank, &result, suite.Desc)

	assert.NotNil(t, err)
}

func (suite *UserManagementTestSuite) Test_UserManagement_Add_GetId() {
	t := suite.T()
	uMan := suite.New()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()

	id, addErr := uMan.AddUser(testUser, testAuth)
	testUser.Id = id
	user, getErr := uMan.GetUserById(id)

	assert.Nil(t, addErr)
	assert.Nil(t, getErr)
	test.AssertProtoEq(t, &testUser, &user, suite.Desc)
}

func (suite *UserManagementTestSuite) Test_UserManagement_Add_GetUserName() {
	t := suite.T()
	uMan := suite.New()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()

	id, addErr := uMan.AddUser(testUser, testAuth)
	testUser.Id = id
	user, getErr := uMan.GetUserByUserName(testUser.UserName)

	assert.Nil(t, addErr)
	assert.Nil(t, getErr)
	test.AssertProtoEq(t, &testUser, &user, suite.Desc)
}

func (suite *UserManagementTestSuite) Test_UserManagement_Add_GetAuthInfoByUserName() {
	t := suite.T()
	uMan := suite.New()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()

	_, addErr := uMan.AddUser(testUser, testAuth)
	auth, authErr := uMan.GetAuthInfoByUserName(testUser.UserName)

	assert.Nil(t, addErr)
	assert.Nil(t, authErr)
	test.AssertProtoEq(t, &testAuth, &auth, suite.Desc)
}

func (suite *UserManagementTestSuite) Test_UserManagement_Add_GetAuthInfoById() {
	t := suite.T()
	uMan := suite.New()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()

	id, addErr := uMan.AddUser(testUser, testAuth)
	auth, authErr := uMan.GetAuthInfoById(id)

	assert.Nil(t, addErr)
	assert.Nil(t, authErr)
	test.AssertProtoEq(t, &testAuth, &auth, suite.Desc)
}

func (suite *UserManagementTestSuite) Test_UserManagement_Add_Remove_Get() {
	t := suite.T()
	uMan := suite.New()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()

	id, addErr := uMan.AddUser(testUser, testAuth)
	remErr := uMan.DeleteUser(id)
	_, getErr := uMan.GetUserByUserName(testUser.UserName)

	assert.Nil(t, addErr)
	assert.Nil(t, remErr)
	assert.NotNil(t, getErr)
}

func (suite *UserManagementTestSuite) Test_UserManagement_GetHistory_UpdateHistory() {
	t := suite.T()
	uMan := suite.New()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()
	fakeCardSet := test.GetFakeCardSet()
	testUpdate := pbuf.CardUpdate{
		CardId: fakeCardSet.CardIds[0],
		Score:  int32(3)}

	id, addErr := uMan.AddUser(testUser, testAuth)
	testUser.Id = id
	oldHistory, hisErr := uMan.GetHistory(testUser, fakeCardSet.Id)
	upErr := uMan.UpdateHistory(testUser, fakeCardSet.Id, testUpdate)
	newHistory, hisErr := uMan.GetHistory(testUser, fakeCardSet.Id)

	assert.Nil(t, addErr)
	assert.Nil(t, hisErr)
	assert.Nil(t, upErr)
	assert.Equal(t, oldHistory.PlayIndex+1, newHistory.PlayIndex)
}
