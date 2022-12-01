package model

import (
	"fmt"
	"math/rand"
)

type Pessoa struct {
	ID     uint
	Nome   string
	Idade  uint
	altura float64
	Peso   float64
}

func NewPessoa(nome string, idade uint, altura, peso float64) Pessoa {
	return Pessoa{
		ID:     uint(rand.Intn(100-1) + 1),
		Nome:   nome,
		Idade:  idade,
		altura: altura,
		Peso:   peso,
	}
}

func (p Pessoa) Altura() string {
	return fmt.Sprintf("%.2f metros", p.altura)
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Nome: %s | Idade: %d anos | Altura: %.2f | Peso: %.2f", p.Nome, p.Idade, p.altura, p.Peso)
}

func (p Pessoa) Andar() string {
	return fmt.Sprintf("%s está andando", p.Nome)
}

func (p Pessoa) Correr() string {
	return fmt.Sprintf("%s está correndo", p.Nome)
}
