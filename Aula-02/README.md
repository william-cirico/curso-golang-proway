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
1) Defina uma struct Carro com os seguintes tipos inteiros:
bateria
consumoBateria
velocidade
distancia

Permita a criação de um carro através do método NewCarro que obtém a velocidade do carro em metros e a porcentagem de consumo de bateria

Defina outra struct chamada de Corrida com o campo de distancia do tipo inteiro. Permita a criação de uma corrida através do método NewCorrida

Implemente o método dirigir para a struct Carro que atualiza o número de metros do carro baseado na velocidade e reduz a bateria de acordo com consumoBateria.
Se não houver mais bateria o carro não vai se mover.

Cheque se o carro pode finalizar uma corrida:

Para finalizar a corrida, um carro precisa ser capaz de dirigir pela distância da corrida. Para isso, implemente a função PodeFinalizar que recebe um Carro e uma Corrida e retorna true se o carro consegue finalizar a corrida; senão, retorna false.

2) Uma empresa de pesquisas online solicitou o desenvolvimento de um software capaz de identificar qual tamanho de pizza apresenta o melhor custo beneficio. 
O software deverá receber diversos tamanhos de pizza e seus respectivos preços e ao final exibir um relatório informando em valores absolutos e relativos (percentual)  qual a diferença de preços entre as pizzas e deverá informar qual pizza tem  o melhor custo beneficio.

Entradas:  1. Nome comercial (broto, baby, pequena, média, grande, exagerada, gigante, etc), o tamanho da pizza (diâmetro em centímetros) e respectivo  preço. O software deverá aceitar tantas entradas quanto o usuário deseja comparar, desde que não haja tamanhos duplicados. 

Saída: relatório contendo todos os nomes e tamanhos de pizza ordenados do melhor para o pior custo benefício.  O relatório deverá informar o percentual  de diferença do preço de um tamanho para o outro. 

Exemplo de relatório: 
<table>
  <tr>
    <th>Nome</th>
    <th>Tamanho</th>
    <th>Preço</th>
    <th>R$ p/ cm2</th>
    <th>Diferença %</th>
  </tr>
  <tr>
    <td>Broto</td>
    <td>15cm</td>
    <td>R$ 25,00</td>
    <td>R$ 1,00</td>
    <td>Melhor Custo Benefício</td>
  </tr>
  <tr>
    <td>Pequena</td>
    <td>25cm</td>
    <td>R$ 35,00</td>
    <td>R$ 1,40</td>
    <td>+40%</td>
  </tr>
  <tr>
    <td>Média</td>
    <td>35cm</td>
    <td>R$ 40,00</td>
    <td>R$ 1,50</td>
    <td>+25%</td>
  </tr>
</table>

*desconsiderar os valores do exemplo (não são valores calculados)
