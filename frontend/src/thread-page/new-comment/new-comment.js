import React from 'react';
import './new-comment.css' 
import axios from "axios";
import {getCookie} from "../../chat/chat/chat.js";

export function NewComment({onNewComment, currUserId}) {
  const handleSubmit = (event) => {
    event.preventDefault()
    let pathArray = window.location.pathname.split('/')
    let post = pathArray[2];
    let commentContent = event.target.elements.CommentContent.value;

    let cookie = getCookie("sessionID");
    let data = {
      commentContent,
      cookie,
      post,
    };
    axios({
      method: "POST",
      url: "http://localhost:8080/createComments",
      data: data,
      headers: { "Content-Type": "text/plain" },
      withCredentials: true,
      credentials: 'include',
    }).then(function (response) {
      console.log(response);
      let date = new Date();
      let formattedDate = date.toLocaleString('en-GB', { 
        day: 'numeric',
        month: 'short',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      });


      onNewComment({comment_id: 65656, comment_content: commentContent, comment_date: formattedDate, post_id: post, user_id: currUserId})
      event.target.elements.CommentContent.value = '';
    })
      .catch(function (error) {
        console.log(error);
      });
  }
  return (
    <div style={{ bottom: 0, display: "flex", alignItems: "center" }}>
      <form onSubmit={handleSubmit} method="post" className="narrow">
        <br /><br />
        <div className="container_newcom">
          <h2>You can create a new comment below!</h2>
        </div>
        <div className="container_newcom">
          <textarea id="contentArea" name="CommentContent" maxLength="250" rows="5" cols="50" placeholder="Enter your message"></textarea>
          <br /><br />
        </div>
        <div className="container_newcom">
          <button className="small" type="submit">Comment</button>
        </div>
      </form>
    </div>
    ); 
} 
