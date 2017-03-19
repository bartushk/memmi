export default function putters(state = {}, action) {
  if (state.cardCache === undefined) {
    state.cardCache = {}
  }

  if (state.cardHistory === undefined) {
    state.cardHistory = []
  }

  switch (action.type) {

  case 'PUT_FETCHED_CARD':
    state.cardCache[action.value.id] = action.value
    state.cardHistory = state.cardHistory.concat(action.value.id)
    console.log(state)
    break

  default:
    break
  }
  return state
}
