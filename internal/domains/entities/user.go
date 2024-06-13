package entities

import "errors"

var (
	ErrUserNotFound       = errors.New("Usuário não encontrado")
	ErrUserConflict       = errors.New("Usuário conflitante")
	ErrInvalidCredentials = errors.New("Credenciais invalidas")
)

type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}
