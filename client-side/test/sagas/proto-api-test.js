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
  let result = api(url, data.getFakeCardOne(), 'card', null)

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
    data.getFakeCardOne(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with user type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getFakeUser(), 'user', null)

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
    data.getFakeUser(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with cardset type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getFakeCardSet(), 'card-set', null)

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
    data.getFakeCardSet(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with card-update type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getFakeCardUpdate(), 'card-update', null)

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
    data.getFakeCardUpdate(),
    'Decoded value should match input.'
  )

  assert.end()
})

test('Request with player-history type.', (assert) => {
  const url = 'api'
  let result = api(url, data.getFakeUserHistory(), 'player-history', null)

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
    data.getFakeUserHistory(),
    'Decoded value should match input.'
  )

  assert.end()
})
