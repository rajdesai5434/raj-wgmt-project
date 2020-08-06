import React from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form'

class WingMateProfile extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      dater_username:"",
      relationship_to_dater:"",
      date_of_birth:"",
      intro_wing_line:"",
      current_city:"",
    }

    this.handleDaterUserChange = this.handleDaterUserChange.bind(this)
    this.handleRelationChange = this.handleRelationChange.bind(this)
    this.handleDobChange = this.handleDobChange.bind(this)
    this.handleIntrohange = this.handleIntrohange.bind(this)
    this.handleCityChange = this.handleCityChange.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)
    this.getWingMateProfile = this.getWingMateProfile.bind(this)
    this.postSaveChanges = this.postSaveChanges.bind(this)
  }

  //make get api call here
  componentDidMount(){
    this.getWingMateProfile()
  }

  handleDaterUserChange(event){
    this.setState({
      dater_username:event.target.value
    })
  }

  handleRelationChange(event){
    this.setState({
      relationship_to_dater:event.target.value
    })
  }

  handleDobChange(event){
    this.setState({
      date_of_birth:event.target.value
    })
  }

  handleIntrohange(event){
    this.setState({
      intro_wing_line:event.target.value
    })
  }

  handleCityChange(event){
    this.setState({
      current_city:event.target.value
    })
  }

  getWingMateProfile(){
    // create a new XMLHttpRequest
    var xhr = new XMLHttpRequest()
    console.log("Inside GET API Amigo")
    // get a callback when the server responds
    xhr.addEventListener('load', () => {
      // update the state of the component with the result here
      if (xhr.status===200){
        var resp = JSON.parse(xhr.responseText);
        this.setState({
          dater_username:resp["msg"]["dater_username"],
          relationship_to_dater:resp["msg"]["relationship_to_dater"],
          date_of_birth:resp["msg"]["date_of_birth"],
          intro_wing_line:resp["msg"]["intro_wing_line"],
          current_city:resp["msg"]["current_city"],
        })
      }
    })
    // open the request with the verb and the url
    const url = `http://localhost:5000/api/wingmate_profile?username=${this.props.userCreds["username"]}`
    xhr.open('GET', url)
    xhr.send(null)
  }

  postSaveChanges() {
      // create a new XMLHttpRequest
      var xhr = new XMLHttpRequest()
      // get a callback when the server responds
      xhr.addEventListener('load', () => {
        // update the state of the component with the result here
        if (xhr.status===200){
          var resp = JSON.parse(xhr.responseText);
          console.log(resp["msg"])
        }
      })
      // open the request with the verb and the url
      xhr.open('POST', 'http://localhost:5000/api/wingmate_profile')
      xhr.send(JSON.stringify({
        username: this.props.userCreds["username"],
        dater_username: this.state.dater_username,
        relationship_to_dater: this.state.relationship_to_dater,
        date_of_birth: this.state.date_of_birth,
        intro_wing_line: this.state.intro_wing_line,
        current_city: this.state.current_city,
      }))
  }

  //make post request here
  handleSubmit(event){
    this.postSaveChanges()
    event.preventDefault()
  }

  render() {
    return (
      <div className="container">
        <center>
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <div>
              <Form>
                <Form.Group controlId="formGroupUsername">
                  <Form.Label>Username</Form.Label>
                  <Form.Control size="lg" type="text" value={this.props.userCreds["username"]} disabled/>
                </Form.Group>
                <Form.Group controlId="formGroupDater">
                  <Form.Label>Dater Username</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.dater_username} onChange={this.handleDaterUserChange} placeholder="Dater Username" />
                </Form.Group>
                <Form.Group controlId="formGroupRelation">
                  <Form.Label>Relationship to Dater</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.relationship_to_dater} onChange={this.handleRelationChange} placeholder="friend, relative or other" />
                </Form.Group>
                <Form.Group controlId="formGroupDOB">
                  <Form.Label>Date of Birth</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.date_of_birth} onChange={this.handleDobChange} placeholder="YYYY-MM-DD" />
                </Form.Group>
                <Form.Group controlId="formGroupIntro">
                  <Form.Label>Intro Wing Line</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.intro_wing_line} onChange={this.handleIntrohange} placeholder="Enter Intro" />
                </Form.Group>
                <Form.Group controlId="formGroupCurCity">
                  <Form.Label>Current City</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.current_city} onChange={this.handleCityChange} placeholder="Your current city" />
                </Form.Group>
              </Form>
              <div>
                <Button variant="secondary" size="lg" onClick={this.handleSubmit}>Save changes!</Button>{' '}
              </div>
            </div>
          </div>
        </center>
      </div>
    )
  }
}

export default WingMateProfile;
