package rest

import (
	"github.com/RafaArauj/study-go/internal/domains/entities"
	"github.com/RafaArauj/study-go/internal/domains/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type UserRestController struct {
	svc repositories.UserService
}

func NewUserRestController(svc repositories.UserService) *UserRestController {
	return &UserRestController{svc: svc}
}

func (n *UserRestController) handleError(c *gin.Context, err error) {
	switch err {
	case entities.ErrInvalidCredentials:
		c.AbortWithError(http.StatusUnauthorized, err)
	case entities.ErrUserNotFound:
		c.AbortWithError(http.StatusNotFound, err)
	case entities.ErrUserConflict:
		c.AbortWithError(http.StatusConflict, err)

	}
}

func (u *UserRestController) CreateUser(c *gin.Context) {
	type CreateUserPayLoad struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
	}
	var payload CreateUserPayLoad
	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	err := u.svc.CreateUser(c, payload.Login, payload.Password, payload.FullName)
	if err != nil {
		u.handleError(c, err)
		return
	}
}
func (u *UserRestController) ListUser(c *gin.Context) {
	ul, err := u.svc.List(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, ul)
}
func (u *UserRestController) AuthenticationUser(c *gin.Context) {
	log.Print("Recebeu uma requisição")
	header := c.GetHeader("Authorization")
	split := strings.Split(header, " ")
	if len(split) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	split = strings.Split(split[1], ":")
	if len(split) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err := u.svc.ValidateUser(c, split[0], split[1])
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.Next()
}
