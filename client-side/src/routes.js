import React                           from 'react'
import { IndexRoute, Route }           from 'react-router'
import App                             from 'components/App'
import Home                            from 'pages/Home'

export default (

  <Route path="/" component={App}>
    <IndexRoute component={Home} />
  </Route>

)
