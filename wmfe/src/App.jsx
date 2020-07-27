import React from 'react';
import logo from './logo.svg';
import './App.css';
import LoggedIn from './LoggedIn'
import Home from './Home'

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user: "",
      loggedIn: false,
    }
    this.handleUserChange = this.handleUserChange.bind(this);
    this.handleLoggedInStatus = this.handleLoggedInStatus.bind(this);
  }

  handleUserChange(name) {
    this.setState({user: name});
  }

  handleLoggedInStatus(status) {
    this.setState({loggedIn:status});
  }

  render() {
    console.log(this.state.loggedIn)
    if (this.state.loggedIn) {
      return (
        <div style={{backgroundColor: 'red'}}>
          <LoggedIn />
        </div>);
    } else {
      return (
        <div style={{backgroundColor: 'red'}}>
          <Home
            userNameChange={this.handleUserChange}
            loggedInStat={this.handleLoggedInStatus}
            />
        </div>);
    }
  }
}

export default App;
