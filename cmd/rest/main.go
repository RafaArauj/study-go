package main

import (
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			log.Printf("OPTIONS request")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {

	log.Print("Iniciando o servidor")
	r := gin.Default()
	r.Use(CORSMiddleware())
	noteControler := NewNotesController()

	api := r.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Qualquer coisa que estou botando aqui")
	})
	notes := api.Group("/notes").Use(CORSMiddleware())
	//o api. cria a estrutura padrão de "api/notes"
	notes.POST("/", func(c *gin.Context) {
		var (
			note Note
			now  = time.Now()
		)
		if err := c.BindJSON(&note); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if note.Text == "" {
			c.AbortWithError(http.StatusBadRequest, ErrInvalidNote)
			return
		}
		note.ID = generateId()
		note.CreatedAt = now
		note.UpdatedAt = now
		err := noteControler.CreateNote(&note)
		if err == ErrNoteConflict {
			c.AbortWithError(http.StatusConflict, err)
			return
		}
	})
	notes.GET("/", func(c *gin.Context) {
		n, err := noteControler.List()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, n)

	})
	notes.DELETE("/:id", func(c *gin.Context) {
		err := noteControler.DeleteById(c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Status(http.StatusOK)

	})
	notes.PATCH("/:id", func(c *gin.Context) {
		var (
			note Note
			now  = time.Now()
		)
		if err := c.BindJSON(&note); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if note.Text == "" {
			c.AbortWithError(http.StatusBadRequest, ErrInvalidNote)
			return
		}

		err := noteControler.EditById(c.Param("id"), note.Text, now)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Status(http.StatusOK)

	})
	r.Run()

}
