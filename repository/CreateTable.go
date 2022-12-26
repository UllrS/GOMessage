package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDataBaseTable() {
	fmt.Println("create table")
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY, body TEXT, sender VARCHAR(64), recipient VARCHAR(64), attachedpath VARCHAR(128), attachedname VARCHAR(128), time INTEGER)"); err != nil {
		fmt.Println(err.Error())
	}
}
