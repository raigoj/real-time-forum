import React from "react";
import './title.css';

export class Title extends React.Component {
  render() {
    return (
    <div className="container gradient-background">
      <br /><br />
        <h1 className="heading">Please register or log in!</h1><br />
        <p className="addition">Without logging in or out, you cannot post or submit likes/dislikes.</p>

        <br /><br />
        <a className="button secondary-button" href='/register'>Register</a>
        <br />
        <a className="button primary-button" href='/signin'>Log in</a>
        <br /><br />
    </div>

    )
  }
}

