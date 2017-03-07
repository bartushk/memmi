import { combineReducers } from 'redux'
import { routerReducer } from 'react-router-redux'
import appReducers from './main/app'

export const reducers = combineReducers({
  routing: routerReducer,
  app: appReducers
})
