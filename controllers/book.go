package controllers

import (
	"arjun/library/models"
	"arjun/library/repo"
	"database/sql"
	"fmt"
	"log"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) ViewBooks(db *sql.DB) []models.Book {
	repo := repo.Repo{}
	books := repo.ViewAllBooks(db)
	return books
}

func (c Controller) ViewBook(db *sql.DB, id string) models.Book {
	fmt.Println(id)
	repo := repo.Repo{}
	book := repo.ViewSingleBook(db, id)
	return book
}

func (c Controller) AddBook(book models.Book, db *sql.DB) int {
	repo := repo.Repo{}
	bookID := repo.AddBook(db, book)

	return bookID
}

func (c Controller) UpdateBook(book models.Book, db *sql.DB) int64 {
	repo := repo.Repo{}
	rowsUpdated := repo.UpdatesBook(db, book)

	return rowsUpdated
}

func (c Controller) DeleteBook(db *sql.DB, id string) int64 {
	repo := repo.Repo{}
	rowsDeleted := repo.DeletesBook(db, id)
	return rowsDeleted

}
