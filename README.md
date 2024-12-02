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
- **Geração de Gráfico**: Um gráfico de dispersão é gerado para visualizar a comparação do tempo de execução dos algoritmos para diferentes tamanhos de arrays.
- **Entradas Aleatórias**: O código gera arrays aleatórios para testar os algoritmos de ordenação.

## Como Executar

### Requisitos

- [Go 1.20+](https://golang.org/) instalado.

### Passos para Executar

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/algoritmos-ordenacao.git
   cd algoritmos-ordenacao
   
2. Instale as dependências necessárias:

   ```bash
    go get -u gonum.org/v1/plot/...
    
3. Compile e execute o código:

**Para usar o Makefile, você precisa do make. Ele geralmente está pré-instalado no Linux e no macOS. No Windows, você pode usar o WSL (Windows Subsystem for Linux) ou uma alternativa como o Make for Windows.**

Compile e rode o programa:

```bash
make run
```
Para limpar os arquivos gerados:
```bash
make clean
```

### Exemplo de Saída
O código exibirá uma tabela semelhante a esta no terminal, onde cada linha representa o desempenho de um algoritmo para um determinado tamanho de array:

```bash
Algoritmo             Tamanho do Array    Tempo (ms)    Comparações    Trocas
Bubble Sort           1000                5             3500           1700
Improved Bubble Sort  1000                4             3000           1500
Insertion Sort        1000                3             2000           1200
Selection Sort        1000                6             4000           2000
Merge Sort            1000                2             1500           800
Quick Sort            1000                1             1200           400
Heap Sort             1000                2             1800           900
Cocktail Sort         1000                5             3500           1800
Pancake Sort          1000                6             4200           2200
```
Além disso, o gráfico de dispersão será gerado e salvo como uma imagem PNG chamada sorting_algorithms_time.png, mostrando o tempo de execução de cada algoritmo.

### Descrição do Código
O código possui as seguintes funções principais:

- Algoritmos de Ordenação: Cada função de ordenação (como bubbleSort, mergeSort, etc.) recebe um array e retorna o número de comparações e trocas feitas durante a execução.
- Função runSortAlgorithm: Executa o algoritmo de ordenação e calcula o tempo de execução.
- Funções de Geração de Gráficos e Tabelas: A função generateGraph gera o gráfico de comparação entre os algoritmos, enquanto a função generateTable gera a tabela com os dados de tempo, comparações e trocas.

### Gráfico
O gráfico é gerado usando o pacote gonum/plot, e o arquivo PNG gerado será salvo no diretório local com o nome sorting_algorithms_time.png. O gráfico exibe a comparação do tempo de execução dos algoritmos de ordenação para diferentes tamanhos de arrays.

### Tabela
A tabela é exibida no terminal com as seguintes colunas:

Algoritmo: O nome do algoritmo de ordenação.
- Tamanho do Array: O tamanho do array usado no teste.
- Tempo (ms): O tempo de execução do algoritmo, em milissegundos.
- Comparações: O número de comparações feitas durante a execução.
- Trocas: O número de trocas realizadas durante a execução.

