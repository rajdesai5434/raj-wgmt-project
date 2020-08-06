import React from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form'

class DaterProfile extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      wing_username:"",
      search_permission:"",
      date_of_birth:"",
      short_intro:"",
      current_city:"",
      job_role:"",
      employment_status:"",
      study_college:""
    }

    this.handleWingUserChange = this.handleWingUserChange.bind(this)
    this.handleSearchPermChange = this.handleSearchPermChange.bind(this)
    this.handleDobChange = this.handleDobChange.bind(this)
    this.handleIntroChange = this.handleIntroChange.bind(this)
    this.handleCityChange = this.handleCityChange.bind(this)
    this.handleJobRoleChange = this.handleJobRoleChange.bind(this)
    this.handleEmployStatChange = this.handleEmployStatChange.bind(this)
    this.handleCollegeChange = this.handleCollegeChange.bind(this)

    this.getDaterProfile = this.getDaterProfile.bind(this)
    this.postSaveChanges = this.postSaveChanges.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)
  }

  //make get api call here
  componentDidMount(){
    this.getDaterProfile()
  }

  handleWingUserChange(event){
    this.setState({
      wing_username:event.target.value
    })
  }

  handleSearchPermChange(event){
    this.setState({
      search_permission:event.target.value
    })
  }

  handleDobChange(event){
    this.setState({
      date_of_birth:event.target.value
    })
  }

  handleIntroChange(event){
    this.setState({
      short_intro:event.target.value
    })
  }

  handleCityChange(event){
    this.setState({
      current_city:event.target.value
    })
  }

  handleJobRoleChange(event){
    this.setState({
      job_role:event.target.value
    })
  }

  handleEmployStatChange(event){
    this.setState({
      employment_status:event.target.value
    })
  }

  handleCollegeChange(event){
    this.setState({
      study_college:event.target.value
    })
  }

  getDaterProfile(){
    // create a new XMLHttpRequest
    var xhr = new XMLHttpRequest()
    console.log("Inside GET API Amigo")
    // get a callback when the server responds
    xhr.addEventListener('load', () => {
      // update the state of the component with the result here
      if (xhr.status===200){
        var resp = JSON.parse(xhr.responseText);
        this.setState({
          wing_username:resp["msg"]["wing_username"],
          search_permission:resp["msg"]["search_permission"],
          date_of_birth:resp["msg"]["date_of_birth"],
          short_intro:resp["msg"]["short_intro"],
          current_city:resp["msg"]["current_city"],
          job_role:resp["msg"]["job_role"],
          employment_status:resp["msg"]["employment_status"],
          study_college:resp["msg"]["study_college"],
        })
      }
    })
    // open the request with the verb and the url
    const url = `http://localhost:5000/api/dater_profile?username=${this.props.userCreds["username"]}`
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
      xhr.open('POST', 'http://localhost:5000/api/dater_profile')
      xhr.send(JSON.stringify({
        username: this.props.userCreds["username"],
        wing_username: this.state.wing_username,
        search_permission: this.state.search_permission,
        date_of_birth: this.state.date_of_birth,
        short_intro: this.state.short_intro,
        current_city: this.state.current_city,
        job_role: this.state.job_role,
        employment_status: this.state.employment_status,
        study_college: this.state.study_college,
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
                <Form.Group controlId="formGroupWing">
                  <Form.Label>Wingmate Username</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.wing_username} onChange={this.handleWingUserChange} placeholder="Wingmate Username" />
                </Form.Group>
                <Form.Group controlId="formGroupSearchPerm">
                  <Form.Label>Search Permission</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.search_permission} onChange={this.handleSearchPermChange} placeholder="true or false" />
                </Form.Group>
                <Form.Group controlId="formGroupDOB">
                  <Form.Label>Date of Birth</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.date_of_birth} onChange={this.handleDobChange} placeholder="YYYY-MM-DD" />
                </Form.Group>
                <Form.Group controlId="formGroupIntro">
                  <Form.Label>Short Intro</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.short_intro} onChange={this.handleIntroChange} placeholder="Enter Intro" />
                </Form.Group>
                <Form.Group controlId="formGroupCurCity">
                  <Form.Label>Current City</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.current_city} onChange={this.handleCityChange} placeholder="Your current city" />
                </Form.Group>
                <Form.Group controlId="formGroupJobRole">
                  <Form.Label>Job Role</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.job_role} onChange={this.handleJobRoleChange} placeholder="Current Job Role" />
                </Form.Group>
                <Form.Group controlId="formGroupEmpStat">
                  <Form.Label>Employment Status</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.employment_status} onChange={this.handleEmployStatChange} placeholder="Employment Status" />
                </Form.Group>
                <Form.Group controlId="formGroupStudyCollege">
                  <Form.Label>College</Form.Label>
                  <Form.Control size="lg" type="text" value={this.state.study_college} onChange={this.handleCollegeChange} placeholder="College" />
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

export default DaterProfile;
