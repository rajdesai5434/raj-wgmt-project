import React from 'react';

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user: "",
      pwd: "",
      authenticated_users:{"r":"r"},
    }
    this.handlePwdChange = this.handlePwdChange.bind(this);
    this.handleUserChange = this.handleUserChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleUserChange(event) {
    this.setState({user: event.target.value});
  }

  handlePwdChange(event) {
    this.setState({pwd: event.target.value});
  }

  handleSubmit(event) {
    if (this.state.authenticated_users[this.state.user]===this.state.pwd){
        //send the info back to app to render LoggedIn
        this.props.userNameChange(this.state.user)
        this.props.loggedInStat(true)
    } else {
      console.log("GAWD HELP!!!")
    }
    event.preventDefault();
  }

  render() {
    return (
      <div className="container">
        <center>
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>WingMate</h1>
            <p>Help your friend get a Date!</p>
            <p>Sign in to get access </p>
            <div>
              <form onSubmit={this.handleSubmit}>
                <label>
                  Username:
                  <input type="text" value={this.state.user} onChange={this.handleUserChange} />
                </label>
                <label>
                  Password:
                  <input type="text" value={this.state.pwd} onChange={this.handlePwdChange} />
                </label>
                <input type="submit" value="Submit" />
              </form>
            </div>
          </div>
        </center>
      </div>
    )
  }
}

export default Home;