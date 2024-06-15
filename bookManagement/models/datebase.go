package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB(filepath string) {
	var err error
	db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	CreateTable()
}

func CreateTable() {
	createBooksTableSQL := `CREATE TABLE IF NOT EXISTS books (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
        "title" TEXT,
        "isbn" TEXT,
        "author" TEXT,
        "year" INTEGER
    );`

	_, err := db.Exec(createBooksTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
