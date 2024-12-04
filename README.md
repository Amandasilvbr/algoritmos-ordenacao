# Algoritmos de Ordenação em Go

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Este repositório contém implementações de vários algoritmos de ordenação em Go, juntamente com a medição do desempenho, como tempo de execução, número de comparações e trocas. O código também gera um gráfico de comparação do tempo de execução de cada algoritmo com base no tamanho do array.

**O gerenciamento do projeto utiliza um Makefile para simplificar os comandos de compilação e execução.**

## Algoritmos Implementados

- **Bubble Sort**
- **Improved Bubble Sort**
- **Insertion Sort**
- **Selection Sort**
- **Merge Sort**
- **Quick Sort**
- **Heap Sort**
- **Cocktail Sort**
- **Pancake Sort**

## Funcionalidades

- **Contagem de Comparações e Trocas**: Cada algoritmo de ordenação conta o número de comparações e trocas realizadas durante a execução.
- **Geração de Tabela**: Uma tabela detalhada é gerada mostrando o tempo de execução, o número de comparações e o número de trocas para diferentes tamanhos de arrays.
- **Entradas Aleatórias**: O código gera arrays aleatórios para testar os algoritmos de ordenação.

## Como Executar

### Requisitos

- [Go 1.20+](https://golang.org/) instalado.

### Passos para Executar

1. Clone o repositório:
   ```bash
   git clone https://github.com/Amandasilvbr/algoritmos-ordenacao.git
   cd algoritmos-ordenacao/cmd

3. Compile e execute o código:

```bash
go run *.go
```

### Exemplo de Saída
O código exibirá uma tabela semelhante a esta no terminal, onde cada linha representa o desempenho de um algoritmo para um determinado tamanho de array:

```bash
Algoritmo               Tamanho do Array        Caso            Tempo (ms)      Comparações     Trocas
Bubble Sort             1000                    Melhor Caso     0               499500          0
Bubble Sort             1000                    Caso Médio      1               499500          246449
Bubble Sort             1000                    Pior Caso       1               499500          499500
Bubble Sort             10000                   Melhor Caso     48              49995000        0
Bubble Sort             10000                   Caso Médio      97              49995000        24996135
Bubble Sort             10000                   Pior Caso       70              49995000        49995000
```

### Descrição do Código
O código possui as seguintes funções principais:

- Algoritmos de Ordenação: Cada função de ordenação (como bubbleSort, mergeSort, etc.) recebe um array e retorna o número de comparações e trocas feitas durante a execução.
- Função runSortAlgorithm: Executa o algoritmo de ordenação e calcula o tempo de execução.
- Funções de Geração de Tabelas: A função generateTable gera a tabela com os dados de tempo, comparações e trocas.

### Tabela
A tabela é exibida no terminal com as seguintes colunas:

Algoritmo: O nome do algoritmo de ordenação.
- Tamanho do Array: O tamanho do array usado no teste.
- Tempo (ms): O tempo de execução do algoritmo, em milissegundos.
- Comparações: O número de comparações feitas durante a execução.
- Trocas: O número de trocas realizadas durante a execução.

