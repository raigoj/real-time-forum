package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	utils "real-time-forum/utils"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

func CookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     utils.COOKIE_NAME,
		Value:    "ok",
		Expires:  time.Now().Add(10 * time.Minute),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
	}
	http.SetCookie(w, cookie)
	fmt.Println(283)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   utils.COOKIE_NAME,
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, &c)

	sessionId, err := r.Cookie(utils.COOKIE_NAME)
	if err != nil {
		// handle error
		return
	}
	err = utils.RemoveSessionFromDB(sessionId.Value)
	if err != nil {
		// handle error
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// map where key is username and value is type client
	clients = map[string]*client{}
)

func WsEndPoint(w http.ResponseWriter, r *http.Request) {
	// TODO: Identify logged in user, if not logged in return error
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	var err error
	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade: %s\n", err.Error())
		return

	}
	c := client{wsConn: wsConn, send: make(chan interface{})}
	// TODO: Add them to clients map
	go wsReader(wsConn)
	go c.wsWriter()

}

func wsReader(wsConn *websocket.Conn) {
	defer wsConn.Close()
	for {
		var msg utils.Message
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break

		}
		//fmt.Println("Message Received: ", msg.Message, "\n", msg.Sender, "\n", msg.Recipient, "\n", msg.Time)
		stmt, err := utils.Db.Prepare(`
						INSERT INTO Chat(message, sender, recipient, time)
						VALUES(?, ?, ?, ?)
					`)
		//fmt.Println(err)
		utils.CheckError(err)
		_, err = stmt.Exec(msg.Message, msg.Sender, msg.Recipient, msg.Time)
		utils.CheckError(err)
		// TODO: Select from clients and send message to channel
		// See https://gobyexample.com/channels

	}

}

type client struct {
	wsConn *websocket.Conn
	send   chan interface{}
}

func (c *client) wsWriter() {
	defer c.wsConn.Close()
	for {
		if err := c.wsConn.WriteJSON(<-c.send); err != nil {
			fmt.Printf("failed write to client: %v\n", err)
			return
		}
	}
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	fmt.Println(000000, username, "\n", password, "\n", ok)
	rows, err := utils.Db.Query("SELECT * FROM session")
	utils.CheckInternalServerError(err, w)

	var ssn utils.Sessions
	var ssns []utils.Sessions
	for rows.Next() {
		err = rows.Scan(&ssn.User_id, &ssn.Token, &ssn.TimeCreated)
		utils.CheckInternalServerError(err, w)
		ssns = append(ssns, ssn)
	}

	err = json.NewEncoder(w).Encode(ssns)

	utils.CheckError(err)

}
