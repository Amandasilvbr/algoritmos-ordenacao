package main

// Bubble Sort: Sorts by comparing and swapping adjacent elements until sorted.
func bubbleSort(arr []int) (comparisons, swaps int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			comparisons++
			if arr[j] > arr[j+1] {
				swaps++
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return comparisons, swaps
}

// Improved Bubble Sort: Optimized version that stops if the list is already sorted.
func improvedBubbleSort(arr []int) (comparisons, swaps int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			comparisons++
			if arr[j] > arr[j+1] {
				swaps++
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return comparisons, swaps
}

// Insertion Sort: Inserts elements into the correct position in the sorted list.
func insertionSort(arr []int) (comparisons, swaps int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			comparisons++
			arr[j+1] = arr[j]
			j--
			swaps++
		}
		arr[j+1] = key
	}
	return comparisons, swaps
}

// Selection Sort: Selects the smallest element and places it in the correct position.
func selectionSort(arr []int) (comparisons, swaps int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			comparisons++
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			swaps++
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
	return comparisons, swaps
}

// Merge Sort: Divides and conquers by splitting the list and merging them sorted.
func mergeSort(arr []int) (comparisons, swaps int) {
	if len(arr) <= 1 {
		return 0, 0
	}

	mid := len(arr) / 2
	leftComparisons, leftSwaps := mergeSort(arr[:mid])
	rightComparisons, rightSwaps := mergeSort(arr[mid:])
	_, mergeComparisons := merge(arr[:mid], arr[mid:])

	return leftComparisons + rightComparisons + mergeComparisons, leftSwaps + rightSwaps
}

// Merge: Merges two sorted arrays.
func merge(left, right []int) ([]int, int) {
	result := []int{}
	i, j := 0, 0
	comparisons := 0

	for i < len(left) && j < len(right) {
		comparisons++
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result, comparisons
}

// Quick Sort: Splits the list into sublists and sorts them recursively.
func quickSort(arr []int) (comparisons, swaps int) {
	if len(arr) < 2 {
		return 0, 0
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}
	comparisons = 0
	swaps = 0

	for _, val := range arr[1:] {
		comparisons++
		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	leftComparisons, leftSwaps := quickSort(left)
	rightComparisons, rightSwaps := quickSort(right)

	return comparisons + leftComparisons + rightComparisons, swaps + leftSwaps + rightSwaps
}

// Heap Sort: Builds a heap and extracts the largest element.
func heapSort(arr []int) (comparisons, swaps int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		comparisons += heapify(arr, n, i, &swaps)
	}

	for i := n - 1; i > 0; i-- {
		swaps++
		arr[i], arr[0] = arr[0], arr[i]
		comparisons += heapify(arr, i, 0, &swaps)
	}

	return comparisons, swaps
}

// Heapify: Adjusts the heap to maintain the max-heap property.
func heapify(arr []int, n, i int, swaps *int) int {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	comparisons := 0

	if left < n && arr[left] > arr[largest] {
		comparisons++
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		comparisons++
		largest = right
	}
	if largest != i {
		*swaps++
		arr[i], arr[largest] = arr[largest], arr[i]
		comparisons += heapify(arr, n, largest, swaps)
	}

	return comparisons
}

// Cocktail Sort: A variation of Bubble Sort that alternates the direction of swaps.
func cocktailSort(arr []int) (comparisons, swaps int) {
	n := len(arr)
	start, end := 0, n-1
	swapped := true

	for swapped {
		swapped = false

		// Passagem da esquerda para a direita
		for i := start; i < end; i++ {
			comparisons++
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swaps++
				swapped = true
			}
		}

		if !swapped {
			break
		}

		end--
		swapped = false

		for i := end - 1; i >= start; i-- {
			comparisons++
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swaps++
				swapped = true
			}
		}

		start++
	}

	return comparisons, swaps
}

// Pancake Sort: Sorts by flipping sublists to place the largest element at the front.
func pancakeSort(arr []int) (comparisons, swaps int) {
	n := len(arr)
	for i := n; i > 1; i-- {
		maxIndex, maxComp := findMax(arr, i)
		comparisons += maxComp

		if maxIndex != i-1 {
			if maxIndex > 0 {
				reverse(arr, maxIndex)
				swaps++
			}
			reverse(arr, i-1)
			swaps++
		}
	}
	return comparisons, swaps
}

// FindMax: Finds the index of the maximum element in the first n elements.
func findMax(arr []int, n int) (int, int) {
	maxIndex := 0
	comparisons := 0
	for i := 1; i < n; i++ {
		comparisons++
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex, comparisons
}

// Reverse: Reverses the elements of the array from 0 to n-1.
func reverse(arr []int, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
