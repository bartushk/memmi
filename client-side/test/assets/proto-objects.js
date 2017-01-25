import pbuf from '../../src/pbuf/pbuf'
const proto = pbuf.pbuf

const fakeUser = {
  id:                   new Buffer('testUser'),
  userName:             'd_cat',
  firstName:            'Doodie',
  lastName:             'Cat',
  email:                'Dooie@cat.com',
  isAuthenticated:      false
}

const fakeCardOne = {
  id:           new Buffer('cardOne'),
  cardIndex:    0,
  title:        'Card One',
  front:        {type: 'html', value: '<h1>Hello</h1>'},
  back:         {type: 'html', value: '<h1>GoodBye</h1>'}
}

const fakeHistoryOne = {
  cardId:       new Buffer('cardOne'),
  currentScore: 1,
  cardIndex:    0,
  scores:       [1],
  indicies:     [0]
}

const fakeCardTwo = {
  id:           new Buffer('cardTwo'),
  cardIndex:    1,
  title:        'Card Two',
  front:        {type: 'html', value: '<h1>Ni Hao</h1>'},
  back:         {type: 'html', value: '<h1>Zai Jian</h1>'}
}

const fakeHistoryTwo = {
  cardId:       new Buffer('cardTwo'),
  currentScore: -2,
  cardIndex:    1,
  scores:       [-2],
  indicies:     [1]
}

const fakeCardSet = {
  id:           new Buffer('testCardSet'),
  version:      0,
  createdDate:  Date.now().toString(),
  authorId:     new Buffer('testUser'),
  setName:      'CatSet',
  title:        'Cat Cards',
  cards:        [fakeCardOne, fakeCardTwo]
}

const fakeCardUpdate = {
  cardId:       new Buffer('coolCard'),
  score:        2
}

const fakeUserHistory = {
  userId:     new Buffer('d_cat'),
  cardSetId:    new Buffer('testCardSet'),
  playIndex:    0,
  history:      [fakeHistoryOne, fakeHistoryTwo]
}

export default {
  getFakeCardSet:       () => { return proto.CardSet.create(fakeCardSet) },
  getFakeUserHistory: () => { return proto.UserHistory.create(fakeUserHistory) },
  getFakeCardUpdate:    () => { return proto.CardUpdate.create(fakeCardUpdate) },
  getFakeUser:          () => { return proto.User.create(fakeUser) },
  getFakeCardOne:       () => { return proto.Card.create(fakeCardOne) },
  getFakeCardTwo:       () => { return proto.Card.create(fakeCardTwo) }
}
