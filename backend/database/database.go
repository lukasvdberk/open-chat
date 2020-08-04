package database

import (
	"database/sql"
	"fmt"
	"os"
)

import _ "github.com/go-sql-driver/mysql"

// TODO add proper error handling
func GetSqlConnection() *sql.DB {
	// Gets connection info from environment variables
	dbHost := "tcp(db-mysql:3306)"
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	databaseName := os.Getenv("MYSQL_DATABASE")

	connectionString := user + ":" + password + "@" + dbHost + "/" + databaseName

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err)
	}

	// To test whether the connection worked
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}
