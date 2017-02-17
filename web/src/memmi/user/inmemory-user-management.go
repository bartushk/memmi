package user

import (
	"errors"
	"fmt"
	"hash/fnv"
	"memmi/card"
	"memmi/pbuf"
	"reflect"
)

func NewInMemoryManagement() *InMemoryUserManagement {
	newMan := &InMemoryUserManagement{
		userIds:       make(map[string][]byte),
		authInfo:      make(map[string]pbuf.UserAuthInfo),
		users:         make(map[string]pbuf.User),
		userHistories: make(map[string]pbuf.UserHistory),
		CardMan:       card.NewInMemoryManagement(),
	}
	return newMan
}

type InMemoryUserManagement struct {
	userIds       map[string][]byte
	authInfo      map[string]pbuf.UserAuthInfo
	users         map[string]pbuf.User
	userHistories map[string]pbuf.UserHistory
	userCounter   uint32
	CardMan       card.CardManagement
}

func (manager *InMemoryUserManagement) getKey(id []byte) string {
	return fmt.Sprintf("%x", id)
}

func (manager *InMemoryUserManagement) getId(input uint32) []byte {
	hash := fnv.New64()
	hash.Write([]byte(fmt.Sprintf("%x", input)))
	return hash.Sum(nil)
}

func (manager *InMemoryUserManagement) GetHistory(user pbuf.User, cardSetId []byte) pbuf.UserHistory {
	fullId := append(user.Id, cardSetId...)
	key := manager.getKey(fullId)
	savedHistory, ok := manager.userHistories[key]
	set, err := manager.CardMan.GetCardSetById(cardSetId)

	if err != nil {
		return pbuf.UserHistory{}
	}

	// If not okay, generate a new blank set for this user
	if !ok {
		newHistory := card.GenerateEmptyHistory(&set)
		manager.userHistories[key] = newHistory
		return newHistory
	}

	// Check if the history is the correct version, generate a new one if it isn't
	if savedHistory.SetVersion != set.Version {
		updatedHistory := card.GenerateEmptyHistory(&set)
		manager.userHistories[key] = updatedHistory
		return updatedHistory
	}

	return savedHistory
}

func (manager *InMemoryUserManagement) GetAuthInfoByUserName(userName string) pbuf.UserAuthInfo {
	userId := manager.userIds[userName]
	return manager.GetAuthInfoById(userId)
}

func (manager *InMemoryUserManagement) GetAuthInfoById(userId []byte) pbuf.UserAuthInfo {
	key := manager.getKey(userId)
	return manager.authInfo[key]
}

func (manager *InMemoryUserManagement) GetUserByUserName(userName string) pbuf.User {
	userId := manager.userIds[userName]
	return manager.GetUserById(userId)
}

func (manager *InMemoryUserManagement) GetUserById(userId []byte) pbuf.User {
	key := manager.getKey(userId)
	return manager.users[key]
}

func (manager *InMemoryUserManagement) UpdateHistory(user pbuf.User, cardSetId []byte, update pbuf.CardUpdate) error {
	fullId := append(user.Id, cardSetId...)
	key := manager.getKey(fullId)
	history, ok := manager.userHistories[key]
	if !ok {
		return errors.New("Could not find history to update.")
	}
	for _, cardHistory := range history.History {
		if reflect.DeepEqual(cardHistory.CardId, update.CardId) {
			cardHistory.CurrentScore += update.Score
			cardHistory.Scores = append(cardHistory.Scores, update.Score)
			cardHistory.Indicies = append(cardHistory.Indicies, history.PlayIndex)
			history.PlayIndex += 1
		}
	}
	return nil
}

func (management *InMemoryUserManagement) AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) error {
	for _, savedUser := range management.users {
		if savedUser.UserName == user.UserName {
			return errors.New("UserNames must be unique.")
		}
	}
	id := management.getId(management.userCounter)
	user.Id = id
	key := management.getKey(id)
	management.users[key] = user
	management.userIds[user.UserName] = id
	management.authInfo[key] = authInfo
	management.userCounter += 1
	return nil
}

func (management *InMemoryUserManagement) DeleteUser(userId []byte) error {
	key := management.getKey(userId)
	user, ok := management.users[key]
	if !ok {
		return errors.New("User did not exisst and could not be deleted,")
	}
	delete(management.userIds, user.UserName)
	delete(management.authInfo, key)
	delete(management.users, key)
	return nil
}
