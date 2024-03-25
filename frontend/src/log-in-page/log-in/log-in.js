import React from "react";
import './log-in.css';
import axios from "axios";
//import Cookie from '../fetch-cookie/cookie.js';
//import Post from '../log-in-post/post.js';


export class LogIn extends React.Component {
    handleSubmit = (event) => {
    event.preventDefault()
    let username = event.target.elements.username.value
    let password = event.target.elements.password.value
    let data = {
      username,
      password,
    };
    axios({
      method: "POST",
      url: "http://localhost:8080/signin",
      data: data,
      headers: { "Content-Type": "text/plain" },
      withCredentials: true,
      credentials: 'include',
    }).then(function (response) {
      let cookieData = response.data.split("\"")
      let cookieName = cookieData[3]
      let cookieId = cookieData[5]
      var now = new Date();
      var time = now.getTime();
      var expireTime = time + 17*36000;
      now.setTime(expireTime);
      document.cookie = cookieName + "=" + cookieId + "; expires=" + now.toUTCString() + ";path=/ ; samesite=strict"
      console.log(document.cookie);
      window.location.replace("/threads")
    })
      .catch(function (error) {
        console.log(error);
      });
  }
  render() {
    return (
    <div className="container">

      <br />
      <h1 className="heading">Please enter your details.</h1>
      <br />
      <div className="form-container">
      <form onSubmit={this.handleSubmit}  >
          <div className="form-group">

          <label htmlFor="username" className="form-label"><b>Username: </b></label>
          <input type="username" className="form-control form-input" id="username" placeholder="Enter Username" name="username" required />        
          </div>

          <div className="form-group">
          <label className="form-label" htmlFor="password"><b>Password: </b></label>
          <input type="password" id="password" className="form-control form-input" placeholder="Enter Password" name="password" required/>
          </div>
          <br />

          <button type="b" name="actionbutton" value="signin" className="button primary-button">Log in</button>
      </form>
      </div>

    </div>
    
    );
  }
}
