package main

import "fmt"

func somar(x, y int) int {
	return x + y
}

func trocarOrdem(x, y int) (int, int) {
	return y, x
}

// Função variádica
func mostraNomes(nomes ...string) {
	for _, nome := range nomes {
		fmt.Println(nome)
	}
}

func main() {
	// Defer adia a chamada da função
	defer fmt.Println("Sou executado depois :d")

	fmt.Println(somar(10, 2))
	fmt.Println(trocarOrdem(10, 2))
	mostraNomes("João", "Maria")
}
