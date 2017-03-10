import test from 'tape'
import {take, select, call, put} from 'redux-saga/effects'
import pbuf from '../../src/pbuf/pbuf'

import data from '../assets/proto-objects'
import {getNext} from '../../src/sagas/index'
import api from '../../src/sagas/proto-api'
import {nextCardRequest} from '../../src/reducers/main/selectors'

const url = 'api/card/get-next'
const proto = pbuf.pbuf

test('getNext saga.', (assert) => {
  const gen = getNext()
  const fakeRequest = data.getNextCardRequest()
  const fakeApiResult = {asdf:1234}
  assert.deepEqual(
    gen.next().value,
    take('GET_NEXT_CARD'),
    'Must take the GET_NEXT_CARD action.'
  )

  assert.deepEqual(
    gen.next().value,
    select(nextCardRequest),
    'Select the next request from the app state.'
  )

  assert.deepEqual(
    gen.next(fakeRequest).value,
    call(api, url, proto.NextCardRequest.create(fakeRequest),
         'next-card-request', 'card'),
    'Must call the proto api with the correct arguments.'
  )

  assert.deepEqual(
    gen.next(fakeApiResult).value,
    put({type: 'PUT_FETCHED_CARD', value: fakeApiResult}),
    'Must put the result of the api call through the right action.'
  )

  assert.end()
})
