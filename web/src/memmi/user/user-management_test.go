package user

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"memmi/card"
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
	return &InMemoryUserManagement{
		CardMan: &card.InMemoryCardManagement{},
	}
}

func Test_UserManagement_Suite(t *testing.T) {
	inMemSuite := &UserManagementTestSuite{
		Desc: "InMemoryManagement", New: NewMan(NewInMemory)}
	suite.Run(t, inMemSuite)
}

func (suite *UserManagementTestSuite) Test_UserManagement_GetHistory_BadId_ReturnError() {
	t := suite.T()
	testMan := suite.New()
	testUser := test.GetFakeUser()
	blank := pbuf.UserHistory{}

	result, err := testMan.GetHistory(testUser, 0)

	test.AssertProtoEq(t, &blank, &result, suite.Desc)

	assert.NotNil(t, err)
}
