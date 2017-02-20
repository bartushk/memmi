package user

import (
	"errors"
	"fmt"
	"memmi/card"
	"memmi/pbuf"
)

func NewInMemoryManagement() *InMemoryUserManagement {
	newMan := &InMemoryUserManagement{
		userIds:       make(map[string]int64),
		authInfo:      make(map[int64]pbuf.UserAuthInfo),
		users:         make(map[int64]pbuf.User),
		userHistories: make(map[string]pbuf.UserHistory),
		CardMan:       card.NewInMemoryManagement(),
	}
	return newMan
}

type InMemoryUserManagement struct {
	userIds       map[string]int64
	authInfo      map[int64]pbuf.UserAuthInfo
	users         map[int64]pbuf.User
	userHistories map[string]pbuf.UserHistory
	userCounter   int64
	CardMan       card.CardManagement
}

func (manager *InMemoryUserManagement) getHistoryKey(id1 int64, id2 int64) string {
	return fmt.Sprintf("%x:%x", id1, id2)
}

func (manager *InMemoryUserManagement) GetHistory(user pbuf.User, cardSetId int64) (pbuf.UserHistory, error) {
	fullId := manager.getHistoryKey(user.Id, cardSetId)
	savedHistory, ok := manager.userHistories[fullId]
	set, err := manager.CardMan.GetCardSetById(cardSetId)

	if err != nil {
		return pbuf.UserHistory{}, err
	}

	if user.IsAnon {
		return card.GenerateEmptyHistory(&set), nil
	}

	// If not okay, generate a new blank set for this user
	if !ok {
		newHistory := card.GenerateEmptyHistory(&set)
		manager.userHistories[fullId] = newHistory
		return newHistory, nil
	}

	// Check if the history is the correct version, generate a new one if it isn't
	if savedHistory.SetVersion != set.Version {
		updatedHistory := card.GenerateEmptyHistory(&set)
		manager.userHistories[fullId] = updatedHistory
		return updatedHistory, nil
	}

	return savedHistory, nil
}

func (manager *InMemoryUserManagement) GetAuthInfoByUserName(userName string) (pbuf.UserAuthInfo, error) {
	userId, ok := manager.userIds[userName]
	if !ok {
		return pbuf.UserAuthInfo{}, errors.New("No auth info found.")
	}
	return manager.GetAuthInfoById(userId)
}

func (manager *InMemoryUserManagement) GetAuthInfoById(userId int64) (pbuf.UserAuthInfo, error) {
	result, ok := manager.authInfo[userId]
	if !ok {
		return pbuf.UserAuthInfo{}, errors.New("No auth info found.")
	}
	return result, nil
}

func (manager *InMemoryUserManagement) GetUserByUserName(userName string) (pbuf.User, error) {
	userId, ok := manager.userIds[userName]
	if !ok {
		return pbuf.User{}, errors.New("No user found.")
	}
	return manager.GetUserById(userId)
}

func (manager *InMemoryUserManagement) GetUserById(userId int64) (pbuf.User, error) {
	result, ok := manager.users[userId]
	if !ok {
		return pbuf.User{}, errors.New("No user found.")
	}
	return result, nil
}

func (manager *InMemoryUserManagement) UpdateHistory(user pbuf.User, cardSetId int64, update pbuf.CardUpdate) error {
	fullId := manager.getHistoryKey(user.Id, cardSetId)
	history, ok := manager.userHistories[fullId]
	if !ok {
		return errors.New("Could not find history to update.")
	}
	for _, cardHistory := range history.History {
		if cardHistory.CardId == update.CardId {
			cardHistory.CurrentScore += update.Score
			cardHistory.Scores = append(cardHistory.Scores, update.Score)
			cardHistory.Indicies = append(cardHistory.Indicies, history.PlayIndex)
			history.PlayIndex += 1
		}
	}
	return nil
}

func (management *InMemoryUserManagement) AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) (int64, error) {
	for _, savedUser := range management.users {
		if savedUser.UserName == user.UserName {
			return 0, errors.New("Tried adding a user that already exists userNames must be unique.")
		}
	}
	id := management.userCounter
	_, ok := management.users[id]
	if ok {
		return 0, errors.New("Tried adding a user at an occupied id.")
	}
	user.Id = id
	management.users[id] = user
	management.userIds[user.UserName] = id
	management.authInfo[id] = authInfo
	management.userCounter += 1
	return id, nil
}

func (management *InMemoryUserManagement) DeleteUser(userId int64) error {
	user, ok := management.users[userId]
	if !ok {
		return errors.New("User did not exisst and could not be deleted,")
	}
	delete(management.userIds, user.UserName)
	delete(management.authInfo, userId)
	delete(management.users, userId)
	return nil
}
