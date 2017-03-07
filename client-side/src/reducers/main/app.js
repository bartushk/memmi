import combineReducers from 'flat-combine-reducers'
import putReducers from './putters'


export default combineReducers(
  putReducers
)
