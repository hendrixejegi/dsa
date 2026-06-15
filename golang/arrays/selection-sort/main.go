package main

import (
	"dsa/utils"
	"fmt"
)

// selectionSort sorts a slice of integers in ascending order using the
// Selection Sort algorithm.
//
// Idea:
// Repeatedly select the smallest element from the unsorted portion
// of the slice and swap it into its correct position.
//
// Time Complexity:
// The algorithm always uses two nested loops:
// - The outer loop goes through each element of the array.
// - The inner loop scans the remaining unsorted part to find the smallest element.
//
// Because of this, every element is compared with most of the others,
// leading to about n * n comparisons.
//
// So the time complexity is O(n^2) in all cases (best, average, and worst),
// even if the array is already sorted.
//
// Space Complexity:
// O(1) because the sorting is done in-place without using extra memory.

func selectionSort(arr []int) {

	// Outer loop: moves the boundary of the sorted portion of the array.
	// At each iteration, position i is where the next smallest element should go.
	for i := 0; i < len(arr)-1; i++ {

		// Assume the current index i holds the minimum value
		// in the unsorted portion.
		minValueIdx := i

		// Inner loop: scan the rest of the array to find the actual minimum.
		// We start from i+1 because elements before i are already sorted.
		for j := i + 1; j < len(arr); j++ {

			// If we find a smaller element, update the index of the minimum.
			if arr[j] < arr[minValueIdx] {
				minValueIdx = j
			}
		}

		// After finding the minimum element in the unsorted part,
		// swap it with the element at position i.
		if i != minValueIdx {
			arr[i], arr[minValueIdx] = arr[minValueIdx], arr[i]
		}
	}
}

func main() {
	arr := utils.CreateArr(10)
	fmt.Println(arr)
	utils.Duration(func() {
		selectionSort(arr)
	})
	fmt.Println(arr)
}
