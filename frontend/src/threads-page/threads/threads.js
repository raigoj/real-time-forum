import {useEffect, useState, useCallback} from 'react';
import axios from "axios";
import './threads.css';
import NewThreads from '../new-thread/new-thread.js';
import {Header} from '../header/header.js';
import { Link } from "react-router-dom";
export default function Threads() { 
 
  const [posts, setPosts] = useState([]);
  const handleLogOut = () => {
    const cookies = document.cookie.split(";"); // split cookies by semicolon
    for (let i = 0; i < cookies.length; i++) {
      const cookieName = cookies[i].split("=")[0]; // get cookie name
      document.cookie = cookieName + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT"; // set expiration to past to delete cookie
    }
    window.location.replace("http://localhost:3000/");
  };
  const [usernameMap, setUsernameMap] = useState({});

  function getCookieThreads(cName) {
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
      const resD = session.find(token => token.token === getCookieThreads("sessionID"));
      if (resD === undefined) {
        window.location.href = '/';
      }
      const usersResponse = await axios.get("http://localhost:8080/users");
      const users = await usersResponse.data;
      const usernamePromises = posts.map(post => userIdToName(post.user_id));
      const usernames = await Promise.all(usernamePromises);
      const usernameMap = {};
      usernames.forEach((username, i) => {
        usernameMap[posts[i].user_id] = username;
      });
      setUsernameMap(usernameMap);
    }
    fetchData();
  }, [posts]);

   const fetchUsernames = useCallback(async () => {
    const postsResponse = await axios.get("http://localhost:8080/home");
    const resPosts = await postsResponse.data;
    setPosts(resPosts)
    const usernamePromises = posts.map(post => userIdToName(post.user_id));
    const usernames = await Promise.all(usernamePromises);
    const usernameMap = {};
    usernames.forEach((username, i) => {
      usernameMap[posts[i].user_id] = username;
    });
    setUsernameMap(usernameMap);
  }, [posts]); 
  useEffect(() => {
    fetchUsernames();
  }, [fetchUsernames]);
  //console.log(posts)
   const addPost = (postData) => {
    //let postDataArr = Object.keys(postData);
    setPosts([...posts, postData]);
  };
 const postsLis = posts.map(post => (
    <li className="row" key={'post_' + post.post_id}>
      <Link to={"/thread/" + post.post_id} >
        <h4 className="title">
          {post.post_name} <br /> 
        </h4>
        <div className="bottom">
          <p className="th-timestamp">
            Created: {post.post_date} <br />
          </p>
          <p className="th-comment-count">
           Post author: {usernameMap[post.user_id]} <br />
          </p>
          <p className="th-tags">
            Tagged as {categoryIdToName(post.category_id)}
          </p>
        </div>
      </Link>
    </li>
  )); 
  const lastPostId = posts.reduce((max, post) => Math.max(max, post.post_id), 0) + 1;
  return (
    
    <div className="threads">
      <Header />
      <ul className="posts">{postsLis}</ul>
      <NewThreads addPost={addPost} lastPostId={lastPostId} />
      {/*<Users /> 
      <Chat />*/}
      
    </div>
    
  );
}
let usersCache;
export async function userIdToName(userId) {
  if (!usersCache) {
    const usersResponse = await axios.get("http://localhost:8080/users");
    usersCache = await usersResponse.data;
  }
  const resData = usersCache.find(token => token.id === userId);
  if (!resData) {
    return userId;
  }
  return resData.username;
}
export function categoryIdToName(categoryId) {
  if (!categoryId) {
    return "";
  }
  const categories = [
    { id: 1, name: "Food & Travel" },
    { id: 2, name: "Cat Pictures" },
    { id: 3, name: "Problematic" },
    { id: 4, name: "Funny" }
  ];
  let categoryName = "";
  let first = true;
    categoryId.toString().split("").forEach((id) => {
    const foundCategory = categories.find((category) => category.id === parseInt(id));
    if (foundCategory) {
      if (first) {
        first = false;
      } else {
        categoryName += ", ";
      }
      categoryName += foundCategory.name;
    }
  });
  return categoryName; 
}
