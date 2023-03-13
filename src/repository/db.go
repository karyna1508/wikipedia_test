package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")
var sslmode = os.Getenv("SSLMODE")
var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

func CreateSchemaAndTable() error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err = db.Exec(`CREATE SCHEMA IF NOT EXIST wikipedia_project;`); err != nil {
		return err
	}
	if _, err = db.Exec(`CREATE TABLE IF NOT EXIST users(ID SERIAL PRIMARY KEY, TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP, USERNAME TEXT, CHAT_ID INT,MESSAGE TEXT, ANSWER TEXT);`); err != nil {
		return err
	}
	return nil
}

func CollectData(username string, chatid int64, message string, answer []string) error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()
	answ := strings.Join(answer, ",")

	data := `INSERT INTO users(username, chat_id, message, answer) VALUES ($1, $2, $3, $4)`

	if _, err = db.Exec(data, `@`, username, chatid, message, answ); err != nil {
		return err
	}
	return err
}

func GetNumberOfUsers() (int64, error) {
	var count int64

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return 0, nil
	}
	row := db.QueryRow("SELECT COUNT(DISTINCT username) FROM users;")
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
