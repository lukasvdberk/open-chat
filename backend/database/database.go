package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

// This file contains various functions for retrieving data from the database and converting values.
func GetSqlConnection() *sql.DB {
	// Gets connection info from environment variables
	dbHost := "tcp(db-mysql:3306)"
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	databaseName := os.Getenv("MYSQL_DATABASE")

	connectionString := user + ":" + password + "@" + dbHost + "/" + databaseName

	db, err := sql.Open("mysql", connectionString+"?parseTime=true")

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
func SelectStatement(prepareStatement string, arguments ...interface{}) []map[string]string {
	// performs a query with a preparestatemnt
	// the arguments is the values you want the ? to replace with.
	// @return A list of maps where each map is a row that come back from the select statement

	db := GetSqlConnection()
	stmtOut, err := db.Prepare(prepareStatement)

	if err != nil {
		return nil
	}

	rows, err := stmtOut.Query(arguments...)

	if err != nil {
		return nil
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// This will be returned
	var rowsMapped []map[string]string

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil
		}
		var value string

		var rowMap map[string]string
		rowMap = make(map[string]string)

		for i, col := range values {
			if col != nil {
				column := columns[i]
				value = string(col)
				rowMap[column] = value
			}
		}
		rowsMapped = append(rowsMapped, rowMap)
	}

	defer db.Close()
	return rowsMapped
}

func InsertStatement(insertStatement string, arguments ...interface{}) int64 {
	// @return will return the id of the created record. If it fails will return -1

	db := GetSqlConnection()
	stmtIns, err := db.Prepare(insertStatement)

	if err != nil {
		return -1
	}

	res, err := stmtIns.Exec(arguments...)
	if err != nil {
		return -1
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1
	}

	defer db.Close()
	return id
}

func DateStringToTimeStamp(dateStr string) int64 {
	layout := "2006-01-02T15:04:05Z07:00"
	convertedTime, err := time.Parse(layout, dateStr)

	if err != nil {
		return -1
	}

	return convertedTime.Unix()
}
