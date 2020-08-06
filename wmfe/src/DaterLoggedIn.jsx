import React from 'react';
import Button from 'react-bootstrap/Button';
import DaterProfile from './DaterProfile'

class DaterLoggedIn extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      creds: {},
      dater_profile:false
    }

    this.logout = this.logout.bind(this)
    this.onDaterProfileClick = this.onDaterProfileClick.bind(this)
  }

  componentDidMount(){
    this.setState({
      creds:this.props.userCreds
    })
  }

  onDaterProfileClick(){
    this.setState({
      dater_profile:true
    })
  }

  logout(){
      this.props.loggedInStat(false)
  }

  render() {
    if (this.state.dater_profile && this.state.creds){
      const displayText = "Welcome "+this.state.creds["username"]+ ", get ready to rumble!!"
      return (
        <div className="container">
          <div className="col-lg-12">
            <center>
              <h2>WingMate</h2>
              <Button variant="secondary" size="lg" onClick={this.logout}>Log out</Button>{' '}
              <Button variant="secondary" size="lg" onClick={this.onDaterProfileClick} disabled>Dater Profile</Button>{' '}
              <br/>
              <div>{displayText}</div>
              <p>Lets setup the profile</p>
              <DaterProfile
                userCreds={this.props.userCreds}
                />
              </center>
          </div>
        </div>
      )
    } else if(this.state.creds){ //can potentially remove this once have loading thingie
      const displayText = "Welcome "+this.state.creds["username"]
      return (
        <div className="container">
          <div className="col-lg-12">
            <center>
              <h2>WingMate</h2>
              <Button variant="secondary" size="lg" onClick={this.logout}>Log out</Button>{' '}
              <Button variant="secondary" size="lg" onClick={this.onDaterProfileClick}>Dater Profile</Button>{' '}
              <br/>
              <div>{displayText}</div>
              </center>
          </div>
        </div>
      )
    } else{
      return (<div>Add a loading thingie here lolol</div>)
    }
  }
}
export default DaterLoggedIn;
