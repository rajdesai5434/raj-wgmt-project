import React from 'react';

class SignUp extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
      first_name: "",
      last_name: "",
      email: "",
      app_use_status: "",
    }

    this.handlePwdChange = this.handlePwdChange.bind(this);
    this.handleUserChange = this.handleUserChange.bind(this);
    this.handleEmailChange = this.handleEmailChange.bind(this);
    this.handleFNameChange = this.handleFNameChange.bind(this);
    this.handleLNameChange = this.handleLNameChange.bind(this);
    this.handleAppUseChange = this.handleAppUseChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.SignUpPosting = this.SignUpPosting.bind(this);
  }

  handleUserChange(event) {
    this.setState({username: event.target.value});
  }

  handlePwdChange(event) {
    this.setState({password: event.target.value});
  }

  handleEmailChange(event) {
    this.setState({email: event.target.value});
  }

  handleFNameChange(event) {
    this.setState({first_name: event.target.value});
  }

  handleLNameChange(event) {
    this.setState({last_name: event.target.value});
  }

  handleAppUseChange(event) {
    this.setState({app_use_status: event.target.value});
  }

  SignUpPosting() {
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
      xhr.open('POST', 'http://localhost:5000/api/signup')
      xhr.send(JSON.stringify({
        username: this.state.username,
        password: this.state.password,
        fname: this.state.first_name,
        lname: this.state.last_name,
        email: this.state.email,
        appUseStatus: this.state.app_use_status,
      }))
  }

  handleSubmit(event) {
    //create a new entry in db
    //send usernameand loggedin stat back to the home page
    if (this.state.username && this.state.password && this.state.email && this.state.first_name && this.state.last_name){
      console.log("Heading to api call... weeeeee....")
      this.SignUpPosting()
    }else {
      console.log("GAWD HELP!!!")
    }
    event.preventDefault();
  }

  render() {
    return (
      <div className="container">
        <center>
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>WingMate SignUp</h1>
            <p>Alright Alright Alright, Lets Get you Signed Up Mates!!</p>
            <div>
              <form onSubmit={this.handleSubmit}>
                <label>
                  Username:
                  <input type="text" value={this.state.user} onChange={this.handleUserChange} />
                </label>
                <label>
                  Password:
                  <input type="password" value={this.state.pwd} onChange={this.handlePwdChange} />
                </label>
                <label>
                  Email:
                  <input type="text" value={this.state.user} onChange={this.handleEmailChange} />
                </label>
                <label>
                  First Name:
                  <input type="text" value={this.state.pwd} onChange={this.handleFNameChange} />
                </label>
                <label>
                  Last Name:
                  <input type="text" value={this.state.user} onChange={this.handleLNameChange} />
                </label>
                <label>
                  App Use Status:
                  <input type="text" value={this.state.user} onChange={this.handleAppUseChange} />
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

export default SignUp;
