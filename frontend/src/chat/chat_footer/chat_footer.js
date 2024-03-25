import React, { useState, useEffect } from 'react';
import {ws, wsSetup} from '../chat/chat.js'

const ChatFooter = ({dataUserChange, onNewMessage, lastMessage}) => {
  
  useEffect(() => {
    wsSetup()
  }, []);
  
  const [currentMessage, setCurrentMessage] = useState('')
  const [messages, setMessages] = useState([]);
   const handleSubmit = (message) => {
  let messagesCopy = [...messages];
  var time = new Date().getTime();
  var date = new Date(time);
  //console.log("??whatsthis??", messagesCopy)
  onNewMessage({ message_id: lastMessage, message: messagesCopy.message, recipient: dataUserChange + '.0', time: messagesCopy.date })
  messagesCopy.push({ sent: message.Sent, message: message.Message, date: date.toString().split("GMT")[0] });
  setMessages(prevMessages => [...prevMessages, { sent: message.Sent, message: message.Message, date: date.toString().split("GMT")[0] }]);
  setCurrentMessage('');
};


ws.onmessage = function(event) {
  let jsonData = JSON.parse(event.data);
  handleSubmit(jsonData);
  //console.log("chatfooter18", jsonData);
};  
  let textfandsubmit = (<div>Choose someone to chat with</div>)
  if (dataUserChange !== 0) {
    textfandsubmit = (
      <div className="chat__footer">
      <form className="form" onSubmit={(event) => {
        event.preventDefault();
        ws.send(JSON.stringify({ message: currentMessage, targetId: dataUserChange, init: false }))
      }}> 
        <input
          type="text"
          placeholder="Write message"
          className="message"
          value={currentMessage}
          onChange={(e) => setCurrentMessage(e.target.value)}
        />
        <button className="sendBtn">SEND</button>
      </form>
    </div>

    )
  }

  return (
    <>
      {textfandsubmit}
    </>
  );
};
export default ChatFooter;

