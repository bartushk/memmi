import React from 'react'
import AppBar from 'material-ui/AppBar'
import logoUrl from './logo.png'
import './Header.scss'


class Header extends React.Component {
  render() {
    return (
      <AppBar title="Memmi" className="header">
        <img src={logoUrl} className="logo" />
      </AppBar>
    )
  }
}

export default Header
