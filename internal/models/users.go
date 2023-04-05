package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Insert(name, email, password string) error {
	return nil
}

func (r *UserRepository) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (r *UserRepository) Exists(id int) (bool, error) {
	return false, nil
}
