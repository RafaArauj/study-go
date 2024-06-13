package repositories

import (
	"github.com/RafaArauj/study-go/internal/domains/entities"
	"time"
)

type NotesStorage interface {
	CreateNote(note *entities.Note) error
	GetById(id string) (*entities.Note, error)
	List() ([]*entities.Note, error)
	DeleteById(id string) error
	EditById(id, text string, updatedAt time.Time) error
}
