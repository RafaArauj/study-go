package services

import (
	"context"
	"github.com/RafaArauj/study-go/internal/domains/entities"
	"github.com/RafaArauj/study-go/internal/domains/repositories"
	"time"
)

type NotesService struct {
	storage repositories.NotesStorage
	id      repositories.IDGen
}

func NewNotesService(storage repositories.NotesStorage, id repositories.IDGen) *NotesService {
	return &NotesService{storage: storage, id: id}
}

func (n *NotesService) CreateNote(ctx context.Context, text string) error {
	var (
		now = time.Now()
	)

	note := entities.Note{
		ID:        n.id.GenerateID(),
		Text:      text,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if note.Text == "" {
		return entities.ErrInvalidNote
	}

	return n.storage.CreateNote(ctx, &note)
}

func (n *NotesService) GetById(ctx context.Context, id string) (*entities.Note, error) {
	return n.storage.GetById(ctx, id)

}

func (n *NotesService) List(ctx context.Context) ([]*entities.Note, error) {
	return n.storage.List(ctx)
}

func (n *NotesService) DeleteById(ctx context.Context, id string) error {
	return n.storage.DeleteById(ctx, id)
}

func (n *NotesService) EditById(ctx context.Context, id, text string) error {
	if text == "" {
		return entities.ErrInvalidNote
	}
	note, err := n.storage.GetById(ctx, id)
	if err != nil {
		return err
	}
	note.Text = text
	note.UpdatedAt = time.Now()
	return n.storage.EditById(ctx, id, note)
}
