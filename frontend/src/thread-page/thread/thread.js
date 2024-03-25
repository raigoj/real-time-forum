import {useEffect, useState} from 'react';
import axios from "axios";
import './thread.css';
import {NewComment} from '../new-comment/new-comment.js';
import {categoryIdToName} from '../../threads-page/threads/threads.js';
function Thread({currUserId}) {
  const [post, setPost] = useState([]);
  const [comments, setComments] = useState([]);
  const [usernames, setUsernames] = useState({});
  let pathArray = window.location.pathname.split('/');
  let url = "//localhost:8080/thread?id=" + pathArray[2];
  useEffect(() => {
    axios.get(url)
      .then(function (response) { 
        setPost(response.data.Post)
        setComments(response.data.Comments) 
      })
      .catch(function (error) { 
        console.log(error)
      });
    }, [url]);
    
  useEffect(() => {
    fetch('http://localhost:8080/users')
      .then(res => res.json())
      .then(data => {
        console.log(data)
        let userDict = {};
        data.forEach(user => {
          userDict[user.id] = user.username;
        });
        setUsernames(userDict);
      });
  }, []);
  function getUsername(userId) {
    return usernames[userId] || 'Unknown User';
  } 
  let postDisplay = null;
  if (!post) {
    postDisplay = <p>Loading post...</p>;
  } else {
    const username = getUsername(post.user_id);
    postDisplay = (
      <div className="thread_row" >
      <div className="post">
        <h4 className="title">
          {post.post_name}
        </h4>
        <div className="bottom">
          <p className="post-content">
            {post.post_content}
          </p>
          <p className="timestamp">
            Created: {post.post_date}
          </p>
          <p className="post-author">
            Post author: {username} 
          </p>
          <p className="tags">
            Tagged as {categoryIdToName(post.category_id)}
          </p>
        </div></div>  
      </div>
    )
  }

  function getCookieThread(cName) {
    const name = cName + "=";
    const cDecoded = decodeURIComponent(document.cookie); //to be careful
    const cArr = cDecoded.split('; ');
    let res;
    cArr.forEach(val => {
      if (val.indexOf(name) === 0) res = val.substring(name.length);
    })
    return res
}

  useEffect(() => {
    const fetchData = async () => {
      const sessionResponse = await axios.get("http://localhost:8080/session");
      const session = sessionResponse.data;
      const resD = session.find(token => token.token === getCookieThread("sessionID"));
      if (resD === undefined) {
        window.location.href = '/';
      }
    }
    fetchData();
  });

  

  const handleNewComment = (value) => {
    if (comments && comments.length) {
      setComments([...comments, value]);
    } else {
      setComments([value]);
    }
  };
  let commentsLis = comments?.map((comment) => {
    console.log(typeof comment.user_id)
    const username = getUsername(comment.user_id);
    return (
    <li className="row" key={'comment_' + comment.comment_id}>
        <p className="comment-content">
          {comment.comment_content}
        </p>
        <p className="comment-timestamp">
          Created: {comment.comment_date} <br />
        </p>
        <p className="comment-author">
          &nbsp; Post author: {username} <br />
        </p>
    </li>
    );
  });

  
  const handleLogOut = () => {
    const cookies = document.cookie.split(";"); // split cookies by semicolon
    for (let i = 0; i < cookies.length; i++) {
      const cookieName = cookies[i].split("=")[0]; // get cookie name
      document.cookie = cookieName + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT"; // set expiration to past to delete cookie
    }
    window.location.replace("http://localhost:3000/");
  };

  
  const handleBackButton = () => {
    window.history.back();
  }
  
  return (
    <div>
        <button className="back-button" onClick={() => { 
  window.location = `${window.location.protocol}//${window.location.host}/threads`
}}>Back</button>

<button className="log-off-button-3" onClick={handleLogOut}>Log off</button>



      <button className="chat-button" onClick={() => { 
  window.location = `${window.location.protocol}//${window.location.host}/chat`
}}>Chat</button>


      <ul>{postDisplay}</ul>
      <ul className="comments-container">{commentsLis}</ul>
      <NewComment onNewComment={handleNewComment} currUserId={currUserId} />
    </div> 
  );
};
export default Thread;


