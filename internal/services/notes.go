package services

import (
	"github.com/RafaArauj/study-go/internal/domains/entities"
	"github.com/RafaArauj/study-go/internal/domains/repositories"
	"time"
)

type NotesService struct {
	storage repositories.NotesStorage
}

func (n *NotesService) CreateNote(note *entities.Note) error {
	return n.storage.CreateNote(note)
}

func (n *NotesService) GetById(id string) (*entities.Note, error) {
	return n.storage.GetById(id)

}

func (n *NotesService) List() ([]*entities.Note, error) {
	return n.storage.List()
}

func (n *NotesService) DeleteById(id string) error {
	return n.storage.DeleteById(id)
}

func (n *NotesService) EditById(id, text string, updatedAt time.Time) error {
	return n.storage.EditById(id, text, updatedAt)
}
