package controllers

import (
	"arjun/library/models"
	"arjun/library/repo"
	"database/sql"
	"encoding/json"
	"fmt"
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
		repo := repo.Repo{}
		books := repo.ViewAllBooks(db)
		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) ViewBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Displays the specific book...")
		//var book models.Book
		params := mux.Vars(r)
		id := params["id"]

		fmt.Println(id)

		repo := repo.Repo{}
		book := repo.ViewSingleBook(db, id)
		// rows := db.QueryRow("select * from books where id = $1", params["id"])

		// err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		// logFatal(err)

		// json.NewEncoder(w).Encode(book)
		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("adds a book...")
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)

		repo := repo.Repo{}
		bookID := repo.AddBook(db, book)

		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Updates the book...")
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		repo := repo.Repo{}
		rowsUpdated := repo.UpdatesBook(db, book)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Deletes the book...")
		params := mux.Vars(r)
		id := params["id"]

		repo := repo.Repo{}
		rowsDeleted := repo.DeletesBook(db, id)
		json.NewEncoder(w).Encode(rowsDeleted)

	}
}
