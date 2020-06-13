package main

import (
	"arjun/library/models"
	"arjun/library/driver"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book

var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db = driver.ConnectDB()
	router := mux.NewRouter()

	router.HandleFunc("/books", viewBooks).Methods("GET")
	router.HandleFunc("/books/{id}", viewBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func viewBooks(w http.ResponseWriter, r *http.Request) {
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

func viewBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Displays the specific book...")
	var book models.Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id = $1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
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

func updateBook(w http.ResponseWriter, r *http.Request) {
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

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Deletes the book...")
	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id = $1", params["id"])
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)

}
