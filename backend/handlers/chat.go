package handlers

import (
	"encoding/json"
  //"fmt"
	"net/http"
  "strings"
	utils "real-time-forum/utils"
	_ "github.com/mattn/go-sqlite3"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
  vars := strings.Split(r.URL.Path, "/")
  userId1, userId2 := vars[2], vars[3] 
  //fmt.Println("what", userId1, "\n no way", userId2)
  fId1 := userId1 + ".0"
  fId2 := userId2 + ".0"
  query := "SELECT * FROM Chat WHERE (sender = ? AND recipient = ?) OR (sender = ? AND recipient = ?)"
  rows, err := utils.Db.Query(query, userId1, fId2, userId2, fId1)
  utils.CheckInternalServerError(err, w)

  var msg utils.Message
  var msgs []utils.Message
  for rows.Next() {
	  err = rows.Scan(&msg.Message_id, &msg.Message, &msg.Sender,
		  &msg.Recipient, &msg.Time, &msg.Read)
	  utils.CheckInternalServerError(err, w)
	  msgs = append(msgs, msg)
  }

  err = json.NewEncoder(w).Encode(msgs)
  utils.CheckError(err)	
}

func AllChatHandler(w http.ResponseWriter, r *http.Request) {
  vars := strings.Split(r.URL.Path, "/")
  userId1:= vars[2]
  ///fmt.Println("what", userId1, "\n no way", userId2)
  fId1 := userId1 + ".0"
  query := "SELECT * FROM Chat WHERE (sender = ?) OR (recipient = ?)"
  rows, err := utils.Db.Query(query, userId1, fId1)
  utils.CheckInternalServerError(err, w)

  var msg utils.Message
  var msgs []utils.Message
  for rows.Next() {
	  err = rows.Scan(&msg.Message_id, &msg.Message, &msg.Sender,
		  &msg.Recipient, &msg.Time, &msg.Read)
	  utils.CheckInternalServerError(err, w)
	  msgs = append(msgs, msg)
  }

  err = json.NewEncoder(w).Encode(msgs)
  utils.CheckError(err)	
}
