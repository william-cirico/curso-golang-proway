# Aula 03 - Pacote OS, Goroutines e Channels

## Goroutines
São métodos e funções que podem ser executados em concorrência, ou seja, utilizando novas threads.

**Thread**: É uma sequência de instruções que pode ser executada em pararelo com outras.
![image](https://user-images.githubusercontent.com/69127474/206001874-96116149-2b77-42fe-be83-1eb2099e087b.png)

## Channels
Channels são a forma de conseguirmos trafegar valores entre goroutines.
![image](https://user-images.githubusercontent.com/69127474/206019421-46438b67-2f4d-4833-ad68-bf45209174d1.png)

## Conteúdo abordado
- [ ] Crição de diretórios e arquivos
- [ ] Renomear diretórios
- [ ] Remover diretórios e arquivos
- [ ] Encontrar arquivos dentro de um diretório
- [ ] Goroutines
- [ ] Channels
- [ ] Waitgroup

## Exercícios

1) Crie uma função maioresArquivos(path string) que recebe o caminho para um diretório no sistema operacional e mostre na tela o nome e o tamanho dos 5 maiores arquivos daquela pasta.

2) Crie uma função que multiplique um vetor de números inteiros e uma função que divida um vetor de números inteiros. Faça as duas funções rodarem paralelamente através da utilização de gorountines e ao final mostre a soma dos resultados das funções.
