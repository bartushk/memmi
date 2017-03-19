import React from 'react'
import AppBar from 'material-ui/AppBar'
import logoUrl from './logo_67x50.svg'
import getAction from '../../action-wrapper'
import './Header.scss'


class Header extends React.Component {
  render() {
    const action = getAction()
    return (
      <AppBar title="Memmi" className="header" onLeftIconButtonTouchTap={() => action('GET_CARD', {test:'asdf'})}>
        <img src={logoUrl} />
      </AppBar>
    )
  }
}

export default Header
