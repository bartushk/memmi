import 'whatwg-fetch'
import {put, call, take} from 'redux-saga/effects'


export function* sendRequest() {
  while (true) {
    const action = yield take('TEST')
    console.log(action)
    const result = yield call(fetch, '/api', { method: 'post', body: JSON.stringify(action.value)})
    console.log('-')
    console.log(result)
    console.log('-')
    yield put({type: 'TEST_DONE', value: {}})
  }
}


export function* sagas() {
  yield [ sendRequest() ]
}
