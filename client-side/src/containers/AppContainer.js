import React, { Component, PropTypes } from 'react'
import { Router }                      from 'react-router'
import { Provider }                    from 'react-redux'
import routes                          from 'routes'

class AppContainer extends Component {
  static propTypes = {
    store   : PropTypes.object.isRequired,
    history : PropTypes.object.isRequired
  }

  shouldComponentUpdate() {
    return false
  }

  render() {
    const { history, store } = this.props

    return (
      <Provider store={store}>
        <Router history={history} children={routes} />
      </Provider>
    )
  }
}

export default AppContainer
