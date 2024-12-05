package main

import (
	"fmt"
)

func main() {
	sizes := []int{1000, 5000, 10000, 50000, 100000}

	algorithms := []struct {
		name   string
		sorter func([]int) (int, int)
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

	generateTable(algorithms, sizes)

	fmt.Println("\nTestes concluídos.")
}
