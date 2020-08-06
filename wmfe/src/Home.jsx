import React from 'react';
import SignUp from './SignUp'
import Button from 'react-bootstrap/Button';

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user: "",
      pwd: "",
      sign_up:false
    }
    this.handlePwdChange = this.handlePwdChange.bind(this);
    this.handleUserChange = this.handleUserChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.onSignUpClick = this.onSignUpClick.bind(this);
  }

  handleUserChange(event) {
    this.setState({user: event.target.value});
  }

  handlePwdChange(event) {
    this.setState({pwd: event.target.value});
  }
  
  onSignUpClick(){
    this.setState({
      sign_up:true
    })
      
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
          var resp = JSON.parse(xhr.responseText);
          this.props.userCreds(resp["msg"])
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
    event.preventDefault();
  }

  render() {
    if (!this.state.sign_up){
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
              <p>Not a user, join us!!</p>
              <Button variant="secondary" size="lg" onClick={this.onSignUpClick}>Sign Up</Button>{' '}
            </div>
          </center>
        </div>
      )
    }else{
      return(
      <SignUp
        userCreds={this.props.userCreds}
        loggedInStat={this.props.loggedInStat}
        />
      )
    }
  }
}

export default Home;
