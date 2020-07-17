package apis

import (
	"arjun/library/controllers"
	"arjun/library/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

type BookApi struct{}

func (b *BookApi) ViewAllBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Displays the books...")
		controller := controllers.Controller{}
		books := controller.ViewBooks(db)
		json.NewEncoder(w).Encode(books)
	}
}

func (b *BookApi) ViewABook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Displays the books...")
		params := mux.Vars(r)
		id := params["id"]
		controller := controllers.Controller{}
		book := controller.ViewBook(db, id)
		json.NewEncoder(w).Encode(book)
	}
}

func (b *BookApi) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("adds a book...")
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)
		controller := controllers.Controller{}
		bookID := controller.AddBook(book, db)
		json.NewEncoder(w).Encode(bookID)
	}
}

func (b *BookApi) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Updates the book...")
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		controller := controllers.Controller{}
		rowsUpdated := controller.UpdateBook(book, db)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (b *BookApi) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Deletes the book...")
		params := mux.Vars(r)
		id := params["id"]

		controller := controllers.Controller{}
		rowsDeleted := controller.DeleteBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)

	}
}
