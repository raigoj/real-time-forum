package utils

type Users struct {
	User_id    int64  `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Age        int64  `json:"age"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Gender     string `json:"gender"`
}

type Posts struct {
	Post_id      int64  `json:"post_id"`
	Post_name    string `json:"post_name"`
	Post_content string `json:"post_content"`
	Post_date    string `json:"post_date"`
	User_id      int64  `json:"user_id"`
	Category_id  int64  `json:"category_id"`
}

type Comments struct {
	Comment_id      int64  `json:"comment_id"`
	Comment_content string `json:"comment_content"`
	Comment_date    string `json:"comment_date"`
	User_id         int64  `json:"user_id"`
	Post_id         int64  `json:"post_id"`
}

type LikesDislikes struct {
	Likedislike_id int64 `json:"likedislike_id"`
	Like_dislike   int64 `json:"like_dislike"`
	User_id        int64 `json:"user_id"`
	Post_id        int64 `json:"post_id"`
	Comment_id     int64 `json:"comment_id"`
}

type Categories struct {
	Category_id int64  `json:"category_id"`
	Category    string `json:"category"`
}

type CommentLikes struct {
	Likedislike_id int64 `json:"likedislike_id"`
	Like_dislike   int64 `json:"like_dislike"`
	User_id        int64 `json:"user_id"`
	Comment_id     int64 `json:"comment_id"`
}

type PostLikes struct {
	Likedislike_id int64 `json:"likedislike_id"`
	Like_dislike   int64 `json:"like_dislike"`
	User_id        int64 `json:"user_id"`
	Post_id        int64 `json:"post_id"`
}

type Messages struct {
	Message_id   int64  `json:"message_id"`
	Message      string `json:"message"`
	From_id      int64  `json:"from_id"`
	To_id        int64  `json:"to_id"`
	Message_date string `json:"message_date"`
}

type Message struct {
	Message_id int64  `json:"message_id"`
	Message    string `json:"message"`
	Sender     string `json:"sender"`
	Recipient  string `json:"recipient"`
	Time       string `json:"time"`
  Read       int64 `json:"read"`
}

type Sessions struct {
  User_id int64 `json:"user_id"`
  Token string `json:"token"`
  TimeCreated string `json:"timeCreated"`
}
