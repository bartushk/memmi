import {put, take, call, select} from 'redux-saga/effects'
import pbuf from '../../src/pbuf/pbuf'
import api from './proto-api'
import {nextCardRequest} from '../reducers/main/selectors'

const url = 'api/card/get-next'
const proto = pbuf.pbuf

function* getNext() {
  for (;;) {
    yield take('GET_NEXT_CARD')
    const cardRequest = yield select(nextCardRequest)
    const requestBody = proto.NextCardRequest.create(cardRequest)
    const result = yield call(api, url, requestBody, 'next-card-request', 'card')
    yield put({type: 'PUT_FETCHED_CARD', value: result})
  }
}


export function* sagas() {
  yield [ getNext() ]
}
