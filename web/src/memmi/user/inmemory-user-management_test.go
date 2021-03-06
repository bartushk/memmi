package user

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"memmi/card"
	"memmi/pbuf"
	"memmi/test"
	"strconv"
	"testing"
)

func getMockedUserManagement() (*InMemoryUserManagement, *card.MockCardManagement) {
	uMan := NewInMemoryManagement()
	cardMan := &card.MockCardManagement{}
	cardMan.On("GetCardSetById", mock.Anything).Return(test.GetFakeCardSet(), nil)
	uMan.CardMan = cardMan
	return uMan, cardMan
}

func Test_InMemoryUserManagement_NewInMemoryUserManagement(t *testing.T) {
	newMan := NewInMemoryManagement()

	assert.NotNil(t, newMan.userIds)
	assert.NotNil(t, newMan.authInfo)
	assert.NotNil(t, newMan.users)
	assert.NotNil(t, newMan.userHistories)
	assert.NotNil(t, newMan.CardMan)
}

func Test_InMemoryUserManagement_GetHistory_CardSetError_ErrorReturned(t *testing.T) {
	uMan := NewInMemoryManagement()
	cMan := &card.MockCardManagement{}
	uMan.CardMan = cMan
	testUser := test.GetFakeUser()
	blankHistory := pbuf.UserHistory{}

	cMan.On("GetCardSetById", mock.Anything).Return(pbuf.CardSet{}, errors.New(""))

	result, err := uMan.GetHistory(testUser, "")

	test.AssertProtoEq(t, &blankHistory, &result, "Should have returned blank Card Set.")
	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_GetHistory_UserIsAnon_ReturnGeneratedHistory(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUser := test.GetFakeUser()
	testUser.IsAnon = true
	fakeSet := test.GetFakeCardSet()
	expectedResult := card.GenerateEmptyHistory(&fakeSet)

	result, err := uMan.GetHistory(testUser, "")

	test.AssertProtoEq(t, &expectedResult, &result, "Should generate empty history for anon.")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_GetHistory_NoSavedHistory_ReturnGeneratedHistory(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUser := test.GetFakeUser()
	fakeSet := test.GetFakeCardSet()
	expectedResult := card.GenerateEmptyHistory(&fakeSet)
	expectedResult.UserId = testUser.Id

	result, err := uMan.GetHistory(testUser, "")

	test.AssertProtoEq(t, &expectedResult, &result, "Should generate empty history when none is saved.")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_GetHistory_HasSavedHistory_ReturnSavedHistory(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUser := test.GetFakeUser()
	fakeSet := test.GetFakeCardSet()
	historyKey := uMan.getHistoryKey(testUser.Id, fakeSet.Id)
	fakeHistory := test.GetFakeHistory()
	fakeHistory.PlayIndex = 123
	uMan.userHistories[historyKey] = fakeHistory

	result, err := uMan.GetHistory(testUser, fakeSet.Id)

	test.AssertProtoEq(t, &fakeHistory, &result, "Should return saved history.")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_GetHistory_MismatchingVersion_ReturnGeneratedHistory(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUser := test.GetFakeUser()
	fakeSet := test.GetFakeCardSet()
	expectedResult := card.GenerateEmptyHistory(&fakeSet)
	expectedResult.UserId = testUser.Id
	historyKey := uMan.getHistoryKey(testUser.Id, fakeSet.Id)
	fakeHistory := test.GetFakeHistory()
	fakeHistory.PlayIndex = 123
	fakeHistory.SetVersion = fakeSet.Version + 1
	uMan.userHistories[historyKey] = fakeHistory

	result, err := uMan.GetHistory(testUser, fakeSet.Id)

	test.AssertProtoEq(t, &expectedResult, &result, "Should return saved history.")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_GetAuthInfoByUserName_NoIdMap_ReturnsError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.UserAuthInfo{}

	result, err := uMan.GetAuthInfoByUserName("I don't exist")

	test.AssertProtoEq(t, &expected, &result, "Should return blank info")
	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_GetAuthInfoByUserName_NoIdInfo_ReturnsError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.UserAuthInfo{}
	userName := "I don't exist :D"
	uMan.userIds[userName] = "kbart"

	result, err := uMan.GetAuthInfoByUserName(userName)

	test.AssertProtoEq(t, &expected, &result, "Should return blank info")
	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_GetAuthInfoByUserName_GoodInfo_ReturnsInfo(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.UserAuthInfo{PassHash: "asda4we4as"}
	userName := "I do exist :P"
	uMan.userIds[userName] = "Asdf"
	uMan.authInfo["Asdf"] = expected

	result, err := uMan.GetAuthInfoByUserName(userName)

	test.AssertProtoEq(t, &expected, &result, "Should return saved info")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_GetAuthInfoById_NoInfo_ReturnsError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.UserAuthInfo{}

	result, err := uMan.GetAuthInfoById("asdf")

	test.AssertProtoEq(t, &expected, &result, "Should return blank info")
	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_GetAuthInfoById_HasInfo_ReturnsInfo(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.UserAuthInfo{PassHash: "asda4we4as"}
	uMan.authInfo["asdf"] = expected

	result, err := uMan.GetAuthInfoById("asdf")

	test.AssertProtoEq(t, &expected, &result, "Should return saved info")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_GetUserByUserName_NoIdMap_ReturnsError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.User{}

	result, err := uMan.GetUserByUserName("I don't exist")

	test.AssertProtoEq(t, &expected, &result, "Should return blank user")
	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_GetUserByUserName_NoIdInfo_ReturnsError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.User{}
	userName := "I don't exist :D"
	uMan.userIds[userName] = "RASDR"

	result, err := uMan.GetUserByUserName(userName)

	test.AssertProtoEq(t, &expected, &result, "Should return blank user")
	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_GetUserByUserName_GoodInfo_ReturnsInfo(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.User{FirstName: "Kyle"}
	userName := "I do exist :P"
	uMan.userIds[userName] = "Rasdr"
	uMan.users["Rasdr"] = expected

	result, err := uMan.GetUserByUserName(userName)

	test.AssertProtoEq(t, &expected, &result, "Should return saved user")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_GetUserById_NoInfo_ReturnsError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.User{}

	result, err := uMan.GetUserById("qqq")

	test.AssertProtoEq(t, &expected, &result, "Should return blank user")
	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_GetUserById_HasInfo_ReturnsInfo(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	expected := pbuf.User{FirstName: "Kyle"}
	uMan.users["kb"] = expected

	result, err := uMan.GetUserById("kb")

	test.AssertProtoEq(t, &expected, &result, "Should return saved user")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_UpdateHistory_NoHistoryFound_ReturnsError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUpdate := pbuf.CardUpdate{CardId: "cardId", Score: int32(1)}
	testUser := test.GetFakeUser()
	testCardSet := test.GetFakeCardSet()

	err := uMan.UpdateHistory(testUser, testCardSet.Id, testUpdate)

	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_UpdateHistory_HistoryExists_GetsUpdated(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUpdate := pbuf.CardUpdate{CardId: "c3", Score: int32(1)}
	testUser := test.GetFakeUser()
	testCardSet := test.GetFakeCardSet()
	testHistory := test.GetFakeHistory()
	fullId := uMan.getHistoryKey(testUser.Id, testCardSet.Id)
	uMan.userHistories[fullId] = testHistory
	expected := test.GetFakeHistory()
	expected.UserId = testUser.Id
	expected.PlayIndex += 1
	expected.History[2].CurrentScore += testUpdate.Score
	expected.History[2].Scores = append(expected.History[2].Scores, testUpdate.Score)
	expected.History[2].Indicies = append(expected.History[2].Indicies, testHistory.PlayIndex)

	err := uMan.UpdateHistory(testUser, testCardSet.Id, testUpdate)

	actual := uMan.userHistories[fullId]

	test.AssertProtoEq(t, &expected, &actual, "Result was not updated.")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_AddUser_UserNameExists_ReturnError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()

	uMan.users["rea354a4"] = testUser

	id, err := uMan.AddUser(testUser, testAuth)

	assert.NotNil(t, err)
	assert.Equal(t, id, "")
}

func Test_InMemoryUserManagement_AddUser_IdExists_ReturnsError(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()

	uMan.users["0"] = pbuf.User{}

	id, err := uMan.AddUser(testUser, testAuth)

	assert.NotNil(t, err)
	assert.Equal(t, id, "")
}

func Test_InMemoryUserManagement_AddUser_NoProblems_UserAdded(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()
	uMan.userCounter = 13

	id, err := uMan.AddUser(testUser, testAuth)

	testUser.Id = id
	actualUser := uMan.users[id]
	actualAuth := uMan.authInfo[id]

	expected, _ := strconv.Atoi(id)

	test.AssertProtoEq(t, &testUser, &actualUser, "User not saved in users map")
	test.AssertProtoEq(t, &testAuth, &actualAuth, "AuthInfo not saved in authInfo map")
	assert.Equal(t, expected, uMan.userCounter-1)
	assert.Equal(t, id, uMan.userIds[testUser.UserName], "Entry not saved in username-id map")
	assert.Nil(t, err)
}

func Test_InMemoryUserManagement_DeleteUser_NoUser_ErrorReturned(t *testing.T) {
	uMan, _ := getMockedUserManagement()

	err := uMan.DeleteUser("asdf")

	assert.NotNil(t, err)
}

func Test_InMemoryUserManagement_DeleteUser_EverythingDeleted(t *testing.T) {
	uMan, _ := getMockedUserManagement()
	testUser := test.GetFakeUser()
	testAuth := test.GetFakeAuthInfo()
	uMan.users[testUser.Id] = testUser
	uMan.userIds[testUser.UserName] = testUser.Id
	uMan.authInfo[testUser.Id] = testAuth

	uMan.DeleteUser(testUser.Id)

	assert.Equal(t, 0, len(uMan.users), "There should be no users left.")
	assert.Equal(t, 0, len(uMan.userIds), "There should be no userIds left.")
	assert.Equal(t, 0, len(uMan.authInfo), "There should be no authInfos left.")
}
