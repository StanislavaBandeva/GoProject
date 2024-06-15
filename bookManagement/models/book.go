package models

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	ISBN   string `json:"isbn"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func createTable() {
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

func GetBooks() ([]Book, error) {
	rows, err := db.Query("SELECT id, title, isbn, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBook(id int) (Book, error) {
	var book Book
	row := db.QueryRow("SELECT id, title, isbn, author, year FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Year)
	if err != nil {
		return book, err
	}
	return book, nil
}

func CreateBook(book Book) error {
	_, err := db.Exec("INSERT INTO books (title, isbn, author, year) VALUES (?, ?, ?, ?)", book.Title, book.ISBN, book.Author, book.Year)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(id int, book Book) error {
	_, err := db.Exec("UPDATE books SET title = ?, isbn = ?, author = ?, year = ? WHERE id = ?", book.Title, book.ISBN, book.Author, book.Year, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(id int) error {
	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
