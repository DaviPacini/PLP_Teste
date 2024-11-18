package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Heroi representa um herói com nome e força
type Heroi struct {
	nome  string
	forca int
}

// SimuladorBatalha simula uma batalha entre dois heróis
type SimuladorBatalha struct{}

// simularBatalha realiza a simulação da batalha entre dois heróis
func (sb SimuladorBatalha) simularBatalha(heroi1, heroi2 Heroi) {
	fmt.Printf("Iniciando batalha entre %s e %s!\n", heroi1.nome, heroi2.nome)

	// Força final de cada herói será a força inicial mais todos os outros fatores
	forcaFinalHeroi1 := heroi1.forca
	forcaFinalHeroi2 := heroi2.forca

	if rand.Intn(10) < 2 {
		forcaFinalHeroi1 += 15
	}
	if rand.Intn(10) < 2 {
		forcaFinalHeroi2 += 15
	}

	// Cálculo da probabilidade de vitória
	forcaTotal := forcaFinalHeroi1 + forcaFinalHeroi2
	chanceHeroi1 := float64(forcaFinalHeroi1) / float64(forcaTotal)

	fmt.Printf("%s força final: %d\n", heroi1.nome, forcaFinalHeroi1)
	fmt.Printf("%s força final: %d\n", heroi2.nome, forcaFinalHeroi2)

	// Determinação do vencedor
	if rand.Float64() < chanceHeroi1 {
		fmt.Printf("%s venceu a batalha!\n", heroi1.nome)
	} else {
		fmt.Printf("%s venceu a batalha!\n", heroi2.nome)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	heroi1 := Heroi{nome: "Homelander", forca: 100}
	heroi2 := Heroi{nome: "Homelander reverso", forca: 50}

	simulador := SimuladorBatalha{}
	simulador.simularBatalha(heroi1, heroi2)
}
