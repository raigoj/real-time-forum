package utils

func CheckCommentLikeBool(user_id, comment_id int64) bool {
	// checks if this like with the same user_id and comment_id already exists
	stmt := "SELECT like_dislike FROM CommentLikes WHERE user_id = ? AND comment_id = ?"
	row := Db.QueryRow(stmt, user_id, comment_id)
	var like int64
	var likeExists bool
	err := row.Scan(&like)
	if err != nil {
		likeExists = false
	} else {
		likeExists = true
	}
	// returns true if exists, false if doesn't
	return likeExists
}

func CheckCommentLikeInt64(user_id, comment_id int64) int64 {
	// checks if this like with the same user_id and comment_id already exists
	stmt := "SELECT like_dislike FROM CommentLikes WHERE user_id = ? AND comment_id = ?"
	row := Db.QueryRow(stmt, user_id, comment_id)
	var like int64
	err := row.Scan(&like)
	CheckError(err)
	// returns -1 or 1
	return like
}

func CheckPostLikeBool(user_id, post_id int64) bool {
	// checks if this like with the same user_id and comment_id already exists
	stmt := "SELECT like_dislike FROM PostLikes WHERE user_id = ? AND post_id = ?"
	row := Db.QueryRow(stmt, user_id, post_id)
	var like int64
	var likeExists bool
	err := row.Scan(&like)
	if err != nil {
		likeExists = false
	} else {
		likeExists = true
	}
	// returns true if exists, false if doesn't
	return likeExists
}

func CheckPostLikeInt64(user_id, post_id int64) int64 {
	// checks if this like with the same user_id and comment_id already exists
	stmt := "SELECT like_dislike FROM PostLikes WHERE user_id = ? AND post_id = ?"
	row := Db.QueryRow(stmt, user_id, post_id)
	var like int64
	err := row.Scan(&like)
	CheckError(err)
	// returns -1 or 1
	return like
}
