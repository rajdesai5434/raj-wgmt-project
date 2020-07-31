import React from 'react';
import './App.css';
import LoggedIn from './LoggedIn'
import Home from './Home'

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      creds: "",
      loggedIn: false,
    }
    this.handleCredsChange = this.handleCredsChange.bind(this);
    this.handleLoggedInStatus = this.handleLoggedInStatus.bind(this);
    this.logout = this.logout.bind(this)
  }

  handleCredsChange(cred) {
    this.setState({creds: cred});
  }

  handleLoggedInStatus(status) {
    this.setState({loggedIn:status});
  }

  logout(){
    this.setState({
      loggedIn: false,
      creds: ""});
  }

  render() {
    if (this.state.loggedIn) {
      if (this.state.creds["appUseStatus"]==="dater"){
        return (
          <div className="container">
            <div className="col-lg-12">
              <br />
              <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                <center>
                  <h2>WingMate</h2>
                  <p>Alright Alright Alright, lets get you started Dater!!</p>
                </center>
            </div>
          </div>
        )
      } else if (this.state.creds["appUseStatus"]==="wing_mate"){
        return (
          <div className="container">
            <div className="col-lg-12">
              <br />
              <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                <center>
                  <h2>WingMate</h2>
                  <p>Alright Alright Alright, lets get you started mah Wing Mate!!</p>
                </center>
            </div>
          </div>
        )
      } else {
        return (
          <div className="container">
            <div className="col-lg-12">
              <br />
              <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                <center>
                  <h2>WingMate</h2>
                  <p>Alright Alright Alright, lets get you started, relative!!</p>
                </center>
            </div>
          </div>
        )
      }
    } else{
      return (
        <div style={{backgroundColor: 'red'}}>
          <Home
            userCreds={this.handleCredsChange}
            loggedInStat={this.handleLoggedInStatus}
            />
        </div>);
    }
  }
}

export default App;
