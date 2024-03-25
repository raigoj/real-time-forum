package utils

import (
	"database/sql"
	"net/http"
	"regexp"
	"time"
	"unicode"
)

var Db *sql.DB

func CheckUuid(w http.ResponseWriter, r *http.Request, uuid *http.Cookie) {
	if uuid == nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
	} else {
		return
	}
}

func SplitInt(n int64) []int64 {
	slc := []int64{}
	for n > 0 {
		if n%10 != 0 {
			slc = append(slc, n%10)
			n = n / 10
		} else {
			n = n / 10
		}
	}
	return slc
}

func CheckInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func CheckBadRequest(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, "Bad Request 400, http: named cookie not present", http.StatusBadRequest)
		panic(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func AddSessionToDB(uuid string, user_id int) error {
	var ins *sql.Stmt
	stmt := `INSERT INTO session (token, user_id, timeCreated) VALUES (?, ?, ?)`
	ins, err := Db.Prepare(stmt)
	CheckError(err)
	defer ins.Close()
	timeCurrent := time.Now()
	if _, err := ins.Exec(uuid, user_id, timeCurrent); err != nil {
		return err
	}
	return nil
}

func IsValidName(username string) bool {
	var nameAlphaNumeric = true
	for _, char := range username {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
	var nameLength bool
	if 5 <= len(username) && len(username) <= 50 {
		nameLength = true
	}
	if nameAlphaNumeric && nameLength {
		return true
	}
	return false

}

func IsValidPassword(password string) bool {
	var pswdLowercase, pswdUppercase, pswdNumber, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			pswdLowercase = true
		case unicode.IsUpper(char):
			pswdUppercase = true
		case unicode.IsNumber(char):
			pswdNumber = true
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 6 < len(password) && len(password) < 60 {
		pswdLength = true
	}
	if pswdLowercase && pswdUppercase && pswdNumber && pswdLength && pswdNoSpaces {
		return true
	}
	return false
}
func IsValidEmail(email string) bool {
	// check email syntax is valid
	emailRegex, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		return false
	}
	rg := emailRegex.MatchString(email)
	if !rg {
		return false
	}
	// check email length
	if len(email) < 4 {
		return false
	}
	if len(email) > 253 {
		return false
	}
	return true
}

func CreateCommentsTable(db *sql.DB) error {
	commentsTable := `CREATE TABLE IF NOT EXISTS Comments (
        comment_id INTEGER PRIMARY KEY AUTOINCREMENT,        
        comment_content TEXT,
        comment_date TEXT,
        user_id INTEGER,
		post_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES Users(user_id),
		FOREIGN KEY (post_id) REFERENCES Posts(post_id)
      );`

	statement, err := db.Prepare(commentsTable)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func CheckCookieFromDb(w http.ResponseWriter, r *http.Request) bool {
	uuid, _ := r.Cookie("sessionID")
	token := uuid.Value
	var user_id int64
	stmt := `SELECT user_id FROM session where token = ?`
	row := Db.QueryRow(stmt, token)
	row.Scan(&user_id)
	var exists bool
	if user_id == 0 {
		exists = false
	} else {
		exists = true
	}
	return exists
}
