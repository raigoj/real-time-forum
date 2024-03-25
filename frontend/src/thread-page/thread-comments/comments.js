import { useEffect, useState } from 'react';
import axios from 'axios';
import './comments.css';
function Comments() {
  const [comments, setComments] = useState([]);
  useEffect(() => {
    let pathArray = window.location.pathname.split('/');
    let url = `//localhost:8080/thread?id=${pathArray[2]}`;
    axios.get(url)
      .then((response) => setComments(response.data));
  }, []);
  const commentsList = comments.length > 0 ? comments.map((comment) => (
    <li className="row" key={`comment_${comment.comment_id}`}>
      <p className="comment-content">
        {comment.comment_content}
      </p>
      <p className="timestamp">
        Created: {comment.comment_date}
      </p>
      <p className="author">
        Author: {comment.user_id}
      </p>
    </li>
  )) : <li className="row">No Comments</li>;
  return (
    <ul className="comments-container">
      {commentsList}
    </ul>
  );
}
export default Comments;
