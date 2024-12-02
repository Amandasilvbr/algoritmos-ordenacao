package main

import (
	"math/rand"
	"time"
)

func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return arr
}

func generateReverseSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	return arr
}

func runSortAlgorithm(sortFunc func([]int) (int, int), arr []int) (time.Duration, int, int) {
	start := time.Now()
	comparisons, swaps := sortFunc(arr)
	duration := time.Since(start)
	return duration, comparisons, swaps
}
