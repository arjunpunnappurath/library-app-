package main

import (
	"arjun/library/apis"
	"arjun/library/driver"
	"arjun/library/models"
	"database/sql"
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

	api := apis.BookApi{}
	userapi := apis.UserApis{}

	router.HandleFunc("/books", api.ViewAllBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", api.ViewABook(db)).Methods("GET")
	router.HandleFunc("/books", api.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", api.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", api.DeleteBook(db)).Methods("DELETE")

	router.HandleFunc("/users", userapi.ViewUsers(db)).Methods("GET")
	router.HandleFunc("/users", userapi.AddUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", userapi.DeleteUser(db)).Methods("DELETE")
	router.HandleFunc("/login", userapi.Login(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
