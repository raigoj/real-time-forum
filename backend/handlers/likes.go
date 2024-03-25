package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	utils "real-time-forum/utils"
)

func CreateCommentLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	uuid, err := r.Cookie("sessionID")
	if uuid != nil {
		var commentlikes utils.CommentLikes
		val := (r.URL.Query().Get("val"))
		cmtId := (r.URL.Query().Get("cmtId"))

		if utils.CheckCookieFromDb(w, r) {
			utils.CheckBadRequest(err, w)
			uuidVal := uuid.Value
			smtt := "SELECT user_id FROM session WHERE token = ?"
			row := utils.Db.QueryRow(smtt, uuidVal)
			var userID int64
			row.Scan(&userID)
			cmtIdInt64, err := strconv.ParseInt(cmtId, 10, 64)
			utils.CheckError(err)
			like_dislike, err := strconv.ParseInt(val, 10, 64)
			utils.CheckError(err)
			id := (r.URL.Query().Get("id"))
			threadUrl := ("/thread?id=" + id)

			commentlikes.User_id = userID
			commentlikes.Comment_id = cmtIdInt64
			if !utils.CheckCommentLikeBool(userID, cmtIdInt64) {
				commentlikes.Like_dislike = like_dislike
				stmt, err := utils.Db.Prepare(`
					INSERT INTO CommentLikes( like_dislike, user_id, comment_id)
					VALUES(?, ?, ?)
					`)
				if err != nil {
					fmt.Println("Prepare query error")
					panic(err)
				}
				_, err = stmt.Exec(commentlikes.Like_dislike,
					commentlikes.User_id, commentlikes.Comment_id)
				if err != nil {
					fmt.Println("Execute query error")
					panic(err)
				}
				http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
			} else {
				if utils.CheckCommentLikeInt64(userID, cmtIdInt64) != like_dislike {
					likedislikeNeg := like_dislike
					commentlikes.Like_dislike = likedislikeNeg
					stmt, err := utils.Db.Prepare(`UPDATE CommentLikes SET like_dislike = $1 WHERE user_id = $2 AND comment_id = $3;`)
					utils.CheckError(err)
					_, err = stmt.Exec(commentlikes.Like_dislike,
						commentlikes.User_id, commentlikes.Comment_id)
					utils.CheckError(err)
					http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
				} else {
					http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
				}
			}
		} else {

			c := http.Cookie{
				Name:   utils.COOKIE_NAME,
				Value:  "",
				MaxAge: -1,
				Path:   "/",
			}
			http.SetCookie(w, &c)

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	} else {
		t, err := template.New("error.html").ParseFiles(filepath.Join("statics/error.html"))
		utils.CheckInternalServerError(err, w)
		err = t.Execute(w, "You have to be logged in to like or dislike")
		utils.CheckInternalServerError(err, w)

	}

}

