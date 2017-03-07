import pbuf from '../../../src/pbuf/pbuf'

const proto = pbuf.pbuf

// TODO: Actually write good logic here.
export const nextCardRequest = state => {
  if (state !== null) {
    const cardRequest = {
      cardSetId: '0',
      previousCardId: '0',
      algorithm: 1
    }
    return proto.NextCardRequest.create(cardRequest)
  }
  return {}
}
