package main

import (
	"fmt"
)

func main() {
	sizes := []int{1000, 10000, 100000}

	// Algoritmos a serem testados
	algorithms := []struct {
		name   string
		sorter func([]int) (int, int) // Função de ordenação que retorna comparações e trocas
	}{
		{"Bubble Sort", bubbleSort},
		{"Improved Bubble Sort", improvedBubbleSort},
		{"Insertion Sort", insertionSort},
		{"Selection Sort", selectionSort},
		{"Merge Sort", mergeSort},
		{"Quick Sort", quickSort},
		{"Heap Sort", heapSort},
		{"Cocktail Sort", cocktailSort},
		{"Pancake Sort", pancakeSort},
	}

	// Gerar gráfico
	generateGraph(algorithms, sizes)

	// Gerar tabela
	generateTable(algorithms, sizes)

	fmt.Println("\nTestes concluídos.")
}
