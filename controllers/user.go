package controllers

import (
	"arjun/library/models"
	"arjun/library/repo"
	"database/sql"
)

func (c *Controller) ViewUsers(db *sql.DB) []models.User {
	repo := repo.Repo{}
	users := repo.ViewUsers(db)
	return users
}

func (c *Controller) AddUser(db *sql.DB, user models.User) int {
	repo := repo.Repo{}
	userId := repo.AddUser(db, user)
	return userId
}

func (c *Controller) DeleteUser(db *sql.DB, id string) int64 {
	repo := repo.Repo{}
	rowsDeleted := repo.DeleteUser(db, id)
	return rowsDeleted
}
