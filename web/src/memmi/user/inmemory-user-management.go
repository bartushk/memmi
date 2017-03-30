package user

import (
	"errors"
	"fmt"
	"github.com/op/go-logging"
	"memmi/card"
	"memmi/pbuf"
	"strconv"
)

var log = logging.MustGetLogger("memmi")

func NewInMemoryManagement() *InMemoryUserManagement {
	newMan := &InMemoryUserManagement{
		userIds:       make(map[string]string),
		authInfo:      make(map[string]pbuf.UserAuthInfo),
		users:         make(map[string]pbuf.User),
		userHistories: make(map[string]pbuf.UserHistory),
		CardMan:       card.NewInMemoryManagement(),
	}
	return newMan
}

type InMemoryUserManagement struct {
	userIds       map[string]string
	authInfo      map[string]pbuf.UserAuthInfo
	users         map[string]pbuf.User
	userHistories map[string]pbuf.UserHistory
	userCounter   int
	CardMan       card.CardManagement
}

func (manager *InMemoryUserManagement) getHistoryKey(id1 string, id2 string) string {
	return fmt.Sprintf("%s:%s", id1, id2)
}

func (manager *InMemoryUserManagement) GetHistory(user pbuf.User, cardSetId string) (pbuf.UserHistory, error) {
	fullId := manager.getHistoryKey(user.Id, cardSetId)
	savedHistory, ok := manager.userHistories[fullId]
	set, err := manager.CardMan.GetCardSetById(cardSetId)

	if err != nil {
		log.Errorf("Could not find card set: %s", cardSetId)
		return pbuf.UserHistory{}, err
	}

	if user.IsAnon {
		return card.GenerateEmptyHistory(&set), nil
	}

	// If not okay, generate a new blank set for this user
	if !ok {
		log.Infof("New history created user %s set %s.", user.Id, cardSetId)
		newHistory := card.GenerateEmptyHistory(&set)
		newHistory.UserId = user.Id
		manager.userHistories[fullId] = newHistory
		return newHistory, nil
	}

	// Check if the history is the correct version, generate a new one if it isn't
	if savedHistory.SetVersion != set.Version {
		log.Infof("Old version history for user %s set %s.", user.Id, cardSetId)
		log.Infof("Old version: %s New Version: %s", savedHistory.SetVersion, set.Version)
		updatedHistory := card.GenerateEmptyHistory(&set)
		updatedHistory.UserId = user.Id
		manager.userHistories[fullId] = updatedHistory
		return updatedHistory, nil
	}

	log.Infof("Clean fetch of user history for user %s set %s", user.Id, cardSetId)
	return savedHistory, nil
}

func (manager *InMemoryUserManagement) GetAuthInfoByUserName(userName string) (pbuf.UserAuthInfo, error) {
	userId, ok := manager.userIds[userName]
	if !ok {
		log.Errorf("Could not find authInfo for %s.", userName)
		return pbuf.UserAuthInfo{}, errors.New("No auth info found.")
	}
	return manager.GetAuthInfoById(userId)
}

func (manager *InMemoryUserManagement) GetAuthInfoById(userId string) (pbuf.UserAuthInfo, error) {
	result, ok := manager.authInfo[userId]
	if !ok {
		log.Errorf("Could not find authInfo for %s.", userId)
		return pbuf.UserAuthInfo{}, errors.New("No auth info found.")
	}
	return result, nil
}

func (manager *InMemoryUserManagement) GetUserByUserName(userName string) (pbuf.User, error) {
	userId, ok := manager.userIds[userName]
	if !ok {
		log.Errorf("Could not find authInfo for %s.", userName)
		return pbuf.User{}, errors.New("No user found.")
	}
	return manager.GetUserById(userId)
}

func (manager *InMemoryUserManagement) GetUserById(userId string) (pbuf.User, error) {
	result, ok := manager.users[userId]
	if !ok {
		log.Errorf("Could not find user for %s.", userId)
		return pbuf.User{}, errors.New("No user found.")
	}
	return result, nil
}

func (manager *InMemoryUserManagement) UpdateHistory(user pbuf.User, cardSetId string, update pbuf.CardUpdate) error {
	fullId := manager.getHistoryKey(user.Id, cardSetId)
	history, ok := manager.userHistories[fullId]
	if !ok {
		log.Errorf("Could not find history for %s %s", cardSetId, user)
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
	manager.userHistories[fullId] = history
	return nil
}

func (management *InMemoryUserManagement) AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) (string, error) {
	for _, savedUser := range management.users {
		if savedUser.UserName == user.UserName {
			log.Errorf("User already exists with username: %s", user)
			return "", errors.New("Tried adding a user that already exists userNames must be unique.")
		}
	}
	id := strconv.Itoa(management.userCounter)
	_, ok := management.users[id]
	if ok {
		log.Errorf("Id collision for user. %s %s", id, user)
		return "", errors.New("Tried adding a user at an occupied id.")
	}
	user.Id = id
	management.users[id] = user
	management.userIds[user.UserName] = id
	management.authInfo[id] = authInfo
	management.userCounter += 1
	log.Infof("User succesfully added: %s", user)
	return id, nil
}

func (management *InMemoryUserManagement) DeleteUser(userId string) error {
	user, ok := management.users[userId]
	if !ok {
		log.Errorf("User did not exist and could not be deleted: %s", user)
		return errors.New("User did not exist and could not be deleted,")
	}
	delete(management.userIds, user.UserName)
	delete(management.authInfo, userId)
	delete(management.users, userId)
	log.Infof("User deleted: %s", user)
	return nil
}
