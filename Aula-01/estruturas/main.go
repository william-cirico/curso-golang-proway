package main

import (
	"fmt"
	"time"
)

func main() {
	// Estruturas condicionais
	x := 20
	y := 18
	if x > y {
		fmt.Println("20 é maior que 18")
	} else if x < y {
		fmt.Println("20 é menor que 18")
	} else {
		fmt.Println("20 é igual a 18")
	}

	diaSemana := time.Now().Weekday()
	switch diaSemana {
	case 1:
		fmt.Println("Segunda-Feira")
	case 2:
		fmt.Println("Terça-Feira")
	case 3:
		fmt.Println("Quarta-Feira")
	case 4:
		fmt.Println("Quinta-Feira")
	case 5:
		fmt.Println("Sextou!")
	case 6, 7:
		fmt.Println("Final de semana")
	default:
		fmt.Println("Dia inválido")
	}

	// Estruturas de repetição
	for i := 0; i < 3; i++ {
		fmt.Println(fmt.Sprintf("Repetiu %d vez", i+1))
	}

	frutas := []string{"acerola", "bergamota", "carambola"}
	for i, fruta := range frutas {
		fmt.Println(fmt.Sprintf("A %d fruta é: %s", i+1, fruta))
	}

	// While
	i := 3
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
