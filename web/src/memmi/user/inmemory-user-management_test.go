package user

import (
	"github.com/golang/protobuf/proto"
	"memmi/pbuf"
	"testing"
)

func Test_InMemoryUserManagement_NewInMemoryUserManagement(t *testing.T) {
	newMan := NewInMemoryManagement()

	if newMan.userIds == nil {
		t.Error("userIds should be intitialized.")
	}

	if newMan.authInfo == nil {
		t.Error("authInfo should be intitialized.")
	}

	if newMan.users == nil {
		t.Error("users should be intitialized.")
	}

	if newMan.userHistories == nil {
		t.Error("userHistories should be intitialized.")
	}

	if newMan.CardMan == nil {
		t.Error("CardMan should be initialized.")
	}
}

func Test_InMemoryUserManagement_GetHistory_BadId_ReturnBlank(t *testing.T) {
	newMan := NewInMemoryManagement()
	testUser := getTestUser()
	blank := pbuf.UserHistory{}

	result, err := newMan.GetHistory(testUser, 0)

	if !proto.Equal(&blank, &result) {
		t.Error("Wrong history returned.",
			"Expected:", blank,
			"Got:", result)
	}

	if err == nil {
		t.Error("An error should have been returned.")
	}
}
