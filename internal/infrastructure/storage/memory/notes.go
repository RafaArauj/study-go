package memory

import (
	"github.com/RafaArauj/study-go/internal/domains/entities"
	"time"
)

/*
NotesMemoryStorage

responsavel por:
- Criar Notas
- Armazenar as notas
- Listar as notas
- Buscar nota especifica
*/

type NotesMemoryStorage struct {
	//Notes armazena a nota usando id como identificador
	Notes map[string]*entities.Note
	// nesse contexto o * é indicação  que estamos usando um ponteiro de noite
	//ponteiro = endereço de memória
	//Notes list armazena o id em ordem de inserção
	NotesList []string
}

func NewNotesController() *NotesMemoryStorage {
	return &NotesMemoryStorage{
		Notes:     make(map[string]*entities.Note),
		NotesList: make([]string, 0, 16),
	}
}

func (n *NotesMemoryStorage) CreateNote(note *entities.Note) error {
	_, exists := n.Notes[note.ID]
	if exists {
		return entities.ErrNoteConflict
	}
	n.Notes[note.ID] = note
	n.NotesList = append(n.NotesList, note.ID)
	return nil
}
func (n *NotesMemoryStorage) GetById(id string) (*entities.Note, error) {
	note, exists := n.Notes[id]
	if !exists {
		return nil, entities.ErrNoteNotFound
	}

	return note, nil
}
func (n *NotesMemoryStorage) List() ([]*entities.Note, error) {
	result := make([]*entities.Note, len(n.NotesList))
	for i, v := range n.NotesList {
		result[i] = n.Notes[v]
	}
	return result, nil
}
func (n *NotesMemoryStorage) DeleteById(id string) error {
	_, exists := n.Notes[id]
	if !exists {
		return entities.ErrNoteNotFound
	}
	delete(n.Notes, id)
	as := make([]string, 0, len(n.Notes))
	for _, v := range n.NotesList {
		if v == id {
			continue
		}
		as = append(as, v)
	}
	n.NotesList = as
	return nil
}

func (n *NotesMemoryStorage) EditById(id, text string, updatedAt time.Time) error {
	_, exists := n.Notes[id]
	if !exists {
		return entities.ErrNoteNotFound
	}
	n.Notes[id].UpdatedAt = updatedAt
	n.Notes[id].Text = text
	return nil
}
