package main

import (
	"math/rand"
	"time"
)

// Melhor Caso: Array ordenado em ordem crescente
func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

// Caso Médio: Array com números aleatórios
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return arr
}

// Pior Caso: Array ordenado em ordem decrescente
func generateReverseSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	return arr
}

// Função para rodar o algoritmo de ordenação e coletar métricas
func runSortAlgorithm(sortFunc func([]int) (int, int), arr []int) (time.Duration, int, int) {
	start := time.Now()
	comparisons, swaps := sortFunc(arr)
	duration := time.Since(start)
	return duration, comparisons, swaps
}
