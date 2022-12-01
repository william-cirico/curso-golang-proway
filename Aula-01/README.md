# Aula 01 - Instalação e Fundamentos do GO

Go é uma linguagem de programação compilada que foi construída com foco em perfomance e programação concorrente.

**Compilação** é o ato de transcrever um código escrito na linguagem de alto-nível (Go) para o executável em binário.

**Programação concorrente** é a capacidade de executar várias ações ao mesmo tempo através do uso de Threads:
![image](https://user-images.githubusercontent.com/69127474/205124852-df7bfce1-301e-40a2-9e90-1488159b58a1.png)

## Instalação do Go

Para instalar o Go no Windows basta acessar o site https://go.dev/dl/,fazer o download do instalador para a sua máquina e seguir os passos de instalação:
![image](https://user-images.githubusercontent.com/69127474/205125866-573dff37-9c3e-4420-bea2-b935e1431845.png)

## Assuntos abordados na aula
**Tipos de dados:**
- [ ] int
- [ ] uint
- [ ] float64
- [ ] string
- [ ] bool
- [ ] matrizes
- [ ] slices
- [ ] maps
- [ ] struct

**Conversao entre tipos**
- [ ] int para float64
- [ ] float64 para int
- [ ] int para string
- [ ] string para int

**Tratamento de erros**
- [ ] Funções que retornam valor e erro

**Entrada de dados**

**Estruturas Condicionais e de Repetição**

**Funções**

**Ponteiros**

## Exercícios
1) Faça uma função que recebe um valor inteiro e verifica se o valor é positivo. 
A função deve retornar um valor booleano. Receba um valor do usuário e mostre o resultado
da função.

2) Crie uma função minimo(a, b) que retorna o menor valor entre dois números a e b.
Receba os dois valores do usuário e mostre o resultado da função.

3) Crie uma função potencia(a, b) que retorna a elevado a b. Receba os dois valores do usuário e 
mostre o resultado da função.

4) Crie uma script que receba o valor de 3 notas, utilize uma função para calcular a média e
mostre ela na tela.

5) Faça um script que receba a média de um aluno e com a ajuda de uma função que recebe 
a média por parâmetro retorne o conceito da média conforme a tabela abaixo:
<table>
  <tr>
    <th>Média</th>
    <th>Conceito</th>
  </tr>
  <tr>
    <td>0.0 à 4.9</td>
    <td>D</td>
  </tr>
   <tr>
    <td>5.0 à 6.9</td>
    <td>C</td>
  </tr>
   <tr>
    <td>7.0 à 8.9</td>
    <td>B</td>
  </tr>
   <tr>
    <td>9.0 à 10.0</td>
    <td>A</td>
  </tr>
</table>

6) Faça uma função que recebe, por parâmetro, a altura e o sexo de uma pessoa e 
retorna o seu peso ideal. Receba a altura e o sexo do usuário e mostre o resultado da função.
Utilize as fórmulas:
Mulher: 62.1 x altura - 44.7
Homem: 72.7 x altura - 58

7) Faça um procedimento que lê 50 valores inteiros e retorna o maior e o menor deles. 
Utilize o pacote math/rand para gerar valores aleatórios entre 1-1000.

8) Faça uma função que recebe a idade de uma pessoa em anos (365 dias), meses e dias e retorna essa idade expressa em dias.

9) Faça uma função que recebe, por parâmetro, um valor inteiro e positivo e retorna o número de divisores 
desse valor.

10) Faça uma função que receba um valor inteiro e positivo e calcule o seu fatorial.

11) Faça uma função que verifique se um valor é perfeito ou não. Um valor é dito perfeito quando ele é igual a 
soma dos seus divisores excetuando ele próprio. (Ex: 6 é perfeito, 6 = 1 + 2 + 3, que são seus divisores). 
A função deve retornar um valor booleano. Receba um valor do usuário e mostre o resultado da função.

12) Faça um procedimento que recebe 50 valores inteiros por parâmetro e retorna-os ordenados em ordem decrescente.
Gere 50 valores aleatórios com o pacote math/rand, adicione-os em um vetor e mostre eles na tela antes da
ordenação e após a ordenação.
