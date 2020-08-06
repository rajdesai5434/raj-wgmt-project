import React from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form'

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
      this.SignUpPosting()
      this.preventDefault()
    }else {
      console.log("GAWD HELP!!!")
    }
  }

  render() {
    return (
      <div className="container">
        <center>
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>WingMate SignUp</h1>
            <p>Alright Alright Alright, Lets Get you Signed Up Mates!!</p>
            <div>
              <Form>
                <Form.Group controlId="formGroupUsername">
                  <Form.Label>Username</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.username} onChange={this.handleUserChange} placeholder="Enter Username" />
                </Form.Group>
                <Form.Group controlId="formGroupPassword">
                  <Form.Label>Password</Form.Label>
                  <Form.Control size="lg" type="password" value={this.state.password} onChange={this.handlePwdChange} placeholder="Password" />
                </Form.Group>
                <Form.Group controlId="formGroupEmail">
                  <Form.Label>Email address</Form.Label>
                  <Form.Control size="lg" type="email" value={this.state.email} onChange={this.handleEmailChange} placeholder="Enter email" />
                </Form.Group>
                <Form.Group controlId="formGroupFname">
                  <Form.Label>First Name</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.first_name} onChange={this.handleFNameChange} placeholder="First Name" />
                </Form.Group>
                <Form.Group controlId="formGroupLname">
                  <Form.Label>Last Name</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.last_name} onChange={this.handleLNameChange} placeholder="Last Name" />
                </Form.Group>
                <Form.Group controlId="formGroupPassword">
                  <Form.Label>App Use Case?</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.app_use_status} onChange={this.handleAppUseChange} placeholder="Enter either dater or wing_mate" />
                </Form.Group>
              </Form>
              <div>
                <Button variant="secondary" size="lg" onClick={this.handleSubmit}>Create Profile!!</Button>{' '}
              </div>
            </div>
          </div>
        </center>
      </div>
    )
  }

}

export default SignUp;
