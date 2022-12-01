package main

import (
	"fmt"

	"github.com/william-cirico/curso-go-proway/Aula-02/oop/model"
)

func main() {
	joao := model.NewPessoa("João", 20, 1.70, 80)

	fmt.Println(joao)
	fmt.Println(joao.Altura())
	fmt.Println(joao.Andar())
	fmt.Println(joao.Correr())

	kwid := model.NewCarro("Jeep Renegade", "MXN-7233", 210, 4)

	fmt.Println(kwid.Ligar())

	acelerou, errAcelerou := kwid.Acelerar()
	if errAcelerou != nil {
		fmt.Println(errAcelerou.Error())
	}

	fmt.Println(acelerou)

	desligou, errDesligou := kwid.Desligar()
	if errDesligou != nil {
		fmt.Println(errDesligou.Error())
	}

	fmt.Println(desligou)

	pessoaFisica := model.PessoaFisica{
		Pessoa: joao,
		RG:     "8271822",
		CPF:    "82781292891",
	}

	fmt.Println(pessoaFisica.Pessoa.Nome)
}
