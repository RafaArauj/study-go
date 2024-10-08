package generators

import (
	"crypto/sha1"
	"fmt"
)

type HashGen struct {
}

func NewHashGen() *HashGen {
	return &HashGen{}
}

func (h *HashGen) GenerateFromString(s string) string {
	sum := sha1.Sum([]byte(s))
	return fmt.Sprintf("%x", sum)
}
