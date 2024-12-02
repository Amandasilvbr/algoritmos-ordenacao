package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Function to generate and display the table
func generateTable(algorithms []struct {
	name   string
	sorter func([]int) (int, int)
}, sizes []int) {

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "Algoritmo\tTamanho do Array\tCaso\tTempo (ms)\tComparações\tTrocas")

	for _, algo := range algorithms {
		for _, size := range sizes {
			bestCase := generateSortedArray(size)
			duration, comparisons, swaps := runSortAlgorithm(algo.sorter, append([]int{}, bestCase...))
			fmt.Fprintf(writer, "%s\t%d\tMelhor Caso\t%d\t%d\t%d\n", algo.name, size, duration.Milliseconds(), comparisons, swaps)

			averageCase := generateRandomArray(size)
			duration, comparisons, swaps = runSortAlgorithm(algo.sorter, append([]int{}, averageCase...))
			fmt.Fprintf(writer, "%s\t%d\tCaso Médio\t%d\t%d\t%d\n", algo.name, size, duration.Milliseconds(), comparisons, swaps)

			worstCase := generateReverseSortedArray(size)
			duration, comparisons, swaps = runSortAlgorithm(algo.sorter, append([]int{}, worstCase...))
			fmt.Fprintf(writer, "%s\t%d\tPior Caso\t%d\t%d\t%d\n", algo.name, size, duration.Milliseconds(), comparisons, swaps)
		}
	}
}
