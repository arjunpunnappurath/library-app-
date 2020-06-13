package main

import (
	"arjun/library/controllers"
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

	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.ViewBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.ViewBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
