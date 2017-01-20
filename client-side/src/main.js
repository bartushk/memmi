import React                                         from 'react'
import ReactDOM                                      from 'react-dom'
import { createStore, applyMiddleware, compose }     from 'redux'
import { browserHistory }                            from 'react-router'
import { syncHistoryWithStore, routerMiddleware }    from 'react-router-redux'
import createSagaMiddleware                          from 'redux-saga'
import { reducers }                                  from './reducers/index'
import { sagas }                                     from './sagas/index'
import AppContainer                                  from 'containers/AppContainer'
import injectTapEventPlugin                          from 'react-tap-event-plugin'
import './components/main.scss'

injectTapEventPlugin()
// create the store
const sagaMiddleware = createSagaMiddleware()
let middleware = applyMiddleware(routerMiddleware(browserHistory), sagaMiddleware)
if (process.env.NODE_ENV !== 'production' && window.devToolsExtension) {
  middleware = compose(middleware, window.devToolsExtension && window.devToolsExtension())
}
const store = createStore(reducers, middleware)
const history = syncHistoryWithStore(browserHistory, store)
sagaMiddleware.run(sagas)

// render the main component
const MOUNT_NODE  = document.getElementById('app')
ReactDOM.render( <AppContainer history={history} store={store}/>, MOUNT_NODE)