func CreatePostDislike(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	if utils.CheckCookieFromDb(w, r) {
		uuid, err := r.Cookie("sessionID")
		if uuid != nil {
			utils.CheckBadRequest(err, w)
			var postlikes utils.PostLikes
			val := (r.URL.Query().Get("val"))
			pstId := (r.URL.Query().Get("pstId"))
			splitPstId := strings.Split(pstId, " ")
			uuidVal := uuid.Value
			smtt := "SELECT user_id FROM session WHERE token = ?"
			row := utils.Db.QueryRow(smtt, uuidVal)
			var userID int64
			row.Scan(&userID)
			pstIdInt64, err := strconv.ParseInt(splitPstId[1], 10, 64)
			utils.CheckError(err)
			like_dislike, err := strconv.ParseInt(val, 10, 64)
			utils.CheckError(err)
			threadUrl := ("/home")
			if utils.CheckPostLikeBool(userID, pstIdInt64) {
				if utils.CheckPostLikeInt64(userID, pstIdInt64) == like_dislike {

					http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
				} else {
					postlikes.User_id = userID
					postlikes.Post_id = pstIdInt64
					likedislikeNeg := like_dislike
					postlikes.Like_dislike = likedislikeNeg
					stmt, err := utils.Db.Prepare(`UPDATE PostLikes SET like_dislike = ? WHERE user_id = ? AND post_id = ?;`)
					utils.CheckError(err)
					_, err = stmt.Exec(postlikes.Like_dislike,
						postlikes.User_id, postlikes.Post_id)
					utils.CheckError(err)
					http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
				}

			} else {
				postlikes.User_id = userID
				postlikes.Post_id = pstIdInt64
				postlikes.Like_dislike = like_dislike
				stmt, err := utils.Db.Prepare(`
					INSERT INTO PostLikes( like_dislike, user_id, post_id)
					VALUES(?, ?, ?)
					`)
				if err != nil {
					fmt.Println("Prepare query error")
					panic(err)
				}
				_, err = stmt.Exec(postlikes.Like_dislike,
					postlikes.User_id, postlikes.Post_id)
				if err != nil {
					fmt.Println("Execute query error")
					panic(err)
				}
				http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusBadRequest)
		}
	} else {
		c := http.Cookie{
			Name:   utils.COOKIE_NAME,
			Value:  "",
			MaxAge: -1,
			Path:   "/",
		}
		http.SetCookie(w, &c)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func CreatePostLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	if utils.CheckCookieFromDb(w, r) {
		uuid, err := r.Cookie("sessionID")
		if uuid != nil {
			utils.CheckBadRequest(err, w)
			var postlikes utils.PostLikes
			val := (r.URL.Query().Get("val"))
			pstId := (r.URL.Query().Get("pstId"))
			uuidVal := uuid.Value
			smtt := "SELECT user_id FROM session WHERE token = ?"
			row := utils.Db.QueryRow(smtt, uuidVal)
			var userID int64
			row.Scan(&userID)
			pstIdInt64, err := strconv.ParseInt(pstId, 10, 64)
			utils.CheckError(err)
			like_dislike, err := strconv.ParseInt(val, 10, 64)
			utils.CheckError(err)
			threadUrl := ("/home")
			if utils.CheckPostLikeBool(userID, pstIdInt64) {
				if utils.CheckPostLikeInt64(userID, pstIdInt64) == like_dislike {
					http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
				} else {
					postlikes.User_id = userID
					postlikes.Post_id = pstIdInt64
					likedislikeNeg := like_dislike
					postlikes.Like_dislike = likedislikeNeg
					stmt, err := utils.Db.Prepare(`UPDATE PostLikes SET like_dislike = ? WHERE user_id = ? AND post_id = ?;`)

					utils.CheckError(err)
					_, err = stmt.Exec(postlikes.Like_dislike,
						postlikes.User_id, postlikes.Post_id)
					utils.CheckError(err)
					http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
					defer stmt.Close()
				}

			} else {
				postlikes.User_id = userID
				postlikes.Post_id = pstIdInt64
				postlikes.Like_dislike = like_dislike
				stmt, err := utils.Db.Prepare(`
					INSERT INTO PostLikes( like_dislike, user_id, post_id)
					VALUES(?, ?, ?)
					`)
				utils.CheckInternalServerError(err, w)
				_, err = stmt.Exec(postlikes.Like_dislike,
					postlikes.User_id, postlikes.Post_id)
				utils.CheckInternalServerError(err, w)
				http.Redirect(w, r, threadUrl, http.StatusMovedPermanently)
				defer stmt.Close()
			}
		} else {
			http.Redirect(w, r, "/", http.StatusBadRequest)
		}
	} else {
		c := http.Cookie{
			Name:   utils.COOKIE_NAME,
			Value:  "",
			MaxAge: -1,
			Path:   "/",
		}
		http.SetCookie(w, &c)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func LikedPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
	var funcMap = template.FuncMap{
		"idToUser": func(n int64) string {
			return utils.IdToUser(n)
		},
		"idToCategory": func(n int64) string {
			return utils.IdToCategory(n)
		},
		"like": func(postId int64) int64 {
			return utils.LikePost(postId)
		},
	}
	if utils.CheckCookieFromDb(w, r) {
		uuidS, err := r.Cookie("sessionID")
		if uuidS != nil {
			utils.CheckBadRequest(err, w)
			uuidVal := uuidS.Value
			smtt := "SELECT user_id FROM session WHERE token = ?"
			row := utils.Db.QueryRow(smtt, uuidVal)
			var userId int64
			row.Scan(&userId)
			smttt := "SELECT post_id FROM PostLikes WHERE user_id = ?"
			var postId int64
			roww, err := utils.Db.Query(smttt, userId)
			utils.CheckError(err)
			var postIds []int64
			for roww.Next() {
				err = roww.Scan(&postId)
				utils.CheckInternalServerError(err, w)
				postIds = append(postIds, postId)
			}
			stmt := "SELECT * FROM Posts where post_id = ?"
			var post utils.Posts
			var posts []utils.Posts
			for x := range postIds {
				rows, err := utils.Db.Query(stmt, postIds[x])
				utils.CheckError(err)
				for rows.Next() {
					err = rows.Scan(&post.Post_id, &post.Post_name, &post.Post_content,
						&post.Post_date, &post.User_id, &post.Category_id)
					utils.CheckInternalServerError(err, w)
					posts = append(posts, post)
				}
			}
			t, err := template.New("likedthreads.html").Funcs(funcMap).ParseFiles(filepath.Join("statics/likedthreads.html"))
			utils.CheckInternalServerError(err, w)
			err = t.Execute(w, posts)
			utils.CheckInternalServerError(err, w)
		} else {
			http.Redirect(w, r, "/", http.StatusBadRequest)
		}
	} else {
		c := http.Cookie{
			Name:   utils.COOKIE_NAME,
			Value:  "",
			MaxAge: -1,
			Path:   "/",
		}
		http.SetCookie(w, &c)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
