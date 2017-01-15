import React from 'react'

class App extends React.Component {
  static propTypes = {
    children: React.PropTypes.object
  }

  render() {
    return (
      <div className="App">
        {this.props.children}
      </div>
    )
  }
}

export default App
