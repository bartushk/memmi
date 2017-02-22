import pbuf from '../../src/pbuf/pbuf'
import {util} from 'protobufjs'
const proto = pbuf.pbuf

const user = {
  id:                   util.Long.fromValue(0),
  userName:             'd_cat',
  firstName:            'Doodie',
  lastName:             'Cat',
  email:                'Dooie@cat.com',
  isAuthenticated:      false,
  joinedDate:  util.Long.fromValue(10)
}

const cardOne = {
  id:           util.Long.fromValue(0),
  title:        'Card One',
  front:        {type: 'html', value: '<h1>Hello</h1>'},
  back:         {type: 'html', value: '<h1>GoodBye</h1>'},
  tags:         ['Cooler', 'Nerder']
}

const historyOne = {
  cardId:       util.Long.fromValue(0),
  currentScore: 1,
  cardIndex:    0,
  scores:       [1],
  indicies:     [0]
}

const cardTwo = {
  id:           util.Long.fromValue(1),
  title:        'Card Two',
  front:        {type: 'html', value: '<h1>Ni Hao</h1>'},
  back:         {type: 'html', value: '<h1>Zai Jian</h1>'},
  tags:         ['Coolest', 'Nerdest']
}

const historyTwo = {
  cardId:       util.Long.fromValue(1),
  currentScore: -2,
  cardIndex:    1,
  scores:       [-2],
  indicies:     [1]
}

const cardSet = {
  id:           util.Long.fromValue(0),
  version:      0,
  createdDate:  util.Long.fromValue(10),
  authorId:     util.Long.fromValue(0),
  title:        'Cat Cards',
  cardIds:      [cardOne.id, cardTwo.id],
  tags:         ['Cool', 'Nerd']
}

const cardUpdate = {
  cardId:       util.Long.fromValue(0),
  score:        2
}

const userHistory = {
  userId:       util.Long.fromValue(0),
  cardSetId:    util.Long.fromValue(0),
  playIndex:    0,
  history:      [historyOne, historyTwo]
}

const nextCardRequest = {
  cardSetId: util.Long.fromValue(1),
  previousCardId: util.Long.fromValue(1),
  algorithm: 0
}

const cardScoreReport  = {
  cardSetId: util.Long.fromValue(1),
  update: cardUpdate
}

const reportAndNext = {
  nextRequest: nextCardRequest,
  report: cardScoreReport
}

const requestError = {
  reason: 'Just didn\'t work'
}

const updateResponse = {
  status: 0
}

const cardSetRequest = {
  id: util.Long.fromValue(0)
}

const cardRequest = {
  id: util.Long.fromValue(0)
}

export default {
  getCardSet:           () => { return proto.CardSet.create(cardSet) },
  getUserHistory:       () => { return proto.UserHistory.create(userHistory) },
  getCardUpdate:        () => { return proto.CardUpdate.create(cardUpdate) },
  getUser:              () => { return proto.User.create(user) },
  getCardOne:           () => { return proto.Card.create(cardOne) },
  getCardTwo:           () => { return proto.Card.create(cardTwo) },
  getNextCardRequest:   () => { return proto.NextCardRequest.create(nextCardRequest)},
  getCardScoreReport:   () => { return proto.CardScoreReport.create(cardScoreReport)},
  getReportAndNext:     () => { return proto.ReportAndNext.create(reportAndNext)},
  getRequestError:      () => { return proto.RequestError.create(requestError)},
  getUpdateResponse:    () => { return proto.UpdateResponse.create(updateResponse)},
  getCardSetRequest:    () => { return proto.CardSetRequest.create(cardSetRequest)},
  getCardRequest:       () => { return proto.CardRequest.create(cardRequest)}
}
