package main

import (
	"fmt"
	"github.com/RafaArauj/study-go/internal/infrastructure/storage/memory"
	"log"
	"testing"
)

func TestPonteiros(t *testing.T) {
	numeroA := 10

	//utilizando o & egaos o endereço de memória do número A, ou seja, o ponteiro de número A
	ponteiroNumeroA := &numeroA

	numeroB := numeroA
	ponteiroNumeroB := &numeroB

	imprimeValores(numeroA, numeroB, ponteiroNumeroA, ponteiroNumeroB)

	numeroB += 10
	imprimeValores(numeroA, numeroB, ponteiroNumeroA, ponteiroNumeroB)
	*ponteiroNumeroB += 10
	imprimeValores(numeroA, numeroB, ponteiroNumeroA, ponteiroNumeroB)
}

func imprimeValores(nA int, nB int, pna *int, pnb *int) {
	fmt.Printf("Numero A: %d | Numero B: %d \nPonteiro Numero A: %v | Ponteiro Numero B: %v \n", nA, nB, pna, pnb)
}
func TestGeneratedId(t *testing.T) {
	log.Print(memory.generateId())
}
