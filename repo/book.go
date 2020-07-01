package repo

import (
	"arjun/library/models"
	"database/sql"
	"log"
)

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

type Repo struct{}

func (r *Repo) ViewAllBooks(db *sql.DB) []models.Book {
	var book models.Book
	books := []models.Book{}

	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	return books
}

func (r *Repo) ViewSingleBook(db *sql.DB, id string) models.Book {
	var book models.Book

	// rows := db.QueryRow("select * from books where id = $1", params["id"])

	// err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	// logFatal(err)
	rows := db.QueryRow("select * from books where id = $1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	return book
}
