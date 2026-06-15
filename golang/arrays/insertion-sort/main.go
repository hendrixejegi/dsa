package main

import (
	"dsa/utils"
	"fmt"
)

// insertionSort sorts the array in-place using the insertion sort algorithm.
// It builds a sorted portion of the array from left to right by taking each
// element and inserting it into its correct position within the already-sorted
// portion. Best case is O(n) when the array is already sorted, and average/worst
// case is O(n^2) due to nested comparisons and shifts.

func insertionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Start from index 1 because we treat index 0 as already "sorted"
	for i := 1; i < len(arr); i++ {

		// This is the value we want to insert into the sorted portion
		key := arr[i]

		// j starts from the element just before i
		j := i - 1

		// Shift elements of the sorted portion to the right
		// as long as they are greater than the key
		for j >= 0 && arr[j] > key {

			// Move the larger element one position to the right
			arr[j+1] = arr[j]

			// Move left in the sorted portion
			j--
		}

		// Insert the key into the correct position
		arr[j+1] = key
	}
}

func main() {
	arr := utils.CreateArr(10)
	fmt.Println(arr)
	utils.Duration(func() { insertionSort(arr) })
	fmt.Println(arr)
}
