package models

import (
	"database/sql"
	"errors"
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
	stmt := "SELECT id, title, content, created FROM notes WHERE id = ?"

	n := &Note{}
	err := r.DB.QueryRow(stmt, id).Scan(&n.ID, &n.Title, &n.Content, &n.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return n, nil
}

func (r *NoteRepository) GetAll() ([]*Note, error) {
	return nil, nil
}
