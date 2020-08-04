package auth

import (
	"database/sql"
	"github.com/lukasvdberk/opensource-discord/database"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ProfilePhoto string `json:"profilePhoto"`
}

// TODO add profile photo
func RegisterUser(user *User) *User {
	// Hashes the password and saves the user to the database.

	db := database.GetSqlConnection()
	stmtIns, err := db.Prepare("INSERT INTO User (username, password) VALUES(?, ?)")

	if err != nil {
		return nil
	}

	_, hashedPassword := hashPassword(user.Password)
	res, err := stmtIns.Exec(user.Username, hashedPassword)
	if err != nil {
		return nil
	}

	user.Id, err = res.LastInsertId()

	if err != nil {
		return nil
	}

	return user
}

func CheckUserCredentials(user **User) bool {
	// Hashes the password and checks if it is in the database.
	db := database.GetSqlConnection()
	stmtOut, err := db.Prepare("SELECT id, password FROM User WHERE username = ?") // ? = placeholder
	if err != nil {
		return false
	}

	rows, err := stmtOut.Query((*user).Username)
	if err != nil {
		return false
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// TODO make this a more generic method if there is another use case for multiple row fetching
	// TODO find out if scan can directly set in a struct
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		var value string
		for i, col := range values {
			if col != nil {
				column := columns[i]

				if column == "id" {
					(*user).Id, err = strconv.ParseInt(string(col), 10, 64)
				}

				if column == "password" {
					value = string(col)

					// user was found with the correct password
					if checkPasswordHash((*user).Password, value) {
						return true
					}
				}
			}
		}
	}
	return false
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (error, string) {
	// TODO make bcrypt cost a variable maybe a settings.yml or something
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return err, string(bytes)
}
