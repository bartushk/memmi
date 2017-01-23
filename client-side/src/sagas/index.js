import {put, take} from 'redux-saga/effects'

export function* sendTest() {
  while (true) {
    yield take('TEST')
    yield put({type: 'TEST_DONE', value: {}})
  }
}


export function* sagas() {
  yield [ sendTest() ]
}
