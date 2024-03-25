import {useEffect, useState} from 'react';
import axios from "axios";
import './user-list.css';

function Users({dataFromChat, onUserChange}) {
  const usersWithInteractions = useUsersWithLastInteraction(dataFromChat, onUserChange);
  const handleUserClick = userId => {
    onUserChange(userId);
  };
  if (!Array.isArray(usersWithInteractions)) {
    return <div>Loading...</div>;
  }
  if (!usersWithInteractions) {
    return <div>Loading...</div>;
  }
   
  const [data] = usersWithInteractions;
  const read = usersWithInteractions[2]; 
  const online = usersWithInteractions[3];
  //console.log("the data", read)

  const matchedData = data.map(obj1 => {
    let matched = false;
    for (let i = 0; i < read.length; i++) {
      if (obj1.id.toString() === read[i].userId) {
        matched = true;
        return { ...obj1, value2: read[i].count };
      }
    }
    if (!matched) return obj1;
  });
  //console.log(matchedData);

  const usersList = matchedData.map(user => (
    <div className="userlist" key={`user_` + user.id} onClick={() => handleUserClick(user.id)}>
      {online.includes(user.id) ? <span className="dot"></span> : null}
      <p>{user.username}</p>
      <p className="readmsgs">{user.value2}</p>
    </div>
  )); 

  return (
    <div className="chat__sidebar">
      <ul>{usersList}</ul>
    </div>
  );
}
export default Users;

function useUsersWithLastInteraction(currentUserId, onUserChange) {
  const [usersWithInteractions, setUsersWithInteractions] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [readArr, setReadArr] = useState({});
  const [online, setOnline] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const usersResponse = await axios.get("http://localhost:8080/users");
      const users = usersResponse.data;
      const messagesResponse = await axios.get("http://localhost:8080/allchat/" + currentUserId);
      const messages = messagesResponse.data;
      const loggedUsersResponse = await axios.get("http://localhost:8080/connected_clients");
      const loggedUsers = loggedUsersResponse.data;
      //console.log(loggedUsers)

      if (!messages || messages.length === 0) {
        const sortedUsers = users
          .filter(user => user.id !== currentUserId)
          .sort((a, b) => a.username.localeCompare(b.username));
        setUsersWithInteractions(sortedUsers);
        setIsLoading(false);
        return;
      }

      const clientIds = Object.keys(loggedUsers).map((id) => parseInt(id));

      setOnline(clientIds);

      const messageData = [];

      for (let i = 0; i < messages.length; i++) {
        const senderId = messages[i].sender;
        const read = messages[i].read;
        // Skip the current user's messages
        if (senderId === currentUserId.toString()) {
          continue;
        }
        if (read === 1) {
          continue;
        }
        // Check if the senderId already exists in the messageData array
        let senderExists = false;
        let senderIndex;
        for (let j = 0; j < messageData.length; j++) {
          if (messageData[j].userId === senderId) {
            senderExists = true;
            senderIndex = j;
            break;
          }
        }

        // If the senderId doesn't exist in the messageData array, add it
        if (!senderExists) {
          messageData.push({ userId: senderId, count: 0 });
          senderIndex = messageData.length - 1;
        }

        // Increment the count for this sender if the read value is 0
        if (read === 0) {
          messageData[senderIndex].count++;
        }
      }

      //console.log(messageData);
      setReadArr(messageData);

      const interactions = messages
        .reduce((acc, message) => {
          const senderId = parseInt(message.sender, 10);
          const recipientId = parseInt(message.recipient, 10);
          if (!acc[senderId]) {
            acc[senderId] = { id: senderId, timestamp: new Date(message.time) };
          } else {
            acc[senderId].timestamp = new Date(message.time);
          }
          if (!acc[recipientId]) {
            acc[recipientId] = { id: recipientId, timestamp: new Date(message.time) };
          } else {
            acc[recipientId].timestamp = new Date(message.time);
          }
            return acc;
        }, {}); 

       const sortedUsers = users
        .filter(user => user.id !== currentUserId)
        .map(user => ({ ...user, ...interactions[user.id] }))
        .sort((a, b) => {
        if (!a.timestamp && !b.timestamp) {
          return a.username.localeCompare(b.username);
        }
        if (!a.timestamp) {
          return 1;
        }
        if (!b.timestamp) {
          return -1;
        }
        return b.timestamp - a.timestamp;
      }); 
        setUsersWithInteractions(sortedUsers);
        setIsLoading(false);
    };

    fetchData();
  }, [currentUserId, onUserChange]);

  return [usersWithInteractions, isLoading, readArr, online];
}
