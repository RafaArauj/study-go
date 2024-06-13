package repositories

import (
	"context"
	"github.com/RafaArauj/study-go/internal/domains/entities"
)

type NotesStorage interface {
	CreateNote(ctx context.Context, note *entities.Note) error
	GetById(ctx context.Context, id string) (*entities.Note, error)
	List(ctx context.Context) ([]*entities.Note, error)
	DeleteById(ctx context.Context, id string) error
	EditById(ctx context.Context, id string, note *entities.Note) error
}
type NotesService interface {
	CreateNote(ctx context.Context, text string) error
	GetById(ctx context.Context, id string) (*entities.Note, error)
	List(ctx context.Context) ([]*entities.Note, error)
	DeleteById(ctx context.Context, id string) error
	EditById(ctx context.Context, id, text string) error
}
