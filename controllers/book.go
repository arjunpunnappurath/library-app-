package controllers

import (
	"arjun/library/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) ViewBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Displays the books...")

		var book models.Book
		books := []models.Book{}

		rows, err := db.Query("select * from books")
		logFatal(err)

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			logFatal(err)
		}

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) ViewBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Displays the specific book...")
		var book models.Book
		params := mux.Vars(r)

		rows := db.QueryRow("select * from books where id = $1", params["id"])

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("adds a book...")
		var book models.Book
		var bookID int
		json.NewDecoder(r.Body).Decode(&book)

		log.Println(book)

		err := db.QueryRow("insert into books(title,author,year) values ($1,$2,$3)RETURNING id;",
			book.Title, book.Author, book.Year).Scan(&bookID)

		logFatal(err)

		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Updates the book...")
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		result, err := db.Exec("update books set title =$1, author = $2, year = $3 where id =$4 RETURNING id",
			&book.Title, &book.Author, &book.Year, &book.ID)
		logFatal(err)

		rowsUpdated, err := result.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Deletes the book...")
		params := mux.Vars(r)

		result, err := db.Exec("delete from books where id = $1", params["id"])
		logFatal(err)

		rowsDeleted, err := result.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsDeleted)

	}
}
