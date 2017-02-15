package user

import (
	"errors"
	"fmt"
	"hash/fnv"
	"memmi/card"
	"memmi/pbuf"
	"reflect"
)

type InMemoryUserManagement struct {
	userIds       map[string][]byte
	authInfo      map[string]pbuf.UserAuthInfo
	userHistories map[string]pbuf.UserHistory
	userCounter   int64
	cardMan       card.CardManagement
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
	result, ok := manager.userHistories[key]
	if !ok {
		set := manager.cardMan.GetCardSetById(cardSetId)
		result = card.GenerateEmptyHistory(&set)
	}
	return result
}

func (manager *InMemoryUserManagement) GetAuthInfoByUserName(userName string) pbuf.UserAuthInfo {
	userId := manager.userIds[userName]
	return manager.GetAuthInfoById(userId)
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

func (manager *InMemoryUserManagement) GetAuthInfoById(userId []byte) pbuf.UserAuthInfo {
	key := manager.getKey(userId)
	return manager.authInfo[key]
}

func (management *InMemoryUserManagement) AddUser(user pbuf.User, authInfo pbuf.UserAuthInfo) error {
	return nil
}

func (management *InMemoryUserManagement) DeleteUser(userId []byte) error {
	return nil
}
