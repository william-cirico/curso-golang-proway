package main

import (
	"fmt"
	"strconv"
)

// Struct - Esse sim é o objeto (Pode ter métodos)
type Pessoa struct {
	Nome   string
	Idade  int
	Altura float64
	Peso   float64
}

// Método de uma Pessoa
func (pessoa Pessoa) Andar() string {
	return "A pessoa " + pessoa.Nome + " está andando."
}

func main() {
	// Tipos de dados
	var numeroInteiro int = 10
	fmt.Println(numeroInteiro)

	var inteiroNaoAssinado uint = 120
	fmt.Println(inteiroNaoAssinado)

	var numeroDecimal float64 = 10.2
	fmt.Println(numeroDecimal)

	var palavra string = "Hoje é terça-feira"
	fmt.Println(palavra)

	var booleano bool = false
	fmt.Println(booleano)

	// Matrizes
	// := é utilizado no momento de criação da variável
	cores := [3]string{"amarelo", "branco", "coral"}
	fmt.Println(cores)

	// Slice - Slice é um vetor que pode aumentar de tamanho
	frutas := []string{"abacate", "banana", "caju"}
	fmt.Println(cap(frutas))
	frutas = append(frutas, "damasco")
	fmt.Println(cap(frutas))
	frutas = append(frutas, "embaúba", "figo")
	fmt.Println(cap(frutas))
	frutas = append(frutas, "embaúba")
	fmt.Println(cap(frutas))
	fmt.Println(len(frutas))
	fmt.Println(frutas)

	// Podemos fatiar a slice - [inicio:fim + 1]
	fmt.Println(frutas[1:])
	fmt.Println(frutas[2:5])
	fmt.Println(frutas[:5])
	fmt.Println(frutas[0:5])

	novoVetor := []string{}
	fmt.Println(cap(novoVetor))

	valorDecimal := 10.2
	fmt.Println(valorDecimal)

	// Conversão entre tipos
	var valorDecimalEmInteiro int
	valorDecimalEmInteiro = int(valorDecimal)
	fmt.Println(valorDecimalEmInteiro)

	// Inteiro para String
	valorInteiroEmString := strconv.Itoa(valorDecimalEmInteiro)
	fmt.Println(valorInteiroEmString)

	valorEmStringTeste := "10.2a"

	// String para inteiro
	valorEmStringParaInteiro, err := strconv.Atoi(valorEmStringTeste)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(valorEmStringParaInteiro)

	// Map - É o objeto do JavaScript
	// { nome: "William", idade: 10 }
	// map[tipo_chave]tipo_valor
	scooby := map[string]string{
		"name":    "Scooby",
		"reino":   "Animalia",
		"genero":  "Canis",
		"especie": "Canis lupus",
	}

	fmt.Println(scooby)
	fmt.Println(scooby["name"])

	// Constructor -> Criar um objeto daquela classe
	pessoa := Pessoa{
		Nome:  "William",
		Idade: 22,
	}

	fmt.Println(pessoa.Nome)
	fmt.Println(pessoa.Altura)
	fmt.Println(pessoa.Andar())

	outraPessoa := Pessoa{
		Nome: "Maria",
	}

	fmt.Println(outraPessoa.Andar())
}
