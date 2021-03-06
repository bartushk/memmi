import 'whatwg-fetch'
import pbuf from '../pbuf/pbuf'

const protos = pbuf.pbuf
// This is a hack as a way to shim in a different fetch function for unit testing.
let _fetch = null

export function SetFetch(newFetch) {
  _fetch = newFetch
}

function decode(data, type) {
  switch (type) {
  case 'user':
    return protos.User.decode(data)
  case 'card-set':
    return protos.CardSet.decode(data)
  case 'card':
    return protos.Card.decode(data)
  case 'card-update':
    return protos.CardUpdate.decode(data)
  case 'player-history':
    return protos.UserHistory.decode(data)
  case 'card-score-report':
    return protos.CardScoreReport.decode(data)
  case 'request-error':
    return protos.RequestError.decode(data)
  case 'update-response':
    return protos.UpdateResponse.decode(data)
  case 'card-set-request':
    return protos.CardSetRequest.decode(data)
  case 'card-request':
    return protos.CardRequest.decode(data)
  default:
    return {}
  }
}

function encode(proto, type) {
  switch (type) {
  case 'user':
    return protos.User.encode(proto).finish()
  case 'card-set':
    return protos.CardSet.encode(proto).finish()
  case 'card':
    return protos.Card.encode(proto).finish()
  case 'card-update':
    return protos.CardUpdate.encode(proto).finish()
  case 'player-history':
    return protos.UserHistory.encode(proto).finish()
  case 'card-score-report':
    return protos.CardScoreReport.encode(proto).finish()
  case 'request-error':
    return protos.RequestError.encode(proto).finish()
  case 'update-response':
    return protos.UpdateResponse.encode(proto).finish()
  case 'card-set-request':
    return protos.CardSetRequest.encode(proto).finish()
  case 'card-request':
    return protos.CardRequest.encode(proto).finish()
  default:
    throw exception('Invalid type to decode with: ' + type)
  }
}

function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response
  }
  const error = new Error(response.statusText)
  error.response = response
  throw error
}

export default function makeProtoRequest(url, requestProto, requestType, responseType) {
  let fetchArg = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-protobuf'
    },
    body: encode(requestProto, requestType)
  }
  // This is a hack as a way to shim in a different fetch function for unit testing.
  if ( _fetch ) {
    let encodedValue = encode(requestProto, requestType)
    let decodedValue = decode(encodedValue, requestType)
    return _fetch(url, fetchArg, decodedValue, encodedValue)
  }
  return fetch(url, fetchArg)
          .then(checkStatus)
          .then(response => response.arrayBuffer())
          .then((buf) => decode(new Buffer(buf), responseType))
          .catch((error) => {
            return decode(new Buffer(error.response.arrayBuffer()), 'request-error')
          })
}


