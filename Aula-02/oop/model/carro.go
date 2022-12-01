package model

import (
	"errors"
	"fmt"

	"github.com/william-cirico/curso-go-proway/Aula-02/oop/interfaces"
)

type Carro struct {
	Modelo           string
	Placa            string
	VelocidadeMaxima float64
	QuantidadePortas uint
	Ligado           bool
}

func NewCarro(modelo, placa string, velocidadeMaxima float64, quantidadePortas uint) interfaces.Veiculo {
	return &Carro{
		Modelo:           modelo,
		Placa:            placa,
		VelocidadeMaxima: velocidadeMaxima,
		QuantidadePortas: quantidadePortas,
		Ligado:           false,
	}
}

func (carro *Carro) Ligar() string {
	carro.Ligado = true

	return fmt.Sprintf("O carro %s ligou 🟢", carro.Modelo)
}

func (carro *Carro) Acelerar() (string, error) {
	if !carro.Ligado {
		return "", errors.New("o carro não está ligado ainda")
	}

	return fmt.Sprintf("O carro %s está acelerando 🚗💨", carro.Modelo), nil
}

func (carro *Carro) Desligar() (string, error) {
	if !carro.Ligado {
		return "", errors.New("o carro não está ligado ainda")
	}

	return fmt.Sprintf("O carro %s desligou 🔴", carro.Modelo), nil
}
