package generators

import "math/rand/v2"

func NewAlphabetIDGen() *AlphabetIDGen {
	return &AlphabetIDGen{}
}

type AlphabetIDGen struct {
}

func (a *AlphabetIDGen) GenerateID() string {
	const (
		alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
		size  = 8
	)
	out := ""
	for len(out) < size {
		next := rand.N(len(alpha) - 1)
		out += string(alpha[next])
	}
	return out
}
