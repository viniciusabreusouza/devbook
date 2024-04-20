package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// CreatingUsers insert a user on DB
func CreatingUsers(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.ConnectDatabase()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.CreateNewUsersRepository(db)
	repository.Create(user)
}

// SearchingUsers search users on DB
func SearchingUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching Users"))
}

// SearchingUser search a users by ID on DB
func SearchingUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching a User by Id"))
}

// UpdatingUser update a users by ID on DB
func UpdatingUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User"))
}

// DeletingUser delete a users by ID on DB
func DeletingUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
