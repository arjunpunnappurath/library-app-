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
	rows := db.QueryRow("select * from books where id = $1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	return book
}

func (r *Repo) AddBook(db *sql.DB, book models.Book) int {

	log.Println("Entered the repo module...")
	var bookID int
	log.Println(book)

	err := db.QueryRow("insert into books(title,author,year) values ($1,$2,$3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)

	logFatal(err)

	log.Println("Exiting the repo module...")

	return bookID

}

func (r *Repo) UpdatesBook(db *sql.DB, book models.Book) int64 {

	result, err := db.Exec("update books set title = $1, author = $2, year = $3 where id = $4 RETURNING id",
		*&book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

func (r *Repo) DeletesBook(db *sql.DB, id string) int64 {

	result, err := db.Exec("delete from books where id = $1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
