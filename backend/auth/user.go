package auth

import (
	"github.com/lukasvdberk/opensource-discord/database"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	ProfilePhoto string `json:"profilePhoto"`
	password     string
}

func RegisterUser(user *User) *User {
	// Only requires username and password to be filled!
	// Hashes the password and saves the user to the database.

	db := database.GetSqlConnection()
	stmtIns, err := db.Prepare("INSERT INTO User (username, password) VALUES(?, ?)")

	if err != nil {
		return nil
	}

	_, hashedPassword := hashPassword(user.password)
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
	usersListMap := database.SelectStatement(
		"SELECT id, password FROM User WHERE username = ?",
		(*user).Username,
	)

	// an error occurred or there was no user found
	if usersListMap == nil {
		return false
	}

	for _, userMap := range usersListMap {
		id := userMap["id"]
		if _, err := strconv.Atoi(id); err == nil {
			// Sets the Id to the pointed argument. So when the function is done the receiver will also have the id of user.
			(*user).Id, _ = strconv.ParseInt(id, 10, 64)
		}
		if checkPasswordHash((*user).password, userMap["password"]) {
			return true
		}
	}

	return false
}

func GetUserById(userId int64) *User {
	userListMap := database.SelectStatement(
		"SELECT * FROM User WHERE id = ? LIMIT 1",
		userId,
	)
	for _, userMap := range userListMap {
		userToReturn := new(User)
		userToReturn.Id = userId
		userToReturn.Username = userMap["username"]
		userToReturn.ProfilePhoto = userMap["profilePhoto"]
		return userToReturn
	}
	return nil
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
