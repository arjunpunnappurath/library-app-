package repo

import (
	"arjun/library/models"
	"database/sql"
)

func (r *Repo) ViewUsers(db *sql.DB) []models.User {
	var user models.User
	users := []models.User{}

	rows, err := db.Query("select * from users")
	logFatal(err)

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		logFatal(err)

		users = append(users, user)
	}
	return users
}

func (r *Repo) AddUser(db *sql.DB, user models.User) int {
	var userId int

	err := db.QueryRow("insert into users (username,password)values ($1,$2)RETURNING id;",
		user.Username, user.Password).Scan(&userId)

	logFatal(err)
	return userId
}

func (r *Repo) DeleteUser(db *sql.DB, id string) int64 {
	result, err := db.Exec("delete from users where id = $1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
