import React, {useState} from 'react';
import './new-thread.css';
import axios from "axios";
import {getCookie} from "../../chat/chat/chat.js";

export default function NewThreads({addPost, currUserId, lastPostId}) { 
  const categories = [
      { id: 1000, name: "Food & Travel" },
      { id: 200, name: "Cat Pictures" },
      { id: 30, name: "Problematic" },
      { id: 4, name: "Funny" }
  ];

  const handleCheckboxChange = event => {
    const { name, checked } = event.target;
    const categoryId = parseInt(name);
    setState(prevState => {
      if (checked) {
        console.log(categoryId)
        return {
          ...prevState,
          categoryIds: [...prevState.categoryIds, categoryId],
        };
      } else {
        return {
          ...prevState,
          categoryIds: prevState.categoryIds.filter(id => id !== categoryId),
        };
      }
    });
  };

  const [state, setState] = useState({
    post_name: "",
    post_content: "",
    categoryIds: [],
  });

  const handleInputChange = event => {
    const { name, value } = event.target;
    setState({ ...state, [name]: value });
  };

  const handleSubmit = event => {
    event.preventDefault();
    let date = new Date();
      let formattedDate = date.toLocaleString('en-GB', { 
        day: 'numeric',
        month: 'short',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      });
    const cookie = getCookie("sessionID");
    const { post_name, post_content, categoryIds } = state;

    const categoryIdsSum = categoryIds.reduce((sum, id) => sum + id, 0).toString();
    const post = {
      post_id: lastPostId,
      post_name: post_name,
      post_content: post_content,
      category_id: categoryIdsSum, 
      post_date: formattedDate,
      user_id: currUserId
    };
    const data = {
      post_name,
      post_content,
      categoryIdsSum,
      cookie
    };

    console.log(post)
    addPost(post);
    setState({
      post_name: "",
      post_content: "",
      categoryIds: [],
    });

    axios({
      method: "POST",
      url: "http://localhost:8080/create",
      data: data,
      headers: { "Content-Type": "text/plain" },
      withCredentials: true,
      credentials: 'include',
    }).then(function (response) {
      console.log(response.config.data); 
    })
      .catch(function (error) {
        console.log(error);
      });
    
  };
  return (
    <div className="new-thread-container">
          <h2>If necessary, create a new thread below!</h2>                <br />
          <form onSubmit={handleSubmit} method="post">
            <div className="thread-form">
              <input
                type="text"
                id="subjectArea"
                name="post_name"
                value={state.post_name}
                onChange={handleInputChange}
                maxLength="50"
                placeholder="Thread title"
                className="subject-area"
                required
              />
              <textarea
                id="contentArea"
                name="post_content"
                value={state.post_content}
                onChange={handleInputChange} 
                maxLength="250"
                placeholder="Enter your message"
                className="content-area"
                required
              />
                              <div style={{ marginBottom: '20px' }} className="checkbox-container">
                  {categories.map((category, index) => (
                    <div key={category.id} className={index < 2 ? "left-column" : "right-column"}>
                      <label>
                        <input
                          type="checkbox"
                          name={category.id}
                          checked={state.categoryIds.includes(category.id)}
                          onChange={handleCheckboxChange}
                        />
                        {category.name}
                      </label>
                    </div>
                  ))}
                </div>
              <button className="new-thread-button" type="submit">
                New thread
              </button>
            </div>
          </form>
        </div>
   
  );
}
