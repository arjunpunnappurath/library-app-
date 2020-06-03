package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

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

	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	log.Println(pgURL)
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

	var book Book
	books := []Book{}

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
	var book Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id = $1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("adds a book...")
	var book Book
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
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Deletes the book...")
}
