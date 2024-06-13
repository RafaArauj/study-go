package entities

import (
	"errors"
	"time"
)

var (
	ErrNoteConflict = errors.New("Já existe uma nota com esse ID")
	ErrNoteNotFound = errors.New("Nota não encontrada")
	ErrInvalidNote  = errors.New("Nota invalida")
)

type Note struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
