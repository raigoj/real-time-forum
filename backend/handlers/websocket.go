package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	utils "real-time-forum/utils"
	"time"

	"github.com/gorilla/websocket"
)

// This article helped me a lot with this project:
// 	https://www.thepolyglotdeveloper.com/2016/12/create-real-time-chat-app-golang-angular-2-websockets/
type ClientManager struct {
    clients    map[idType]*websocket.Conn
    register   chan *Client
    unregister chan *Client
}

type Client struct {
	id     idType
    socket *websocket.Conn
}

type idType int

var manager = ClientManager{
    register:   make(chan *Client),
    unregister: make(chan *Client),
    clients:    make(map[idType]*websocket.Conn),
}

func (manager *ClientManager) start() {
    for {
        select {
        case conn := <-manager.register:
            fmt.Println("clients", manager.clients)
            manager.clients[conn.id] = conn.socket
        case conn := <-manager.unregister:
            if _, ok := manager.clients[conn.id]; ok {
                delete(manager.clients, conn.id)
            }
        }
    }
}
// ToDo:
//	think of a better name for this function
func (c *Client) reader(conn *websocket.Conn) {
	unregister := func() {
		manager.unregister <- c
		conn.Close()
		log.Println("Client Connection Terminated")
	}
	for {
		// reads the message from a clients connection
		messageType, p, err := c.socket.ReadMessage()
		if err != nil {
			unregister()
			return
		}
    //fmt.Println(63, messageType, "\n", string(p))
		type toClient struct {
			Message  string
			Sent	 bool
			SenderId idType
		}
		// Clients message should be a json {
		// 	message string
		// 	senderId float64
		// 	targetId float64
		// 	init bool
		// }
		var v map[string]interface{}
		json.Unmarshal([]byte(p), &v)
		message  := v["message"]
		targetId := v["targetId"]
		init 	 := v["init"]
    //fmt.Println("this is the thing:", targetId)
	  if targetId == 234982293.0 {
      err = Readmsg(message)
      utils.CheckError(err)
      continue
    }	
		// registers the user with their ID
		if init.(bool) {
			username := v["username"]
			user, err := FromUsers("user_id", username)
      //fmt.Println(85, user, username)
			if err != nil || len(user) == 0 {
				unregister()
				return
			}
			c.id = idType(user[0].User_id)
			manager.register <- c
			continue
		}
		var to toClient
		to.Message = message.(string)
		to.Sent = false
		to.SenderId = c.id
		jsonTo, err := json.Marshal(to)
		if err != nil {
			continue
		}
		//look through global variable clientArray
		//find targetId and write message to the senders and the targets connection
		
		target := idType(targetId.(float64))
		value, isValid := manager.clients[target]
		// If target's connection is valid then WriteMessage to his connection
		if isValid {
			err = value.WriteMessage(messageType, []byte(jsonTo))
      //fmt.Println("???", string(jsonTo))
			if err != nil {
				utils.CheckError(err)
			}
		}
		to.Sent = true
		jsonTo, err = json.Marshal(to)
		if err != nil {
			utils.CheckError(err)
			continue
		}
		err = Message(c.id, targetId, message)
    utils.CheckError(err)
		
		
		err = conn.WriteMessage(messageType,[]byte(jsonTo))

		if err != nil {
			utils.CheckError(err)
		}
	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request){
	go manager.start()
	
	connection, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
  //fmt.Println(135, connection)
	client := &Client{id: 0, socket: connection}
  //fmt.Println(147, client)
	log.Println("Client Connection Established")
	client.reader(connection)
}

func Readmsg(messageId interface{}) error {
  stmt, err := utils.Db.Prepare("UPDATE Chat SET read = '1' WHERE msg_id = ?")

  if err != nil {
    return err
  }
  defer stmt.Close()
  res, err := stmt.Exec(messageId)
  if err != nil {
    utils.CheckError(err)
  } else {
    fmt.Println(res)
  }
  return nil
}

func Message(senderId, targetId, message interface{}) error {
  //targetIdInt := int64(targetId)
	stmt, err := utils.Db.Prepare("INSERT INTO Chat (message, sender, recipient, time, read) VALUES (?, ?, ?, ?, ?)")
	
	if err != nil {
		return err
	}
	defer stmt.Close()
  res, err := stmt.Exec(message, senderId, targetId, time.Now().Format("02 Jan 2006, 15:04"), 0)
  if err != nil {
    utils.CheckError(err)
  } else {
    fmt.Println(res)
  }
	return nil
}

func FromUsers(condition string, value interface{}) ([]utils.Users, error){
  stmt := "SELECT user_id FROM session WHERE token = ?"
	row := utils.Db.QueryRow(stmt, value)
	var userID int64
  row.Scan(&userID)
	rows, err := doRows("Users", condition, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var returnUser []utils.Users
	for rows.Next() {
		var user utils.Users
		err = rows.Scan(
			&user.User_id, 
			&user.Email, 
			&user.Username, 
			&user.Password,
			&user.Age,
			&user.First_name,
			&user.Last_name,
			&user.Gender,
    )
		if err != nil {
			return nil, err
		}
		returnUser = append(returnUser, user)
	}
	return returnUser, nil
}

func doRows(FROM, condition string, value interface{}) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	// Check if there is a condition
	switch condition {
	// No condition
	case "":
		rows, err = utils.Db.Query("SELECT * FROM " + FROM)
	
	// With condition
	default:
		rows, err = utils.Db.Query("SELECT * FROM " + FROM + " WHERE " + condition + " = $1", value)
	}
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func GetConnectedClients(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(manager.clients)
}
