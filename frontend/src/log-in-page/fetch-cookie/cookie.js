import {useEffect, useState} from 'react';
import axios from "axios";
  
async function Cookie() {
  const [cookie, setCookie] = useState([]);
  useEffect(() => {
    axios.get("//localhost:8080/signin")
      .then((response) => setCookie(response.data));
    }, []);
  let cookieName = cookie?.Name;
  let cookieVal  = cookie?.Value;
  console.log(cookieName);
  console.log(cookieVal);
}

export default Cookie;


