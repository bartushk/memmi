import test from 'tape'
import {take, select, call, put} from 'redux-saga/effects'
import pbuf from '../../src/pbuf/pbuf'

import data from '../assets/proto-objects'
import {getCard} from '../../src/sagas/index'
import api from '../../src/sagas/proto-api'
import {selectCardId} from '../../src/reducers/main/selectors'

const url = 'api/card/get-next'
const proto = pbuf.pbuf

test('getCard saga.', (assert) => {
  const gen = getCard()
  const fakeRequest = data.getCardRequest()
  const fakeApiResult = {asdf:1234}
  assert.deepEqual(
    gen.next().value,
    take('GET_CARD'),
    'Must take the GET_NEXT_CARD action.'
  )

  assert.deepEqual(
    gen.next().value,
    select(selectCardId),
    'Select the next request from the app state.'
  )

  assert.deepEqual(
    gen.next(fakeRequest).value,
    call(api, url, proto.CardRequest.create(fakeRequest),
         'card-request', 'card'),
    'Must call the proto api with the correct arguments.'
  )

  assert.deepEqual(
    gen.next(fakeApiResult).value,
    put({type: 'PUT_FETCHED_CARD', value: fakeApiResult}),
    'Must put the result of the api call through the right action.'
  )

  assert.end()
})
