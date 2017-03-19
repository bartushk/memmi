import pbuf from '../../../src/pbuf/pbuf'

const proto = pbuf.pbuf

// TODO: Actually write good logic here.
export const selectCardId = state => {
  if (state !== null) {
    const cardRequest = {
      id: '0'
    }
    return proto.CardRequest.create(cardRequest)
  }
  return {}
}
