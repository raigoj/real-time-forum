import React from 'react';
import './header.css';

export class Header extends React.Component {
 handleLogOut = () => {
    const cookies = document.cookie.split(";"); // split cookies by semicolon
    for (let i = 0; i < cookies.length; i++) {
      const cookieName = cookies[i].split("=")[0]; // get cookie name
      document.cookie = cookieName + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT"; // set expiration to past to delete cookie
    }
    window.location.replace("http://localhost:3000/");
  };

  render() {
    return (
      <div className="header-container">
        <br />
        <img src="http://localhost:8080/images/genie.gif" alt="Genie" />
        <br />
        <h1>Welcome to the Forum!</h1>
        <button className="log-out-button-2" onClick={this.handleLogOut}>Log Off</button>
        <button className="chat-button" onClick={() => { 
  window.location = `${window.location.protocol}//${window.location.host}/chat`
}}>Chat</button>

      </div>
    );
  }
}
