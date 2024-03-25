import React, {useState, setState} from "react";
import './registration.css';
import axios from "axios";
import { Link } from "react-router-dom";

function Registration() {
  const [respData, setRespData] = useState(null);
  const handleSubmit = (event) => {
    event.preventDefault()
    let username = event.target.elements.username.value;
    let password = event.target.elements.password.value;
    let email = event.target.elements.email.value;
    let age = event.target.elements.age.value;
    let gender = event.target.elements.gender.value;
    let firstname = event.target.elements.firstname.value;
    let lastname = event.target.elements.lastname.value;
    let data = {
      username,
      password,
      email,
      age,
      gender,
      firstname,
      lastname,
    };
    console.log(data)
    axios({
      method: "POST",
      url: "http://localhost:8080/register",
      data: data,
      headers: { "Content-Type": "text/plain" },
      withCredentials: true,
      credentials: 'include',
    }).then(function (response) {
        console.log("this is the response", response);
        if (response.data === "") {
          window.location.href = '/signin';
        } else {
          setRespData(response.data)
        }
      })
      .catch(function (error) {
        console.error("Error fetching data: ", error);
      });
  }
    return (
      
      <div className="container">
<br /><br /><br /><br /><br /><br /><br /><br />
        
        <h1 className="heading">Please enter your details.</h1>
        <br /><br />
        <div className="form-container">
          <form onSubmit={handleSubmit} >
            
          <div className="form-group">

            <label className="form-label" htmlFor="username"><b>Username</b></label><br />
            <div className="comment">Pick a username that is at least five characters long.</div>
            <input type="username" className="form-control form-input" placeholder="Enter Username" name="username" required />
            
            <br />
            </div>

            <div className="form-group">

              <label className="form-label" htmlFor="password"><b>Password</b></label><br />
              <div className="comment">More than six characters, containing numbers and both upper and lower case letters.</div>
              <input type="password" className="form-control form-input" id="InputPassword1" placeholder="Enter Password" title="The password needs to be at least six characters long, and contain numbers and both upper and lower case letters" name="password" required />
              <br />
              </div>

              <div className="form-group">
              <label className="form-label" htmlFor="email"><b>Email</b></label>
              <div className="comment">Please enter a valid email address so you can receive our weekly newsletter.</div>
              <input type="email" className="form-control form-input" placeholder="Enter Email" name="email" required />
              </div>

              <div className="form-group">
              <label className="form-label" htmlFor="age"><b>Age</b></label>
              <div className="comment">We do not accept people younger than 10 and older than 100!</div>
              <input type="number" className="form-control form-input" placeholder="Enter Age" name="age" min="10" max="100" required />
              </div>

              <div className="form-group">
              <label className="form-label" htmlFor="gender"><b>Gender</b></label>
              <div className="comment">We need to know your gender for science: </div>
              <select name="gender" className = "form-select" id="gender">
                <option value="Male">Male</option>
                <option value="Female">Female</option>
                <option value="?">Opel</option>
              </select>
              </div>

              <div className="form-group">
              <label className="form-label" htmlFor="firstname"><b>Firstname</b></label>
              <div className="comment">My name is:</div>
              <input type="text" className="form-control form-input" placeholder="Enter Firstname" name="firstname" required />
              </div>

              <div className="form-group">
              <label className="form-label" htmlFor="lastname"><b>Lastname</b></label>
              <div className="comment">.... and last name is:</div>
              <input type="text" className="form-control form-input" placeholder="Enter Lastname" name="lastname" required />
              </div>
<br /><br />
            <button className="button primary-button" type="submit" name="actionbutton" value="register" >Register</button>
            <h2>{respData}</h2>
          </form>
        </div>
      </div>
    )
}

export default Registration;
