import {put, take, call, select} from 'redux-saga/effects'
import pbuf from '../../src/pbuf/pbuf'
import api from './proto-api'
import {selectCardId} from '../reducers/main/selectors'

const url = 'api/card/get-next'
const proto = pbuf.pbuf

export function* getCard() {
  for (;;) {
    yield take('GET_CARD')
    const cardRequest = yield select(selectCardId)
    const requestBody = proto.CardRequest.create(cardRequest)
    const result = yield call(api, url, requestBody, 'card-request', 'card')
    yield put({type: 'PUT_FETCHED_CARD', value: result})
  }
}


export function* sagas() {
  yield [ getCard() ]
}
