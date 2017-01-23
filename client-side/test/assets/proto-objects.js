import pbuf from '../../src/pbuf/pbuf'
const proto = pbuf.pbuf

const fakeUser = {
  id:                   'testUser',
  userName:             'd_cat',
  firstName:            'Doodie',
  lastName:             'Cat',
  email:                'Dooie@cat.com',
  isAuthenticated:      false
}

const fakeCardOne = {
  id:           'cardOne',
  cardIndex:    0,
  title:        'Card One',
  front:        {type: 'html', value: '<h1>Hello</h1>'},
  back:         {type: 'html', value: '<h1>GoodBye</h1>'}
}

const fakeHistoryOne = {
  cardId:       'cardOne',
  currentScore: 1,
  cardIndex:    0,
  scores:       [1],
  indicies:     [0]
}

const fakeCardTwo = {
  id:           'cardTwo',
  cardIndex:    1,
  title:        'Card Two',
  front:        {type: 'html', value: '<h1>Ni Hao</h1>'},
  back:         {type: 'html', value: '<h1>Zai Jian</h1>'}
}

const fakeHistoryTwo = {
  cardId:       'cardTwo',
  currentScore: -2,
  cardIndex:    1,
  scores:       [-2],
  indicies:     [1]
}

const fakeCardSet = {
  id:           'testCardSet',
  version:      0,
  createdDate:  Date.now().toString(),
  authorId:     'testUser',
  setName:      'CatSet',
  title:        'Cat Cards',
  cards:        [fakeCardOne, fakeCardTwo]
}

const fakeCardUpdate = {
  cardId:       'coolCard',
  score:        2
}

const fakePlayerHistory = {
  playerId:     'd_cat',
  cardSetId:    'testCardSet',
  playIndex:    0,
  history:      [fakeHistoryOne, fakeHistoryTwo]
}

export default {
  getFakeCardSet:       () => { return proto.CardSet.create(fakeCardSet) },
  getFakePlayerHistory: () => { return proto.PlayerHistory.create(fakePlayerHistory) },
  getFakeCardUpdate:    () => { return proto.CardUpdate.create(fakeCardUpdate) },
  getFakeUser:          () => { return proto.User.create(fakeUser) },
  getFakeCardOne:       () => { return proto.Card.create(fakeCardOne) },
  getFakeCardTwo:       () => { return proto.Card.create(fakeCardTwo) }
}
