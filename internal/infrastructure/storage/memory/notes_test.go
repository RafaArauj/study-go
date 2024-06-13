package memory

import (
	"github.com/RafaArauj/study-go/internal/domains/entities"
	"testing"
)

func TestNotesController(t *testing.T) {

	c := NewNotesController()

	t.Run("CreateNote", func(t *testing.T) {
		err := c.CreateNote(nil, &entities.Note{
			ID:   "1",
			Text: "Exemplo",
		})

		if err != nil {
			t.FailNow()
		}

		err = c.CreateNote(nil, &entities.Note{
			ID:   "1",
			Text: "Exemplo",
		})
		if err == nil {
			t.FailNow()
		}

	})
}
