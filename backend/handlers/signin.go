package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	utils "real-time-forum/utils"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var loginErr string
	err := r.ParseForm()
	utils.CheckError(err)
	body, err := ioutil.ReadAll(r.Body)
	utils.CheckError(err)
	fmt.Println(string(body))
	bodyStr := string(body)
	spltBody := strings.Split(bodyStr, "\"")
	username := string(spltBody[3])
	password := string(spltBody[7])
	//fmt.Println(username)
	//fmt.Println(password)
	// retrieve password from db to compare (hash) with user supplied password's hash
	var hash string
	var userID int
	stmt := "SELECT Hash, user_id FROM Users WHERE username = ?"
	row := utils.Db.QueryRow(stmt, username)
	err = row.Scan(&hash, &userID)
	if err != nil {
		fmt.Println("error selecting Hash in db by Username")
		loginErr = "username and password do not match"
		err = json.NewEncoder(w).Encode(loginErr)
		utils.CheckError(err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil { // returns nil on success
		err = json.NewEncoder(w).Encode("")
		utils.CheckError(err)
		utils.InMemorySession = utils.NewSession() //creates a new session
		sessionId := utils.InMemorySession.Init(username)

		if err := utils.AddSessionToDB(sessionId, userID); err != nil {
			http.Error(w, "something wrong", http.StatusInternalServerError)
			fmt.Println("Error: ", err)
			return
		}
		var sessionIdDb string
		var cookieVals []string
		stmt = `SELECT token FROM session WHERE user_id = ?`
		row = utils.Db.QueryRow(stmt, userID)
		row.Scan(&sessionIdDb)
		cookie := &http.Cookie{
			Name:     utils.COOKIE_NAME,
			Value:    sessionIdDb,
			Expires:  time.Now().Add(10 * time.Minute),
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Secure:   true,
		}
		fmt.Println(cookie)
		http.SetCookie(w, cookie)
		cookieVals = []string{cookie.Name, cookie.Value}
		fmt.Println(260, cookieVals)
		err = json.NewEncoder(w).Encode(cookieVals)
		utils.CheckError(err)
		return
	}

	//t, err := template.New("login.html").ParseFiles(filepath.Join("statics/login.html"))
	//utils.CheckInternalServerError(err, w)
	loginErr = "username and password do not match"

	err = json.NewEncoder(w).Encode(loginErr)
	utils.CheckError(err)

}
