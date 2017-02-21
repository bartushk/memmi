import pbuf from '../../src/pbuf/pbuf'
import {util} from 'protobufjs'
const proto = pbuf.pbuf

const fakeUser = {
  id:                   util.Long.fromValue(0),
  userName:             'd_cat',
  firstName:            'Doodie',
  lastName:             'Cat',
  email:                'Dooie@cat.com',
  isAuthenticated:      false,
  joinedDate:  util.Long.fromValue(10)
}

const fakeCardOne = {
  id:           util.Long.fromValue(0),
  title:        'Card One',
  front:        {type: 'html', value: '<h1>Hello</h1>'},
  back:         {type: 'html', value: '<h1>GoodBye</h1>'},
  tags:         ['Cooler', 'Nerder']
}

const fakeHistoryOne = {
  cardId:       util.Long.fromValue(0),
  currentScore: 1,
  cardIndex:    0,
  scores:       [1],
  indicies:     [0]
}

const fakeCardTwo = {
  id:           util.Long.fromValue(1),
  title:        'Card Two',
  front:        {type: 'html', value: '<h1>Ni Hao</h1>'},
  back:         {type: 'html', value: '<h1>Zai Jian</h1>'},
  tags:         ['Coolest', 'Nerdest']
}

const fakeHistoryTwo = {
  cardId:       util.Long.fromValue(1),
  currentScore: -2,
  cardIndex:    1,
  scores:       [-2],
  indicies:     [1]
}

const fakeCardSet = {
  id:           util.Long.fromValue(0),
  version:      0,
  createdDate:  util.Long.fromValue(10),
  authorId:     util.Long.fromValue(0),
  title:        'Cat Cards',
  cardIds:      [fakeCardOne.id, fakeCardTwo.id],
  tags:         ['Cool', 'Nerd']
}

const fakeCardUpdate = {
  cardId:       util.Long.fromValue(0),
  score:        2
}

const fakeUserHistory = {
  userId:       util.Long.fromValue(0),
  cardSetId:    util.Long.fromValue(0),
  playIndex:    0,
  history:      [fakeHistoryOne, fakeHistoryTwo]
}

const nextCardRequest = {
  cardSetId: util.Long.fromValue(1),
  previousCardId: util.Long.fromValue(1),
  algorithm: 0
}

export default {
  getFakeCardSet:       () => { return proto.CardSet.create(fakeCardSet) },
  getFakeUserHistory:   () => { return proto.UserHistory.create(fakeUserHistory) },
  getFakeCardUpdate:    () => { return proto.CardUpdate.create(fakeCardUpdate) },
  getFakeUser:          () => { return proto.User.create(fakeUser) },
  getFakeCardOne:       () => { return proto.Card.create(fakeCardOne) },
  getFakeCardTwo:       () => { return proto.Card.create(fakeCardTwo) },
  getNextCardRequest:   () => { return proto.NextCardRequest.create(nextCardRequest)}
}
