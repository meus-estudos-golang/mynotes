package mocks

import (
	"time"

	"github.com/pauloa.junior/mynotes/internal/models"
)

var mockNote = &models.Note{
	ID:      1,
	Title:   "Supermercado",
	Content: "Comprar pão",
	Created: time.Now(),
}

type FakeNoteRepository struct{}

func (r *FakeNoteRepository) Insert(title, content string) (int, error) {
	return 2, nil
}

func (r *FakeNoteRepository) GetById(id int) (*models.Note, error) {
	switch id {
	case 1:
		return mockNote, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (r *FakeNoteRepository) GetAll() ([]*models.Note, error) {
	return []*models.Note{mockNote}, nil
}
