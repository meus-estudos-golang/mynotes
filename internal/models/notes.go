package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type NoteRepository struct {
	DB *sql.DB
}

func (r *NoteRepository) Insert(title, content string) (int, error) {
	stmt := "INSERT INTO notes (title, content, created) VALUES (?, ?, UTC_TIMESTAMP())"

	result, err := r.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *NoteRepository) GetById(id int) (*Note, error) {
	return nil, nil
}

func (r *NoteRepository) GetAll() ([]*Note, error) {
	return nil, nil
}
