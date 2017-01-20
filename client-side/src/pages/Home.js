import React from 'react'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import getMuiTheme from 'material-ui/styles/getMuiTheme'
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme'
import colors from '../components/Colors'
import Header from '../components/Header/Header'


const theme = getMuiTheme(darkBaseTheme)
theme.appBar.color = colors.darkPurple
theme.appBar.textColor = colors.midGreen

class Home extends React.Component {
  render() {
    return (
      <MuiThemeProvider muiTheme={theme} >
        <Header/>
      </MuiThemeProvider>
    )
  }
}

export default Home
