import React, { useState, useEffect, useRef } from 'react';
import axios from "axios";

export let ws = {}
const Chat = ({onDataChange, dataUserChange, newMessage, onLastMessage, onCurrUser}) => {
  const [currentUserId, setCurrentUserId] = useState('');
  const [currentUserName, setCurrentUserName] = useState('');
  const [currentRecipientName, setCurrentRecipientName] = useState('');
  const [msgs, setMsgs] = useState([]);
  const messagesEndRef = useRef(null)
  const [limit, setLimit] = useState(10);
  const messagesContainerRef = useRef(null);
  const [scrollPos, setScrollPos] = useState(0);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" })
  }
  useEffect(() => {
    const observer = new IntersectionObserver(entries => {
      if (entries[0].isIntersecting) {
        scrollToBottom();
      }
    });
    observer.observe(messagesEndRef.current);

    return () => {
      observer.disconnect();
    };
  }, []);

  useEffect(() => {
     const fetchData = async () => {
      const sessionResponse = await axios.get("http://localhost:8080/session");
      const session = sessionResponse.data;
      const resD = session.find(token => token.token === getCookie("sessionID"));
      if (resD === undefined) {
        window.location.href = '/';
      }
      setCurrentUserId(resD.user_id);
      onCurrUser(currentUserId) 
      onDataChange(resD.user_id)
      const usersResponse = await axios.get("http://localhost:8080/users");
      const users = await usersResponse.data;
      const resData = users.find(token => token.id === currentUserId);
      setCurrentUserName(resData.username);
      let recipientId = dataUserChange
      if (recipientId === 0) {
        return;
      }
      const recName = users.find(token => token.id === dataUserChange);
      setCurrentRecipientName(recName.username);
      const chatResponse = await axios.get(`http://localhost:8080/chat/${currentUserId}/${recipientId}`);
      const chat = chatResponse.data;
      setMsgs(chat)
      //console.log(chat)
    }
    fetchData();
  }, [currentUserId, onDataChange, dataUserChange, currentRecipientName, onCurrUser]);

  const handleScroll = () => {
    setScrollPos(messagesContainerRef.current.scrollTop);

    if (
      messagesContainerRef.current.scrollTop === 0
    ) {
      setLimit(prevLimit => prevLimit + 10);
      const prevScrollHeight = messagesContainerRef.current.scrollHeight;
      setTimeout(() => {
        messagesContainerRef.current.scrollTop =
          messagesContainerRef.current.scrollHeight - prevScrollHeight;
      }, 0);
    } else if (
      messagesEndRef.current &&
      messagesContainerRef.current.scrollTop + messagesContainerRef.current.offsetHeight >=
      messagesContainerRef.current.scrollHeight
    ) {
      messagesEndRef.current.scrollIntoView({ behavior: "smooth" });
    }
  };

  useEffect(() => {
    messagesContainerRef.current.addEventListener("scroll", handleScroll);
    return () => {
      messagesContainerRef.current.removeEventListener("scroll", handleScroll);
    };
  }, [handleScroll, messagesEndRef]);

   useEffect(() => {
    const messagesContainer = messagesContainerRef.current;
    messagesContainer.scrollTop = messagesContainer.scrollHeight;
  }, [msgs]); 
  
//console.log(131111, msgs) 
  const [msgsLength, setMsgsLength] = useState(0);
  useEffect(() => {
    if (msgs) {
      setMsgsLength(msgs.length);
    }
  }, [msgs]); 
  const [count, setCount] = useState(0)

  const readMessageRef = useRef(null);

  //useEffect(() => {
    //const observer = new IntersectionObserver((entries) => {
      //entries.forEach((entry) => {
        //if (entry.isIntersecting && !msgs.read) {
          // send websocket message to let the backend know the message has been read
          //console.log(msgs)
        //}
      //});
    //});
    //observer.observe(readMessageRef.current);
    //return () => {
      //observer.unobserve(readMessageRef.current);
    //};
  //}, [msgs]);
  const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        //console.log("this the type of the read value", entry.target)
        if (entry.target.dataset.read === "0") {
          // Send a websocket message to the backend to change the read value to 1 in the database
          ws.send(JSON.stringify({ message: entry.target.dataset.msgid, targetId: 234982293.0, init: false }))

          console.log(entry.target.dataset.read, entry.target.dataset.msgid)
          observer.unobserve(entry.target);

        }
        observer.unobserve(entry.target);
      }
    });
  });

  let msgsLis = (<div>Welcome to the chatroom!</div>)
    if (dataUserChange !== 0) {
      if (msgs && msgs.length !== 0) {
        if (msgs && msgs.length !== msgsLength) {
          setMsgsLength(msgs.length);
          if (count === 0) {
            setCount(count + 1)
          } else {
            setLimit(limit + 1)
          }
        }
        const displayedMsgs = msgs.slice(Math.max(msgs.length - limit, 0), msgs.length);
        msgsLis = displayedMsgs.map((msg) => { 
          if (currentUserId.toString() === msg.sender) {
            return (
              <div className="message__chats" key={'msg_' + msg.message_id}>
                <p className="sender__name">{currentUserName}</p>
                <p className='sender__name'>{msg.time}</p>
                <div className="message__sender">
                  <p>{msg.message}</p>
                </div>
              </div>
            );
          } else {
            return (
              <div className="message__chats" key={'msg_' + msg.message_id} data-read={msg.read} data-msgid={msg.message_id} ref={el => el && observer.observe(el)}>
                <p className="recipient__name">{currentRecipientName}</p>
                <p className='recipient__name'>{msg.time}</p>
                <div className="message__recipient">
                  <p>{msg.message}</p>
                </div>
              </div>
            );
          }
        }); 
      } else {
        msgsLis = (<div> You have no message history with {currentRecipientName}</div>)
      }
    }

  const msgsDiv = (
    <div className="message__container" ref={messagesContainerRef}>
      {msgsLis}
    </div>
  ) 
  return (
    <>
      <header className="chat__mainHeader">
      </header>
      <div key={JSON.stringify(msgs)}>
        {msgsDiv} 
        <div ref={messagesEndRef} />
      </div>
    </>
  );
};

export default Chat;

export const wsSetup = () => { 

  if (ws.readyState !== 0) { 
    ws = new WebSocket("ws://localhost:8080/socket") 
  } 
  let cookie = getCookie("sessionID");
  ws.onopen = function() {
    ws.send(JSON.stringify({username: cookie, init: true}))
  }

}


export function getCookie(cName) {
      const name = cName + "=";
      const cDecoded = decodeURIComponent(document.cookie); //to be careful
      const cArr = cDecoded.split('; ');
      let res;
      cArr.forEach(val => {
        if (val.indexOf(name) === 0) res = val.substring(name.length);
      })
      return res
}

