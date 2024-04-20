package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// CreateNewUsers create a user Repository
func CreateNewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insert a new user on DB
func (u users) Create(user models.User) (uint64, error) {
	return 0, nil
}
