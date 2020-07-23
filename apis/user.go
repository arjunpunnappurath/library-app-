package apis

import (
	"arjun/library/controllers"
	"arjun/library/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserApis struct{}

func (u *UserApis) ViewUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controller := controllers.Controller{}
		users := controller.ViewUsers(db)

		json.NewEncoder(w).Encode(users)
	}
}

func (u *UserApis) AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
		controller := controllers.Controller{}
		userID := controller.AddUser(db, user)

		json.NewEncoder(w).Encode(userID)
	}
}

func (u *UserApis) DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		controller := controllers.Controller{}
		rowsDeleted := controller.DeleteUser(db, id)
		json.NewEncoder(w).Encode(rowsDeleted)
	}
}