import pbuf from '../../src/pbuf/pbuf'
const proto = pbuf.pbuf

const user = {
  id:                   'user',
  userName:             'd_cat',
  firstName:            'Doodie',
  lastName:             'Cat',
  email:                'Doodie@cat.com',
  isAuthenticated:      false,
  joinedDate:           new Date().toISOString()
}

const cardOne = {
  id:           'c1',
  title:        'Card One',
  front:        {type: 'html', value: '<h1>Hello</h1>'},
  back:         {type: 'html', value: '<h1>GoodBye</h1>'},
  tags:         ['Cooler', 'Nerder']
}

const historyOne = {
  cardId:       'c1',
  currentScore: 1,
  cardIndex:    0,
  scores:       [1],
  indicies:     [0]
}

const cardTwo = {
  id:           'c2',
  title:        'Card Two',
  front:        {type: 'html', value: '<h1>Ni Hao</h1>'},
  back:         {type: 'html', value: '<h1>Zai Jian</h1>'},
  tags:         ['Coolest', 'Nerdest']
}

const historyTwo = {
  cardId:       'c2',
  currentScore: -2,
  cardIndex:    1,
  scores:       [-2],
  indicies:     [1]
}

const cardSet = {
  id:           'cs1',
  version:      0,
  createdDate:  new Date().toISOString(),
  authorId:     'user',
  title:        'Cat Cards',
  cardIds:      [cardOne.id, cardTwo.id],
  tags:         ['Cool', 'Nerd']
}

const cardUpdate = {
  cardId:       'c1',
  score:        2
}

const userHistory = {
  userId:       'user',
  cardSetId:    'cs1',
  playIndex:    0,
  history:      [historyOne, historyTwo]
}

const nextCardRequest = {
  cardSetId: 'cs1',
  previousCardId: 'c1',
  algorithm: 0
}

const cardScoreReport  = {
  cardSetId: 'cs1',
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
  id: 'cs1'
}

const cardRequest = {
  id: 'c2'
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
