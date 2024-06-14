package main

import (
	"github.com/RafaArauj/study-go/internal/controllers/rest"
	"github.com/RafaArauj/study-go/internal/infrastructure/generators"
	"github.com/RafaArauj/study-go/internal/infrastructure/storage/memory"
	"github.com/RafaArauj/study-go/internal/services"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	noteControler := rest.NewNoteRestController(services.NewNotesService(memory.NewNotesController(), generators.NewAlphabetIDGen()))
	userControler := rest.NewUserRestController(services.NewUserService(memory.NewUserStorage(), generators.NewAlphabetIDGen(), generators.NewHashGen()))

	api := r.Group("/api")
	users := api.Group("/users").Use(CORSMiddleware())
	notes := api.Group("/notes").Use(CORSMiddleware(), userControler.AuthenticationUser)
	//o api. cria a estrutura padrão de "api/notes"
	users.GET("/", userControler.ListUser)
	users.POST("/", userControler.CreateUser)
	notes.POST("/", noteControler.CreateNote)
	notes.GET("/", noteControler.ListNotes)
	notes.DELETE("/:id", noteControler.DeleteNotes)
	notes.PATCH("/:id", noteControler.EditNotes)
	r.Run()

}
