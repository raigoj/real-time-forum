package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	utils "real-time-forum/utils"

	_ "github.com/mattn/go-sqlite3"
)

type ThreadResponse struct {
    Post     utils.Posts
    Comments []utils.Comments
}

func ThreadHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	stmtt := "Select * From Posts WHERE post_id = ?"
	rowss, errr := utils.Db.Query(stmtt, id)
	utils.CheckError(errr)
	var post utils.Posts
	if rowss.Next() {
		err = rowss.Scan(&post.Post_id, &post.Post_name, &post.Post_content,
			&post.Post_date, &post.User_id, &post.Category_id)
		utils.CheckError(err)
	} else {
		http.NotFound(w, r)
		return
	}
  defer rowss.Close()
	stmt := "SELECT * FROM Comments WHERE post_id = ?"
	rows, err := utils.Db.Query(stmt, id)
	utils.CheckInternalServerError(err, w)
	var comments []utils.Comments
	var comment utils.Comments
	for rows.Next() {
		err = rows.Scan(&comment.Comment_id, &comment.Comment_content,
			&comment.Comment_date, &comment.User_id, &comment.Post_id)
		utils.CheckInternalServerError(err, w)
		comments = append(comments, comment)
	}
  defer rows.Close()
	response := ThreadResponse{
		Post:     post,
		Comments: comments,
	}
	err = json.NewEncoder(w).Encode(response)
	utils.CheckError(err)
}

func CreateCommentsHandlerRefactored(w http.ResponseWriter, r *http.Request) {

	// function refactored to use a map to parse the request body and
	// store the values in a Comment struct. Also using json.Unmarshal to
	// parse the request body and store the values in the map.

	// Define a Comment struct
	type Comment struct {
		CommentContent string
		CommentDate    string
		UserID         int64
		PostID         int64
	}

	// Create a new Comment struct
	comment := Comment{}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Parse the request body into a map
	bodyMap := make(map[string]string)
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		panic(err)
	}

	// Get the values from the map
	cookie := bodyMap["cookie"]
	commentContent := bodyMap["commentContent"]
	postId := bodyMap["postId"]

	// Get the user ID from the database
	stmt := "SELECT user_id FROM session WHERE token = ?"
	row := utils.Db.QueryRow(stmt, cookie)
	var userID int64
	err = row.Scan(&userID)
	if err != nil {
		panic(err)
	}

	// Set the values of the Comment struct
	comment.CommentContent = commentContent
	if comment.CommentContent != "" {
		comment.CommentDate = time.Now().Format("02 Jan 2006, 15:04")
		comment.UserID = userID
		comment.PostID, _ = strconv.ParseInt(postId, 10, 64)

		// Save to database
		stmt, err := utils.Db.Prepare(`
                INSERT INTO Comments(comment_content, comment_date, user_id, post_id)
                VALUES(?, ?, ?, ?)
            `)
		if err != nil {
			fmt.Println("Prepare query error")
			panic(err)
		}
		_, err = stmt.Exec(comment.CommentContent, comment.CommentDate,
			comment.UserID, comment.PostID)
		if err != nil {
			fmt.Println("Execute query error")
			panic(err)
		}
	}
}

func CreateCommentsHandler(w http.ResponseWriter, r *http.Request) {

	var comment utils.Comments

	body, err := ioutil.ReadAll(r.Body)
	utils.CheckError(err)
	fmt.Println(string(body))
	bodyStr := string(body)
	spltBody := strings.Split(bodyStr, "\"")
	commentContent := string(spltBody[3])
	postId := string(spltBody[11])
	cookie := string(spltBody[7])
	fmt.Println(commentContent, postId, cookie)

	//uuidVal := uuidS.Value
	stmt := "SELECT user_id FROM session WHERE token = ?"
	row := utils.Db.QueryRow(stmt, cookie)
	var userID int64
	row.Scan(&userID)
	comment.Comment_content = commentContent
	if comment.Comment_content != "" {
		comment.Comment_date = time.Now().Format("02 Jan 2006, 15:04")
		comment.User_id = userID
		fmt.Println(comment.User_id)
		comment.Post_id, _ = strconv.ParseInt(postId, 10, 64)
		fmt.Println(comment.Post_id)

		// Save to database
		stmt, err := utils.Db.Prepare(`
					INSERT INTO Comments(comment_content, comment_date, user_id, post_id)
					VALUES(?, ?, ?, ?)
				`)
		if err != nil {
			fmt.Println("Prepare query error")
			panic(err)
		}
		_, err = stmt.Exec(comment.Comment_content, comment.Comment_date,
			comment.User_id, comment.Post_id)
		if err != nil {
			fmt.Println("Execute query error")
			panic(err)
		}

	} else {

	}

}
