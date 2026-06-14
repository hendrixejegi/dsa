package main

import (
	"dsa/utils"
	"fmt"
	"math/rand"
)

// Each pass places the largest unsorted element at the end of the array,
// so the sorted portion grows from the back.
// The outer loop controls how much of the array is still unsorted,
// shrinking the range after each iteration.
// The inner loop performs adjacent comparisons within the unsorted range
// and swaps elements to move larger values toward the end.

// Time Complexity:
// Worst/Average: O(n^2) (array is in reverse order or unsorted)
// Best: O(n) when already sorted (optimized with early termination)

func bubbleSort(arr []int) {
	// After each pass, the largest unsorted element is placed at index i,
	// so we only need to scan up to i-1 on the next pass.
	for i := len(arr) - 1; i > 0; i-- {
		swapped := false

		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}

func main() {
	arr := make([]int, 300)

	for i := range arr {
		arr[i] = rand.Intn(1000)
	}

	utils.Duration(func() {
		bubbleSort(arr)
		fmt.Println(arr)
	})
}
