package main

import "github.com/pauloa.junior/mynotes/internal/models"

type templateData struct {
	Note  *models.Note
	Notes []*models.Note
}
