import {put, take, call} from 'redux-saga/effects'
import pbuf from '../../src/pbuf/pbuf'
import api from './proto-api'

const url = 'api/card/get-next'
const proto = pbuf.pbuf
const cardRequest = {
  cardSetId: '0',
  previousCardId: '0',
  algorithm: 1
}

export function* sendTest() {
  while (true) {
    yield take('TEST')
    console.log(cardRequest)
    const requestBody = proto.NextCardRequest.create(cardRequest)
    console.log(requestBody)
    const result = yield call(api, url, requestBody, 'next-card-request', 'card')
    console.log(result)
    yield put({type: 'TEST_DONE', value: {}})
  }
}


export function* sagas() {
  yield [ sendTest() ]
}
