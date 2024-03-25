package handlers

import (
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	utils "real-time-forum/utils"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := utils.Db.Query("SELECT * FROM Users")
	utils.CheckInternalServerError(err, w)

	var user utils.Users
	var users []utils.Users
	for rows.Next() {
		err = rows.Scan(&user.User_id, &user.Username, &user.Password, &user.Email,
			&user.Age, &user.First_name, &user.Last_name, &user.Gender)
		utils.CheckInternalServerError(err, w)
		users = append(users, user)
	}

	err = json.NewEncoder(w).Encode(users)

	utils.CheckError(err)
}
