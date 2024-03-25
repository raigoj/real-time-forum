import React, {useState} from "react";
//import './App.css';
import Threads from './threads-page/threads/threads.js';
import Thread from './thread-page/thread/thread.js';
import { Routes, Route, BrowserRouter } from "react-router-dom";
import {Title} from './title-page/title.js';
import {LogIn} from './log-in-page/log-in/log-in.js'
import Registration from './registration-page/registration/registration.js';
import {Errors} from './error-page/errors/errors.js';
import ChatPage from './chat/chatbox/chatbox.js';

function App() {
  const [currUserId, setCurrUserId] = useState('');
  const handleCurrUser = (value) => {
    setCurrUserId(value)
  }
    return (
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Title />} />
          <Route path="/signin" element={<LogIn />} />
          <Route path="/register" element={<Registration />} />
          <Route path="/threads" element={<Threads currUserId={currUserId} />} />
          <Route path="/thread/:threadid" element={<Thread currUserId={currUserId} />} />
          <Route path="/chat" element={<ChatPage onCurrUser={handleCurrUser} />} />
          <Route path="*" element={<Errors />} />
        </Routes>
      </BrowserRouter>
    );
}

export default App;
