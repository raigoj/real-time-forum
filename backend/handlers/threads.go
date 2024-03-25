package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	utils "real-time-forum/utils"
)

func ThreadsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := utils.Db.Query("SELECT * FROM Posts")
	utils.CheckInternalServerError(err, w)

	var post utils.Posts
	var posts []utils.Posts
	for rows.Next() {
		err = rows.Scan(&post.Post_id, &post.Post_name, &post.Post_content,
			&post.Post_date, &post.User_id, &post.Category_id)
		utils.CheckInternalServerError(err, w)
		posts = append(posts, post)
	}

	err = json.NewEncoder(w).Encode(posts)

	utils.CheckError(err)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
  utils.CheckError(err)
  fmt.Println(string(body))
  bodyStr := string(body)
  spltBody := strings.Split(bodyStr, "\"")
  postName := string(spltBody[3])
  postContent := string(spltBody[7])
  cat := string(spltBody[11])
  cookie := string(spltBody[15])
  fmt.Println(postName, postContent, cat, cookie)
  //email := string(spltBody[11])
//if utils.CheckCookieFromDb(w, r) {
		var post utils.Posts
		
		//if uuidS != nil {
			utils.CheckBadRequest(err, w)
			//uuidVal := uuidS.Value
			stmt := "SELECT user_id FROM session WHERE token = ?"
			row := utils.Db.QueryRow(stmt, cookie)
			var userID int64
			row.Scan(&userID)
			post.Post_name = postName
			post.Post_content = postContent
			post.Post_date = time.Now().Format("02 Jan 2006, 15:04")
			post.User_id = userID
			CategoryId164, err2 := strconv.ParseInt(cat, 10, 64)
			utils.CheckError(err2)
      post.Category_id = CategoryId164
			if post.Post_name != "" && post.Post_content != "" && post.Category_id != 0 {
				// Save to database
				stmt, err1 := utils.Db.Prepare(`
						INSERT INTO Posts(post_name, post_content, post_date, user_id, category_id)
						VALUES(?, ?, ?, ?, ?)
					`)
				if err1 != nil {
					fmt.Println("Prepare query error")
					panic(err)
				}
				_, err = stmt.Exec(post.Post_name, post.Post_content, post.Post_date,
					post.User_id, post.Category_id)
				if err != nil {
					fmt.Println("Execute query error")
					panic(err)
				}
			} else {	
			  err = json.NewEncoder(w).Encode("Thread title, content or category fields were empty and the thread was not created")
      }
}
