package auth

import (
	"fmt"
	"github.com/lukasvdberk/opensource-discord/database"
	"golang.org/x/crypto/bcrypt"
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
		fmt.Println("Error executing query")
		return nil
	}

	_, hashedPassword := hashPassword(user.Password)
	res, err := stmtIns.Exec(user.Username, hashedPassword)
	if err != nil {
		fmt.Println("Error executing query")
		return nil
	}

	user.Id, err = res.LastInsertId()

	if err != nil {
		fmt.Println("Failed to get id")
		return nil
	}

	return user
}

func hashPassword(password string) (error, string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return err, string(bytes)
}
