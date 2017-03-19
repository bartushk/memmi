import test from 'tape'
import data from '../assets/proto-objects'
import {SetFetch}  from '../../src/sagas/proto-api'
import api  from '../../src/sagas/proto-api'

function mockFetch(url, args, decode, encode) {
  return {url, args, decode, encode}
}

// Sets a different fetch function to help with testing.
SetFetch(mockFetch)

test('Request with card type.', (assert) => {
  const url = '/api'
  let result = api(url, data.getCardOne(), 'card', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getCardOne(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with user type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getUser(), 'user', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getUser(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with cardset type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getCardSet(), 'card-set', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getCardSet(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with card-update type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getCardUpdate(), 'card-update', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getCardUpdate(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with player-history type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getUserHistory(), 'player-history', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getUserHistory(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with CardScoreReport type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getCardScoreReport(), 'card-score-report', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getCardScoreReport(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with RequestError type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getRequestError(), 'request-error', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getRequestError(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with UpdateResponse type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getUpdateResponse(), 'update-response', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getUpdateResponse(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with CardSetRequest type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getCardSetRequest(), 'card-set-request', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getCardSetRequest(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with CardRequest type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getCardRequest(), 'card-request', null)

  assert.equal(result.url, url, 'Correct url passed to fetch.')
  assert.equal(result.args.method, 'POST', 'Method should be post.')

  assert.equal(
    result.args.headers['Content-Type'],
    'application/x-protobuf',
    'Content type should be "application/x-protobuf"'
  )

  assert.deepEqual(
    result.encode,
    result.args.body,
    'Request body should be encoded protobuf.'
  )

  assert.deepEqual(
    result.decode,
    data.getCardRequest(),
    'Decoded value should match input.'
  )

  assert.end()
})
