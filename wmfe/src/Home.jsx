import React from 'react';
import SignUp from './SignUp'

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user: "",
      pwd: "",
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

  SignInPosting() {
      // create a new XMLHttpRequest
      var xhr = new XMLHttpRequest()
      console.log("Inside API Amigo")
      // get a callback when the server responds
      xhr.addEventListener('load', () => {
        // update the state of the component with the result here
        //xhr.responseText
        if (xhr.status===200){
          this.props.userNameChange(xhr.responseText)
          this.props.loggedInStat(true)
        }
      })
      // open the request with the verb and the url
      xhr.open('POST', 'http://localhost:5000/api/signin')
      xhr.send(JSON.stringify({
        username: this.state.user,
        password: this.state.pwd,
      }))
  }

  handleSubmit(event) {
    //check with backend if authentication is correct
    this.SignInPosting()
    console.log("Here i gooo, weeeee....")
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
            <p>Not a user, sign up here mate!</p>
              <div>
                <SignUp
                  userNameChange={this.props.userNameChange}
                  loggedInStat={this.props.loggedInStat}
                  />
              </div>
          </div>
        </center>
      </div>
    )
  }
}

export default Home;
