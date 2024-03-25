package main

import (
	"database/sql" // base driver for using sql
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	handler "real-time-forum/handlers"
	utils "real-time-forum/utils"
)

func main() {
	var err error
	utils.Db, err = sql.Open("sqlite3", "./users.db") // opens the database
	utils.CheckError(err)
	defer utils.Db.Close()
	http.Handle("/statics", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))))
	http.Handle("/js", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.HandleFunc("/register", Middleware(handler.RegistrationHandler))
	http.HandleFunc("/signin", Middleware(handler.Signin))
	http.HandleFunc("/home", Middleware(handler.ThreadsHandler))
	http.HandleFunc("/thread", Middleware(handler.ThreadHandler))
	http.HandleFunc("/createComments", Middleware(handler.CreateCommentsHandler))
	http.HandleFunc("/create", Middleware(handler.CreateHandler))
	http.HandleFunc("/createCommentLike", Middleware(handler.CreateCommentLike))
	http.HandleFunc("/createPostLike", Middleware(handler.CreatePostLike))
	http.HandleFunc("/createPostDislike", Middleware(handler.CreatePostDislike))
	http.HandleFunc("/likedposts", Middleware(handler.LikedPosts))
	http.HandleFunc("/socket", handler.WsEndpoint)
	http.HandleFunc("/logout", Middleware(handler.Logout))
	http.HandleFunc("/cookie", Middleware(handler.CookieHandler))
	http.HandleFunc("/users", Middleware(handler.UsersHandler))
	http.HandleFunc("/chat/", Middleware(handler.ChatHandler))
  http.HandleFunc("/allchat/", Middleware(handler.AllChatHandler))
  http.HandleFunc("/session", Middleware(handler.SessionHandler))
  http.HandleFunc("/connected_clients", Middleware(handler.GetConnectedClients))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.HandleFunc("/newthread", Middleware(handler.ThreadHandler))
	fmt.Println("Welcome to the Forum! Go to: http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // only ever returns in case of unexpected error
	}
}

func Middleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authentication")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.WriteHeader(http.StatusOK)
		fn.ServeHTTP(w, r)
	}
}
