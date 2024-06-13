package rest

import (
	"github.com/RafaArauj/study-go/internal/domains/entities"
	"github.com/RafaArauj/study-go/internal/domains/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NoteRestController struct {
	svc repositories.NotesService
}

func NewNoteRestController(svc repositories.NotesService) *NoteRestController {
	return &NoteRestController{svc: svc}
}

func (n *NoteRestController) handleError(c *gin.Context, err error) {
	switch err {
	case entities.ErrInvalidNote:
		c.AbortWithError(http.StatusBadRequest, err)
	case entities.ErrNoteNotFound:
		c.AbortWithError(http.StatusNotFound, err)
	case entities.ErrNoteConflict:
		c.AbortWithError(http.StatusConflict, err)
	}
}

func (n *NoteRestController) CreateNote(c *gin.Context) {
	type CreateNotePayLoad struct {
		Text string `json:"text"`
	}

	var payload CreateNotePayLoad
	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := n.svc.CreateNote(c, payload.Text)
	if err != nil {
		n.handleError(c, err)
		return
	}
}
func (n *NoteRestController) ListNotes(c *gin.Context) {
	nl, err := n.svc.List(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nl)
}
func (n *NoteRestController) EditNotes(c *gin.Context) {
	type EditNotePayload struct {
		Text string `json:"text"`
	}
	var (
		payload EditNotePayload
	)
	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := n.svc.EditById(c, c.Param("id"), payload.Text)
	if err != nil {
		n.handleError(c, err)
		return
	}
	c.Status(http.StatusOK)

}
func (n *NoteRestController) DeleteNotes(c *gin.Context) {
	err := n.svc.DeleteById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)

}
