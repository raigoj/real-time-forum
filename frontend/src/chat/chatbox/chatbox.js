import React, { useState } from 'react';
import './chatbox.css';
import Users from '../user-list/user-list';
import Chat from '../chat/chat.js';
import ChatFooter from '../chat_footer/chat_footer.js';

function ChatPage({ onCurrUser }) {
  const [dataFromChat, setDataFromChat] = useState('');
  const handleDataChange = (value) => {
    setDataFromChat(value);
  };
  const [newMessage, setNewMessage] = useState([]);
  const handleNewMessage = (value) => {
    setNewMessage(value);
  };
  const [dataUserChange, setDataFromUserClick] = useState(0);
  const handleUserChange = (value) => {
    setDataFromUserClick(value);
  };
  const [lastMessage, setLastMessage] = useState('');
  const handleLastMessage = (value) => {
    setLastMessage(value);
  };
  const handleLogOut = () => {
    const cookies = document.cookie.split(";"); // split cookies by semicolon
    for (let i = 0; i < cookies.length; i++) {
      const cookieName = cookies[i].split("=")[0]; // get cookie name
      document.cookie = cookieName + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT"; // set expiration to past to delete cookie
    }
    window.location.replace("http://localhost:3000/");
  };
    
  return (
    <div className="chat">
      <button
        className="back-button-chat"
        onClick={() => {
          window.location = `${window.location.protocol}//${window.location.host}/threads`;
        }}
      >
        Back to the Forum
      </button>

      <button className="back-button-chat-2" onClick={handleLogOut}>
        Log off
      </button>

      <Users dataFromChat={dataFromChat} onUserChange={handleUserChange} />
      <div className="chat__main">
        <Chat
          onDataChange={handleDataChange}
          dataUserChange={dataUserChange}
          newMessage={newMessage}
          onLastMessage={handleLastMessage}
          onCurrUser={onCurrUser}
        />
        <ChatFooter
          dataUserChange={dataUserChange}
          dataFromChat={dataFromChat}
          onNewMessage={handleNewMessage}
          lastMessage={lastMessage}
        />
      </div>
    </div>
  );
}

export default ChatPage;
