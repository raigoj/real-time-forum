package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	utils "real-time-forum/utils"
	"strconv"
	"strings"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	var regErr string
	body, err := ioutil.ReadAll(r.Body)
	utils.CheckError(err)
	fmt.Println(string(body))
	bodyStr := string(body)
	spltBody := strings.Split(bodyStr, "\"")
	username := string(spltBody[3])
	password := string(spltBody[7])
	age := string(spltBody[15])
	iAge, err := strconv.Atoi(age)
	utils.CheckError(err)
	gender := string(spltBody[19])
	firstname := string(spltBody[23])
	lastname := string(spltBody[27])
	email := string(spltBody[11])
	fmt.Println(username, "\n", password, "\n", iAge, "\n", gender, "\n", firstname, "\n", lastname, "\n", email)
	//r.ParseForm()
	//email := r.FormValue("email")
	validEmail := utils.IsValidEmail(email)
	//username := r.FormValue("username")
	validUsername := utils.IsValidName(username)
	//password := r.FormValue("password")
	validPassword := utils.IsValidPassword(password)
	if !validEmail || !validPassword || !validUsername {
		regErr = "please check username and password criteria"
		err = json.NewEncoder(w).Encode(regErr)
		utils.CheckError(err)
		return
	}
	stmt := "SELECT user_id FROM Users WHERE email = ?"
	row := utils.Db.QueryRow(stmt, email)
	var uID string
	err = row.Scan(&uID)
	if err != sql.ErrNoRows {
		regErr = "email already used"
		err = json.NewEncoder(w).Encode(regErr)
		utils.CheckError(err)
		return
	}
	stmt = "SELECT user_id FROM users WHERE username = ?"
	row = utils.Db.QueryRow(stmt, username)
	err = row.Scan(&uID)
	if err != sql.ErrNoRows {
		regErr = "username already taken"
		err = json.NewEncoder(w).Encode(regErr)
		utils.CheckError(err)
		return
	}
	// create hash from password
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println(string(hash))
	if err != nil {
		regErr = "there was a problem registering account"
		err = json.NewEncoder(w).Encode(regErr)
		utils.CheckError(err)
		return
	}
	var insertStmt *sql.Stmt
	insertStmt, err = utils.Db.Prepare("INSERT INTO Users (username, Hash, email, age, gender, firstname, lastname) VALUES (?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		regErr = "there was a problem registering account"
		err = json.NewEncoder(w).Encode(regErr)
		utils.CheckError(err)
		return
	}
	defer insertStmt.Close()
	_, err = insertStmt.Exec(username, hash, email, iAge, gender, firstname, lastname)
	if err != nil {
		regErr = "there was a problem registering account"
		err = json.NewEncoder(w).Encode(regErr)
		utils.CheckError(err)
		return
	}

}
