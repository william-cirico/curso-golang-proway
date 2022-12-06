# Aula 02 - Orientação a Objetos com Go

Em Go conseguimos aplicar os princípios da OOP utilizando Structs.

## Conteúdo abordado
- [ ] packages
- [ ] construtores
- [ ] métodos
- [ ] método String()
- [ ] herança
- [ ] interfaces

## Exercícios
1) Seguindo os princípios de encapsulamento crie a classe Funcionario:        
  - Atributos: matricula, nome e salario.
 Regras para validações dos atributos:
  - matrícula: Deve ser um número inteiro positivo.
  - nome: deve conter de 4 a 50 caracteres.
  - salario: deve ser um valor não inferior a R$ 465,00.

2) Escreva a classe Aluno. Cada objeto dessa classe deve guardar os seguintes dados:
  - matrícula,
  - nome
  - 2 notas de prova
  - 1 nota de trabalho
Essa classe deve ter 2 métodos:
  - media: calcula a média final do aluno (cada prova tem peso 3 e o trabalho tem peso 2)
  - final: calcula quanto o aluno precisa para a prova final (retorna zero se ele não for para a final)

3) https://exercism.org/tracks/go/exercises/need-for-speed
4) https://exercism.org/tracks/go/exercises/elons-toys
