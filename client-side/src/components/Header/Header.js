import React from 'react'
import logoUrl from './logo.png'
import './Header.scss'


class Header extends React.Component {
  render() {
    return (
      <div className="header">
        <img src={logoUrl} className="logo" />
      </div>
    )
  }
}

export default Header
