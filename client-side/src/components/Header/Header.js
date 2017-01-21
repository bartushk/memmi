import React from 'react'
import AppBar from 'material-ui/AppBar'
import logoUrl from './logo_67x50.svg'
import './Header.scss'


class Header extends React.Component {
  render() {
    return (
      <AppBar title="Memmi" className="header">
        <img src={logoUrl} />
      </AppBar>
    )
  }
}

export default Header
