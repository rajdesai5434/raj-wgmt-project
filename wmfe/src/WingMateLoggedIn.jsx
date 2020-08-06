import React from 'react';
import Button from 'react-bootstrap/Button';
import WingMateProfile from './WingMateProfile'

class WingMateLoggedIn extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      creds: {},
      wing_profile:false,
      go_winging:false
    }

    this.logout = this.logout.bind(this)
    this.onWingProfileClick = this.onWingProfileClick.bind(this)
    this.onGoWingingClick = this.onGoWingingClick.bind(this)
  }

  componentDidMount(){
    this.setState({
      creds:this.props.userCreds
    })
  }

  onWingProfileClick(){
    this.setState({
      wing_profile:true,
      go_winging:false,
    })
  }

  onGoWingingClick(){
    this.setState({
      go_winging:true,
      wing_profile:false
    })
  }

  logout(){
      this.props.loggedInStat(false)
  }

  render() {
    if (this.state.wing_profile && this.state.creds){
      const displayText = "Welcome "+this.state.creds["username"]+ ", you are one true mate!!"
      return (
        <div className="container">
          <div className="col-lg-12">
            <center>
              <h2>WingMate</h2>
              <Button variant="secondary" size="lg" onClick={this.logout}>Log out</Button>{' '}
              <Button variant="secondary" size="lg" onClick={this.onWingProfileClick} disabled>Wing Profile</Button>{' '}
              <Button variant="secondary" size="lg" onClick={this.onGoWingingClick}>Go Winging</Button>{' '}
              <br/>
              <div>{displayText}</div>
              <p>Lets setup the profile</p>
                <WingMateProfile
                  userCreds={this.props.userCreds}
                  />
              </center>
          </div>
        </div>
      )
    } else if (this.state.go_winging && this.state.creds){
      const displayText = "Welcome "+this.state.creds["username"]+ ", you are one true mate!!"
      return (
        <div className="container">
          <div className="col-lg-12">
            <center>
              <h2>WingMate</h2>
              <Button variant="secondary" size="lg" onClick={this.logout}>Log out</Button>{' '}
              <Button variant="secondary" size="lg" onClick={this.onWingProfileClick}>Wing Profile</Button>{' '}
              <Button variant="secondary" size="lg" onClick={this.onGoWingingClick} disabled>Go Winging</Button>{' '}
              <br/>
              <div>{displayText}</div>
              <p>You can start winging soon</p>
              </center>
          </div>
        </div>
      )
    } else if(this.state.creds){ //can potentially remove this once have loading thingie
      const displayText = "Welcome "+this.state.creds["username"]+ ", you are one true mate!!"
      return (
        <div className="container">
          <div className="col-lg-12">
            <center>
              <h2>WingMate</h2>
              <Button variant="secondary" size="lg" onClick={this.logout}>Log out</Button>{' '}
              <Button variant="secondary" size="lg" onClick={this.onWingProfileClick}>Wing Profile</Button>{' '}
              <Button variant="secondary" size="lg" onClick={this.onGoWingingClick}>Go Winging</Button>{' '}
              <br/>
              <div>{displayText}</div>
              <p>Alright Alright Alright, lets get you started mah mate</p>
              </center>
          </div>
        </div>
      )
    } else{
      return (<div>Add a loading thingie here lolol</div>)
    }
  }
}
export default WingMateLoggedIn;
