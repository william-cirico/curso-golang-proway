package main

import "fmt"

func main() {
	// Ponteiros guardam no memória o endereço de uma variável
	valor := 10

	ponteiroValor := &valor
	fmt.Println(ponteiroValor)
	fmt.Println(fmt.Sprintf("O valor do ponteiro é: %d", *ponteiroValor))

	valorASerModificado := 5
	modificaPonteiro(&valorASerModificado)
	fmt.Println(valorASerModificado)
}

func modificaPonteiro(p *int) {
	*p = 20
}
