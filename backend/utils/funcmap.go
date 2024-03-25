package utils

import (
	"net/http"
	"strconv"
	"strings"
)

func Like(commentId int64) int64 {
	stmt := "SELECT like_dislike FROM CommentLikes WHERE comment_id = ?"
	row, err := Db.Query(stmt, commentId)
	CheckError(err)
	var likedislike string
	var likesdislikes []string
	for row.Next() {
		err = row.Scan(&likedislike)
		CheckError(err)
		likesdislikes = append(likesdislikes, likedislike)
	}
	var likes int64
	var i int
	for _, r := range likesdislikes {
		i, err = strconv.Atoi(r)
		CheckError(err)
		if i == 1 {
			likes += 1
		} else {
			likes -= 1
		}
	}
	return likes
}

func IdToUser(n int64) string {
	stmt := "SELECT username FROM Users WHERE user_id = ?"
	row := Db.QueryRow(stmt, n)
	var user string
	row.Scan(&user)
	return user
}

func PostName(w http.ResponseWriter, r *http.Request) string {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}
	stmt := "SELECT post_name FROM Posts WHERE post_id = ?"
	row := Db.QueryRow(stmt, id)
	var post_name string
	row.Scan(&post_name)
	return post_name
}

func PostContent(w http.ResponseWriter, r *http.Request) string {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}
	stmt := "SELECT post_content FROM Posts WHERE post_id = ?"
	row := Db.QueryRow(stmt, id)
	var user string
	row.Scan(&user)
	return user
}

func PostDate(w http.ResponseWriter, r *http.Request) string {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}
	stmt := "SELECT post_date FROM Posts WHERE post_id = ?"
	row := Db.QueryRow(stmt, id)
	var user string
	row.Scan(&user)
	return user
}

func PostAuthor(w http.ResponseWriter, r *http.Request) string {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}
	stmt := "SELECT user_id FROM Posts WHERE post_id = ?"
	row := Db.QueryRow(stmt, id)
	var user_id int64
	row.Scan(&user_id)
	stmt = "SELECT username FROM Users WHERE user_id = ?"
	row = Db.QueryRow(stmt, user_id)
	var user string
	row.Scan(&user)
	return user
}

func LikePost(postId int64) int64 {
	stmt := "SELECT like_dislike FROM PostLikes WHERE post_id = ?"
	row, err := Db.Query(stmt, postId)
	CheckError(err)
	var likedislike string
	var likesdislikes []string
	for row.Next() {
		err = row.Scan(&likedislike)
		CheckError(err)
		likesdislikes = append(likesdislikes, likedislike)
	}
	var likes int64
	var i int
	for _, r := range likesdislikes {
		i, err = strconv.Atoi(r)
		CheckError(err)
		if i == 1 {
			likes += 1
		} else {
			likes -= 1
		}
	}
	return likes
}

func IdToCategory(n int64) string {
	splN := SplitInt(n)
	stmt := "SELECT category FROM Categories WHERE category_id = ?"
	var category string
	var categories []string
	for i := range splN {
		row := Db.QueryRow(stmt, splN[i])
		row.Scan(&category)
		categories = append(categories, category)
	}
	strCategories := strings.Join(categories, ", ")
	return strCategories
}

func LoggedInId(w http.ResponseWriter, r *http.Request) string {
	uuid, err := r.Cookie("sessionID")
	CheckUuid(w, r, uuid)
	CheckBadRequest(err, w)
	uuidVal := uuid.Value
	uuidFromDb := "SELECT user_id FROM session WHERE token = ?"
	row := Db.QueryRow(uuidFromDb, uuidVal)
	var userID int64
	row.Scan(&userID)
	stmt := "SELECT username FROM Users WHERE user_id = ?"
	row = Db.QueryRow(stmt, userID)
	var user string
	row.Scan(&user)
	return user
}

func CheckCookie(w http.ResponseWriter, r *http.Request) string {
	_, err := r.Cookie("sessionID")
	if err != nil {
		return "/"
	} else {
		return "/home"
	}
}
