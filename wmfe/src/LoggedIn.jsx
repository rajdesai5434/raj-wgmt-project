import React from 'react';

class LoggedIn extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      creds: {}
    }
    
    this.logout = this.logout.bind(this)
  }
  
  logout(){
      this.props.loggedInStat(false)
  }

  render() {
    console.log(this.props.userCreds)
    if (this.props.userCreds["appUseStatus"]==="dater"){
      return (
        <div className="container">
          <div className="col-lg-12">
            <br />
            <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
              <center>
                <h2>WingMate</h2>
                <p>Alright Alright Alright, lets get you started mah Dater!!</p>
              </center>
          </div>
        </div>
      )
    } else if (this.props.creds["appUseStatus"]==="wing_mate"){
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
  }
}

export default LoggedIn;