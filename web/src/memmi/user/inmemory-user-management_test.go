package user

import (
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
